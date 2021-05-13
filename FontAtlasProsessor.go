package giu

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/AllenDang/go-findfont"
	"github.com/AllenDang/imgui-go"
)

var (
	stringMap    map[rune]bool // key is rune, value indicates whether it's a new rune.
	defaultFonts []FontInfo
	extraFonts   []FontInfo
	extraFontMap map[string]*imgui.Font
)

type FontInfo struct {
	fontName string
	size     float32
}

func (f *FontInfo) String() string {
	return fmt.Sprintf("%s:%.2f", f.fontName, f.size)
}

func init() {
	stringMap = make(map[rune]bool)
	extraFontMap = make(map[string]*imgui.Font)

	// Pre-register fonts
	os := runtime.GOOS
	switch os {
	case "darwin":
		// English font
		registerDefaultFont("Menlo", 14)
		// Chinese font
		registerDefaultFont("PingFang", 15)
		// Jananese font
		registerDefaultFont("ヒラギノ角ゴシック W0", 15)
		// Korean font
		registerDefaultFont("AppleSDGothicNeo", 15)
		// TODO add more fonts for different languages.
	case "windows":
		// English font
		registerDefaultFont("Calibri", 14)
		// Chinese font
		registerDefaultFont("MSYH", 15)
		// TODO add more fonts for different languages.
	}
}

func AddFont(fontName string, size float32) *FontInfo {
	fi := FontInfo{
		fontName: fontName,
		size:     size,
	}
	extraFonts = append(extraFonts, fi)

	return &fi
}

func registerDefaultFont(fontName string, size float32) {
	_, err := findfont.Find(fontName)
	if err != nil {
		fmt.Printf("[Warning]Cannot find font %s at system, related text will not be rendered.\n", fontName)
		return
	}

	fontInfo := FontInfo{fontName: fontName, size: size}
	defaultFonts = append(defaultFonts, fontInfo)
}

// Register string to font atlas builder.
// Note only register strings that will be displayed on the UI.
func tStr(str string) string {
	for _, s := range str {
		if _, ok := stringMap[s]; !ok {
			stringMap[s] = false
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
		}
	}

	return str
}

// Rebuild font atlas when necessary.
func rebuildFontAtlas() {
	if len(defaultFonts) == 0 {
		return
	}

	shouldRebuild := false

	for _, v := range stringMap {
		if !v {
			shouldRebuild = true
			break
		}
	}

	if shouldRebuild {
		fonts := Context.IO().Fonts()

		var sb strings.Builder

		for k := range stringMap {
			stringMap[k] = true
			sb.WriteRune(k)
		}

		fonts.Clear()

		ranges := imgui.NewGlyphRanges()
		builder := imgui.NewFontGlyphRangesBuilder()
		builder.AddRanges(fonts.GlyphRangesDefault())
		builder.AddText(sb.String())
		builder.BuildRanges(ranges)

		for i, fontInfo := range defaultFonts {
			fontName := fontInfo.fontName
			size := fontInfo.size

			fontPath := findFontPath(fontName)

			if i == 0 {
				fonts.AddFontFromFileTTFV(fontPath, size, imgui.DefaultFontConfig, ranges.Data())
			} else {
				fontConfig := imgui.NewFontConfig()
				fontConfig.SetMergeMode(true)
				fonts.AddFontFromFileTTFV(fontPath, float32(size), fontConfig, ranges.Data())
			}
		}

		// Add extra fonts
		for _, fontInfo := range extraFonts {
			fontPath := findFontPath(fontInfo.fontName)

			// Store imgui.Font for PushFont
			f := fonts.AddFontFromFileTTFV(fontPath, fontInfo.size, imgui.DefaultFontConfig, ranges.Data())
			extraFontMap[fontInfo.String()] = &f
		}

		fontTextureImg := fonts.TextureDataRGBA32()
		Context.renderer.SetFontTexture(fontTextureImg)
	}
}

func findFontPath(fontName string) string {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot find font %s", fontName))
	}

	return fontPath
}
