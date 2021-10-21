package giu

import (
	"fmt"
	"image"
	"image/color"

	"github.com/AllenDang/imgui-go"
	"golang.org/x/image/colornames"
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

// Build implements Widget interface.
func (b *ButtonWidget) Build() {
	if b.disabled {
		imgui.BeginDisabled(true)
		defer imgui.EndDisabled()
	}

	if imgui.ButtonV(Context.FontAtlas.tStr(b.id), imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
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
func Buttonf(format string, args ...interface{}) *ButtonWidget {
	return Button(fmt.Sprintf(format, args...))
}

var _ Widget = &ArrowButtonWidget{}

// ArrowButtonWidget represents a square button with an arrow.
type ArrowButtonWidget struct {
	id      string
	dir     Direction
	onClick func()
}

// OnClick adds callback called when button is clicked.
func (b *ArrowButtonWidget) OnClick(onClick func()) *ArrowButtonWidget {
	b.onClick = onClick
	return b
}

// ArrowButton creates ArrowButtonWidget.
func ArrowButton(dir Direction) *ArrowButtonWidget {
	return &ArrowButtonWidget{
		id:      GenAutoID("ArrowButton"),
		dir:     dir,
		onClick: nil,
	}
}

// ID allows to manually set widget's id.
func (b *ArrowButtonWidget) ID(id string) *ArrowButtonWidget {
	b.id = id
	return b
}

// Build implements Widget interface.
func (b *ArrowButtonWidget) Build() {
	if imgui.ArrowButton(b.id, uint8(b.dir)) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &SmallButtonWidget{}

type SmallButtonWidget struct {
	id      string
	onClick func()
}

func (b *SmallButtonWidget) OnClick(onClick func()) *SmallButtonWidget {
	b.onClick = onClick
	return b
}

func SmallButton(id string) *SmallButtonWidget {
	return &SmallButtonWidget{
		id:      GenAutoID(id),
		onClick: nil,
	}
}

func SmallButtonf(format string, args ...interface{}) *SmallButtonWidget {
	return SmallButton(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (b *SmallButtonWidget) Build() {
	if imgui.SmallButton(Context.FontAtlas.tStr(b.id)) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &InvisibleButtonWidget{}

type InvisibleButtonWidget struct {
	id      string
	width   float32
	height  float32
	onClick func()
}

func (b *InvisibleButtonWidget) Size(width, height float32) *InvisibleButtonWidget {
	b.width, b.height = width, height
	return b
}

func (b *InvisibleButtonWidget) OnClick(onClick func()) *InvisibleButtonWidget {
	b.onClick = onClick
	return b
}

func (b *InvisibleButtonWidget) ID(id string) *InvisibleButtonWidget {
	b.id = id
	return b
}

func InvisibleButton() *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      GenAutoID("InvisibleButton"),
		width:   0,
		height:  0,
		onClick: nil,
	}
}

// Build implements Widget interface.
func (b *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButton(Context.FontAtlas.tStr(b.id), imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
}

var _ Widget = &ImageButtonWidget{}

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
}

// Build implements Widget interface.
func (b *ImageButtonWidget) Build() {
	if b.texture == nil && b.texture.id == 0 {
		return
	}

	if imgui.ImageButtonV(
		b.texture.id,
		imgui.Vec2{X: b.width, Y: b.height},
		ToVec2(b.uv0), ToVec2(b.uv1),
		b.framePadding, ToVec4Color(b.bgColor),
		ToVec4Color(b.tintColor),
	) && b.onClick != nil {
		b.onClick()
	}
}

func (b *ImageButtonWidget) Size(width, height float32) *ImageButtonWidget {
	b.width, b.height = width, height
	return b
}

func (b *ImageButtonWidget) OnClick(onClick func()) *ImageButtonWidget {
	b.onClick = onClick
	return b
}

func (b *ImageButtonWidget) UV(uv0, uv1 image.Point) *ImageButtonWidget {
	b.uv0, b.uv1 = uv0, uv1
	return b
}

func (b *ImageButtonWidget) BgColor(bgColor color.Color) *ImageButtonWidget {
	b.bgColor = bgColor
	return b
}

func (b *ImageButtonWidget) TintColor(tintColor color.Color) *ImageButtonWidget {
	b.tintColor = tintColor
	return b
}

func (b *ImageButtonWidget) FramePadding(padding int) *ImageButtonWidget {
	b.framePadding = padding
	return b
}

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
	}
}

var _ Widget = &ImageButtonWithRgbaWidget{}

type ImageButtonWithRgbaWidget struct {
	*ImageButtonWidget
	rgba image.Image
	id   string
}

func ImageButtonWithRgba(rgba image.Image) *ImageButtonWithRgbaWidget {
	return &ImageButtonWithRgbaWidget{
		id:                GenAutoID("ImageButtonWithRgba"),
		ImageButtonWidget: ImageButton(nil),
		rgba:              rgba,
	}
}

func (b *ImageButtonWithRgbaWidget) Size(width, height float32) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.Size(width, height)
	return b
}

func (b *ImageButtonWithRgbaWidget) OnClick(onClick func()) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.OnClick(onClick)
	return b
}

func (b *ImageButtonWithRgbaWidget) UV(uv0, uv1 image.Point) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.UV(uv0, uv1)
	return b
}

func (b *ImageButtonWithRgbaWidget) BgColor(bgColor color.Color) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.BgColor(bgColor)
	return b
}

func (b *ImageButtonWithRgbaWidget) TintColor(tintColor color.Color) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.TintColor(tintColor)
	return b
}

func (b *ImageButtonWithRgbaWidget) FramePadding(padding int) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.FramePadding(padding)
	return b
}

// Build implements Widget interface.
func (b *ImageButtonWithRgbaWidget) Build() {
	if state := Context.GetState(b.id); state == nil {
		Context.SetState(b.id, &ImageState{})

		NewTextureFromRgba(b.rgba, func(tex *Texture) {
			Context.SetState(b.id, &ImageState{texture: tex})
		})
	} else {
		var isOk bool
		imgState, isOk := state.(*ImageState)
		Assert(isOk, "ImageButtonWithRgbaWidget", "Build", "got unexpected type of widget's state")
		b.ImageButtonWidget.texture = imgState.texture
	}

	b.ImageButtonWidget.Build()
}

var _ Widget = &CheckboxWidget{}

// CheckboxWidget adds a checkbox.
type CheckboxWidget struct {
	text     string
	selected *bool
	onChange func()
}

// Build implements Widget interface.
func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(Context.FontAtlas.tStr(c.text), c.selected) && c.onChange != nil {
		c.onChange()
	}
}

// OnChange adds callback called when checkbox's state was changed.
func (c *CheckboxWidget) OnChange(onChange func()) *CheckboxWidget {
	c.onChange = onChange
	return c
}

// Checkbox creates a new CheckboxWidget.
func Checkbox(text string, selected *bool) *CheckboxWidget {
	return &CheckboxWidget{
		text:     GenAutoID(text),
		selected: selected,
		onChange: nil,
	}
}

var _ Widget = &RadioButtonWidget{}

type RadioButtonWidget struct {
	text     string
	active   bool
	onChange func()
}

// Build implements Widget interface.
func (r *RadioButtonWidget) Build() {
	if imgui.RadioButton(Context.FontAtlas.tStr(r.text), r.active) && r.onChange != nil {
		r.onChange()
	}
}

func (r *RadioButtonWidget) OnChange(onChange func()) *RadioButtonWidget {
	r.onChange = onChange
	return r
}

func RadioButton(text string, active bool) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:     GenAutoID(text),
		active:   active,
		onChange: nil,
	}
}

var _ Widget = &SelectableWidget{}

type SelectableWidget struct {
	label    string
	selected bool
	flags    SelectableFlags
	width    float32
	height   float32
	onClick  func()
	onDClick func()
}

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

func Selectablef(format string, args ...interface{}) *SelectableWidget {
	return Selectable(fmt.Sprintf(format, args...))
}

func (s *SelectableWidget) Selected(selected bool) *SelectableWidget {
	s.selected = selected
	return s
}

func (s *SelectableWidget) Flags(flags SelectableFlags) *SelectableWidget {
	s.flags = flags
	return s
}

func (s *SelectableWidget) Size(width, height float32) *SelectableWidget {
	s.width, s.height = width, height
	return s
}

func (s *SelectableWidget) OnClick(onClick func()) *SelectableWidget {
	s.onClick = onClick
	return s
}

// OnDClick handles mouse left button's double click event.
// SelectableFlagsAllowDoubleClick will set once tonDClick callback is notnull.
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

	if imgui.SelectableV(Context.FontAtlas.tStr(s.label), s.selected, int(s.flags), imgui.Vec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}

	if s.onDClick != nil && IsItemActive() && IsMouseDoubleClicked(MouseButtonLeft) {
		s.onDClick()
	}
}

var _ Widget = &TreeNodeWidget{}

type TreeNodeWidget struct {
	label        string
	flags        TreeNodeFlags
	layout       Layout
	eventHandler func()
}

func TreeNode(label string) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:        Context.FontAtlas.tStr(label),
		flags:        0,
		layout:       nil,
		eventHandler: nil,
	}
}

func TreeNodef(format string, args ...interface{}) *TreeNodeWidget {
	return TreeNode(fmt.Sprintf(format, args...))
}

func (t *TreeNodeWidget) Flags(flags TreeNodeFlags) *TreeNodeWidget {
	t.flags = flags
	return t
}

// Event create TreeNode with eventHandler
// You could detect events (e.g. IsItemClicked IsMouseDoubleClicked etc...) and handle them for TreeNode inside eventHandler.
func (t *TreeNodeWidget) Event(handler func()) *TreeNodeWidget {
	t.eventHandler = handler
	return t
}

func (t *TreeNodeWidget) Layout(widgets ...Widget) *TreeNodeWidget {
	t.layout = Layout(widgets)
	return t
}

// Build implements Widget interface.
func (t *TreeNodeWidget) Build() {
	open := imgui.TreeNodeV(t.label, int(t.flags))

	if t.eventHandler != nil {
		t.eventHandler()
	}

	if open {
		t.layout.Build()
		if (t.flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}
