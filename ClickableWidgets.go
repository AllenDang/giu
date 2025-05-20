package giu

import (
	"fmt"
	"image"
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
	"golang.org/x/image/colornames"
)

var _ Widget = &ButtonWidget{}

// ButtonWidget represents a ImGui button widget.
type ButtonWidget struct {
	id       ID
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

// Buttonf creates button with formatted label
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

// ID allows to manually set widget's id.
func (b *ButtonWidget) ID(id ID) *ButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *ButtonWidget) Build() {
	if b.disabled {
		imgui.BeginDisabled()
		defer imgui.EndDisabled()
	}

	if imgui.ButtonV(Context.PrepareString(b.id.String()), imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &ArrowButtonWidget{}

// ArrowButtonWidget represents a square button with an arrow.
type ArrowButtonWidget struct {
	id      ID
	dir     Direction
	onClick func()
}

// ArrowButton creates ArrowButtonWidget.
func ArrowButton(dir Direction) *ArrowButtonWidget {
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
func (b *ArrowButtonWidget) ID(id ID) *ArrowButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *ArrowButtonWidget) Build() {
	if imgui.ArrowButton(b.id.String(), imgui.Dir(b.dir)) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &SmallButtonWidget{}

// SmallButtonWidget is like a button but without frame padding.
type SmallButtonWidget struct {
	id      ID
	onClick func()
}

// SmallButton constructs a new small button widget.
func SmallButton(id string) *SmallButtonWidget {
	return &SmallButtonWidget{
		id:      GenAutoID(id),
		onClick: nil,
	}
}

// SmallButtonf allows to set formatted label for small button.
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
	if imgui.SmallButton(Context.PrepareString(b.id.String())) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &InvisibleButtonWidget{}

// InvisibleButtonWidget is a clickable region.
// NOTE: you may want to display other widgets on this button.
// to do so, you may move drawing cursor back by Get/SetCursor(Screen)Pos.
type InvisibleButtonWidget struct {
	id      ID
	width   float32
	height  float32
	onClick func()
}

// InvisibleButton constructs a new invisible button widget.
func InvisibleButton() *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      GenAutoID("InvisibleButton"),
		width:   Auto,
		height:  Auto,
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
func (b *InvisibleButtonWidget) ID(id ID) *InvisibleButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButton(Context.FontAtlas.RegisterString(b.id.String()), imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &ImageButtonWidget{}

// ImageButtonWidget is similar to ButtonWidget but with image texture instead of text label.
type ImageButtonWidget struct {
	texture      *Texture
	width        float32
	height       float32
	uv0          image.Point
	uv1          image.Point
	framePadding int
	bgColor      color.Color
	tintColor    color.Color
	onClick      func()
	id           ID
}

// ImageButton  constructs image button widget.
func ImageButton(texture *Texture) *ImageButtonWidget {
	return &ImageButtonWidget{
		texture:      texture,
		width:        50,
		height:       50,
		uv0:          image.Point{X: 0, Y: 0},
		uv1:          image.Point{X: 1, Y: 1},
		framePadding: -1,
		bgColor:      colornames.Black,
		tintColor:    colornames.White,
		onClick:      nil,
		id:           GenAutoID("ImageButton"),
	}
}

// ID allows to manually set widget's id.
func (b *ImageButtonWidget) ID(id ID) *ImageButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *ImageButtonWidget) Build() {
	if b.texture == nil || b.texture.tex == nil {
		return
	}

	imgui.PushIDStr(b.id.String())

	if imgui.ImageButtonV(
		fmt.Sprintf("%v", b.texture.tex.ID),
		b.texture.tex.ID,
		imgui.Vec2{X: b.width, Y: b.height},
		ToVec2(b.uv0), ToVec2(b.uv1),
		ToVec4Color(b.bgColor),
		ToVec4Color(b.tintColor),
	) && b.onClick != nil {
		b.onClick()
	}

	imgui.PopID()
}

// Size sets BUTTONS size.
// NOTE: image size is button size - 2 * frame padding.
func (b *ImageButtonWidget) Size(width, height float32) *ImageButtonWidget {
	b.width, b.height = width, height
	return b
}

// OnClick sets click event.
func (b *ImageButtonWidget) OnClick(onClick func()) *ImageButtonWidget {
	b.onClick = onClick
	return b
}

// UV sets image's uv.
func (b *ImageButtonWidget) UV(uv0, uv1 image.Point) *ImageButtonWidget {
	b.uv0, b.uv1 = uv0, uv1
	return b
}

// BgColor sets button's background color.
func (b *ImageButtonWidget) BgColor(bgColor color.Color) *ImageButtonWidget {
	b.bgColor = bgColor
	return b
}

// TintColor sets tit color for image.
func (b *ImageButtonWidget) TintColor(tintColor color.Color) *ImageButtonWidget {
	b.tintColor = tintColor
	return b
}

// FramePadding sets button's frame padding (set 0 to fit image to the frame).
func (b *ImageButtonWidget) FramePadding(padding int) *ImageButtonWidget {
	b.framePadding = padding
	return b
}

var _ Widget = &ImageButtonWithRgbaWidget{}

// ImageButtonWithRgbaWidget does similar to ImageButtonWidget,
// but implements image.Image instead of giu.Texture. It is probably
// more useful than the original ImageButtonWidget.
type ImageButtonWithRgbaWidget struct {
	*ImageButtonWidget
	rgba image.Image
	id   ID
}

// ImageButtonWithRgba creates a new widget.
func ImageButtonWithRgba(rgba image.Image) *ImageButtonWithRgbaWidget {
	return &ImageButtonWithRgbaWidget{
		id:                GenAutoID("ImageButtonWithRgba"),
		ImageButtonWidget: ImageButton(nil),
		rgba:              rgba,
	}
}

// Size sets button's size.
func (b *ImageButtonWithRgbaWidget) Size(width, height float32) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.Size(width, height)
	return b
}

// OnClick sets click events.
func (b *ImageButtonWithRgbaWidget) OnClick(onClick func()) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.OnClick(onClick)
	return b
}

// UV sets image's uv color.
func (b *ImageButtonWithRgbaWidget) UV(uv0, uv1 image.Point) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.UV(uv0, uv1)
	return b
}

// BgColor sets button's background color.
func (b *ImageButtonWithRgbaWidget) BgColor(bgColor color.Color) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.BgColor(bgColor)
	return b
}

// TintColor sets image's tint color.
func (b *ImageButtonWithRgbaWidget) TintColor(tintColor color.Color) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.TintColor(tintColor)
	return b
}

// FramePadding sets frame padding (see (*ImageButtonWidget).TintColor).
func (b *ImageButtonWithRgbaWidget) FramePadding(padding int) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.FramePadding(padding)
	return b
}

// Build implements Widget interface.
func (b *ImageButtonWithRgbaWidget) Build() {
	if state := GetState[imageState](Context, b.id); state == nil {
		SetState(Context, b.id, &imageState{})

		NewTextureFromRgba(b.rgba, func(tex *Texture) {
			SetState(Context, b.id, &imageState{texture: tex})
		})
	} else {
		b.texture = state.texture
	}

	b.ImageButtonWidget.Build()
}

var _ Widget = &CheckboxWidget{}

// CheckboxWidget adds a checkbox.
type CheckboxWidget struct {
	text     ID
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

// ID sets widget's id (overrides text).
func (c *CheckboxWidget) ID(id ID) *CheckboxWidget {
	c.text = id
	return c
}

// OnChange adds callback called when checkbox's state was changed.
func (c *CheckboxWidget) OnChange(onChange func()) *CheckboxWidget {
	c.onChange = onChange
	return c
}

// Build implements Widget interface.
func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(Context.PrepareString(c.text.String()), c.selected) && c.onChange != nil {
		c.onChange()
	}
}

var _ Widget = &RadioButtonWidget{}

// RadioButtonWidget is a small, round button.
// It is common to use it for single-choice questions.
// see examples/widgets.
type RadioButtonWidget struct {
	text     ID
	active   bool
	onChange func()
}

// RadioButton creates a radio button.
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
	if imgui.RadioButtonBool(Context.PrepareString(r.text.String()), r.active) && r.onChange != nil {
		r.onChange()
	}
}

var _ Widget = &SelectableWidget{}

// SelectableWidget is a window-width button with a label which can get selected (highlighted).
// useful for certain lists.
type SelectableWidget struct {
	label    ID
	selected bool
	flags    SelectableFlags
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

// Selectablef creates a selectable widget with formatted label.
func Selectablef(format string, args ...any) *SelectableWidget {
	return Selectable(fmt.Sprintf(format, args...))
}

// Selected sets if selectable widget is selected.
func (s *SelectableWidget) Selected(selected bool) *SelectableWidget {
	s.selected = selected
	return s
}

// Flags add flags.
func (s *SelectableWidget) Flags(flags SelectableFlags) *SelectableWidget {
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
	if s.onDClick != nil && s.flags&SelectableFlagsAllowDoubleClick != 0 {
		s.flags |= SelectableFlagsAllowDoubleClick
	}

	if imgui.SelectableBoolV(Context.PrepareString(s.label.String()), s.selected, imgui.SelectableFlags(s.flags), imgui.Vec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}

	if s.onDClick != nil && IsItemActive() && IsMouseDoubleClicked(MouseButtonLeft) {
		s.onDClick()
	}
}

var _ Widget = &TreeNodeWidget{}

// TreeNodeWidget is a a wide button with open/close state.
// if is opened, the `layout` is displayed below the widget.
// It can be used to create certain lists, advanced settings sections e.t.c.
type TreeNodeWidget struct {
	label        string
	flags        TreeNodeFlags
	layout       Layout
	event        func()
	eventHandler *EventHandler
}

// TreeNode creates a new tree node widget.
func TreeNode(label string) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:        label,
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
func (t *TreeNodeWidget) Flags(flags TreeNodeFlags) *TreeNodeWidget {
	t.flags = flags
	return t
}

// Event create TreeNode with event handling function.
// You could detect events (e.g. IsItemClicked IsMouseDoubleClicked etc...) and handle them for TreeNode inside eventHandler.
// Deprecated: Use (*TreeNodeWidget).EventHandler instead!
func (t *TreeNodeWidget) Event(handler func()) *TreeNodeWidget {
	t.event = handler
	return t
}

// EventHandler allows to set *EventHandler instance for the actual TreeNode.
func (t *TreeNodeWidget) EventHandler(handler *EventHandler) *TreeNodeWidget {
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
	open := imgui.TreeNodeExStrV(Context.PrepareString(t.label), imgui.TreeNodeFlags(t.flags))

	if t.event != nil {
		t.event()
	}

	if t.eventHandler != nil {
		t.eventHandler.Build()
	}

	if open {
		t.layout.Build()

		if (imgui.TreeNodeFlags(t.flags) & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}

var _ Widget = &LinkWidget{}

// LinkWidget is a clickable text fragment.
type LinkWidget struct {
	text    ID
	onClick func()
}

// Link constructs link widget.
func Link(text string) *LinkWidget {
	return &LinkWidget{
		text: GenAutoID(text),
	}
}

// Linkf allows to add formatted link.
func Linkf(format string, args ...any) *LinkWidget {
	return Link(fmt.Sprintf(format, args...))
}

// ID allows to manually set widget's id (in this case - text). Baypasses GenAutoID mechanism in cae this is needed.
func (l *LinkWidget) ID(id ID) *LinkWidget {
	l.text = id
	return l
}

// OnClick sets click callback.
func (l *LinkWidget) OnClick(onClick func()) *LinkWidget {
	l.onClick = onClick
	return l
}

// Build implements Widget interface.
func (l *LinkWidget) Build() {
	if imgui.TextLink(Context.FontAtlas.RegisterString(l.text.String())) && l.onClick != nil {
		l.onClick()
	}
}
