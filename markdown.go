package giu

import "github.com/AllenDang/imgui-go"

// MarkdownWidget implements DearImGui markdown extension
// https://github.com/juliettef/imgui_markdown
// It is like LabelWidget but with md formatting.
type MarkdownWidget struct {
	md     *string
	linkCb func(url string)
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

// Build implements Widget interface.
func (m *MarkdownWidget) Build() {
	imgui.Markdown(m.md, m.linkCb)
}
