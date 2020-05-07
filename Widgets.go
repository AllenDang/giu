package giu

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"time"

	"github.com/AllenDang/giu/imgui"
	resty "github.com/go-resty/resty/v2"
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
		_, isPopup := w.(*PopupWidget)
		_, isTabItem := w.(*TabItemWidget)
		_, isLabel := w.(*LabelWidget)
		_, isCustom := w.(*CustomWidget)

		if isLabel {
			AlignTextToFramePadding()
		}

		if index > 0 && !isTooltip && !isContextMenu && !isPopup && !isTabItem && !isCustom {
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

func InputTextMultiline(label string, text *string, width, height float32, flags InputTextFlags, cb imgui.InputTextCallback, onChange func()) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		label:    label,
		text:     text,
		width:    width * Context.platform.GetContentScale(),
		height:   height * Context.platform.GetContentScale(),
		flags:    flags,
		cb:       cb,
		onChange: onChange,
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

func Button(id string, onClick func()) *ButtonWidget {
	return ButtonV(id, 0, 0, onClick)
}

func ButtonV(id string, width, height float32, onClick func()) *ButtonWidget {
	return &ButtonWidget{
		id:      id,
		width:   width * Context.platform.GetContentScale(),
		height:  height * Context.platform.GetContentScale(),
		onClick: onClick,
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
	return PlotLinesV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, imgui.Vec2{})
}

func PlotLinesV(label string, values []float32, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize imgui.Vec2) *PlotLinesWidget {
	return &PlotLinesWidget{
		label:        label,
		values:       values,
		valuesOffset: valuesOffset,
		overlayText:  overlayText,
		scaleMin:     scaleMin,
		scaleMax:     scaleMax,
		graphSize:    graphSize,
	}
}

type InvisibleButtonWidget struct {
	id      string
	width   float32
	height  float32
	onClick func()
}

func InvisibleButton(id string, width, height float32, onClick func()) *InvisibleButtonWidget {
	return &InvisibleButtonWidget{
		id:      id,
		width:   width * Context.platform.GetContentScale(),
		height:  height * Context.platform.GetContentScale(),
		onClick: onClick,
	}
}

func (ib *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButton(ib.id, imgui.Vec2{X: ib.width, Y: ib.height}) && ib.onClick != nil {
		ib.onClick()
	}
}

type ImageButtonWidget struct {
	texture *Texture
	width   float32
	height  float32
	onClick func()
}

func (i *ImageButtonWidget) Build() {
	if i.texture != nil && i.texture.id != 0 {
		if imgui.ImageButton(i.texture.id, imgui.Vec2{X: i.width, Y: i.height}) && i.onClick != nil {
			i.onClick()
		}
	}
}

func ImageButton(texture *Texture, width, height float32, onClick func()) *ImageButtonWidget {
	return &ImageButtonWidget{
		texture: texture,
		width:   width * Context.platform.GetContentScale(),
		height:  height * Context.platform.GetContentScale(),
		onClick: onClick,
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

func Checkbox(text string, selected *bool, onChange func()) *CheckboxWidget {
	return &CheckboxWidget{
		text:     text,
		selected: selected,
		onChange: onChange,
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

func RadioButton(text string, active bool, onChange func()) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:     text,
		active:   active,
		onChange: onChange,
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

func Child(id string, border bool, width, height float32, flags WindowFlags, layout Layout) *ChildWidget {
	return &ChildWidget{
		id:     id,
		width:  width * Context.platform.GetContentScale(),
		height: height * Context.platform.GetContentScale(),
		border: border,
		flags:  flags,
		layout: layout,
	}
}

type ComboCustomWidget struct {
	label        string
	previewValue string
	width        float32
	flags        ComboFlags
	layout       Layout
}

func ComboCustom(label, previewValue string, width float32, flags ComboFlags, layout Layout) *ComboCustomWidget {
	return &ComboCustomWidget{
		label:        label,
		previewValue: previewValue,
		width:        width * Context.GetPlatform().GetContentScale(),
		flags:        flags,
		layout:       layout,
	}
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

func Combo(label, previewValue string, items []string, selected *int32, width float32, flags ComboFlags, onChange func()) *ComboWidget {
	return &ComboWidget{
		label:        label,
		previewValue: previewValue,
		items:        items,
		selected:     selected,
		flags:        flags,
		width:        width * Context.platform.GetContentScale(),
		onChange:     onChange,
	}
}

type ContextMenuWidget struct {
	label       string
	mouseButton MouseButton
	layout      Layout
}

func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.label, int(c.mouseButton)) {
		if c.layout != nil {
			c.layout.Build()
		}
		imgui.EndPopup()
	}
}

func ContextMenu(layout Layout) *ContextMenuWidget {
	return ContextMenuV("", MouseButtonRight, layout)
}

func ContextMenuV(label string, mouseButton MouseButton, layout Layout) *ContextMenuWidget {
	return &ContextMenuWidget{
		label:       label,
		mouseButton: mouseButton,
		layout:      layout,
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

func (d *DragIntWidget) Build() {
	imgui.DragIntV(d.label, d.value, d.speed, d.min, d.max, d.format)
}

func DragInt(label string, value *int32) *DragIntWidget {
	return DragIntV(label, value, 1.0, 0, 0, "%d")
}

func DragIntV(label string, value *int32, speed float32, min, max int32, format string) *DragIntWidget {
	return &DragIntWidget{
		label:  label,
		value:  value,
		speed:  speed,
		min:    min,
		max:    max,
		format: format,
	}
}

type GroupWidget struct {
	layout Layout
}

func (g *GroupWidget) Build() {
	imgui.BeginGroup()
	if g.layout != nil {
		g.layout.Build()
	}
	imgui.EndGroup()
}

func Group(layout Layout) *GroupWidget {
	return &GroupWidget{
		layout: layout,
	}
}

type ImageWidget struct {
	texture *Texture
	width   float32
	height  float32
}

func (i *ImageWidget) Build() {
	size := imgui.Vec2{X: i.width, Y: i.height}
	if i.texture != nil && i.texture.id != 0 {
		rect := imgui.ContentRegionAvail()
		if size.X == (-1 * Context.GetPlatform().GetContentScale()) {
			size.X = rect.X
		}
		if size.Y == (-1 * Context.GetPlatform().GetContentScale()) {
			size.Y = rect.Y
		}
		imgui.Image(i.texture.id, size)
	}
}

func Image(texture *Texture, width, height float32) *ImageWidget {
	return &ImageWidget{
		texture: texture,
		width:   width * Context.platform.GetContentScale(),
		height:  height * Context.platform.GetContentScale(),
	}
}

type ImageState struct {
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

func ImageWithFile(imgPath string, width, height float32) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		imgPath: imgPath,
		width:   width * Context.platform.GetContentScale(),
		height:  height * Context.platform.GetContentScale(),
	}
}

func (i *ImageWithFileWidget) Build() {
	stateId := fmt.Sprintf("ImageWithFile_%s", i.imgPath)
	state := Context.GetState(stateId)

	var widget *ImageWidget

	if state == nil {
		widget = Image(nil, i.width, i.height)

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
		widget = Image(imgState.texture, i.width, i.height)
	}

	widget.Build()
}

type ImageWithUrlWidget struct {
	imgUrl          string
	downloadTimeout time.Duration
	width           float32
	height          float32
}

func ImageWithUrl(url string, downloadTimeout time.Duration, width, height float32) *ImageWithUrlWidget {
	return &ImageWithUrlWidget{
		imgUrl:          url,
		downloadTimeout: downloadTimeout,
		width:           width * Context.platform.GetContentScale(),
		height:          height * Context.platform.GetContentScale(),
	}
}

func (i *ImageWithUrlWidget) Build() {
	stateId := fmt.Sprintf("ImageWithUrl_%s", i.imgUrl)
	state := Context.GetState(stateId)

	var widget *ImageWidget

	if state == nil {
		widget = Image(nil, i.width, i.height)

		go func() {
			// Load image from url
			client := resty.New()
			client.SetTimeout(i.downloadTimeout)

			resp, err := client.R().Get(i.imgUrl)
			if err == nil {
				img, _, err := image.Decode(bytes.NewReader(resp.Body()))
				if err == nil {
					rgba := image.NewRGBA(img.Bounds())
					draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)

					texture, err := NewTextureFromRgba(rgba)
					if err == nil {
						Context.SetState(stateId, &ImageState{texture: texture})
					}
				}
			}

		}()
	} else {
		imgState := state.(*ImageState)
		widget = Image(imgState.texture, i.width, i.height)
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

func InputText(label string, width float32, value *string) *InputTextWidget {
	return InputTextV(label, width, value, 0, nil, nil)
}

func InputTextV(label string, width float32, value *string, flags InputTextFlags, cb imgui.InputTextCallback, onChange func()) *InputTextWidget {
	return &InputTextWidget{
		label:    label,
		value:    value,
		width:    width * Context.platform.GetContentScale(),
		flags:    flags,
		cb:       cb,
		onChange: onChange,
	}
}

type InputIntWidget struct {
	label    string
	value    *int32
	width    float32
	flags    InputTextFlags
	onChange func()
}

func InputInt(label string, width float32, value *int32) *InputIntWidget {
	return InputIntV(label, width, value, 0, nil)
}

func InputIntV(label string, width float32, value *int32, flags InputTextFlags, onChange func()) *InputIntWidget {
	return &InputIntWidget{
		label:    label,
		value:    value,
		width:    width,
		flags:    flags,
		onChange: onChange,
	}
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

func InputFloat(label string, width float32, value *float32) *InputFloatWidget {
	return InputFloatV(label, width, value, "%.3f", 0, nil)
}

func InputFloatV(label string, width float32, value *float32, format string, flags InputTextFlags, onChange func()) *InputFloatWidget {
	return &InputFloatWidget{
		label:    label,
		width:    width,
		value:    value,
		format:   format,
		flags:    flags,
		onChange: onChange,
	}
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

func Label(label string) *LabelWidget {
	return LabelV(label, false, nil, nil)
}

func LabelWrapped(label string) *LabelWidget {
	return LabelV(label, true, nil, nil)
}

func LabelV(label string, wrapped bool, color *color.RGBA, font *imgui.Font) *LabelWidget {
	return &LabelWidget{
		label:   label,
		wrapped: wrapped,
		color:   color,
		font:    font,
	}
}

type MainMenuBarWidget struct {
	layout Layout
}

func (m *MainMenuBarWidget) Build() {
	if imgui.BeginMainMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMainMenuBar()
	}
}

func MainMenuBar(layout Layout) *MainMenuBarWidget {
	return &MainMenuBarWidget{
		layout: layout,
	}
}

type MenuBarWidget struct {
	layout Layout
}

func (m *MenuBarWidget) Build() {
	if imgui.BeginMenuBar() {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenuBar()
	}
}

func MenuBar(layout Layout) *MenuBarWidget {
	return &MenuBarWidget{
		layout: layout,
	}
}

type MenuItemWidget struct {
	label    string
	selected bool
	enabled  bool
	onClick  func()
}

func (m *MenuItemWidget) Build() {
	if imgui.MenuItemV(m.label, "", m.selected, m.enabled) && m.onClick != nil {
		m.onClick()
	}
}

func MenuItem(label string, onClick func()) *MenuItemWidget {
	return MenuItemV(label, false, true, onClick)
}

func MenuItemV(label string, selected, enabled bool, onClick func()) *MenuItemWidget {
	return &MenuItemWidget{
		label:    label,
		selected: selected,
		enabled:  enabled,
		onClick:  onClick,
	}
}

type MenuWidget struct {
	label   string
	enabled bool
	layout  Layout
}

func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(m.label, m.enabled) {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenu()
	}
}

func Menu(label string, layout Layout) *MenuWidget {
	return MenuV(label, true, layout)
}

func MenuV(label string, enabled bool, layout Layout) *MenuWidget {
	return &MenuWidget{
		label:   label,
		enabled: enabled,
		layout:  layout,
	}
}

type PopupWidget struct {
	name   string
	open   *bool
	flags  WindowFlags
	layout Layout
}

func PopupModal(name string, layout Layout) *PopupWidget {
	return PopupModalV(name, nil, WindowFlagsNoResize, layout)
}

func PopupModalV(name string, open *bool, flags WindowFlags, layout Layout) *PopupWidget {
	return &PopupWidget{
		name:   name,
		open:   open,
		flags:  flags,
		layout: layout,
	}
}

func (p *PopupWidget) Build() {
	if imgui.BeginPopupModalV(p.name, p.open, int(p.flags)) {
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

func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, imgui.Vec2{X: p.width, Y: p.height}, p.overlay)
}

func ProgressBar(fraction float32, width, height float32, overlay string) *ProgressBarWidget {
	return &ProgressBarWidget{
		fraction: fraction,
		width:    width * Context.platform.GetContentScale(),
		height:   height * Context.platform.GetContentScale(),
		overlay:  overlay,
	}
}

type SelectableWidget struct {
	label    string
	selected bool
	flags    SelectableFlags
	width    float32
	height   float32
	onClick  func()
}

func (s *SelectableWidget) Build() {
	if imgui.SelectableV(s.label, s.selected, int(s.flags), imgui.Vec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}
}

func Selectable(label string, onClick func()) *SelectableWidget {
	return SelectableV(label, false, 0, 0, 0, onClick)
}

func SelectableV(label string, selected bool, flags SelectableFlags, width, height float32, onClick func()) *SelectableWidget {
	return &SelectableWidget{
		label:    label,
		selected: selected,
		flags:    flags,
		width:    width * Context.platform.GetContentScale(),
		height:   height * Context.platform.GetContentScale(),
		onClick:  onClick,
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

func (s *SliderIntWidget) Build() {
	imgui.SliderIntV(s.label, s.value, s.min, s.max, s.format)
}

func SliderInt(label string, value *int32, min, max int32, format string) *SliderIntWidget {
	return &SliderIntWidget{
		label:  label,
		value:  value,
		min:    min,
		max:    max,
		format: format,
	}
}

type SliderFloatWidget struct {
	label  string
	value  *float32
	min    float32
	max    float32
	format string
}

func SliderFloat(label string, value *float32, min, max float32, format string) *SliderFloatWidget {
	return &SliderFloatWidget{
		label:  label,
		value:  value,
		min:    min,
		max:    max,
		format: format,
	}
}

func (sf *SliderFloatWidget) Build() {
	imgui.SliderFloatV(sf.label, sf.value, sf.min, sf.max, sf.format, 1.0)
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

func HSplitter(id string, width, height float32, delta *float32) *HSplitterWidget {
	aw, ah := GetAvaiableRegion()
	if width == 0 {
		width = aw / Context.GetPlatform().GetContentScale()
	}

	if height == 0 {
		height = ah / Context.GetPlatform().GetContentScale()
	}

	return &HSplitterWidget{
		id:     id,
		width:  width * Context.platform.GetContentScale(),
		height: height * Context.platform.GetContentScale(),
		delta:  delta,
	}
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
	color := Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButton(h.id, imgui.Vec2{X: h.width, Y: h.height})
	if imgui.IsItemActive() {
		*(h.delta) = imgui.CurrentIO().GetMouseDelta().Y / Context.platform.GetContentScale()
	} else {
		*(h.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeNS)
		color = Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrabActive))
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

func VSplitter(id string, width, height float32, delta *float32) *VSplitterWidget {
	aw, ah := GetAvaiableRegion()
	if width == 0 {
		width = aw / Context.GetPlatform().GetContentScale()
	}
	if height == 0 {
		height = ah / Context.GetPlatform().GetContentScale()
	}
	return &VSplitterWidget{
		id:     id,
		width:  width * Context.platform.GetContentScale(),
		height: height * Context.platform.GetContentScale(),
		delta:  delta,
	}
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
	color := Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButton(v.id, imgui.Vec2{X: v.width, Y: v.height})
	if imgui.IsItemActive() {
		*(v.delta) = imgui.CurrentIO().GetMouseDelta().X / Context.platform.GetContentScale()
	} else {
		*(v.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeEW)
		color = Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrabActive))
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

func (t *TabItemWidget) Build() {
	if imgui.BeginTabItemV(t.label, t.open, int(t.flags)) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabItem()
	}
}

func TabItem(label string, layout Layout) *TabItemWidget {
	return TabItemV(label, nil, 0, layout)
}

func TabItemV(label string, open *bool, flags TabItemFlags, layout Layout) *TabItemWidget {
	return &TabItemWidget{
		label:  label,
		open:   open,
		flags:  flags,
		layout: layout,
	}
}

type TabBarWidget struct {
	id     string
	flags  TabBarFlags
	layout Layout
}

func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id, int(t.flags)) {
		if t.layout != nil {
			t.layout.Build()
		}
		imgui.EndTabBar()
	}
}

func TabBar(id string, layout Layout) *TabBarWidget {
	return TabBarV(id, 0, layout)
}

func TabBarV(id string, flags TabBarFlags, layout Layout) *TabBarWidget {
	return &TabBarWidget{
		id:     id,
		flags:  flags,
		layout: layout,
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
		_, isPopup := w.(*PopupWidget)

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

func Table(label string, border bool, rows Rows) *TabelWidget {
	return &TabelWidget{
		label:  label,
		border: border,
		rows:   rows,
	}
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
func FastTable(label string, border bool, rows Rows) *FastTabelWidget {
	return &FastTabelWidget{
		label:  label,
		border: border,
		rows:   rows,
	}
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
	tip string
}

func (t *TooltipWidget) Build() {
	if imgui.IsItemHovered() {
		imgui.SetTooltip(t.tip)
	}
}

func Tooltip(tip string) *TooltipWidget {
	return &TooltipWidget{
		tip: tip,
	}
}

type TreeNodeWidget struct {
	label  string
	flags  TreeNodeFlags
	layout Layout
}

func (t *TreeNodeWidget) Build() {
	if imgui.TreeNodeV(t.label, int(t.flags)) {
		if t.layout != nil {
			t.layout.Build()
		}
		if (t.flags & imgui.TreeNodeFlagsNoTreePushOnOpen) == 0 {
			imgui.TreePop()
		}
	}
}

func TreeNode(label string, flags TreeNodeFlags, layout Layout) *TreeNodeWidget {
	return &TreeNodeWidget{
		label:  label,
		flags:  flags,
		layout: layout,
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

func ListBox(id string, items []string, onChange func(selectedIndex int), onDClick func(selectedIndex int)) *ListBoxWidget {
	return ListBoxV(id, 0, 0, true, items, nil, onChange, onDClick, nil)
}

func ListBoxV(id string, width, height float32, border bool, items []string, menus []string, onChange func(selectedIndex int), onDClick func(selectedIndex int), onMenu func(selectedIndex int, menu string)) *ListBoxWidget {
	return &ListBoxWidget{
		id:       id,
		width:    width,
		height:   height,
		border:   border,
		items:    items,
		menus:    menus,
		onChange: onChange,
		onDClick: onDClick,
		onMenu:   onMenu,
	}
}

func (l *ListBoxWidget) Build() {
	var state *ListBoxState
	if s := Context.GetState(l.id); s == nil {
		state = &ListBoxState{selectedIndex: 0}
		Context.SetState(l.id, state)
	} else {
		state = s.(*ListBoxState)
	}

	child := Child(l.id, l.border, l.width, l.height, 0, Layout{
		Custom(func() {
			var clipper imgui.ListClipper
			clipper.Begin(len(l.items))

			for clipper.Step() {
				for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++ {
					selected := i == state.selectedIndex
					item := l.items[i]
					SelectableV(item, selected, SelectableFlagsAllowDoubleClick, 0, 0, func() {
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
						menus = append(menus, MenuItem(fmt.Sprintf("%s##%d", menu, index), func() {
							if l.onMenu != nil {
								l.onMenu(index, menu)
							}
						}))
					}

					if len(menus) > 0 {
						ContextMenuV(fmt.Sprintf("%d_contextmenu", i), MouseButtonRight, menus).Build()
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

func DatePicker(id string, date *time.Time, width float32, onChange func()) *DatePickerWidget {
	return &DatePickerWidget{
		id:       id,
		date:     date,
		width:    width * Context.GetPlatform().GetContentScale(),
		onChange: onChange,
	}
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
			highlightColor := style.GetColor(imgui.StyleColorPlotHistogram)
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

								SelectableV(fmt.Sprintf("%02d", day), day == int(d.date.Day()), 0, 0, 0, func() {
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

			Table("DayTable", true, rows).Build()

			imgui.EndCombo()
		}

		if d.width > 0 {
			PopItemWidth()
		}

		imgui.PopID()
	}
}
