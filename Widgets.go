package giu

import "github.com/AllenDang/giu/imgui"

func SameLine() Widget {
	return func() {
		imgui.SameLine()
	}
}

func InputTextMultiline(label string, text *string) Widget {
	return InputTextMultilineV(label, text, 0, 0, 0, nil)
}

func InputTextMultilineV(label string, text *string, width, height float32, flags int, cb imgui.InputTextCallback) Widget {
	return func() {
		imgui.InputTextMultilineV(label, text, imgui.Vec2{X: width, Y: height}, flags, cb)
	}
}

func Button(id string, clicked func()) Widget {
	return ButtonV(id, 0, 0, clicked)
}

func ButtonV(id string, width, height float32, clicked func()) Widget {
	return func() {
		if imgui.ButtonV(id, imgui.Vec2{X: width, Y: height}) && clicked != nil {
			clicked()
		}
	}
}

func ImageButton(textureId imgui.TextureID, clicked func()) Widget {
	return func() {
		if imgui.ImageButton(textureId, imgui.Vec2{}) && clicked != nil {
			clicked()
		}
	}
}

func Checkbox(text string, selected *bool, changed func()) Widget {
	return func() {
		if imgui.Checkbox(text, selected) && changed != nil {
			changed()
		}
	}
}

func Child(id string, border bool, widgets ...Widget) Widget {
	return ChildV(id, border, 0, 0, 0, widgets...)
}

func ChildV(id string, border bool, width, height float32, flags int, widgets ...Widget) Widget {
	return func() {
		imgui.BeginChildV(id, imgui.Vec2{X: width, Y: height}, border, flags)
		Layout(widgets...)()
		imgui.EndChild()
	}
}

func Combo(label, previewValue string, items []string, selected *int32, flags int, changed func()) Widget {
	return func() {
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
}

func ContextMenu(widgets ...Widget) Widget {
	return ContextMenuV("", 1, widgets...)
}

func ContextMenuV(label string, mouseButton int, widgets ...Widget) Widget {
	return func() {
		if imgui.BeginPopupContextItemV(label, mouseButton) {
			Layout(widgets...)()
			imgui.EndPopup()
		}
	}
}

func DragInt(label string, value *int32) Widget {
	return DragIntV(label, value, 1.0, 0, 0, "%d")
}

func DragIntV(label string, value *int32, speed float32, min, max int32, format string) Widget {
	return func() {
		imgui.DragIntV(label, value, speed, min, max, format)
	}
}

func Group(widgets ...Widget) Widget {
	return func() {
		imgui.BeginGroup()
		Layout(widgets...)()
		imgui.EndGroup()
	}
}

func Image(id imgui.TextureID, width, height float32) Widget {
	size := imgui.Vec2{X: width, Y: height}
	return func() {
		if id != 0 {
			rect := imgui.ContentRegionAvail()
			if size.X == -1 {
				size.X = rect.X
			}
			if size.Y == -1 {
				size.Y = rect.Y
			}
			imgui.Image(id, size)
		}
	}
}

func InputText(label string, value *string) Widget {
	return InputTextV(label, value, 0, nil, nil)
}

func InputTextV(label string, value *string, flags int, cb imgui.InputTextCallback, changed func()) Widget {
	return func() {
		if imgui.InputTextV(label, value, flags, cb) && changed != nil {
			changed()
		}
	}
}

func Label(label string) Widget {
	return func() {
		imgui.Text(label)
	}
}

func ListBox(label string, selected *int32, items []string, itemHeight int, width float32, changed func()) Widget {
	return func() {
		if imgui.ListBoxV(label, selected, items, itemHeight) && changed != nil {
			changed()
		}
	}
}

func MainMenuBar(widgets ...Widget) Widget {
	return func() {
		if imgui.BeginMainMenuBar() {
			Layout(widgets...)()
			imgui.EndMainMenuBar()
		}
	}
}

func MenuBar(widgets ...Widget) Widget {
	return func() {
		if imgui.BeginMenuBar() {
			Layout(widgets...)()
			imgui.EndMenuBar()
		}
	}
}

func MenuItem(label string) Widget {
	return MenuItemV(label, false, true, nil)
}

func MenuItemV(label string, selected, enabled bool, clicked func()) Widget {
	return func() {
		if imgui.MenuItemV(label, "", selected, enabled) && clicked != nil {
			clicked()
		}
	}
}

func Menu(label string, widgets ...Widget) Widget {
	return MenuV(label, true, widgets...)
}

func MenuV(label string, enabled bool, widgets ...Widget) Widget {
	return func() {
		if imgui.BeginMenuV(label, enabled) {
			Layout(widgets...)()
			imgui.EndMenu()
		}
	}
}

func Popup(name string, widgets ...Widget) Widget {
	return PopupV(name, nil, 0, widgets...)
}

func PopupV(name string, open *bool, flags int, widgets ...Widget) Widget {
	return func() {
		if imgui.BeginPopupModalV(name, open, flags) {
			Layout(widgets...)()
			imgui.EndPopup()
		}
	}
}

func ProgressBar(fraction float32, width, height float32, overlay string) Widget {
	return func() {
		imgui.ProgressBarV(fraction, imgui.Vec2{X: width, Y: height}, overlay)
	}
}

func Selectable(label string, clicked func()) Widget {
	return SelectableV(label, false, 0, 0, 0, clicked)
}

func SelectableV(label string, selected bool, flags int, width, height float32, clicked func()) Widget {
	return func() {
		if imgui.SelectableV(label, selected, flags, imgui.Vec2{X: width, Y: height}) && clicked != nil {
			clicked()
		}
	}
}

func Separator() Widget {
	return func() {
		imgui.Separator()
	}
}

func SliderInt(label string, value *int32, min, max int32, format string) Widget {
	return func() {
		imgui.SliderIntV(label, value, min, max, format)
	}
}

func Dummy(width, height float32) Widget {
	return func() {
		imgui.Dummy(imgui.Vec2{X: width, Y: height})
	}
}

func HSplitter(id string, width, height float32, delta *float32) Widget {
	return func() {
		imgui.InvisibleButton(id, imgui.Vec2{X: width, Y: height})
		if imgui.IsItemActive() {
			*(delta) = imgui.CurrentIO().GetMouseDelta().Y
		} else {
			*(delta) = 0
		}
	}
}

func VSplitter(id string, width, height float32, delta *float32) Widget {
	return func() {
		imgui.InvisibleButton(id, imgui.Vec2{X: width, Y: height})
		if imgui.IsItemActive() {
			*(delta) = imgui.CurrentIO().GetMouseDelta().X
		} else {
			*(delta) = 0
		}
	}
}

func TabItem(label string, widgets ...Widget) Widget {
	return TabItemV(label, nil, 0, widgets...)
}

func TabItemV(label string, open *bool, flags int, widgets ...Widget) Widget {
	return func() {
		if imgui.BeginTabItemV(label, open, flags) {
			Layout(widgets...)()
			imgui.EndTabItem()
		}
	}
}

func TabBar(id string, tabs ...Widget) Widget {
	return TabBarV(id, 0, tabs...)
}

func TabBarV(id string, flags int, tabs ...Widget) Widget {
	return func() {
		if imgui.BeginTabBarV(id, flags) {
			for _, tab := range tabs {
				tab()
			}
			imgui.EndTabBar()
		}
	}
}

func Row(widgets ...Widget) Widget {
	return func() {
		for _, w := range widgets {
			w()
			imgui.NextColumn()
		}
	}
}

func Table(label string, border bool, columnCount int, rows ...Widget) Widget {
	return func() {
		imgui.ColumnsV(columnCount, label, border)

		for _, r := range rows {
			if border {
				imgui.Separator()
			}
			r()
		}

		imgui.Columns()
		if border {
			imgui.Separator()
		}
	}
}

func Tooltip(tip string) Widget {
	return func() {
		if imgui.IsItemHovered() {
			imgui.SetTooltip(tip)
		}
	}
}

func TreeNode(label string, flags int, widgets ...Widget) Widget {
	return func() {
		if imgui.TreeNodeV(label, flags) {
			Layout(widgets...)()
			if (flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
				imgui.TreePop()
			}
		}
	}
}
