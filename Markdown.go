package giu

import "github.com/AllenDang/imgui-go"

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
// NOTE: level (counting from 0!) is header level. (for instance, header `# H1` will have level 0)
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
	imgui.Markdown(m.md, m.linkCb, m.headers)
}
