package giu

import (
	"image"
	"image/color"
	"net/http"
	"strings"
	"time"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/faiface/mainthread"
)

// MarkdownWidget implements DearImGui markdown extension
// https://github.com/juliettef/imgui_markdown
// It is like LabelWidget but with md formatting.
type MarkdownWidget struct {
	md      *string
	linkCb  func(url string)
	headers []imgui.MarkdownHeadingFormat
}

// Markdown creates new markdown widget.
func Markdown(md *string) *MarkdownWidget {
	return &MarkdownWidget{
		md:     md,
		linkCb: OpenURL,
	}
}

// OnLink sets another than default link callback.
func (m *MarkdownWidget) OnLink(cb func(url string)) *MarkdownWidget {
	m.linkCb = cb
	return m
}

// Header sets header formatting
// NOTE: level (counting from 0!) is header level. (for instance, header `# H1` will have level 0).
func (m *MarkdownWidget) Header(level int, font *FontInfo, separator bool) *MarkdownWidget {
	// ensure if header data are at least as long as level
	if m.headers == nil {
		m.headers = make([]imgui.MarkdownHeadingFormat, level)
	}

	if level <= len(m.headers) {
		m.headers = append(m.headers, make([]imgui.MarkdownHeadingFormat, len(m.headers)-level+1)...)
	}

	if font != nil {
		if f, ok := Context.FontAtlas.extraFontMap[font.String()]; ok {
			m.headers[level].FieldFont = f
		}
	}

	m.headers[level].FieldSeparator = separator

	return m
}

// Build implements Widget interface.
func (m *MarkdownWidget) Build() {
	cfg := imgui.MarkdownConfig{}
	imgui.Markdown(
		*Context.FontAtlas.RegisterStringPointer(m.md),
		uint64(len(*m.md)),
		cfg, // TODO: implement callbacks
	)
}

func loadImage(path string) imgui.MarkdownImageData {
	var img *image.RGBA

	var err error

	switch {
	case strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://"):
		// Load image from url
		client := &http.Client{Timeout: 5 * time.Second}
		resp, respErr := client.Get(path)

		if respErr != nil {
			return imgui.MarkdownImageData{}
		}

		defer func() {
			closeErr := resp.Body.Close()
			Assert((closeErr == nil), "MarkdownWidget", "loadImage", "Could not close http request!")
		}()

		rgba, _, imgErr := image.Decode(resp.Body)
		if imgErr != nil {
			return imgui.MarkdownImageData{}
		}

		img = ImageToRgba(rgba)
	default:
		img, err = LoadImage(path)
		if err != nil {
			return imgui.MarkdownImageData{}
		}
	}

	size := img.Bounds()

	//nolint:gocritic // TODO/BUG: figure out, why it doesn't work as expected and consider
	//if current workaround is save
	/*
		tex := &Texture{}
		NewTextureFromRgba(img, func(t *Texture) {
			fmt.Println("creating texture")
			tex.id = t.id
		})
	*/

	var id imgui.TextureID

	mainthread.Call(func() {
		id = Context.backend.CreateTextureRgba(img, img.Bounds().Dx(), img.Bounds().Dy())
		if id == nil {
			return
		}
	})

	return imgui.MarkdownImageData{
		FieldUser_texture_id: id,
		FieldSize: imgui.Vec2{
			X: float32(size.Dx()),
			Y: float32(size.Dy()),
		},
		FieldUseLinkCallback: true,
		// default values
		FieldUv0:        ToVec2(image.Point{0, 0}),
		FieldUv1:        ToVec2(image.Point{1, 1}),
		FieldTint_col:   ToVec4Color(color.RGBA{255, 255, 255, 255}),
		FieldBorder_col: ToVec4Color(color.RGBA{0, 0, 0, 0}),
		FieldIsValid:    true,
	}
}
