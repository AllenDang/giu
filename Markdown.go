package giu

import (
	"image"
	"image/color"
	"net/http"
	"strings"
	"time"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/immarkdown"
)

type markdownState struct {
	cfg    immarkdown.MarkdownConfig
	images map[string]immarkdown.MarkdownImageData
}

func (m *markdownState) Dispose() {
	// noop
}

// MarkdownWidget implements DearImGui markdown extension
// https://github.com/juliettef/imgui_markdown
// It is like LabelWidget but with md formatting.
type MarkdownWidget struct {
	md      string
	id      ID
	headers [3]immarkdown.MarkdownHeadingFormat
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
	cfg := immarkdown.NewEmptyMarkdownConfig()
	fmtCb := immarkdown.MarkdownFormalCallback(func(data *immarkdown.MarkdownFormatInfo, start bool) {
		immarkdown.DefaultMarkdownFormatCallback(*data, start)
	})

	cfg.SetFormatCallback(&fmtCb)

	imgCb := immarkdown.MarkdownImageCallback(func(data immarkdown.MarkdownLinkCallbackData) immarkdown.MarkdownImageData {
		link := data.Link()[:data.LinkLength()] // this is because imgui_markdown returns the whole text starting on link and returns link length (for some reason)
		if existing, ok := m.getState().images[link]; ok {
			return existing
		}

		result := mdLoadImage(link)
		m.getState().images[link] = result

		return result
	})

	cfg.SetImageCallback(&imgCb)

	return &markdownState{
		cfg:    *cfg,
		images: make(map[string]immarkdown.MarkdownImageData),
	}
}

// Markdown creates new markdown widget.
func Markdown(md string) *MarkdownWidget {
	return (&MarkdownWidget{
		md: md,
		id: GenAutoID("MarkdownWidget"),
		headers: [3]immarkdown.MarkdownHeadingFormat{
			*immarkdown.NewEmptyMarkdownHeadingFormat(),
			*immarkdown.NewEmptyMarkdownHeadingFormat(),
			*immarkdown.NewEmptyMarkdownHeadingFormat(),
		},
	}).OnLink(OpenURL)
}

// OnLink sets another than default link callback.
// NOTE: due to cimgui-go's limitation https://github.com/AllenDang/cimgui-go?tab=readme-ov-file#callbacks
// we clear MarkdownLinkCallback pool every frame. No further action from you should be required (just feel informed).
// ref (*MasterWindow).beforeRender.
func (m *MarkdownWidget) OnLink(cb func(url string)) *MarkdownWidget {
	igCb := immarkdown.MarkdownLinkCallback(func(data immarkdown.MarkdownLinkCallbackData) {
		link := data.Link()[:data.LinkLength()]
		cb(link)
	})

	m.getState().cfg.SetLinkCallback(&igCb)

	return m
}

// Header sets header formatting
// NOTE: level (counting from 0!) is header level. (for instance, header `# H1` will have level 0).
// NOTE: since cimgui-go there are only 3 levels (so level < 3 here). This will panic if level >= 3!
// TODO: it actually doesn't work.
func (m *MarkdownWidget) Header(level int, font *FontInfo, separator bool) *MarkdownWidget {
	// ensure level is in range
	Assert(level < 3, "MarkdownWidget", "Header", "Header level must be less than 3!")

	m.headers[level] = *immarkdown.NewEmptyMarkdownHeadingFormat()

	if font != nil {
		if f, ok := Context.FontAtlas.extraFontMap[font.String()]; ok {
			m.headers[level].SetFont(f)
		}
	}

	m.headers[level].SetSeparator(separator)

	state := m.getState()
	state.cfg.SetHeadingFormats(&m.headers)

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
}

func mdLoadImage(path string) immarkdown.MarkdownImageData {
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
			return *immarkdown.NewEmptyMarkdownImageData()
		}

		defer func() {
			closeErr := resp.Body.Close()
			Assert((closeErr == nil), "MarkdownWidget", "mdLoadImage", "Could not close http request!")
		}()

		rgba, _, imgErr := image.Decode(resp.Body)
		if imgErr != nil {
			return *immarkdown.NewEmptyMarkdownImageData()
		}

		img = ImageToRgba(rgba)
	default:
		img, err = LoadImage(path)
		if err != nil {
			return *immarkdown.NewEmptyMarkdownImageData()
		}
	}

	size := img.Bounds()
	id := backend.NewTextureFromRgba(img).ID

	result := immarkdown.NewEmptyMarkdownImageData()
	result.SetUsertextureid(id)
	result.SetSize(imgui.Vec2{
		X: float32(size.Dx()),
		Y: float32(size.Dy()),
	})
	result.SetUseLinkCallback(true)
	result.SetUv0(ToVec2(image.Point{0, 0}))
	result.SetUv1(ToVec2(image.Point{1, 1}))
	result.SetTintcol(ToVec4Color(color.RGBA{255, 255, 255, 255}))
	result.SetBordercol(ToVec4Color(color.RGBA{0, 0, 0, 0}))

	result.SetIsValid(true)

	return *result
}
