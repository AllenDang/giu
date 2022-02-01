package giu

import (
	"image"
	"image/color"
	"net/http"
	"strings"
	"time"

	"github.com/AllenDang/imgui-go"
	"github.com/faiface/mainthread"
)

// MarkdownWidget implements DearImGui markdown extension
// https://github.com/juliettef/imgui_markdown
// It is like LabelWidget but with md formatting.
type MarkdownWidget struct {
	md      *string
	linkCb  func(url string)
	headers []imgui.MarkdownHeaderData
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
		m.headers = make([]imgui.MarkdownHeaderData, level)
	}

	if level <= len(m.headers) {
		m.headers = append(m.headers, make([]imgui.MarkdownHeaderData, len(m.headers)-level+1)...)
	}

	if font != nil {
		if f, ok := extraFontMap[font.String()]; ok {
			m.headers[level].Font = *f
		}
	}

	m.headers[level].HasSeparator = separator

	return m
}

// Build implements Widget interface.
func (m *MarkdownWidget) Build() {
	imgui.Markdown(tStrPtr(m.md), m.linkCb, loadImage, m.headers)
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

	// nolint:gocritic // TODO/BUG: figure out, why it doesn't work as expected and consider
	// if current workaround is save
	/*
		tex := &Texture{}
		NewTextureFromRgba(img, func(t *Texture) {
			fmt.Println("creating texture")
			tex.id = t.id
		})
	*/

	var id imgui.TextureID
	mainthread.Call(func() {
		var err error
		id, err = Context.renderer.LoadImage(img)
		if err != nil {
			return
		}
	})

	return imgui.MarkdownImageData{
		TextureID: &id,
		Scale:     true,
		Size: imgui.Vec2{
			X: float32(size.Dx()),
			Y: float32(size.Dy()),
		},
		UseLinkCallback: true,
		// default values
		Uv0:         ToVec2(image.Point{0, 0}),
		Uv1:         ToVec2(image.Point{1, 1}),
		TintColor:   ToVec4Color(color.RGBA{255, 255, 255, 255}),
		BorderColor: ToVec4Color(color.RGBA{0, 0, 0, 0}),
	}
}
