package giu

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/AllenDang/go-findfont"
	"github.com/AllenDang/imgui-go"
)

var (
	shouldRebuildFontAtlas bool
	stringMap              map[rune]bool // key is rune, value indicates whether it's a new rune.
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
	fontAtlasMutex         = &sync.Mutex{}
)

const (
	preRegisterString = "01234567890.*"
)

type FontInfo struct {
	fontName string
	fontPath string
	fontByte []byte
	size     float32
}

func (f *FontInfo) String() string {
	return fmt.Sprintf("%s:%.2f", f.fontName, f.size)
}

func init() {
	stringMap = make(map[rune]bool)
	extraFontMap = make(map[string]*imgui.Font)

	// Pre register numbers
	tStr(preRegisterString)

	// Pre-register fonts
	os := runtime.GOOS
	switch os {
	case "darwin":
		// English font
		registerDefaultFont("Menlo", 14)
		// Chinese font
		registerDefaultFont("STHeiti", 13)
		// Jananese font
		registerDefaultFont("ヒラギノ角ゴシック W0", 17)
		// Korean font
		registerDefaultFont("AppleSDGothicNeo", 16)
		// TODO add more fonts for different languages.
	case "windows":
		// English font
		registerDefaultFont("Calibri", 16)
		// Chinese font
		registerDefaultFont("MSYH", 16)
		// Japanese font
		registerDefaultFont("MSGOTHIC", 16)
		// Korean font
		registerDefaultFont("MALGUNSL", 16)
		// TODO add more fonts for different languages.
	case "linux":
		// English fonts
		registerDefaultFonts([]FontInfo{
			{
				fontName: "FiraCode-Medium",
				size:     15,
			},
			{
				fontName: "Cantarell",
				size:     15,
			},
		})
	}
}

// Add font by name, if the font is found, return *FontInfo, otherwise return nil.
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
		fontAtlasMutex.Lock()

		if _, ok := stringMap[s]; !ok {
			stringMap[s] = false
			shouldRebuildFontAtlas = true
		}

		fontAtlasMutex.Unlock()
	}

	return str
}

// Register string pointer to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStrPtr(str *string) *string {
	tStr(*str)

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

	for k := range stringMap {
		stringMap[k] = true
		sb.WriteRune(k)
	}

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

			fonts.AddFontFromFileTTFV(fontInfo.fontPath, fontInfo.size, fontConfig, ranges.Data())
		}
	} else {
		fonts.AddFontDefault()
	}

	// Add extra fonts
	for _, fontInfo := range extraFonts {
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
