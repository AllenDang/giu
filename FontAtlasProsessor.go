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

var (
	shouldRebuildFontAtlas bool
	defaultFontSize        float32  = 14
	stringMap              sync.Map // key is rune, value indicates whether it's a new rune.
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
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

func (f *FontInfo) SetSize(size float32) *FontInfo {
	result := *f
	result.size = size

	for _, i := range extraFonts {
		if i.String() == result.String() {
			return &result
		}
	}

	extraFonts = append(extraFonts, result)
	shouldRebuildFontAtlas = true

	return &result
}

func initFontAtlasProcessor() {
	extraFontMap = make(map[string]*imgui.Font)

	// Pre register numbers
	tStr(preRegisterString)

	// Pre-register fonts
	os := runtime.GOOS
	switch os {
	case "darwin":
		// English font
		registerDefaultFont("Menlo", defaultFontSize)
		// Chinese font
		registerDefaultFont("STHeiti", defaultFontSize-1)
		// Jananese font
		registerDefaultFont("ヒラギノ角ゴシック W0", defaultFontSize+3)
		// Korean font
		registerDefaultFont("AppleSDGothicNeo", defaultFontSize+2)
	case windows:
		// English font
		registerDefaultFont("Calibri", defaultFontSize+2)
		// Chinese font
		registerDefaultFont("MSYH", defaultFontSize+2)
		// Japanese font
		registerDefaultFont("MSGOTHIC", defaultFontSize+2)
		// Korean font
		registerDefaultFont("MALGUNSL", defaultFontSize+2)
	case "linux":
		// English fonts
		registerDefaultFonts([]FontInfo{
			{
				fontName: "FreeSans.ttf",
				size:     defaultFontSize + 1,
			},
			{
				fontName: "FiraCode-Medium",
				size:     defaultFontSize + 1,
			},
			{
				fontName: "sans",
				size:     defaultFontSize + 1,
			},
		})
	}
}

// SetDefaultFontSize sets the default font size. Invoke this before MasterWindow.NewMasterWindow(..).
func SetDefaultFontSize(size float32) {
	defaultFontSize = size
}

// SetDefaultFont changes default font.
func SetDefaultFont(fontName string, size float32) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		log.Fatalf("Cannot find font %s", fontName)
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath, size: size}
	defaultFonts = append([]FontInfo{fontInfo}, defaultFonts...)
}

// SetDefaultFontFromBytes changes default font by bytes of the font file.
func SetDefaultFontFromBytes(fontBytes []byte, size float32) {
	defaultFonts = append([]FontInfo{
		{
			fontByte: fontBytes,
			size:     size,
		},
	}, defaultFonts...)
}

func GetDefaultFonts() []FontInfo {
	return defaultFonts
}

// AddFont adds font by name, if the font is found, return *FontInfo, otherwise return nil.
// To use added font, use giu.Style().SetFont(...).
func AddFont(fontName string, size float32) *FontInfo {
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

	extraFonts = append(extraFonts, fi)

	return &fi
}

// AddFontFromBytes does similar to AddFont, but using data from memory.
func AddFontFromBytes(fontName string, fontBytes []byte, size float32) *FontInfo {
	fi := FontInfo{
		fontName: fontName,
		fontByte: fontBytes,
		size:     size,
	}

	extraFonts = append(extraFonts, fi)

	return &fi
}

func registerDefaultFont(fontName string, size float32) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath, size: size}
	defaultFonts = append(defaultFonts, fontInfo)
}

func registerDefaultFonts(fontInfos []FontInfo) {
	var firstFoundFont *FontInfo
	for _, fi := range fontInfos {
		fontPath, err := findfont.Find(fi.fontName)
		if err == nil {
			firstFoundFont = &FontInfo{fontName: fi.fontName, fontPath: fontPath, size: fi.size}
			break
		}
	}

	if firstFoundFont != nil {
		defaultFonts = append(defaultFonts, *firstFoundFont)
	}
}

// Register string to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStr(str string) string {
	for _, s := range str {
		if _, ok := stringMap.Load(s); !ok {
			stringMap.Store(s, false)
			shouldRebuildFontAtlas = true
		}
	}

	return str
}

// Register string pointer to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStrPtr(str *string) *string {
	tStr(*str)
	return str
}

func tStrSlice(str []string) []string {
	for _, s := range str {
		tStr(s)
	}

	return str
}

// Rebuild font atlas when necessary.
func rebuildFontAtlas() {
	if !shouldRebuildFontAtlas {
		return
	}

	fonts := Context.IO().Fonts()
	fonts.Clear()

	var sb strings.Builder

	stringMap.Range(func(k, v any) bool {
		stringMap.Store(k, true)
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

	if len(defaultFonts) > 0 {
		fontConfig := imgui.NewFontConfig()
		fontConfig.SetOversampleH(2)
		fontConfig.SetOversampleV(2)
		fontConfig.SetRasterizerMultiply(1.5)

		for i, fontInfo := range defaultFonts {
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
	for _, fontInfo := range extraFonts {
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

		extraFontMap[fontInfo.String()] = &f
	}

	fontTextureImg := fonts.TextureDataRGBA32()
	Context.renderer.SetFontTexture(fontTextureImg)

	shouldRebuildFontAtlas = false
}
