package giu

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ianling/imgui-go"
)

type LineWidget struct {
	widgets []Widget
}

func Line(widgets ...Widget) *LineWidget {
	return &LineWidget{
		widgets: widgets,
	}
}

func (l *LineWidget) Build() {
	index := 0

	for _, w := range l.widgets {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopupModal := w.(*PopupModalWidget)
		_, isPopup := w.(*PopupWidget)
		_, isTabItem := w.(*TabItemWidget)
		_, isLabel := w.(*LabelWidget)
		_, isCustom := w.(*CustomWidget)

		if isLabel {
			AlignTextToFramePadding()
		}

		if index > 0 && !isTooltip && !isContextMenu && !isPopupModal && !isPopup && !isTabItem && !isCustom {
			imgui.SameLine()
		}

		if !isCustom {
			index += 1
		}

		w.Build()
	}
}

func SameLine() {
	imgui.SameLine()
}

type InputTextMultilineWidget struct {
	label         string
	text          *string
	width, height float32
	flags         InputTextFlags
	cb            imgui.InputTextCallback
	onChange      func()
}

func (i *InputTextMultilineWidget) Build() {
	if imgui.InputTextMultilineV(i.label, i.text, imgui.Vec2{X: i.width, Y: i.height}, int(i.flags), i.cb) && i.onChange != nil {
		i.onChange()
	}
}

func (i *InputTextMultilineWidget) Flags(flags InputTextFlags) *InputTextMultilineWidget {
	i.flags = flags
	return i
}

func (i *InputTextMultilineWidget) Callback(cb imgui.InputTextCallback) *InputTextMultilineWidget {
	i.cb = cb
	return i
}

func (i *InputTextMultilineWidget) OnChange(onChange func()) *InputTextMultilineWidget {
	i.onChange = onChange
	return i
}

func (i *InputTextMultilineWidget) Size(width, height float32) *InputTextMultilineWidget {
	scale := Context.platform.GetContentScale()
	i.width, i.height = width*scale, height*scale
	return i
}

func InputTextMultiline(label string, text *string) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		label:    label,
		text:     text,
		width:    0,
		height:   0,
		flags:    0,
		cb:       nil,
		onChange: nil,
	}
}

type ButtonWidget struct {
	id      string
	width   float32
	height  float32
	onClick func()
}

func (b *ButtonWidget) Build() {
	if imgui.ButtonV(b.id, imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}
}

func (b *ButtonWidget) OnClick(onClick func()) *ButtonWidget {
	b.onClick = onClick
	return b
}

func (b *ButtonWidget) Size(width, height float32) *ButtonWidget {
	scale := Context.platform.GetContentScale()
	b.width, b.height = width*scale, height*scale
	return b
}

func Button(id string) *ButtonWidget {
	return &ButtonWidget{
		id:      id,
		width:   0,
		height:  0,
		onClick: nil,
	}
}

type PlotLinesWidget struct {
	label        string
	values       []float32
	valuesOffset int
	overlayText  string
	scaleMin     float32
	scaleMax     float32
	graphSize    imgui.Vec2
}

func (p *PlotLinesWidget) Build() {
	imgui.PlotLinesV(p.label, p.values, p.valuesOffset, p.overlayText, p.scaleMin, p.scaleMax, p.graphSize)
}

func PlotLines(label string, values []float32) *PlotLinesWidget {
	return PlotLinesV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, 0, 0)
}

func PlotLinesV(label string, values []float32, valuesOffset int, overlayText string, scaleMin, scaleMax, width, height float32) *PlotLinesWidget {
	return &PlotLinesWidget{
		label:        label,
		values:       values,
		valuesOffset: valuesOffset,
		overlayText:  overlayText,
		scaleMin:     scaleMin,
		scaleMax:     scaleMax,
		graphSize:    imgui.Vec2{X: width, Y: height},
	}
}

type BulletWidget struct{}

func Bullet() *BulletWidget {
	return &BulletWidget{}
}

func (b *BulletWidget) Build() {
	imgui.Bullet()
}

type BulletTextWidget struct {
	text string
}

func BulletText(text string) *BulletTextWidget {
	return &BulletTextWidget{
		text: text,
	}
}

func (bt *BulletTextWidget) Build() {
	imgui.BulletText(bt.text)
}

type ArrowButtonWidget struct {
	id      string
	dir     Direction
	onClick func()
}

func (b *ArrowButtonWidget) OnClick(onClick func()) *ArrowButtonWidget {
	b.onClick = onClick
	return b
}

func ArrowButton(id string, dir Direction) *ArrowButtonWidget {
	return &ArrowButtonWidget{
		id:      id,
		dir:     dir,
		onClick: nil,
	}
}

func (b *ArrowButtonWidget) Build() {
	if imgui.ArrowButton(b.id, uint8(b.dir)) && b.onClick != nil {
		b.onClick()
	}
}

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
		id:      id,
		onClick: nil,
	}
}

func (b *SmallButtonWidget) Build() {
	if imgui.SmallButton(b.id) && b.onClick != nil {
		b.onClick()
	}
}

type InvisibleButtonWidget struct {
	id      string
	width   float32
	height  float32
	onClick func()
}

func (b *InvisibleButtonWidget) Size(width, height float32) *InvisibleButtonWidget {
	scale := Context.platform.GetContentScale()
	b.width, b.height = width*scale, height*scale
	return b
}

func (b *InvisibleButtonWidget) OnClick(onClick func()) *InvisibleButtonWidget {
	b.onClick = onClick
	return b
}

func InvisibleButton(id string) *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      id,
		width:   0,
		height:  0,
		onClick: nil,
	}
}

func (b *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButtonV(b.id, imgui.Vec2{X: b.width, Y: b.height}, imgui.ButtonFlagsNone) && b.onClick != nil {
		b.onClick()
	}
}

type ImageButtonWidget struct {
	texture      *Texture
	width        float32
	height       float32
	uv0          image.Point
	uv1          image.Point
	framePadding int
	bgColor      color.RGBA
	tintColor    color.RGBA
	onClick      func()
}

func (b *ImageButtonWidget) Build() {
	if b.texture != nil && b.texture.id != 0 {
		if imgui.ImageButtonV(b.texture.id, imgui.Vec2{X: b.width, Y: b.height}, ToVec2(b.uv0), ToVec2(b.uv1), b.framePadding, ToVec4Color(b.bgColor), ToVec4Color(b.tintColor)) && b.onClick != nil {
			b.onClick()
		}
	}
}

func (b *ImageButtonWidget) Size(width, height float32) *ImageButtonWidget {
	scale := Context.platform.GetContentScale()
	b.width, b.height = width*scale, height*scale
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

func (b *ImageButtonWidget) BgColor(bgColor color.RGBA) *ImageButtonWidget {
	b.bgColor = bgColor
	return b
}

func (b *ImageButtonWidget) TintColor(tintColor color.RGBA) *ImageButtonWidget {
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
		width:        50 * Context.platform.GetContentScale(),
		height:       50 * Context.platform.GetContentScale(),
		uv0:          image.Point{X: 0, Y: 0},
		uv1:          image.Point{X: 1, Y: 1},
		framePadding: -1,
		bgColor:      color.RGBA{0, 0, 0, 0},
		tintColor:    color.RGBA{255, 255, 255, 255},
		onClick:      nil,
	}
}

type CheckboxWidget struct {
	text     string
	selected *bool
	onChange func()
}

func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(c.text, c.selected) && c.onChange != nil {
		c.onChange()
	}
}

func (c *CheckboxWidget) OnChange(onChange func()) *CheckboxWidget {
	c.onChange = onChange
	return c
}

func Checkbox(text string, selected *bool) *CheckboxWidget {
	return &CheckboxWidget{
		text:     text,
		selected: selected,
		onChange: nil,
	}
}

type RadioButtonWidget struct {
	text     string
	active   bool
	onChange func()
}

func (r *RadioButtonWidget) Build() {
	if imgui.RadioButton(r.text, r.active) && r.onChange != nil {
		r.onChange()
	}
}

func (r *RadioButtonWidget) OnChange(onChange func()) *RadioButtonWidget {
	r.onChange = onChange
	return r
}

func RadioButton(text string, active bool) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:     text,
		active:   active,
		onChange: nil,
	}
}

type ChildWidget struct {
	id     string
	width  float32
	height float32
	border bool
	flags  WindowFlags
	layout Layout
}

func (c *ChildWidget) Build() {
	imgui.BeginChildV(c.id, imgui.Vec2{X: c.width, Y: c.height}, c.border, int(c.flags))
	if c.layout != nil {
		c.layout.Build()
	}
	imgui.EndChild()
}

func (c *ChildWidget) Border(border bool) *ChildWidget {
	c.border = border
	return c
}

func (c *ChildWidget) Size(width, height float32) *ChildWidget {
	scale := Context.platform.GetContentScale()
	c.width, c.height = width*scale, height*scale
	return c
}

func (c *ChildWidget) Flags(flags WindowFlags) *ChildWidget {
	c.flags = flags
	return c
}

func (c *ChildWidget) Layout(layout Layout) *ChildWidget {
	c.layout = layout
	return c
}

func Child(id string) *ChildWidget {
	return &ChildWidget{
		id:     id,
		width:  0,
		height: 0,
		border: true,
		flags:  0,
		layout: nil,
	}
}

type ComboCustomWidget struct {
	label        string
	previewValue string
	width        float32
	flags        ComboFlags
	layout       Layout
}

func ComboCustom(label, previewValue string) *ComboCustomWidget {
	return &ComboCustomWidget{
		label:        label,
		previewValue: previewValue,
		width:        0,
		flags:        0,
		layout:       nil,
	}
}

func (cc *ComboCustomWidget) Layout(layout Layout) *ComboCustomWidget {
	cc.layout = layout
	return cc
}

func (cc *ComboCustomWidget) Flags(flags ComboFlags) *ComboCustomWidget {
	cc.flags = flags
	return cc
}

func (cc *ComboCustomWidget) Size(width float32) *ComboCustomWidget {
	cc.width = width * Context.platform.GetContentScale()
	return cc
}

func (cc *ComboCustomWidget) Build() {
	if cc.width > 0 {
		imgui.PushItemWidth(cc.width)
	}

	if imgui.BeginComboV(cc.label, cc.previewValue, int(cc.flags)) {
		if cc.layout != nil {
			cc.layout.Build()
		}
		imgui.EndCombo()
	}

	if cc.width > 0 {
		imgui.PopItemWidth()
	}
}

type ComboWidget struct {
	label        string
	previewValue string
	items        []string
	selected     *int32
	width        float32
	flags        ComboFlags
	onChange     func()
}

func Combo(label, previewValue string, items []string, selected *int32) *ComboWidget {
	return &ComboWidget{
		label:        label,
		previewValue: previewValue,
		items:        items,
		selected:     selected,
		flags:        0,
		width:        0,
		onChange:     nil,
	}
}

func (c *ComboWidget) Flags(flags ComboFlags) *ComboWidget {
	c.flags = flags
	return c
}

func (c *ComboWidget) Build() {
	if c.width > 0 {
		imgui.PushItemWidth(c.width)
	}

	if imgui.BeginComboV(c.label, c.previewValue, int(c.flags)) {
		for i, item := range c.items {
			if imgui.Selectable(item) {
				*c.selected = int32(i)
				if c.onChange != nil {
					c.onChange()
				}
			}
		}

		imgui.EndCombo()
	}

	if c.width > 0 {
		imgui.PopItemWidth()
	}
}

func (c *ComboWidget) Size(width float32) *ComboWidget {
	c.width = width * Context.platform.GetContentScale()
	return c
}

func (c *ComboWidget) OnChange(onChange func()) *ComboWidget {
	c.onChange = onChange
	return c
}

type ContextMenuWidget struct {
	label       string
	mouseButton MouseButton
	layout      Layout
}

func ContextMenu(label string) *ContextMenuWidget {
	return &ContextMenuWidget{
		label:       label,
		mouseButton: MouseButtonRight,
		layout:      nil,
	}
}

func (c *ContextMenuWidget) Layout(layout Layout) *ContextMenuWidget {
	c.layout = layout
	return c
}

func (c *ContextMenuWidget) MouseButton(mouseButton MouseButton) *ContextMenuWidget {
	c.mouseButton = mouseButton
	return c
}

func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.label, imgui.PopupFlags(c.mouseButton)) {
		if c.layout != nil {
			c.layout.Build()
		}
		imgui.EndPopup()
	}
}

type DragIntWidget struct {
	label  string
	value  *int32
	speed  float32
	min    int32
	max    int32
	format string
}

func DragInt(label string, value *int32, min, max int32) *DragIntWidget {
	return &DragIntWidget{
		label:  label,
		value:  value,
		speed:  1.0,
		min:    min,
		max:    max,
		format: "%d",
	}
}

func (d *DragIntWidget) Speed(speed float32) *DragIntWidget {
	d.speed = speed
	return d
}

func (d *DragIntWidget) Format(format string) *DragIntWidget {
	d.format = format
	return d
}

func (d *DragIntWidget) Build() {
	imgui.DragIntV(d.label, d.value, d.speed, d.min, d.max, d.format, imgui.SlidersFlagsNone)
}

type GroupWidget struct {
	layout Layout
}

func Group() *GroupWidget {
	return &GroupWidget{
		layout: nil,
	}
}

func (g *GroupWidget) Layout(layout Layout) *GroupWidget {
	g.layout = layout
	return g
}

func (g *GroupWidget) Build() {
	imgui.BeginGroup()
	if g.layout != nil {
		g.layout.Build()
	}
	imgui.EndGroup()
}

type ImageWidget struct {
	texture                *Texture
	width                  float32
	height                 float32
	uv0, uv1               image.Point
	tintColor, borderColor color.RGBA
}

func Image(texture *Texture) *ImageWidget {
	return &ImageWidget{
		texture:     texture,
		width:       100 * Context.platform.GetContentScale(),
		height:      100 * Context.platform.GetContentScale(),
		uv0:         image.Point{X: 0, Y: 0},
		uv1:         image.Point{X: 1, Y: 1},
		tintColor:   color.RGBA{255, 255, 255, 255},
		borderColor: color.RGBA{0, 0, 0, 0},
	}
}

func (i *ImageWidget) Uv(uv0, uv1 image.Point) *ImageWidget {
	i.uv0, i.uv1 = uv0, uv1
	return i
}

func (i *ImageWidget) TintColor(tintColor color.RGBA) *ImageWidget {
	i.tintColor = tintColor
	return i
}

func (i *ImageWidget) BorderCol(borderColor color.RGBA) *ImageWidget {
	i.borderColor = borderColor
	return i
}

func (i *ImageWidget) Size(width, height float32) *ImageWidget {
	scale := Context.platform.GetContentScale()
	i.width, i.height = width*scale, height*scale
	return i
}

func (i *ImageWidget) Build() {
	size := imgui.Vec2{X: i.width, Y: i.height}
	rect := imgui.ContentRegionAvail()
	if size.X == (-1 * Context.GetPlatform().GetContentScale()) {
		size.X = rect.X
	}
	if size.Y == (-1 * Context.GetPlatform().GetContentScale()) {
		size.Y = rect.Y
	}
	if i.texture != nil && i.texture.id != 0 {
		imgui.ImageV(i.texture.id, size, ToVec2(i.uv0), ToVec2(i.uv1), ToVec4Color(i.tintColor), ToVec4Color(i.borderColor))
	} else {
		Dummy(i.width, i.height).Build()
	}
}

type ImageState struct {
	loading bool
	failure bool
	texture *Texture
}

func (is *ImageState) Dispose() {
	is.texture = nil
}

type ImageWithFileWidget struct {
	imgPath string
	width   float32
	height  float32
}

func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		imgPath: imgPath,
		width:   100,
		height:  100,
	}
}

func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithFileWidget) Build() {
	stateId := fmt.Sprintf("ImageWithFile_%s", i.imgPath)
	state := Context.GetState(stateId)

	var widget *ImageWidget

	if state == nil {
		widget = Image(nil).Size(i.width, i.height)

		//Prevent multiple invocation to LoadImage.
		Context.SetState(stateId, &ImageState{})

		img, err := LoadImage(i.imgPath)
		if err == nil {
			go func() {
				texture, err := NewTextureFromRgba(img)
				if err == nil {
					Context.SetState(stateId, &ImageState{texture: texture})
				}
			}()
		}
	} else {
		imgState := state.(*ImageState)
		widget = Image(imgState.texture).Size(i.width, i.height)
	}

	widget.Build()
}

type ImageWithUrlWidget struct {
	imgUrl          string
	downloadTimeout time.Duration
	width           float32
	height          float32
	whenLoading     Layout
	whenFailure     Layout
}

func ImageWithUrl(url string) *ImageWithUrlWidget {
	return &ImageWithUrlWidget{
		imgUrl:          url,
		downloadTimeout: 10 * time.Second,
		width:           100,
		height:          100,
		whenLoading:     Layout{Dummy(100, 100)},
		whenFailure:     Layout{Dummy(100, 100)},
	}
}

func (i *ImageWithUrlWidget) Timeout(downloadTimeout time.Duration) *ImageWithUrlWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

func (i *ImageWithUrlWidget) Size(width, height float32) *ImageWithUrlWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithUrlWidget) LayoutForLoading(layout Layout) *ImageWithUrlWidget {
	i.whenLoading = layout
	return i
}

func (i *ImageWithUrlWidget) LayoutForFailure(layout Layout) *ImageWithUrlWidget {
	i.whenFailure = layout
	return i
}

func (i *ImageWithUrlWidget) Build() {
	stateId := fmt.Sprintf("ImageWithUrl_%s", i.imgUrl)
	state := Context.GetState(stateId)

	var widget *ImageWidget

	if state == nil {
		widget = Image(nil).Size(i.width, i.height)

		//Prevent multiple invocation to download image.
		Context.SetState(stateId, &ImageState{loading: true})

		go func() {
			// Load image from url
			client := resty.New()
			client.SetTimeout(i.downloadTimeout)

			resp, err := client.R().Get(i.imgUrl)
			Context.SetState(stateId, &ImageState{loading: false})
			if err != nil {
				Context.SetState(stateId, &ImageState{failure: true})
				return
			}

			img, _, err := image.Decode(bytes.NewReader(resp.Body()))
			if err != nil {
				Context.SetState(stateId, &ImageState{failure: true})
				return
			}

			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)

			texture, err := NewTextureFromRgba(rgba)
			if err != nil {
				Context.SetState(stateId, &ImageState{failure: true})
				return
			}
			Context.SetState(stateId, &ImageState{loading: false, texture: texture})
		}()
	} else {
		imgState := state.(*ImageState)
		if imgState.failure {
			i.whenFailure.Build()
			return
		}

		if imgState.loading {
			i.whenLoading.Build()
			return
		}

		widget = Image(imgState.texture).Size(i.width, i.height)
	}

	widget.Build()
}

type InputTextWidget struct {
	label    string
	value    *string
	width    float32
	flags    InputTextFlags
	cb       imgui.InputTextCallback
	onChange func()
}

func InputText(label string, value *string) *InputTextWidget {
	return &InputTextWidget{
		label:    label,
		value:    value,
		width:    0,
		flags:    0,
		cb:       nil,
		onChange: nil,
	}
}

func (i *InputTextWidget) Size(width float32) *InputTextWidget {
	i.width = width * Context.platform.GetContentScale()
	return i
}

func (i *InputTextWidget) Flags(flags InputTextFlags) *InputTextWidget {
	i.flags = flags
	return i
}

func (i *InputTextWidget) Callback(cb imgui.InputTextCallback) *InputTextWidget {
	i.cb = cb
	return i
}

func (i *InputTextWidget) OnChange(onChange func()) *InputTextWidget {
	i.onChange = onChange
	return i
}

func (i *InputTextWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)
	}

	if imgui.InputTextV(i.label, i.value, int(i.flags), i.cb) && i.onChange != nil {
		i.onChange()
	}

	if i.width != 0 {
		PopItemWidth()
	}
}

type InputIntWidget struct {
	label    string
	value    *int32
	width    float32
	flags    InputTextFlags
	onChange func()
}

func InputInt(label string, value *int32) *InputIntWidget {
	return &InputIntWidget{
		label:    label,
		value:    value,
		width:    0,
		flags:    0,
		onChange: nil,
	}
}

func (i *InputIntWidget) Size(width float32) *InputIntWidget {
	i.width = width * Context.platform.GetContentScale()
	return i
}

func (i *InputIntWidget) Flags(flags InputTextFlags) *InputIntWidget {
	i.flags = flags
	return i
}

func (i *InputIntWidget) OnChange(onChange func()) *InputIntWidget {
	i.onChange = onChange
	return i
}

func (i *InputIntWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)
	}

	if imgui.InputIntV(i.label, i.value, 0, 100, int(i.flags)) && i.onChange != nil {
		i.onChange()
	}

	if i.width != 0 {
		PopItemWidth()
	}
}

type InputFloatWidget struct {
	label    string
	value    *float32
	width    float32
	flags    InputTextFlags
	format   string
	onChange func()
}

func InputFloatV(label string, value *float32) *InputFloatWidget {
	return &InputFloatWidget{
		label:    label,
		width:    0,
		value:    value,
		format:   "%.3f",
		flags:    0,
		onChange: nil,
	}
}

func (i *InputFloatWidget) Size(width float32) *InputFloatWidget {
	i.width = width * Context.platform.GetContentScale()
	return i
}

func (i *InputFloatWidget) Flags(flags InputTextFlags) *InputFloatWidget {
	i.flags = flags
	return i
}

func (i *InputFloatWidget) Format(format string) *InputFloatWidget {
	i.format = format
	return i
}

func (i *InputFloatWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)
	}

	if imgui.InputFloatV(i.label, i.value, 0, 0, i.format, int(i.flags)) && i.onChange != nil {
		i.onChange()
	}

	if i.width != 0 {
		PopItemWidth()
	}
}

type LabelWidget struct {
	label   string
	wrapped bool
	color   *color.RGBA
	font    *imgui.Font
}

func Label(label string) *LabelWidget {
	return &LabelWidget{
		label:   label,
		wrapped: false,
		color:   nil,
		font:    nil,
	}
}

func (l *LabelWidget) Wrapped(wrapped bool) *LabelWidget {
	l.wrapped = wrapped
	return l
}

func (l *LabelWidget) Color(color *color.RGBA) *LabelWidget {
	l.color = color
	return l
}

func (l *LabelWidget) Font(font *imgui.Font) *LabelWidget {
	l.font = font
	return l
}

func (l *LabelWidget) Build() {
	if l.color != nil {
		PushColorText(*l.color)
	}

	if l.font != nil {
		PushFont(*l.font)
	}

	if l.wrapped {
		PushTextWrapPos()
	}

	imgui.Text(l.label)

	if l.wrapped {
		PopTextWrapPos()
	}

	if l.font != nil {
		PopFont()
	}

	if l.color != nil {
		PopStyleColor()
	}
}

type MainMenuBarWidget struct {
	layout Layout
}

func MainMenuBar() *MainMenuBarWidget {
	return &MainMenuBarWidget{
		layout: nil,
	}
}

func (m *MainMenuBarWidget) Layout(layout Layout) *MainMenuBarWidget {
	m.layout = layout
	return m
}

func (m *MainMenuBarWidget) Build() {
	if imgui.BeginMainMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMainMenuBar()
	}
}

type MenuBarWidget struct {
	layout Layout
}

func MenuBar() *MenuBarWidget {
	return &MenuBarWidget{
		layout: nil,
	}
}

func (m *MenuBarWidget) Layout(layout Layout) *MenuBarWidget {
	m.layout = layout
	return m
}

func (m *MenuBarWidget) Build() {
	if imgui.BeginMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenuBar()
	}
}

type MenuItemWidget struct {
	label    string
	selected bool
	enabled  bool
	onClick  func()
}

func MenuItem(label string) *MenuItemWidget {
	return &MenuItemWidget{
		label:    label,
		selected: false,
		enabled:  true,
		onClick:  nil,
	}
}

func (m *MenuItemWidget) Selected(s bool) *MenuItemWidget {
	m.selected = s
	return m
}

func (m *MenuItemWidget) Enabled(e bool) *MenuItemWidget {
	m.enabled = e
	return m
}

func (m *MenuItemWidget) OnClick(onClick func()) *MenuItemWidget {
	m.onClick = onClick
	return m
}

func (m *MenuItemWidget) Build() {
	if imgui.MenuItemV(m.label, "", m.selected, m.enabled) && m.onClick != nil {
		m.onClick()
	}
}

type MenuWidget struct {
	label   string
	enabled bool
	layout  Layout
}

func Menu(label string) *MenuWidget {
	return &MenuWidget{
		label:   label,
		enabled: true,
		layout:  nil,
	}
}

func (m *MenuWidget) Enabled(e bool) *MenuWidget {
	m.enabled = e
	return m
}

func (m *MenuWidget) Layout(layout Layout) *MenuWidget {
	m.layout = layout
	return m
}

func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(m.label, m.enabled) {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenu()
	}
}

type PopupWidget struct {
	name   string
	flags  imgui.PopupFlags
	layout Layout
}

func Popup(name string) *PopupWidget {
	return &PopupWidget{
		name:   name,
		flags:  0,
		layout: nil,
	}
}

func (p *PopupWidget) Flags(flags imgui.PopupFlags) *PopupWidget {
	p.flags = flags
	return p
}

func (p *PopupWidget) Layout(layout Layout) *PopupWidget {
	p.layout = layout
	return p
}

func (p *PopupWidget) Build() {
	if imgui.BeginPopupV(p.name, p.flags) {
		if p.layout != nil {
			Update()
			p.layout.Build()
		}
		imgui.EndPopup()
	}
}

type PopupModalWidget struct {
	name   string
	open   *bool
	flags  imgui.PopupFlags
	layout Layout
}

func PopupModal(name string) *PopupModalWidget {
	return &PopupModalWidget{
		name:   name,
		open:   nil,
		flags:  imgui.PopupFlagsMouseButtonMiddle,
		layout: nil,
	}
}

func (p *PopupModalWidget) IsOpen(open *bool) *PopupModalWidget {
	p.open = open
	return p
}

func (p *PopupModalWidget) Flags(flags imgui.PopupFlags) *PopupModalWidget {
	p.flags = flags
	return p
}

func (p *PopupModalWidget) Layout(layout Layout) *PopupModalWidget {
	p.layout = layout
	return p
}

func (p *PopupModalWidget) Build() {
	if *p.open {
		imgui.OpenPopup(p.name)
	}

	if imgui.BeginPopupModalV(p.name, p.open, p.flags) {
		if p.layout != nil {
			Update()
			p.layout.Build()
		}
		imgui.EndPopup()
	}
}

func OpenPopup(name string) {
	imgui.OpenPopup(name)
}

func CloseCurrentPopup() {
	imgui.CloseCurrentPopup()
}

type ProgressBarWidget struct {
	fraction float32
	width    float32
	height   float32
	overlay  string
}

func ProgressBar(fraction float32) *ProgressBarWidget {
	return &ProgressBarWidget{
		fraction: fraction,
		width:    0,
		height:   0,
		overlay:  "",
	}
}

func (p *ProgressBarWidget) Size(width, height float32) *ProgressBarWidget {
	scale := Context.platform.GetContentScale()
	p.width, p.height = width*scale, height*scale
	return p
}

func (p *ProgressBarWidget) Overlay(overlay string) *ProgressBarWidget {
	p.overlay = overlay
	return p
}

func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, imgui.Vec2{X: p.width, Y: p.height}, p.overlay)
}

type SelectableWidget struct {
	label    string
	selected bool
	flags    SelectableFlags
	width    float32
	height   float32
	onClick  func()
}

func Selectable(label string) *SelectableWidget {
	return &SelectableWidget{
		label:    label,
		selected: false,
		flags:    0,
		width:    0,
		height:   0,
		onClick:  nil,
	}
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
	scale := Context.platform.GetContentScale()
	s.width, s.height = width*scale, height*scale
	return s
}

func (s *SelectableWidget) OnClick(onClick func()) *SelectableWidget {
	s.onClick = onClick
	return s
}

func (s *SelectableWidget) Build() {
	if imgui.SelectableV(s.label, s.selected, int(s.flags), imgui.Vec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}
}

type SeparatorWidget struct{}

func (s *SeparatorWidget) Build() {
	imgui.Separator()
}

func Separator() *SeparatorWidget {
	return &SeparatorWidget{}
}

type SliderIntWidget struct {
	label  string
	value  *int32
	min    int32
	max    int32
	format string
}

func SliderInt(label string, value *int32, min, max int32) *SliderIntWidget {
	return &SliderIntWidget{
		label:  label,
		value:  value,
		min:    min,
		max:    max,
		format: "%d",
	}
}

func (s *SliderIntWidget) Format(format string) *SliderIntWidget {
	s.format = format
	return s
}

func (s *SliderIntWidget) Build() {
	imgui.SliderIntV(s.label, s.value, s.min, s.max, s.format, imgui.SlidersFlagsNone)
}

type SliderFloatWidget struct {
	label  string
	value  *float32
	min    float32
	max    float32
	format string
}

func SliderFloat(label string, value *float32, min, max float32) *SliderFloatWidget {
	return &SliderFloatWidget{
		label:  label,
		value:  value,
		min:    min,
		max:    max,
		format: "%.3f",
	}
}

func (s *SliderFloatWidget) Format(format string) *SliderFloatWidget {
	s.format = format
	return s
}

func (s *SliderFloatWidget) Build() {
	imgui.SliderFloatV(s.label, s.value, s.min, s.max, s.format, 1.0)
}

type DummyWidget struct {
	width  float32
	height float32
}

func (d *DummyWidget) Build() {
	w, h := GetAvaiableRegion()

	if d.width < 0 {
		d.width = w + d.width
	}

	if d.height < 0 {
		d.height = h + d.height
	}

	imgui.Dummy(imgui.Vec2{X: d.width, Y: d.height})
}

func Dummy(width, height float32) *DummyWidget {
	return &DummyWidget{
		width:  width * Context.platform.GetContentScale(),
		height: height * Context.platform.GetContentScale(),
	}
}

type HSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func HSplitter(id string, delta *float32) *HSplitterWidget {

	return &HSplitterWidget{
		id:     id,
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (h *HSplitterWidget) Size(width, height float32) *HSplitterWidget {
	scale := Context.platform.GetContentScale()
	aw, ah := GetAvaiableRegion()

	if width == 0 {
		h.width = aw / scale
	} else {
		h.width = width * scale
	}

	if height == 0 {
		h.height = ah / scale
	} else {
		h.height = height * scale
	}

	return h
}

func (h *HSplitterWidget) Build() {
	// Calc line position.
	width := int(40 * Context.GetPlatform().GetContentScale())
	height := int(2 * Context.GetPlatform().GetContentScale())

	pt := GetCursorScreenPos()

	centerX := int(h.width / 2)
	centerY := int(h.height / 2)

	ptMin := image.Pt(centerX-width/2, centerY-height/2)
	ptMax := image.Pt(centerX+width/2, centerY+height/2)

	style := imgui.CurrentStyle()
	color := Vec4ToRGBA(style.Color(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButtonV(h.id, imgui.Vec2{X: h.width, Y: h.height}, imgui.ButtonFlagsNone)
	if imgui.IsItemActive() {
		*(h.delta) = imgui.CurrentIO().MouseDelta().Y / Context.platform.GetContentScale()
	} else {
		*(h.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeNS)
		color = Vec4ToRGBA(style.Color(imgui.StyleColorScrollbarGrabActive))
	}

	// Draw a line in the very center
	canvas := GetCanvas()
	canvas.AddRectFilled(pt.Add(ptMin), pt.Add(ptMax), color, 0, 0)
}

type VSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func VSplitter(id string, delta *float32) *VSplitterWidget {
	return &VSplitterWidget{
		id:     id,
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (v *VSplitterWidget) Size(width, height float32) *VSplitterWidget {
	aw, ah := GetAvaiableRegion()
	scale := Context.platform.GetContentScale()

	if width == 0 {
		v.width = aw / scale
	} else {
		v.width = width * scale
	}

	if height == 0 {
		v.height = ah / scale
	} else {
		v.height = height * scale
	}

	return v
}

func (v *VSplitterWidget) Build() {
	// Calc line position.
	width := int(2 * Context.GetPlatform().GetContentScale())
	height := int(40 * Context.GetPlatform().GetContentScale())

	pt := GetCursorScreenPos()

	centerX := int(v.width / 2)
	centerY := int(v.height / 2)

	ptMin := image.Pt(centerX-width/2, centerY-height/2)
	ptMax := image.Pt(centerX+width/2, centerY+height/2)

	style := imgui.CurrentStyle()
	color := Vec4ToRGBA(style.Color(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButtonV(v.id, imgui.Vec2{X: float32(width), Y: float32(height)}, imgui.ButtonFlagsNone)
	if imgui.IsItemActive() {
		*(v.delta) = imgui.CurrentIO().MouseDelta().X / Context.platform.GetContentScale()
	} else {
		*(v.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeEW)
		color = Vec4ToRGBA(style.Color(imgui.StyleColorScrollbarGrabActive))
	}

	// Draw a line in the very center
	canvas := GetCanvas()
	canvas.AddRectFilled(pt.Add(ptMin), pt.Add(ptMax), color, 0, 0)
}

type TabItemWidget struct {
	label  string
	open   *bool
	flags  TabItemFlags
	layout Layout
}

func TabItem(label string) *TabItemWidget {
	return &TabItemWidget{
		label:  label,
		open:   nil,
		flags:  0,
		layout: nil,
	}
}

func (t *TabItemWidget) IsOpen(open *bool) *TabItemWidget {
	t.open = open
	return t
}

func (t *TabItemWidget) Flags(flags TabItemFlags) *TabItemWidget {
	t.flags = flags
	return t
}

func (t *TabItemWidget) Layout(layout Layout) *TabItemWidget {
	t.layout = layout
	return t
}

func (t *TabItemWidget) Build() {
	if imgui.BeginTabItemV(t.label, t.open, int(t.flags)) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabItem()
	}
}

type TabBarWidget struct {
	id     string
	flags  TabBarFlags
	layout Layout
}

func TabBar(id string) *TabBarWidget {
	return &TabBarWidget{
		id:     id,
		flags:  0,
		layout: nil,
	}
}

func (t *TabBarWidget) Flags(flags TabBarFlags) *TabBarWidget {
	t.flags = flags
	return t
}

func (t *TabBarWidget) Layout(layout Layout) *TabBarWidget {
	t.layout = layout
	return t
}

func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id, int(t.flags)) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabBar()
	}
}

type RowWidget struct {
	layout Layout
}

func Row(widgets ...Widget) *RowWidget {
	return &RowWidget{
		layout: widgets,
	}
}

func (r *RowWidget) Build() {
	for i, w := range r.layout {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopup := w.(*PopupModalWidget)

		if i > 0 && !isTooltip && !isContextMenu && !isPopup {
			imgui.NextColumn()
		}
		w.Build()
	}
}

type Rows []*RowWidget

type TabelWidget struct {
	label  string
	border bool
	rows   Rows
}

func Table(label string) *TabelWidget {
	return &TabelWidget{
		label:  label,
		border: true,
		rows:   nil,
	}
}

func (t *TabelWidget) Border(b bool) *TabelWidget {
	t.border = b
	return t
}

func (t *TabelWidget) Rows(rows Rows) *TabelWidget {
	t.rows = rows
	return t
}

func (t *TabelWidget) Build() {
	if len(t.rows) > 0 && len(t.rows[0].layout) > 0 {
		imgui.ColumnsV(len(t.rows[0].layout), t.label, t.border)

		for i, r := range t.rows {
			if i > 0 {
				imgui.NextColumn()
			}

			if t.border {
				imgui.Separator()
			}

			r.Build()
		}

		imgui.Columns()

		if t.border {
			imgui.Separator()
		}
	}
}

type FastTabelWidget struct {
	label  string
	border bool
	rows   Rows
}

// Create a fast table which only render visible rows.
// Note this only works with all rows have same height.
func FastTable(label string) *FastTabelWidget {
	return &FastTabelWidget{
		label:  label,
		border: true,
		rows:   nil,
	}
}

func (t *FastTabelWidget) Border(b bool) *FastTabelWidget {
	t.border = b
	return t
}

func (t *FastTabelWidget) Rows(rows Rows) *FastTabelWidget {
	t.rows = rows
	return t
}

func (t *FastTabelWidget) Build() {
	if len(t.rows) > 0 && len(t.rows[0].layout) > 0 {
		imgui.ColumnsV(len(t.rows[0].layout), t.label, t.border)

		var clipper imgui.ListClipper
		clipper.Begin(len(t.rows))

		for clipper.Step() {
			for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++ {
				r := t.rows[i]

				if i > 0 {
					imgui.NextColumn()
				}

				if t.border {
					imgui.Separator()
				}

				r.Build()
			}
		}

		clipper.End()

		imgui.Columns()

		if t.border {
			imgui.Separator()
		}
	}
}

type TooltipWidget struct {
	tip    string
	layout Layout
}

func (t *TooltipWidget) Build() {
	if imgui.IsItemHovered() {
		if t.layout != nil {
			imgui.BeginTooltip()
			t.layout.Build()
			imgui.EndTooltip()
		} else {
			imgui.SetTooltip(t.tip)
		}
	}
}

func Tooltip(tip string) *TooltipWidget {
	return &TooltipWidget{
		tip:    tip,
		layout: nil,
	}
}

func (t *TooltipWidget) Layout(layout Layout) *TooltipWidget {
	t.layout = layout
	return t
}

type TreeNodeWidget struct {
	label        string
	flags        TreeNodeFlags
	layout       Layout
	eventHandler func()
}

func TreeNode(label string) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:        label,
		flags:        0,
		layout:       nil,
		eventHandler: nil,
	}
}

func (t *TreeNodeWidget) Flags(flags TreeNodeFlags) *TreeNodeWidget {
	t.flags = flags
	return t
}

// Create TreeNode with eventHandler
// You could detect events (e.g. IsItemClicked IsMouseDoubleClicked etc...) and handle them for TreeNode inside eventHandler
func (t *TreeNodeWidget) Event(handler func()) *TreeNodeWidget {
	t.eventHandler = handler
	return t
}

func (t *TreeNodeWidget) Layout(layout Layout) *TreeNodeWidget {
	t.layout = layout
	return t
}

func (t *TreeNodeWidget) Build() {
	open := imgui.TreeNodeV(t.label, int(t.flags))

	if t.eventHandler != nil {
		t.eventHandler()
	}

	if open {
		if t.layout != nil {
			t.layout.Build()
		}
		if (t.flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}

type SpacingWidget struct{}

func (s *SpacingWidget) Build() {
	imgui.Spacing()
}

func Spacing() *SpacingWidget {
	return &SpacingWidget{}
}

type CustomWidget struct {
	builder func()
}

func (c *CustomWidget) Build() {
	if c.builder != nil {
		c.builder()
	}
}

func Custom(builder func()) *CustomWidget {
	return &CustomWidget{
		builder: builder,
	}
}

type ConditionWidget struct {
	cond       bool
	layoutIf   Layout
	layoutElse Layout
}

func Condition(cond bool, layoutIf Layout, layoutElse Layout) *ConditionWidget {
	return &ConditionWidget{
		cond:       cond,
		layoutIf:   layoutIf,
		layoutElse: layoutElse,
	}
}

func (c *ConditionWidget) Build() {
	if c.cond {
		if c.layoutIf != nil {
			c.layoutIf.Build()
		}
	} else {
		if c.layoutElse != nil {
			c.layoutElse.Build()
		}
	}
}

// Batch create widgets and render only which is visible.
func RangeBuilder(id string, values []interface{}, builder func(int, interface{}) Widget) Layout {
	var layout Layout

	layout = append(layout, Custom(func() { imgui.PushID(id) }))

	if len(values) > 0 && builder != nil {
		for i, v := range values {
			valueRef := v
			widget := builder(i, valueRef)
			layout = append(layout, widget)
		}
	}

	layout = append(layout, Custom(func() { imgui.PopID() }))

	return layout
}

type ListBoxState struct {
	selectedIndex int
}

func (s *ListBoxState) Dispose() {
	// Nothing to do here.
}

type ListBoxWidget struct {
	id       string
	width    float32
	height   float32
	border   bool
	items    []string
	menus    []string
	onChange func(selectedIndex int)
	onDClick func(selectedIndex int)
	onMenu   func(selectedIndex int, menu string)
}

func ListBox(id string, items []string) *ListBoxWidget {
	return &ListBoxWidget{
		id:       id,
		width:    0,
		height:   0,
		border:   true,
		items:    items,
		menus:    nil,
		onChange: nil,
		onDClick: nil,
		onMenu:   nil,
	}
}

func (l *ListBoxWidget) Size(width, height float32) *ListBoxWidget {
	scale := Context.platform.GetContentScale()
	l.width, l.height = width*scale, height*scale
	return l
}

func (l *ListBoxWidget) Border(b bool) *ListBoxWidget {
	l.border = b
	return l
}

func (l *ListBoxWidget) ContextMenu(menuItems []string) *ListBoxWidget {
	l.menus = menuItems
	return l
}

func (l *ListBoxWidget) OnChange(onChange func(selectedIndex int)) *ListBoxWidget {
	l.onChange = onChange
	return l
}

func (l *ListBoxWidget) OnDClick(onDClick func(selectedIndex int)) *ListBoxWidget {
	l.onDClick = onDClick
	return l
}

func (l *ListBoxWidget) OnMenu(onMenu func(selectedIndex int, menu string)) *ListBoxWidget {
	l.onMenu = onMenu
	return l
}

func (l *ListBoxWidget) Build() {
	var state *ListBoxState
	if s := Context.GetState(l.id); s == nil {
		state = &ListBoxState{selectedIndex: 0}
		Context.SetState(l.id, state)
	} else {
		state = s.(*ListBoxState)
	}

	child := Child(l.id).Border(l.border).Size(l.width, l.height).Layout(Layout{
		Custom(func() {
			var clipper imgui.ListClipper
			clipper.Begin(len(l.items))

			for clipper.Step() {
				for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++ {
					selected := i == state.selectedIndex
					item := l.items[i]
					Selectable(item).Selected(selected).Flags(SelectableFlagsAllowDoubleClick).OnClick(func() {
						if state.selectedIndex != i {
							state.selectedIndex = i
							if l.onChange != nil {
								l.onChange(i)
							}
						}
					}).Build()

					if IsItemHovered() && IsMouseDoubleClicked(MouseButtonLeft) && l.onDClick != nil {
						l.onDClick(state.selectedIndex)
					}

					// Build context menus
					var menus Layout
					for _, m := range l.menus {
						index := i
						menu := m
						menus = append(menus, MenuItem(fmt.Sprintf("%s##%d", menu, index)).OnClick(func() {
							if l.onMenu != nil {
								l.onMenu(index, menu)
							}
						}))
					}

					if len(menus) > 0 {
						ContextMenu(fmt.Sprintf("%d_contextmenu", i)).Layout(menus).Build()
					}
				}
			}

			clipper.End()
		}),
	})

	child.Build()
}

type DatePickerWidget struct {
	id       string
	date     *time.Time
	width    float32
	onChange func()
}

func DatePicker(id string, date *time.Time) *DatePickerWidget {
	return &DatePickerWidget{
		id:       id,
		date:     date,
		width:    100 * Context.GetPlatform().GetContentScale(),
		onChange: nil,
	}
}

func (d *DatePickerWidget) Size(width float32) *DatePickerWidget {
	d.width = width * Context.platform.GetContentScale()
	return d
}

func (d *DatePickerWidget) OnChange(onChange func()) *DatePickerWidget {
	d.onChange = onChange
	return d
}

func (d *DatePickerWidget) Build() {
	if d.date != nil {
		imgui.PushID(d.id)

		if d.width > 0 {
			PushItemWidth(d.width)
		}

		evtTrigger := func() {
			if d.onChange != nil {
				d.onChange()
			}
		}

		if imgui.BeginComboV(d.id, d.date.Format("2006-01-02"), imgui.ComboFlagHeightLargest) {
			// Build year widget
			imgui.AlignTextToFramePadding()
			imgui.Text(" Year")
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("%14d", d.date.Year()))
			imgui.SameLine()
			if imgui.Button("-##year") {
				*d.date = d.date.AddDate(-1, 0, 0)
				evtTrigger()
			}
			imgui.SameLine()
			if imgui.Button("+##year") {
				*d.date = d.date.AddDate(1, 0, 0)
				evtTrigger()
			}

			// Build month widgets
			imgui.Text("Month")
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("%10s(%02d)", d.date.Month().String(), d.date.Month()))
			imgui.SameLine()
			if imgui.Button("-##month") {
				*d.date = d.date.AddDate(0, -1, 0)
				evtTrigger()
			}
			imgui.SameLine()
			if imgui.Button("+##month") {
				*d.date = d.date.AddDate(0, 1, 0)
				evtTrigger()
			}

			// Build day widgets
			firstDay := time.Date(d.date.Year(), d.date.Month(), 1, 0, 0, 0, 0, time.Local)
			lastDay := firstDay.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

			var days [][]int

			// Build first row
			days = append(days, []int{})
			j := 1
			for i := 0; i < 7; i++ {
				if i < int(firstDay.Weekday()) {
					days[0] = append(days[0], 0)
				} else {
					days[0] = append(days[0], j)
					j += 1
				}
			}

			// Build rest rows
			for ; j <= lastDay.Day(); j++ {
				if len(days[len(days)-1]) == 7 {
					days = append(days, []int{})
				}

				days[len(days)-1] = append(days[len(days)-1], j)
			}

			// Pad last row
			lastRowLen := len(days[len(days)-1])
			if lastRowLen < 7 {
				for i := lastRowLen; i < 7; i++ {
					days[len(days)-1] = append(days[len(days)-1], 0)
				}
			}

			// Build day widgets
			var rows Rows

			// Build week names
			rows = append(rows, Row(
				Label("S"),
				Label("M"),
				Label("T"),
				Label("W"),
				Label("T"),
				Label("F"),
				Label("S"),
			))

			today := time.Now()
			style := imgui.CurrentStyle()
			highlightColor := style.Color(imgui.StyleColorPlotHistogram)
			for r := 0; r < len(days); r++ {
				var row []Widget

				for c := 0; c < 7; c++ {
					day := days[r][c]
					if day == 0 {
						row = append(row, Label(" "))
					} else {
						row = append(row,
							Custom(func() {
								if d.date.Year() == today.Year() && d.date.Month() == today.Month() && day == today.Day() {
									imgui.PushStyleColor(imgui.StyleColorText, highlightColor)
								}

								Selectable(fmt.Sprintf("%02d", day)).Selected(day == int(d.date.Day())).OnClick(func() {
									*d.date, _ = time.ParseInLocation(
										"2006-01-02",
										fmt.Sprintf("%d-%02d-%02d",
											d.date.Year(),
											d.date.Month(),
											day,
										),
										time.Local,
									)

									evtTrigger()
								}).Build()

								if d.date.Year() == today.Year() && d.date.Month() == today.Month() && day == today.Day() {
									imgui.PopStyleColor()
								}
							}),
						)
					}
				}

				rows = append(rows, Row(row...))
			}

			Table("DayTable").Rows(rows).Build()

			imgui.EndCombo()
		}

		if d.width > 0 {
			PopItemWidth()
		}

		imgui.PopID()
	}
}
