package main

import (
	"fmt"
	"image/color"
	"time"

	g "github.com/AllenDang/giu"
)

var (
	name                   string
	items                  []string
	itemSelected           int32
	checked                bool
	checked2               bool
	dragInt                int32
	multiline              string
	radioOp                int
	autoCompleteCandidates           = []string{"hello", "hello world"}
	date                   time.Time = time.Now()
	col                              = &color.RGBA{}
)

func btnClickMeClicked() {
	fmt.Println("Click me is clicked")
}

func comboChanged() {
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

func loop() {
	g.SingleWindowWithMenuBar().Layout(
		g.MenuBar().Layout(
			g.Menu("File").Layout(
				g.MenuItem("Open"),
				g.MenuItem("Save"),
				// You could add any kind of widget here, not just menu item.
				g.Menu("Save as ...").Layout(
					g.MenuItem("Excel file"),
					g.MenuItem("CSV file"),
					g.Button("Button inside menu"),
				),
			),
		),
		g.Label("One line label"),
		g.Label("Auto wrapped label with very long line...............................................this line should be wrapped.").Wrapped(true),
		g.Row(
			g.InputText(&name),
			g.Button("Click Me").OnClick(btnClickMeClicked),
			g.Tooltip("I'm a tooltip"),
			g.Button("More tooltip"),
			g.Tooltip("Advance Tooltip").Layout(
				g.Label("I'm a label"),
				g.Selectable("I'm a selectable"),
				g.Button("I'm a button"),
				g.BulletText("I could be any widgets"),
			),
		),
		g.InputText(&name).Label("Input text with auto complete, input hw and press enter").Size(300).AutoComplete(autoCompleteCandidates),
		g.DatePicker("Date Picker", &date).OnChange(func() {
			fmt.Println(date)
		}),
		g.Row(
			g.Checkbox("Checkbox", &checked).OnChange(func() {
				fmt.Println(checked)
			}),
			g.Checkbox("Checkbox 2", &checked2).OnChange(func() {
				fmt.Println(checked2)
			}),
			g.Dummy(30, 0),
			g.RadioButton("Radio 1", radioOp == 0).OnChange(func() { radioOp = 0 }),
			g.RadioButton("Radio 2", radioOp == 1).OnChange(func() { radioOp = 1 }),
			g.RadioButton("Radio 3", radioOp == 2).OnChange(func() { radioOp = 2 }),
		),

		g.ProgressBar(0.8).Size(-1, 0).Overlay("Progress"),
		g.DragInt("DragInt", &dragInt, 0, 100),
		g.SliderInt("Slider", &dragInt, 0, 100),

		g.Label("Vertical sliders"),
		g.Row(
			g.VSliderInt("##VSlider1", &dragInt, 0, 100).OnChange(func() { fmt.Println(dragInt) }),
			g.VSliderInt("##VSlider2", &dragInt, 0, 100),
			g.VSliderInt("##VSlider3", &dragInt, 0, 100),
		),

		g.Combo("Combo", items[itemSelected], items, &itemSelected).OnChange(comboChanged),

		g.ColorEdit("<- Click the black square. I'm changing a color for you##colorChanger", col).
			Size(100).
			Flags(g.ColorEditFlagsHEX).
			OnChange(func() {
				fmt.Println(col)
			}),

		g.Row(
			g.Button("Button"),
			g.SmallButton("SmallButton"),
		),

		g.BulletText("Bullet1"),
		g.BulletText("Bullet2"),
		g.BulletText("Bullet3"),

		g.Row(
			g.Label("Arrow buttons: "),

			g.ArrowButton("arrow left", g.DirectionLeft),
			g.ArrowButton("arrow right", g.DirectionRight),
			g.ArrowButton("arrow up", g.DirectionUp),
			g.ArrowButton("arrow down", g.DirectionDown),
		),

		g.Row(
			g.Button("Popup Modal").OnClick(btnPopupCLicked),
			g.PopupModal("Confirm").Layout(
				g.Label("Confirm to close me?"),
				g.Row(
					g.Button("Yes").OnClick(func() { g.CloseCurrentPopup() }),
					g.Button("No"),
				),
			),
			g.Label("Right click me to see the context menu"),
			g.ContextMenu().Layout(
				g.Selectable("Context menu 1").OnClick(contextMenu1Clicked),
				g.Selectable("Context menu 2").OnClick(contextMenu2Clicked),
			),
		),

		g.TabBar().Layout(
			g.TabItem("Multiline Input").Layout(
				g.Label("This is first tab with a multiline input text field"),
				g.InputTextMultiline(&multiline).Size(-1, -1),
			),
			g.TabItem("Tree").Layout(
				g.TreeNode("TreeNode1").Flags(g.TreeNodeFlagsCollapsingHeader|g.TreeNodeFlagsDefaultOpen).Layout(
					g.Custom(func() {
						if g.IsItemActive() && g.IsMouseClicked(g.MouseButtonLeft) {
							fmt.Println("Tree node clicked")
						}
					}),
					g.Selectable("Tree node 1").OnClick(func() {
						fmt.Println("Click tree node 1")
					}),
					g.Label("Tree node 2"),
					g.Label("Tree node 3"),
					g.Button("Button inside tree"),
				),
				g.TreeNode("TreeNode with event handler").Layout(
					g.Selectable("Selectable 1").OnClick(func() { fmt.Println(1) }),
					g.Selectable("Selectable 2").OnClick(func() { fmt.Println(2) }),
				).Event(func() {
					if g.IsItemClicked(g.MouseButtonLeft) {
						fmt.Println("Clicked")
					}
				}),
			),
			g.TabItem("ListBox").Layout(
				g.ListBox("ListBox1", []string{"List item 1", "List item 2", "List item 3"}),
			),
			g.TabItem("Table").Layout(
				g.Table().
					Columns(
						g.TableColumn("Column 1"),
						g.TableColumn("Column 2"),
						g.TableColumn("Column 3"),
					).
					Rows(
						g.TableRow(g.Label("Loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog").Wrapped(true), g.Label("Age"), g.Label("Loc")),
						g.TableRow(g.Label("Second Loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog").Wrapped(true), g.Label("Age"), g.Label("Loc")),
						g.TableRow(g.Label("Name"), g.Label("Age"), g.Label("Location")),
						g.TableRow(g.Label("Allen"), g.Label("33"), g.Label("Shanghai/China")),
						g.TableRow(g.Checkbox("check me", &checked), g.Button("click me"), g.Label("Anything")),
					),
			),
			g.TabItem("Group").Layout(
				g.Row(
					g.Column(
						g.Label("I'm inside group 1"),
					),
					g.Column(
						g.Label("I'm inside group 2"),
					),
				),
			),
		),
	)
}

func main() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("Overview", 1000, 800, 0)
	w.Run(loop)
}
