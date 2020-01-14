package giu

import "github.com/AllenDang/giu/imgui"

type Builder func()

func SameLine() {
	imgui.SameLine()
}

func InputTextMultiline(label string, text *string) {
	InputTextMultilineV(label, text, 0, 0, 0, nil)
}

func InputTextMultilineV(label string, text *string, width, height float32, flags int, cb imgui.InputTextCallback) {
	imgui.InputTextMultilineV(label, text, imgui.Vec2{X: width, Y: height}, flags, cb)
}

func Button(id string, clicked func()) {
	ButtonV(id, 0, 0, clicked)
}

func ButtonV(id string, width, height float32, clicked func()) {
	if imgui.ButtonV(id, imgui.Vec2{X: width, Y: height}) && clicked != nil {
		clicked()
	}
}

func ImageButton(textureId imgui.TextureID, clicked func()) {
	if imgui.ImageButton(textureId, imgui.Vec2{}) && clicked != nil {
		clicked()
	}
}

func Checkbox(text string, selected *bool, changed func()) {
	if imgui.Checkbox(text, selected) && changed != nil {
		changed()
	}
}

func RadioButton(text string, active bool, changed func()) {
	if imgui.RadioButton(text, active) && changed != nil {
		changed()
	}
}

func Child(id string, border bool, builder Builder) {
	ChildV(id, border, 0, 0, 0, builder)
}

func ChildV(id string, border bool, width, height float32, flags int, builder Builder) {
	imgui.BeginChildV(id, imgui.Vec2{X: width, Y: height}, border, flags)
	if builder != nil {
		builder()
	}
	imgui.EndChild()
}

func Combo(label, previewValue string, items []string, selected *int32, flags int, changed func()) {
	if imgui.BeginComboV(label, previewValue, flags) {
		for i, item := range items {
			if imgui.Selectable(item) {
				*selected = int32(i)
				if changed != nil {
					changed()
				}
			}
		}

		imgui.EndCombo()
	}
}

func ContextMenu(builder Builder) {
	ContextMenuV("", 1, builder)
}

func ContextMenuV(label string, mouseButton int, builder Builder) {
	if imgui.BeginPopupContextItemV(label, mouseButton) {
		if builder != nil {
			builder()
		}
		imgui.EndPopup()
	}
}

func DragInt(label string, value *int32) {
	DragIntV(label, value, 1.0, 0, 0, "%d")
}

func DragIntV(label string, value *int32, speed float32, min, max int32, format string) {
	imgui.DragIntV(label, value, speed, min, max, format)
}

func Group(builder Builder) {
	imgui.BeginGroup()
	if builder != nil {
		builder()
	}
	imgui.EndGroup()
}

func Image(texture *Texture, width, height float32) {
	size := imgui.Vec2{X: width, Y: height}
	if texture != nil && texture.id != 0 {
		rect := imgui.ContentRegionAvail()
		if size.X == -1 {
			size.X = rect.X
		}
		if size.Y == -1 {
			size.Y = rect.Y
		}
		imgui.Image(texture.id, size)
	}
}

func InputText(label string, value *string) {
	InputTextV(label, value, 0, nil, nil)
}

func InputTextV(label string, value *string, flags int, cb imgui.InputTextCallback, changed func()) {
	if imgui.InputTextV(label, value, flags, cb) && changed != nil {
		changed()
	}
}

func Label(label string) {
	imgui.Text(label)
}

func ListBox(label string, selected *int32, items []string, itemHeight int, width float32, changed func()) {
	if imgui.ListBoxV(label, selected, items, itemHeight) && changed != nil {
		changed()
	}
}

func MainMenuBar(builder Builder) {
	if imgui.BeginMainMenuBar() {
		if builder != nil {
			builder()
		}
		imgui.EndMainMenuBar()
	}
}

func MenuBar(builder Builder) {
	if imgui.BeginMenuBar() {
		if builder != nil {
			builder()
		}
		imgui.EndMenuBar()
	}
}

func MenuItem(label string) {
	MenuItemV(label, false, true, nil)
}

func MenuItemV(label string, selected, enabled bool, clicked func()) {
	if imgui.MenuItemV(label, "", selected, enabled) && clicked != nil {
		clicked()
	}
}

func Menu(label string, builder Builder) {
	MenuV(label, true, builder)
}

func MenuV(label string, enabled bool, builder Builder) {
	if imgui.BeginMenuV(label, enabled) {
		if builder != nil {
			builder()
		}
		imgui.EndMenu()
	}
}

func Popup(name string, builder Builder) {
	PopupV(name, nil, 0, builder)
}

func PopupV(name string, open *bool, flags int, builder Builder) {
	if imgui.BeginPopupModalV(name, open, flags) {
		if builder != nil {
			builder()
		}
		imgui.EndPopup()
	}
}

func ProgressBar(fraction float32, width, height float32, overlay string) {
	imgui.ProgressBarV(fraction, imgui.Vec2{X: width, Y: height}, overlay)
}

func Selectable(label string, clicked func()) {
	SelectableV(label, false, 0, 0, 0, clicked)
}

func SelectableV(label string, selected bool, flags int, width, height float32, clicked func()) {
	if imgui.SelectableV(label, selected, flags, imgui.Vec2{X: width, Y: height}) && clicked != nil {
		clicked()
	}
}

func Separator() {
	imgui.Separator()
}

func SliderInt(label string, value *int32, min, max int32, format string) {
	imgui.SliderIntV(label, value, min, max, format)
}

func Dummy(width, height float32) {
	imgui.Dummy(imgui.Vec2{X: width, Y: height})
}

func HSplitter(id string, width, height float32, delta *float32) {
	imgui.InvisibleButton(id, imgui.Vec2{X: width, Y: height})
	if imgui.IsItemActive() {
		*(delta) = imgui.CurrentIO().GetMouseDelta().Y
	} else {
		*(delta) = 0
	}
}

func VSplitter(id string, width, height float32, delta *float32) {
	imgui.InvisibleButton(id, imgui.Vec2{X: width, Y: height})
	if imgui.IsItemActive() {
		*(delta) = imgui.CurrentIO().GetMouseDelta().X
	} else {
		*(delta) = 0
	}
}

func TabItem(label string, builder Builder) {
	TabItemV(label, nil, 0, builder)
}

func TabItemV(label string, open *bool, flags int, builder Builder) {
	if imgui.BeginTabItemV(label, open, flags) {
		if builder != nil {
			builder()
		}
		imgui.EndTabItem()
	}
}

func TabBar(id string, builder Builder) {
	TabBarV(id, 0, builder)
}

func TabBarV(id string, flags int, builder Builder) {
	if imgui.BeginTabBarV(id, flags) {
		if builder != nil {
			builder()
		}
		imgui.EndTabBar()
	}
}

func NextColumn() {
	imgui.NextColumn()
}

func Table(label string, border bool, columnCount int, builder Builder) {
	imgui.ColumnsV(columnCount, label, border)

	if builder != nil {
		builder()
	}

	imgui.Columns()
	if border {
		imgui.Separator()
	}
}

func Tooltip(tip string) {
	if imgui.IsItemHovered() {
		imgui.SetTooltip(tip)
	}
}

func TreeNode(label string, flags int, builder Builder) {
	if imgui.TreeNodeV(label, flags) {
		if builder != nil {
			builder()
		}
		if (flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}

func Spacing() {
	imgui.Spacing()
}

// Creates a widget block with given ID. IDs are hash of the entire stack!
func ID(id string, builder Builder) {
	imgui.PushID(id)
	if builder != nil {
		builder()
	}
	imgui.PopID()
}
