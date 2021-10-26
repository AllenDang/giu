package main

import (
	"strings"

	"github.com/pkg/browser"

	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
)

var markdown string = getExampleMarkdownText()

func getExampleMarkdownText() string {
	return strings.Join([]string{
		"Wrapping:",
		"Text wraps automatically. To add a new line, use 'Return'.",
		"",
		"Headers:",
		"# H1",
		"## H2",
		"### H3",
		"",
		"Emphasis:",
		"*emphasis*",
		"_emphasis_",
		"**strong emphasis**",
		"__strong emphasis__",
		"",
		"Indents:",
		"On a new line, at the start of the line, add two spaces per indent.",
		"  Indent level 1",
		"    Indent level 2",
		"",
		"Unordered lists:",
		"On a new line, at the start of the line, add two spaces, an asterisks and a space.",
		"For nested lists, add two additional spaces in front of the asterisk per list level increment.",
		"  * Unordered List level 1",
		"    * Unordered List level 2",
		"",
		"Link:",
		"Here is [a link to some cool website!](https://github.com/AllenDang/giu) you must click it!",
		"Image:",
		"![gopher image](./gopher.png)",
		"",
		"Horizontal Rule:",
		"***",
		"___",
	}, "\n")
}

func loop() {
	giu.SingleWindow().Layout(
		giu.SplitLayout(giu.DirectionHorizontal, 320,
			giu.Layout{
				giu.Row(
					giu.Label("Markdown Edition:"),
					giu.Button("Reset").OnClick(func() {
						markdown = getExampleMarkdownText()
					}),
				),
				giu.Custom(func() {
					availableW, availableH := giu.GetAvailableRegion()
					giu.InputTextMultiline(&markdown).Size(availableW, availableH).Build()
				}),
			},
			giu.Custom(func() {
				imgui.Markdown(&markdown, func(s string) {
					browser.OpenURL(s)
				})
			}),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("ImGui Markdown [Demo]", 640, 480, 0)
	wnd.Run(loop)
}
