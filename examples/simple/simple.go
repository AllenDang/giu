package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/giu/imgui"
)

var (
	name         string
	items        []string
	itemSelected int32
	checked      bool
	checked2     bool
	dragInt      int32
	multiline    string
)

func btnClickMeClicked() {
	fmt.Println("Click me is clicked")
}

func comboChanged() {
	fmt.Println(items[itemSelected])
}

func listboxChanged() {
	fmt.Println(items[itemSelected])
}

func contextMenu1Clicked() {
	fmt.Println("Context menu 1 is clicked")
}

func contextMenu2Clicked() {
	fmt.Println("Context menu 2 is clicked")
}

func btnPopupCLicked() {
	g.OpenPopup("Confirm")
}

func loop(w *g.MasterWindow) {
	// Create main menu bar for master window.
	g.MainMenuBar(
		g.Layout{
			g.Menu("File", g.Layout{
				g.MenuItem("Open", nil),
				g.MenuItem("Save", nil),
				// You could add any kind of widget here, not just menu item.
				g.Menu("Save as ...", g.Layout{
					g.MenuItem("Excel file", nil),
					g.MenuItem("CSV file", nil),
					g.Button("Button inside menu", nil),
				}),
			}),
		},
	)

	// Build a new window
	width, height := w.GetSize()
	g.Window("Overview", 0, 20, float32(width), float32(height)-20,
		g.Layout{
			g.Label("One line label"),
			g.SameLine(
				g.InputTextV("##name", &name, -80, 0, nil, nil),
				g.Button("Click Me", btnClickMeClicked),
				g.Tooltip("I'm a tooltip"),
			),
			g.SameLine(
				g.Checkbox("Checkbox", &checked, func() {
					fmt.Println(checked)
				}),
				g.Checkbox("Checkbox 2", &checked2, func() {
					fmt.Println(checked2)
				}),
			),
			g.ProgressBarV(0.8, -1, 0, "Progress"),
			g.DragIntV("DragInt", &dragInt, 1.0, 0, 100, "%d", 0),
			g.SliderInt("Slider", &dragInt, 0, 100),
			g.SameLine(
				g.ComboV("Combo", items[itemSelected], items, &itemSelected, 100, comboChanged),
				g.Label("Right click me"),
				g.ContextMenu(
					g.Layout{
						g.Selectable("Context menu 1", contextMenu1Clicked),
						g.Selectable("Context menu 2", contextMenu2Clicked),
					},
				),
			),
			g.ListBoxV("ListBox", &itemSelected, items, 5, -1, listboxChanged),
			g.SameLine(
				g.Button("Popup Modal", btnPopupCLicked),
				g.Popup("Confirm",
					g.Layout{
						g.Label("Confirm to close me?"),
						g.SameLine(
							g.Button("Yes", func() { g.CloseCurrentPopup() }),
							g.Button("No", nil)),
					},
				),
			),
			g.TabBar("Tabbar Input",
				[]*g.TabItemWidget{
					g.TabItem("Multiline Input",
						g.Layout{
							g.Label("This is first tab with a multiline input text field"),
							g.InputTextMultilineV("##multiline", &multiline, -1, 0, 0, nil),
						},
					),
					g.TabItem("Tree",
						g.Layout{
							g.TreeNodeV("TreeNode1", imgui.TreeNodeFlagsCollapsingHeader|imgui.TreeNodeFlagsDefaultOpen,
								g.Layout{
									g.Label("Tree node 1"),
									g.Label("Tree node 1"),
									g.Label("Tree node 1"),
									g.Button("Button inside tree", nil),
								},
							),
							g.TreeNode("TreeNode2",
								g.Layout{
									g.Label("Tree node 2"),
									g.Label("Tree node 2"),
									g.Label("Tree node 2"),
									g.Button("Button inside tree", nil),
								}),
						},
					),
					g.TabItem("Table",
						g.Layout{
							g.Table("Table", true,
								g.Row(g.Label("Name"), g.Label("Age"), g.Label("Location")),
								g.Row(g.Label("Allen"), g.Label("33"), g.Label("China")),
								g.Row(g.Label("Allen"), g.Label("33"), g.Label("China")),
								g.Row(g.Label("Allen"), g.Label("33"), g.Label("China")),
								g.Row(g.Checkbox("check me", &checked, nil), g.Button("click me", nil), g.Label("...")),
								g.Row(g.Label("Allen"), g.Label("33"), g.Label("China")),
							),
						},
					),
					g.TabItem("Group",
						g.Layout{
							g.SameLine(
								g.Group(g.Layout{
									g.Label("I'm inside group 1"),
								}),
								g.Group(g.Layout{
									g.Label("I'm inside group 2"),
								}),
							),
						},
					),
				},
			),
		},
	)
}

func main() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("Overview", 800, 600, true, nil)
	w.Main(loop)
}
