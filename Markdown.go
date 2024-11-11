//nolint:gocritic,govet,wsl,revive // this file is TODO. We don't want commentedOutCode lint issues here.
package giu

import (
	"image"
	"net/http"
	"strings"
	"time"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/immarkdown"
)

type markdownState struct {
	cfg immarkdown.MarkdownConfig
}

func (m *markdownState) Dispose() {
	// noop
}

func (m *MarkdownWidget) getState() *markdownState {
	if s := GetState[markdownState](Context, m.id); s != nil {
		return s
	}

	newState := m.newState()
	SetState[markdownState](Context, m.id, newState)
	return newState
}

func (m *MarkdownWidget) newState() *markdownState {
	cfg := immarkdown.NewMarkdownConfigEmpty()
	fmtCb := immarkdown.MarkdownFormalCallback(mdFormatCallback)
	cfg.SetFormatCallback(&fmtCb)
	return &markdownState{
		cfg: *cfg,
	}
}

// MarkdownWidget implements DearImGui markdown extension
// https://github.com/juliettef/imgui_markdown
// It is like LabelWidget but with md formatting.
type MarkdownWidget struct {
	md      string
	id      ID
	linkCb  func(url string)
	headers []immarkdown.MarkdownHeadingFormat
}

// Markdown creates new markdown widget.
func Markdown(md string) *MarkdownWidget {
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
		m.headers = make([]immarkdown.MarkdownHeadingFormat, level)
	}

	if level <= len(m.headers) {
		m.headers = append(m.headers, make([]immarkdown.MarkdownHeadingFormat, len(m.headers)-level+1)...)
	}

	m.headers[level] = *immarkdown.NewMarkdownHeadingFormatEmpty()
	if font != nil {
		if f, ok := Context.FontAtlas.extraFontMap[font.String()]; ok {
			m.headers[level].SetFont(f)
		}
	}

	m.headers[level].SetSeparator(separator)

	return m
}

// Build implements Widget interface.
func (m *MarkdownWidget) Build() {
	state := m.getState()
	immarkdown.Markdown(
		Context.FontAtlas.RegisterString(m.md),
		uint64(len(m.md)),
		state.cfg,
	)
	// m.linkCb, loadImage, m.headers)
}

func loadImage(path string) immarkdown.MarkdownImageData {
	var (
		img *image.RGBA
		err error
	)

	switch {
	case strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://"):
		// Load image from url

		client := &http.Client{Timeout: 5 * time.Second}
		resp, respErr := client.Get(path)
		if respErr != nil {
			return *immarkdown.NewMarkdownImageDataEmpty()
		}

		defer func() {
			closeErr := resp.Body.Close()
			Assert((closeErr == nil), "MarkdownWidget", "loadImage", "Could not close http request!")
		}()

		rgba, _, imgErr := image.Decode(resp.Body)
		if imgErr != nil {
			return *immarkdown.NewMarkdownImageDataEmpty()
		}

		img = ImageToRgba(rgba)
	default:
		img, err = LoadImage(path)
		if err != nil {
			return *immarkdown.NewMarkdownImageDataEmpty()
		}
	}

	size := img.Bounds()

	// if current workaround is save
	/*
		tex := &Texture{}
		NewTextureFromRgba(img, func(t *Texture) {
			fmt.Println("creating texture")
			tex.id = t.id
		})
	*/

	var id imgui.TextureID

	mainthreadCallPlatform(func() {
		// TODO: actually load this hehe
		/*
			var err error
			id, err = Context.renderer.LoadImage(img)
			if err != nil {
				return
			}
		*/
	})

	result := immarkdown.NewMarkdownImageDataEmpty()
	result.SetUsertextureid(id)
	//		Scale:     true,
	result.SetSize(imgui.Vec2{
		X: float32(size.Dx()),
		Y: float32(size.Dy()),
	})
	//		UseLinkCallback: true,
	//		default values
	//
	// Uv0:         ToVec2(image.Point{0, 0}),
	// Uv1:         ToVec2(image.Point{1, 1}),
	// TintColor:   ToVec4Color(color.RGBA{255, 255, 255, 255}),
	// BorderColor: ToVec4Color(color.RGBA{0, 0, 0, 0}),
	return *result
}

func mdFormatCallback(f immarkdown.MarkdownFormatInfo, start bool) {
	// noop
}
