package giu

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/utils"
	"github.com/AllenDang/go-findfont"
)

const (
	preRegisterString = " \"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	darwin            = "darwin"
	windows           = "windows"
	defaultFontSize   = 14
)

// FontInfo represents a giu implementation of imgui font.
type FontInfo struct {
	fontName string
	fontPath string
	fontByte []byte
}

// String returns a string representation of the FontInfo. It is intended to be unique for each FontInfo.
func (f *FontInfo) String() string {
	return fmt.Sprintf("%s", f.fontName)
}

// FontAtlas is a mechanism to automatically manage fonts in giu.
// When you add a string in your app, it is registered inside the FontAtlas.
// Then, font data are built based on the registered strings.
// for more details, see: https://github.com/ocornut/imgui/blob/master/docs/FONTS.md
// DefaultFont = font that is used for normal rendering.
// ExtraFont = font that can be set and then it'll be used for rendering things.
type FontAtlas struct {
	shouldRebuildFontAtlas bool
	stringMap              sync.Map // key is rune, value indicates whether it's a new rune.
	defaultFonts           []FontInfo
	extraFonts             []FontInfo
	extraFontMap           map[string]*imgui.Font
	autoRegisterStrings    bool
}

func newFontAtlas() *FontAtlas {
	result := FontAtlas{
		extraFontMap:        make(map[string]*imgui.Font),
		autoRegisterStrings: true,
	}

	result.SetDefaultFontSize(defaultFontSize)

	// Pre register numbers
	result.RegisterString(preRegisterString)

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

// AutoRegisterStrings if enabled, all strings visible in the UI will be automatically registered and the font atlas will be rebuilt accordingly.
// Generally it is recommended to keep this on as long as you are not using some giant strings (e.g. 23k lines in CodeEditor).
// If you disable this, make sure to use PreRegisterString to register all runes you need (all calls to RegisterString* will be ignored!).
func (a *FontAtlas) AutoRegisterStrings(b bool) {
	a.autoRegisterStrings = b
}

// SetDefaultFontSize sets the default font size.
func (a *FontAtlas) SetDefaultFontSize(size float32) {
	// imgui.CurrentContext().SetFontSizeBase(size)
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

// RegisterString is mainly used by widgets to register strings.
// It could be disabled by AutoRegisterStrings.
func (a *FontAtlas) RegisterString(str string) string {
	if !a.autoRegisterStrings {
		return str
	}

	return a.PreRegisterString(str)
}

// PreRegisterString register string to font atlas builder.
// NOTE only register strings that will be displayed on the UI.
func (a *FontAtlas) PreRegisterString(str string) string {
	for _, s := range str {
		if _, ok := a.stringMap.Load(s); !ok {
			a.stringMap.Store(s, false)
			a.shouldRebuildFontAtlas = true
		}
	}

	return str
}

// RegisterStringPointer registers string pointer to font atlas builder.
// Note only register strings that will be displayed on the UI.
func (a *FontAtlas) RegisterStringPointer(str *string) *string {
	a.RegisterString(*str)
	return str
}

// RegisterStringSlice calls RegisterString for each slice element.
func (a *FontAtlas) RegisterStringSlice(str []string) []string {
	for _, s := range str {
		a.RegisterString(s)
	}

	return str
}

// Rebuild font atlas when necessary.
// The whole magic happens here.
func (a *FontAtlas) rebuildFontAtlas() {
	if !a.shouldRebuildFontAtlas {
		return
	}

	fonts := Context.IO().Fonts()
	fonts.Clear()

	var sb strings.Builder

	a.stringMap.Range(func(k, _ any) bool {
		a.stringMap.Store(k, true)

		if ks, ok := k.(rune); ok {
			sb.WriteRune(ks)
		}

		return true
	})

	ranges := imgui.NewGlyphRange()
	builder := imgui.NewFontGlyphRangesBuilder()

	// Because we pre-registered numbers, so default string map's length should greater then 11.
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

			if len(fontInfo.fontByte) == 0 {
				fonts.AddFontFromFileTTFV(
					fontInfo.fontPath,
					0,
					fontConfig,
					ranges.Data(),
				)
			} else {
				fontConfig.SetFontDataOwnedByAtlas(false)
				fonts.AddFontFromMemoryTTFV(
					uintptr(unsafe.Pointer(utils.SliceToPtr(fontInfo.fontByte))),
					int32(len(fontInfo.fontByte)),
					0,
					fontConfig,
					ranges.Data(),
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
				ranges.Data(),
			)
		} else {
			fontConfig := imgui.NewFontConfig()
			fontConfig.SetFontDataOwnedByAtlas(false)
			f = fonts.AddFontFromMemoryTTFV(
				uintptr(unsafe.Pointer(utils.SliceToPtr(fontInfo.fontByte))),
				int32(len(fontInfo.fontByte)),
				0,
				fontConfig,
				ranges.Data(),
			)
		}

		a.extraFontMap[fontInfo.String()] = f
	}

	a.shouldRebuildFontAtlas = false
}
