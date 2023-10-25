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
				g.MenuItem("Open").Shortcut("Ctrl+O"),
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
		g.Label("right/left click me"),
		g.Event().
			OnClick(g.MouseButtonLeft, func() { fmt.Println("I was left-clicked") }).
			OnClick(g.MouseButtonRight, func() { fmt.Println("I was right-clicked") }).
			OnDClick(g.MouseButtonLeft, func() { fmt.Println("I was left-double-clicked") }).
			OnDClick(g.MouseButtonRight, func() { fmt.Println("I was right-double-clicked") }),
		g.Row(
			g.InputText(&name),
			g.Event().OnActivate(func() { fmt.Println("input text focused") }).
				OnDeactivate(func() { fmt.Println("input text unfocused") }),
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
			g.Label("Do you like giu?"),
			g.RadioButton("Yes, of course", radioOp == 0).OnChange(func() { radioOp = 0 }),
			g.RadioButton("I'm going to test it now", radioOp == 1).OnChange(func() { radioOp = 1 }),
			g.RadioButton("No", false),
		),

		g.ProgressBar(0.8).Size(g.Auto, 0).Overlay("Progress"),
		g.DragInt("DragInt", &dragInt, 0, 100),
		g.SliderInt(&dragInt, 0, 100).Label("Slider"),

		g.Label("Vertical sliders"),
		g.Row(
			g.VSliderInt(&dragInt, 0, 100).OnChange(func() { fmt.Println(dragInt) }),
			g.VSliderInt(&dragInt, 0, 100),
			g.VSliderInt(&dragInt, 0, 100),
		),

		g.Combo("Combo", items[itemSelected], items, &itemSelected).OnChange(comboChanged),

		g.ColorEdit("<- Click the black square. I'm changing a color for you##colorChanger", col).
			Size(100).
			Flags(g.ColorEditFlagsDisplayHex).
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

			g.ArrowButton(g.DirectionLeft),
			g.ArrowButton(g.DirectionRight),
			g.ArrowButton(g.DirectionUp),
			g.ArrowButton(g.DirectionDown),
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

		g.TabBar().TabItems(
			g.TabItem("Multiline Input").Layout(
				g.Label("This is first tab with a multiline input text field"),
				g.InputTextMultiline(&multiline).Size(g.Auto, g.Auto),
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
			g.TabItem("TreeTable").Layout(
				g.TreeTable().
					Columns(g.TableColumn("Name"), g.TableColumn("Size")).
					Rows(
						[]*g.TreeTableRowWidget{
							g.TreeTableRow("Folder1", g.Label("")).Children(
								g.TreeTableRow("File1", g.Label("1MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
							),
							g.TreeTableRow("Folder2", g.Label("")).Children(
								g.TreeTableRow("File1", g.Label("1MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
							),
							g.TreeTableRow("Folder3", g.Label("")).Flags(g.TreeNodeFlagsDefaultOpen).Children(
								g.TreeTableRow("File1", g.Label("1MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
								g.TreeTableRow("File2", g.Label("2MB")),
							),
						}...,
					).
					Size(g.Auto, g.Auto),
			),
			g.TabItem("ListBox").Layout(
				g.ListBox([]string{"List item 1", "List item 2", "List item 3"}),
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
			g.TabItem("List").Layout(
				g.Child().Layout(
					g.ListClipper().Layout(
						g.Label("these labels"),
						g.Label("uses ListClipper"),
						g.Label("and are rendered only when visible."),
						g.BulletText("I'm a bullet"),
						g.Label("label 1"),
						g.Label("label 2"),
						g.Label("label 3"),
						g.Label("label 4"),
						g.Label("label 5"),
						g.Label("label 6"),
						g.Label("label 7"),
						g.Label("label 8"),
						g.Label("label 9"),
						g.Button("I'm a button 1"),
						g.Row(
							g.Button("we're the buttons"),
							g.Button("in row"),
						),
						g.InputTextMultiline(&multiline),
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
