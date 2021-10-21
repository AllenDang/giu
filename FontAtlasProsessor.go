package giu

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"

	"github.com/AllenDang/go-findfont"
	"github.com/AllenDang/imgui-go"
)

const (
	preRegisterString = " \"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	windows           = "windows"
)

// FontInfo represents a giu implementation of imgui font.
type FontInfo struct {
	fontName string
	fontPath string
	fontByte []byte
	size     float32
}

func (f *FontInfo) String() string {
	return fmt.Sprintf("%s:%.2f", f.fontName, f.size)
}

type FontAtlas struct {
	shouldRebuildFontAtlas bool
	stringMap              sync.Map // key is rune, value indicates whether it's a new rune.
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
}

func newFontAtlas() FontAtlas {
	result := FontAtlas{
		extraFontMap: make(map[string]*imgui.Font),
	}

	// Pre register numbers
	result.tStr(preRegisterString)

	// Pre-register fonts
	switch runtime.GOOS {
	case "darwin":
		// English font
		result.registerDefaultFont("Menlo", 14)
		// Chinese font
		result.registerDefaultFont("STHeiti", 13)
		// Jananese font
		result.registerDefaultFont("ヒラギノ角ゴシック W0", 17)
		// Korean font
		result.registerDefaultFont("AppleSDGothicNeo", 16)
	case windows:
		// English font
		result.registerDefaultFont("Calibri", 16)
		// Chinese font
		result.registerDefaultFont("MSYH", 16)
		// Japanese font
		result.registerDefaultFont("MSGOTHIC", 16)
		// Korean font
		result.registerDefaultFont("MALGUNSL", 16)
	case "linux":
		// English fonts
		result.registerDefaultFonts([]FontInfo{
			{
				fontName: "FreeSans.ttf",
				size:     15,
			},
			{
				fontName: "FiraCode-Medium",
				size:     15,
			},
		})
	}

	return result
}

// SetDefaultFont changes default font.
func (a *FontAtlas) SetDefaultFont(fontName string, size float32) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		log.Fatalf("Cannot find font %s", fontName)
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath, size: size}
	a.defaultFonts = append([]FontInfo{fontInfo}, a.defaultFonts...)
}

// SetDefaultFontFromBytes changes default font by bytes of the font file.
func (a *FontAtlas) SetDefaultFontFromBytes(fontBytes []byte, size float32) {
	a.defaultFonts = append([]FontInfo{
		{
			fontByte: fontBytes,
			size:     size,
		},
	}, a.defaultFonts...)
}

// AddFont adds font by name, if the font is found, return *FontInfo, otherwise return nil.
// To use added font, use giu.Style().SetFont(...).
func (a *FontAtlas) AddFont(fontName string, size float32) *FontInfo {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		fmt.Printf("[Warning]Cannot find font %s at system, related text will not be rendered.\n", fontName)
		return nil
	}

	fi := FontInfo{
		fontName: fontName,
		fontPath: fontPath,
		size:     size,
	}

	a.extraFonts = append(a.extraFonts, fi)

	return &fi
}

// AddFontFromBytes does similar to AddFont, but using data from memory.
func (a *FontAtlas) AddFontFromBytes(fontName string, fontBytes []byte, size float32) *FontInfo {
	fi := FontInfo{
		fontName: fontName,
		fontByte: fontBytes,
		size:     size,
	}

	a.extraFonts = append(a.extraFonts, fi)

	return &fi
}

func (a *FontAtlas) registerDefaultFont(fontName string, size float32) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath, size: size}
	a.defaultFonts = append(a.defaultFonts, fontInfo)
}

func (a *FontAtlas) registerDefaultFonts(fontInfos []FontInfo) {
	var firstFoundFont *FontInfo
	for _, fi := range fontInfos {
		fontPath, err := findfont.Find(fi.fontName)
		if err == nil {
			firstFoundFont = &FontInfo{fontName: fi.fontName, fontPath: fontPath, size: fi.size}
			break
		}
	}

	if firstFoundFont != nil {
		a.defaultFonts = append(a.defaultFonts, *firstFoundFont)
	}
}

// Register string to font atlas builder.
// Note only register strings that will be displayed on the UI.
func (a *FontAtlas) tStr(str string) string {
	for _, s := range str {
		if _, ok := a.stringMap.Load(s); !ok {
			a.stringMap.Store(s, false)
			a.shouldRebuildFontAtlas = true
		}
	}

	return str
}

// Register string pointer to font atlas builder.
// Note only register strings that will be displayed on the UI.
func (a *FontAtlas) tStrPtr(str *string) *string {
	a.tStr(*str)
	return str
}

func (a *FontAtlas) tStrSlice(str []string) []string {
	for _, s := range str {
		a.tStr(s)
	}

	return str
}

// Rebuild font atlas when necessary.
func (a *FontAtlas) rebuildFontAtlas() {
	if !a.shouldRebuildFontAtlas {
		return
	}

	fonts := Context.IO().Fonts()
	fonts.Clear()

	var sb strings.Builder

	a.stringMap.Range(func(k, v interface{}) bool {
		a.stringMap.Store(k, true)
		if ks, ok := k.(rune); ok {
			sb.WriteRune(ks)
		}

		return true
	})

	ranges := imgui.NewGlyphRanges()
	builder := imgui.NewFontGlyphRangesBuilder()

	// Because we pre-regestered numbers, so default string map's length should greater then 11.
	if sb.Len() > len(preRegisterString) {
		builder.AddText(sb.String())
	} else {
		builder.AddRanges(fonts.GlyphRangesDefault())
	}

	builder.BuildRanges(ranges)

	if len(a.defaultFonts) > 0 {
		fontConfig := imgui.NewFontConfig()
		fontConfig.SetOversampleH(2)
		fontConfig.SetOversampleV(2)
		fontConfig.SetRasterizerMultiply(1.5)

		for i, fontInfo := range a.defaultFonts {
			if i > 0 {
				fontConfig.SetMergeMode(true)
			}

			// Scale font size with DPI scale factor
			if runtime.GOOS == windows {
				fontInfo.size *= Context.GetPlatform().GetContentScale()
			}

			if len(fontInfo.fontByte) == 0 {
				fonts.AddFontFromFileTTFV(fontInfo.fontPath, fontInfo.size, fontConfig, ranges.Data())
			} else {
				fonts.AddFontFromMemoryTTFV(fontInfo.fontByte, fontInfo.size, fontConfig, ranges.Data())
			}
		}

		// Fall back if no font is added
		if fonts.GetFontCount() == 0 {
			fonts.AddFontDefault()
		}
	} else {
		fonts.AddFontDefault()
	}

	// Add extra fonts
	for _, fontInfo := range a.extraFonts {
		// Scale font size with DPI scale factor
		if runtime.GOOS == windows {
			fontInfo.size *= Context.GetPlatform().GetContentScale()
		}

		// Store imgui.Font for PushFont
		var f imgui.Font
		if len(fontInfo.fontByte) == 0 {
			f = fonts.AddFontFromFileTTFV(fontInfo.fontPath, fontInfo.size, imgui.DefaultFontConfig, ranges.Data())
		} else {
			f = fonts.AddFontFromMemoryTTFV(fontInfo.fontByte, fontInfo.size, imgui.DefaultFontConfig, ranges.Data())
		}
		a.extraFontMap[fontInfo.String()] = &f
	}

	fontTextureImg := fonts.TextureDataRGBA32()
	Context.renderer.SetFontTexture(fontTextureImg)

	a.shouldRebuildFontAtlas = false
}
