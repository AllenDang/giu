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
	imgui.OpenPopup("Confirm")
}

func loop(w *g.MasterWindow) {
	// Create main menu bar for master window.
	g.MainMenuBar(
		func() {
			g.Menu("File", func() {
				g.MenuItem("Open")
				g.MenuItem("Save")
				// You could add any kind of widget here, not just menu item.
				g.Menu("Save as ...", func() {
					g.MenuItem("Excel file")
					g.MenuItem("CSV file")
					g.Button("Button inside menu", nil)
				},
				)
			},
			)
		},
	)

	// Build a new window
	width, height := w.GetSize()
	g.Window("Overview", 0, 20, float32(width), float32(height)-20,
		func() {
			g.Label("One line label")
			g.InputText("##name", &name)
			g.SameLine()
			g.Button("Click Me", btnClickMeClicked)
			g.Tooltip("I'm a tooltip")
			g.Checkbox("Checkbox", &checked, func() {
				fmt.Println(checked)
			})
			g.SameLine()
			g.Checkbox("Checkbox 2", &checked2, func() {
				fmt.Println(checked2)
			})
			g.ProgressBar(0.8, -1, 0, "Progress")
			g.DragInt("DragInt", &dragInt)
			g.SliderInt("Slider", &dragInt, 0, 100, "")

			g.Combo("Combo", items[itemSelected], items, &itemSelected, 0, comboChanged)
			g.SameLine()
			g.Label("Right click me")
			g.ContextMenu(func() {
				g.Selectable("Context menu 1", contextMenu1Clicked)
				g.Selectable("Context menu 2", contextMenu2Clicked)
			})
			g.ListBox("ListBox", &itemSelected, items, 5, -1, listboxChanged)
			g.Button("Popup Modal", btnPopupCLicked)
			g.Popup("Confirm", func() {
				g.Label("Confirm to close me?")
				g.Button("Yes", func() { imgui.CloseCurrentPopup() })
				g.SameLine()
				g.Button("No", nil)
			})
			g.TabBar("Tabbar Input", func() {
				g.TabItem("Multiline Input", func() {
					g.Label("This is first tab with a multiline input text field")
					g.InputTextMultiline("##multiline", &multiline)
				})
				g.TabItem("Tree", func() {
					g.TreeNode("TreeNode1", imgui.TreeNodeFlagsCollapsingHeader|imgui.TreeNodeFlagsDefaultOpen, func() {
						g.Label("Tree node 1")
						g.Label("Tree node 1")
						g.Label("Tree node 1")
						g.Button("Button inside tree", nil)
					})
					g.TreeNode("TreeNode2", 0, func() {
						g.Label("Tree node 2")
						g.Label("Tree node 2")
						g.Label("Tree node 2")
						g.Button("Button inside tree", nil)
					})
				})
				g.TabItem("Table", func() {
					g.Table("Table", true, 3, func() {
						g.Separator()
						g.Label("Name")
						g.NextColumn()
						g.Label("Age")
						g.NextColumn()
						g.Label("Location")
						g.Separator()

						g.NextColumn()
						g.Label("Allen")
						g.NextColumn()
						g.Label("33")
						g.NextColumn()
						g.Label("China")
						g.Separator()

						g.NextColumn()
						g.Checkbox("check me", &checked, nil)
						g.NextColumn()
						g.Button("click me", nil)
						g.NextColumn()
						g.Label("...")
					})
				})
				g.TabItem("Group", func() {
					g.Group(func() {
						g.Label("I'm inside group 1")
					})
					g.SameLine()
					g.Group(func() {
						g.Label("I'm inside group 2")
					})
				})
			})
		})
}

func main() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("Overview", 800, 600, true, nil)
	w.Main(loop)
}
