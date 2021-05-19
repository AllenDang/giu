package giu

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/AllenDang/go-findfont"
	"github.com/AllenDang/imgui-go"
)

var (
	shouldRebuildFontAtlas bool
	stringMap              map[rune]bool // key is rune, value indicates whether it's a new rune.
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
)

type FontInfo struct {
	fontName string
	fontPath string
	size     float32
}

func (f *FontInfo) String() string {
	return fmt.Sprintf("%s:%.2f", f.fontName, f.size)
}

func init() {
	stringMap = make(map[rune]bool)
	extraFontMap = make(map[string]*imgui.Font)

	// Pre register numbers
	tStr("01234567890.")

	// Pre-register fonts
	os := runtime.GOOS
	switch os {
	case "darwin":
		// English font
		registerDefaultFont("Menlo", 14)
		// Chinese font
		registerDefaultFont("PingFang", 17)
		// Jananese font
		registerDefaultFont("ヒラギノ角ゴシック W0", 17)
		// Korean font
		registerDefaultFont("AppleSDGothicNeo", 16)
		// TODO add more fonts for different languages.
	case "windows":
		// English font
		registerDefaultFont("Calibri", 15)
		// Chinese font
		registerDefaultFont("MSYH", 15)
		// Japanese font
		registerDefaultFont("Meiryo", 15)
		// TODO add more fonts for different languages.
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

func registerDefaultFont(fontName string, size float32) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		fmt.Printf("[Warning]Cannot find font %s at system, related text will not be rendered.\n", fontName)
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath, size: size}
	defaultFonts = append(defaultFonts, fontInfo)
}

// Register string to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStr(str string) string {
	for _, s := range str {
		if _, ok := stringMap[s]; !ok {
			stringMap[s] = false
			shouldRebuildFontAtlas = true
		}
	}

	return str
}

// Register string to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStrPtr(str *string) *string {
	for _, s := range *str {
		if _, ok := stringMap[s]; !ok {
			stringMap[s] = false
			shouldRebuildFontAtlas = true
		}
	}

	return str
}

// Rebuild font atlas when necessary.
func rebuildFontAtlas() {
	if len(defaultFonts) == 0 {
		return
	}

	if shouldRebuildFontAtlas {
		fonts := Context.IO().Fonts()

		var sb strings.Builder

		for k := range stringMap {
			stringMap[k] = true
			sb.WriteRune(k)
		}

		fonts.Clear()

		ranges := imgui.NewGlyphRanges()
		builder := imgui.NewFontGlyphRangesBuilder()
		if sb.Len() == 0 {
			builder.AddRanges(fonts.GlyphRangesDefault())
		}

		builder.AddText(sb.String())
		builder.BuildRanges(ranges)

		for i, fontInfo := range defaultFonts {
			size := fontInfo.size

			if runtime.GOOS != "darwin" {
				// Scale font size based on DPI scaling.
				size *= Context.platform.GetContentScale()
			}

			fontConfig := imgui.NewFontConfig()
			fontConfig.SetOversampleH(2)
			fontConfig.SetOversampleV(2)
			if i == 0 {
				fontConfig.SetMergeMode(false)
				fonts.AddFontFromFileTTFV(fontInfo.fontPath, size, fontConfig, ranges.Data())
			} else {

				fontConfig.SetMergeMode(true)
				fonts.AddFontFromFileTTFV(fontInfo.fontPath, size, fontConfig, ranges.Data())
			}
		}

		// Add extra fonts
		for _, fontInfo := range extraFonts {
			// Store imgui.Font for PushFont
			f := fonts.AddFontFromFileTTFV(fontInfo.fontPath, fontInfo.size, imgui.DefaultFontConfig, ranges.Data())
			extraFontMap[fontInfo.String()] = &f
		}

		fontTextureImg := fonts.TextureDataRGBA32()
		Context.renderer.SetFontTexture(fontTextureImg)

		shouldRebuildFontAtlas = false
	}
}
