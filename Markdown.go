package giu

import (
	"image"

	"github.com/AllenDang/imgui-go"
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
	img, err := LoadImage(path)
	if err != nil {
		return imgui.MarkdownImageData{}
	}

	size := img.Bounds()
	// scale image to not exceed available region
	availableW, _ := GetAvailableRegion()
	if x := float32(size.Dx()); x > availableW {
		size = image.Rect(0, 0,
			int(availableW),
			int(float32(size.Dy())*availableW/x),
		)
	}

	// nolint:gocritic // TODO: figure out, why it doesn't work as expected and consider
	// if current workaround is save
	/*
		tex := &Texture{}
		NewTextureFromRgba(img, func(t *Texture) {
			fmt.Println("creating texture")
			tex.id = t.id
		})
	*/

	id, err := Context.renderer.LoadImage(img)
	if err != nil {
		return imgui.MarkdownImageData{}
	}

	return imgui.MarkdownImageData{
		TextureID: &id,
		Size: imgui.Vec2{
			X: float32(size.Dx()),
			Y: float32(size.Dy()),
		},
		UseLinkCallback: true,
	}
}
