package giu

import (
	"fmt"
	"log"
	"runtime"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/utils"
	"github.com/AllenDang/go-findfont"
)

const (
	darwin  = "darwin"
	windows = "windows"
	// DefaultFontSize is the default font size used in giu.
	DefaultFontSize = 14
)

// FontInfo represents a giu implementation of imgui font.
type FontInfo struct {
	fontName string
	fontPath string
	fontByte []byte
}

// String returns a string representation of the FontInfo. It is intended to be unique for each FontInfo.
func (f *FontInfo) String() string {
	return f.fontName
}

// FontAtlas is a mechanism to automatically manage fonts in giu.
// When you add a string in your app, it is registered inside the FontAtlas.
// Then, font data are built based on the registered strings.
// for more details, see: https://github.com/ocornut/imgui/blob/master/docs/FONTS.md
// DefaultFont = font that is used for normal rendering.
// ExtraFont = font that can be set and then it'll be used for rendering things.
type FontAtlas struct {
	shouldRebuildFontAtlas bool
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
}

func newFontAtlas() *FontAtlas {
	result := FontAtlas{
		extraFontMap: make(map[string]*imgui.Font),
	}

	result.SetDefaultFontSize(DefaultFontSize)

	// Pre-register fonts
	switch runtime.GOOS {
	case darwin:
		// English font
		result.registerDefaultFont("Menlo")
		// Chinese font
		result.registerDefaultFont("STHeiti")
		// Jananese font
		result.registerDefaultFont("ヒラギノ角ゴシック W0")
		// Korean font
		result.registerDefaultFont("AppleSDGothicNeo")
	case windows:
		// English font
		result.registerDefaultFont("Calibri")
		// Chinese font
		result.registerDefaultFont("MSYH")
		// Japanese font
		result.registerDefaultFont("MSGOTHIC")
		// Korean font
		result.registerDefaultFont("MALGUNSL")
	case "linux":
		// English fonts
		result.registerDefaultFonts([]FontInfo{
			{
				fontName: "FreeSans.ttf",
			},
			{
				fontName: "FiraCode-Medium",
			},
			{
				fontName: "sans",
			},
		})
		// Chinese fonts
		result.registerDefaultFonts([]FontInfo{
			{
				fontName: "wqy-microhei",
			},
			{
				fontName: "SourceHanSansCN",
			},
		})
	}

	return &result
}

// SetDefaultFontSize sets the default font size.
func (a *FontAtlas) SetDefaultFontSize(size float32) {
	imgui.CurrentStyle().SetFontSizeBase(size)
}

// SetDefaultFont changes default font.
func (a *FontAtlas) SetDefaultFont(fontName string) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		log.Fatalf("Cannot find font %s", fontName)
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath}
	a.defaultFonts = append([]FontInfo{fontInfo}, a.defaultFonts...)
}

// SetDefaultFontFromBytes changes default font by bytes of the font file.
func (a *FontAtlas) SetDefaultFontFromBytes(fontBytes []byte) {
	a.defaultFonts = append([]FontInfo{
		{
			fontByte: fontBytes,
		},
	}, a.defaultFonts...)
}

// GetDefaultFonts returns a list of currently loaded default fonts.
func (a *FontAtlas) GetDefaultFonts() []FontInfo {
	return a.defaultFonts
}

// AddFont adds font by name, if the font is found, return *FontInfo, otherwise return nil.
// To use added font, use giu.Style().SetFont(...).
func (a *FontAtlas) AddFont(fontName string) *FontInfo {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		fmt.Printf("[Warning]Cannot find font %s at system, related text will not be rendered.\n", fontName)
		return nil
	}

	fi := FontInfo{
		fontName: fontName,
		fontPath: fontPath,
	}

	a.extraFonts = append(a.extraFonts, fi)

	return &fi
}

// AddFontFromBytes does similar to AddFont, but using data from memory.
func (a *FontAtlas) AddFontFromBytes(fontName string, fontBytes []byte) *FontInfo {
	fi := FontInfo{
		fontName: fontName,
		fontByte: fontBytes,
	}

	a.extraFonts = append(a.extraFonts, fi)

	return &fi
}

func (a *FontAtlas) registerDefaultFont(fontName string) {
	fontPath, err := findfont.Find(fontName)
	if err != nil {
		return
	}

	fontInfo := FontInfo{fontName: fontName, fontPath: fontPath}
	a.defaultFonts = append(a.defaultFonts, fontInfo)
}

func (a *FontAtlas) registerDefaultFonts(fontInfos []FontInfo) {
	var firstFoundFont *FontInfo

	for _, fi := range fontInfos {
		fontPath, err := findfont.Find(fi.fontName)
		if err == nil {
			firstFoundFont = &FontInfo{fontName: fi.fontName, fontPath: fontPath}
			break
		}
	}

	if firstFoundFont != nil {
		a.defaultFonts = append(a.defaultFonts, *firstFoundFont)
	}
}

// Rebuild font atlas when necessary.
// The whole magic happens here.
func (a *FontAtlas) rebuildFontAtlas() {
	if !a.shouldRebuildFontAtlas {
		return
	}

	fonts := Context.IO().Fonts()
	fonts.Clear()

	if len(a.defaultFonts) > 0 {
		fontConfig := imgui.NewFontConfig()
		fontConfig.SetOversampleH(2)
		fontConfig.SetOversampleV(2)
		fontConfig.SetRasterizerMultiply(1.5)

		for i, fontInfo := range a.defaultFonts {
			if i > 0 {
				fontConfig.SetMergeMode(true)
			}

			if len(fontInfo.fontByte) == 0 {
				fonts.AddFontFromFileTTFV(
					fontInfo.fontPath,
					0,
					fontConfig,
					nil,
				)
			} else {
				fontConfig.SetFontDataOwnedByAtlas(false)
				fonts.AddFontFromMemoryTTFV(
					uintptr(unsafe.Pointer(utils.SliceToPtr(fontInfo.fontByte))),
					int32(len(fontInfo.fontByte)),
					0,
					fontConfig,
					nil,
				)
			}
		}

		// Fall back if no font is added
		if fonts.FontCount() == 0 {
			fonts.AddFontDefault()
		}
	} else {
		fonts.AddFontDefault()
	}

	// Add extra fonts
	for _, fontInfo := range a.extraFonts {
		// Store imgui.Font for PushFont
		var f *imgui.Font
		if len(fontInfo.fontByte) == 0 {
			f = fonts.AddFontFromFileTTFV(
				fontInfo.fontPath,
				0,
				imgui.NewFontConfig(),
				nil,
			)
		} else {
			fontConfig := imgui.NewFontConfig()
			fontConfig.SetFontDataOwnedByAtlas(false)
			f = fonts.AddFontFromMemoryTTFV(
				uintptr(unsafe.Pointer(utils.SliceToPtr(fontInfo.fontByte))),
				int32(len(fontInfo.fontByte)),
				0,
				fontConfig,
				nil,
			)
		}

		a.extraFontMap[fontInfo.String()] = f
	}

	a.shouldRebuildFontAtlas = false
}
