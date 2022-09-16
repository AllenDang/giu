package giu

import (
	"fmt"

	imgui "github.com/AllenDang/cimgui-go"
)

var _ Widget = &ButtonWidget{}

// ButtonWidget represents a ImGui button widget.
type ButtonWidget struct {
	id       string
	width    float32
	height   float32
	disabled bool
	onClick  func()
}

// Button creates a new button widget.
func Button(label string) *ButtonWidget {
	return &ButtonWidget{
		id:      GenAutoID(label),
		width:   0,
		height:  0,
		onClick: nil,
	}
}

// Buttonf creates button with formated label
// NOTE: works like fmt.Sprintf (see `go doc fmt`).
func Buttonf(format string, args ...any) *ButtonWidget {
	return Button(fmt.Sprintf(format, args...))
}

// OnClick sets callback called when button is clicked
// NOTE: to set double click, see EventHandler.go.
func (b *ButtonWidget) OnClick(onClick func()) *ButtonWidget {
	b.onClick = onClick
	return b
}

// Disabled sets button's disabled state
// NOTE: same effect as Style().SetDisabled.
func (b *ButtonWidget) Disabled(d bool) *ButtonWidget {
	b.disabled = d
	return b
}

// Size sets button's size.
func (b *ButtonWidget) Size(width, height float32) *ButtonWidget {
	b.width, b.height = width, height
	return b
}

// Build implements Widget interface.
func (b *ButtonWidget) Build() {
	if b.disabled {
		imgui.BeginDisabled()
		defer imgui.EndDisabled()
	}

	if imgui.ButtonV(Context.FontAtlas.RegisterString(b.id), imgui.ImVec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &ArrowButtonWidget{}

// ArrowButtonWidget represents a square button with an arrow.
type ArrowButtonWidget struct {
	id      string
	dir     imgui.ImGuiDir
	onClick func()
}

// ArrowButton creates ArrowButtonWidget.
func ArrowButton(dir imgui.ImGuiDir) *ArrowButtonWidget {
	return &ArrowButtonWidget{
		id:      GenAutoID("ArrowButton"),
		dir:     dir,
		onClick: nil,
	}
}

// OnClick adds callback called when button is clicked.
func (b *ArrowButtonWidget) OnClick(onClick func()) *ArrowButtonWidget {
	b.onClick = onClick
	return b
}

// ID allows to manually set widget's id.
func (b *ArrowButtonWidget) ID(id string) *ArrowButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *ArrowButtonWidget) Build() {
	if imgui.ArrowButton(b.id, b.dir) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &SmallButtonWidget{}

// SmallButtonWidget is like a button but without frame padding.
type SmallButtonWidget struct {
	id      string
	onClick func()
}

// SmallButton constructs a new small button widget.
func SmallButton(id string) *SmallButtonWidget {
	return &SmallButtonWidget{
		id:      GenAutoID(id),
		onClick: nil,
	}
}

// SmallButtonf allows to set formated label for small button.
// It calls SmallButton(fmt.Sprintf(label, args...)).
func SmallButtonf(format string, args ...any) *SmallButtonWidget {
	return SmallButton(fmt.Sprintf(format, args...))
}

// OnClick adds OnClick event.
func (b *SmallButtonWidget) OnClick(onClick func()) *SmallButtonWidget {
	b.onClick = onClick
	return b
}

// Build implements Widget interface.
func (b *SmallButtonWidget) Build() {
	if imgui.SmallButton(Context.FontAtlas.RegisterString(b.id)) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &InvisibleButtonWidget{}

// InvisibleButtonWidget is a clickable region.
// NOTE: you may want to display other widgets on this button.
// to do so, you may move drawing cursor back by Get/SetCursor(Screen)Pos.
type InvisibleButtonWidget struct {
	id      string
	width   float32
	height  float32
	onClick func()
}

// InvisibleButton constructs a new invisible button widget.
func InvisibleButton() *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      GenAutoID("InvisibleButton"),
		width:   0,
		height:  0,
		onClick: nil,
	}
}

// Size sets button's size.
func (b *InvisibleButtonWidget) Size(width, height float32) *InvisibleButtonWidget {
	b.width, b.height = width, height
	return b
}

// OnClick sets click event.
func (b *InvisibleButtonWidget) OnClick(onClick func()) *InvisibleButtonWidget {
	b.onClick = onClick
	return b
}

// ID allows to manually set widget's id (no need to use in normal conditions).
func (b *InvisibleButtonWidget) ID(id string) *InvisibleButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButtonV(Context.FontAtlas.RegisterString(b.id), imgui.ImVec2{X: b.width, Y: b.height}, 0) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &CheckboxWidget{}

// CheckboxWidget adds a checkbox.
type CheckboxWidget struct {
	text     string
	selected *bool
	onChange func()
}

// Checkbox creates a new CheckboxWidget.
func Checkbox(text string, selected *bool) *CheckboxWidget {
	return &CheckboxWidget{
		text:     GenAutoID(text),
		selected: selected,
		onChange: nil,
	}
}

// OnChange adds callback called when checkbox's state was changed.
func (c *CheckboxWidget) OnChange(onChange func()) *CheckboxWidget {
	c.onChange = onChange
	return c
}

// Build implements Widget interface.
func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(Context.FontAtlas.RegisterString(c.text), c.selected) && c.onChange != nil {
		c.onChange()
	}
}

var _ Widget = &RadioButtonWidget{}

// RadioButtonWidget is a small, round button.
// It is common to use it for single-choice questions.
// see examples/widgets.
type RadioButtonWidget struct {
	text     string
	active   bool
	onChange func()
}

// RadioButton creates a radio buton.
func RadioButton(text string, active bool) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:     GenAutoID(text),
		active:   active,
		onChange: nil,
	}
}

// OnChange adds callback when button's state gets changed.
func (r *RadioButtonWidget) OnChange(onChange func()) *RadioButtonWidget {
	r.onChange = onChange
	return r
}

// Build implements Widget interface.
func (r *RadioButtonWidget) Build() {
	if imgui.RadioButton_Bool(Context.FontAtlas.RegisterString(r.text), r.active) && r.onChange != nil {
		r.onChange()
	}
}

var _ Widget = &SelectableWidget{}

// SelectableWidget is a window-width button with a label which can get selected (highlighted).
// useful for certain lists.
type SelectableWidget struct {
	label    string
	selected bool
	flags    imgui.ImGuiSelectableFlags
	width    float32
	height   float32
	onClick  func()
	onDClick func()
}

// Selectable constructs a selectable widget.
func Selectable(label string) *SelectableWidget {
	return &SelectableWidget{
		label:    GenAutoID(label),
		selected: false,
		flags:    0,
		width:    0,
		height:   0,
		onClick:  nil,
	}
}

// Selectablef creates a selectable widget with formated label.
func Selectablef(format string, args ...any) *SelectableWidget {
	return Selectable(fmt.Sprintf(format, args...))
}

// Selected sets if selectable widget is selected.
func (s *SelectableWidget) Selected(selected bool) *SelectableWidget {
	s.selected = selected
	return s
}

// Flags add flags.
func (s *SelectableWidget) Flags(flags imgui.ImGuiSelectableFlags) *SelectableWidget {
	s.flags = flags
	return s
}

// Size sets selectable's size.
func (s *SelectableWidget) Size(width, height float32) *SelectableWidget {
	s.width, s.height = width, height
	return s
}

// OnClick sets on click event.
func (s *SelectableWidget) OnClick(onClick func()) *SelectableWidget {
	s.onClick = onClick
	return s
}

// OnDClick handles mouse left button's double click event.
// SelectableFlagsAllowDoubleClick will set once tonDClick callback is notnull.
// NOTE: IT IS DEPRECATED and could be removed. Use EventHandler instead.
func (s *SelectableWidget) OnDClick(onDClick func()) *SelectableWidget {
	s.onDClick = onDClick
	return s
}

// Build implements Widget interface.
func (s *SelectableWidget) Build() {
	// If onDClick is set, check flags and set related flag when necessary
	if s.onDClick != nil && s.flags&imgui.ImGuiSelectableFlags_AllowDoubleClick != 0 {
		s.flags |= imgui.ImGuiSelectableFlags_AllowDoubleClick
	}

	if imgui.Selectable_BoolV(Context.FontAtlas.RegisterString(s.label), s.selected, s.flags, imgui.ImVec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}

	if s.onDClick != nil && imgui.IsItemActive() && imgui.IsMouseDoubleClicked(imgui.ImGuiMouseButton_Left) {
		s.onDClick()
	}
}

var _ Widget = &TreeNodeWidget{}

// TreeNodeWidget is a a wide button with open/close state.
// if is opened, the `layout` is displayed below the widget.
// It can be used to create certain lists, advanced settings sections e.t.c.
type TreeNodeWidget struct {
	label        string
	flags        imgui.ImGuiTreeNodeFlags
	layout       Layout
	eventHandler func()
}

// TreeNode creates a new tree node widget.
func TreeNode(label string) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:        Context.FontAtlas.RegisterString(label),
		flags:        0,
		layout:       nil,
		eventHandler: nil,
	}
}

// TreeNodef adds TreeNode with formatted label.
func TreeNodef(format string, args ...any) *TreeNodeWidget {
	return TreeNode(fmt.Sprintf(format, args...))
}

// Flags sets flags.
func (t *TreeNodeWidget) Flags(flags imgui.ImGuiTreeNodeFlags) *TreeNodeWidget {
	t.flags = flags
	return t
}

// Event create TreeNode with eventHandler
// You could detect events (e.g. IsItemClicked IsMouseDoubleClicked etc...) and handle them for TreeNode inside eventHandler.
// Deprecated: Use EventHandler instead!
func (t *TreeNodeWidget) Event(handler func()) *TreeNodeWidget {
	t.eventHandler = handler
	return t
}

// Layout sets layout to be displayed when tree node is opened.
func (t *TreeNodeWidget) Layout(widgets ...Widget) *TreeNodeWidget {
	t.layout = Layout(widgets)
	return t
}

// Build implements Widget interface.
func (t *TreeNodeWidget) Build() {
	open := imgui.TreeNodeEx_StrV(t.label, t.flags)

	if t.eventHandler != nil {
		t.eventHandler()
	}

	if open {
		t.layout.Build()
		if (t.flags & imgui.ImGuiTreeNodeFlags_NoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}
