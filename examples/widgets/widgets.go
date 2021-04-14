package main

import (
	"fmt"
	"time"

	g "github.com/ianling/giu"
)

var (
	name         string
	items        []string
	itemSelected int32
	checked      bool
	checked2     bool
	dragInt      int32
	multiline    string
	radioOp      int
	date         time.Time = time.Now()
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
	g.SingleWindowWithMenuBar("Overview").Layout(
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
		g.Line(
			g.InputText("##name", &name),
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
		g.DatePicker("Date Picker", &date).OnChange(func() {
			fmt.Println(date)
		}),
		g.Line(
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
		g.Line(
			g.VSliderInt("##VSlider1", &dragInt, 0, 100).OnChange(func() { fmt.Println(dragInt) }),
			g.VSliderInt("##VSlider2", &dragInt, 0, 100),
			g.VSliderInt("##VSlider3", &dragInt, 0, 100),
		),

		g.Combo("Combo", items[itemSelected], items, &itemSelected).OnChange(comboChanged),

		g.Line(
			g.Button("Button"),
			g.SmallButton("SmallButton"),
		),

		g.BulletText("Bullet1"),
		g.BulletText("Bullet2"),
		g.BulletText("Bullet3"),

		g.Line(
			g.Label("Arrow buttons: "),

			g.ArrowButton("arrow left", g.DirectionLeft),
			g.ArrowButton("arrow right", g.DirectionRight),
			g.ArrowButton("arrow up", g.DirectionUp),
			g.ArrowButton("arrow down", g.DirectionDown),
		),

		g.Line(
			g.Button("Popup Modal").OnClick(btnPopupCLicked),
			g.PopupModal("Confirm").Layout(
				g.Label("Confirm to close me?"),
				g.Line(
					g.Button("Yes").OnClick(func() { g.CloseCurrentPopup() }),
					g.Button("No"),
				),
			),
			g.Label("Right click me to see the context menu"),
			g.ContextMenu("Context menu").Layout(
				g.Selectable("Context menu 1").OnClick(contextMenu1Clicked),
				g.Selectable("Context menu 2").OnClick(contextMenu2Clicked),
			),
		),

		g.TabBar("Tabbar Input").Layout(
			g.TabItem("Multiline Input").Layout(
				g.Label("This is first tab with a multiline input text field"),
				g.InputTextMultiline("##multiline", &multiline).Size(-1, -1),
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
					if g.IsItemClicked() {
						fmt.Println("Clicked")
					}
				}),
			),
			g.TabItem("ListBox").Layout(
				g.ListBox("ListBox1", []string{"List item 1", "List item 2", "List item 3"}),
			),
			g.TabItem("Table").Layout(
				g.Table("Table").
					Columns(
						g.Column("Column 1"),
						g.Column("Column 2"),
						g.Column("Column 3"),
					).
					Rows(
						g.Row(g.Label("Loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog").Wrapped(true), g.Label("Age"), g.Label("Loc")),
						g.Row(g.Label("Second Loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog").Wrapped(true), g.Label("Age"), g.Label("Loc")),
						g.Row(g.Label("Name"), g.Label("Age"), g.Label("Location")),
						g.Row(g.Label("Allen"), g.Label("33"), g.Label("Shanghai/China")),
						g.Row(g.Checkbox("check me", &checked), g.Button("click me"), g.Label("Anything")),
					),
			),
			g.TabItem("Group").Layout(
				g.Line(
					g.Group().Layout(
						g.Label("I'm inside group 1"),
					),
					g.Group().Layout(
						g.Label("I'm inside group 2"),
					),
				),
			),
		),
	).Build()
}

func main() {
	items = make([]string, 100)
	for i := range items {
		items[i] = fmt.Sprintf("Item %d", i)
	}

	w := g.NewMasterWindow("Overview", 800, 600, 0, nil)
	w.Run(loop)
}
