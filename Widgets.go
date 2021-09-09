package giu

import (
	"bytes"
	ctx "context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"time"

	"github.com/AllenDang/imgui-go"
	resty "github.com/go-resty/resty/v2"
	"github.com/sahilm/fuzzy"
)

func GenAutoID(id string) string {
	return fmt.Sprintf("%s##%d", id, Context.GetWidgetIndex())
}

type RowWidget struct {
	widgets Layout
}

func Row(widgets ...Widget) *RowWidget {
	return &RowWidget{
		widgets: widgets,
	}
}

func (l *RowWidget) Build() {
	index := 0

	for _, w := range l.widgets {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopupModal := w.(*PopupModalWidget)
		_, isPopup := w.(*PopupWidget)
		_, isTabItem := w.(*TabItemWidget)
		_, isLabel := w.(*LabelWidget)

		if isLabel {
			AlignTextToFramePadding()
		}

		if index > 0 && !isTooltip && !isContextMenu && !isPopupModal && !isPopup && !isTabItem {
			imgui.SameLine()
		}

		index += 1

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

func InputTextMultiline(text *string) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		text:     text,
		width:    0,
		height:   0,
		flags:    0,
		cb:       nil,
		onChange: nil,
	}
}

func (i *InputTextMultilineWidget) Label(label string) *InputTextMultilineWidget {
	i.label = tStr(label)
	return i
}

func (i *InputTextMultilineWidget) Build() {
	if len(i.label) == 0 {
		i.label = GenAutoID(i.label)
	}

	if imgui.InputTextMultilineV(i.label, tStrPtr(i.text), imgui.Vec2{X: i.width, Y: i.height}, int(i.flags), i.cb) && i.onChange != nil {
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

type ButtonWidget struct {
	id       string
	width    float32
	height   float32
	disabled bool
	onClick  func()
}

func (b *ButtonWidget) Build() {
	if b.disabled {
		imgui.BeginDisabled(true)
	}

	if imgui.ButtonV(GenAutoID(b.id), imgui.Vec2{X: b.width, Y: b.height}) && b.onClick != nil {
		b.onClick()
	}

	if b.disabled {
		imgui.EndDisabled()
	}
}

func (b *ButtonWidget) OnClick(onClick func()) *ButtonWidget {
	b.onClick = onClick
	return b
}

func (b *ButtonWidget) Disabled(d bool) *ButtonWidget {
	b.disabled = d
	return b
}

func (b *ButtonWidget) Size(width, height float32) *ButtonWidget {
	scale := Context.platform.GetContentScale()
	b.width, b.height = width*scale, height*scale
	return b
}

func Button(id string) *ButtonWidget {
	return &ButtonWidget{
		id:      tStr(id),
		width:   0,
		height:  0,
		onClick: nil,
	}
}

func Buttonf(format string, args ...interface{}) *ButtonWidget {
	return Button(fmt.Sprintf(format, args...))
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
		text: tStr(text),
	}
}

func BulletTextf(format string, args ...interface{}) *BulletTextWidget {
	return BulletTextf(fmt.Sprintf(format, args...))
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

func (ab *ArrowButtonWidget) Build() {
	if imgui.ArrowButton(GenAutoID(ab.id), uint8(ab.dir)) && ab.onClick != nil {
		ab.onClick()
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
		id:      tStr(id),
		onClick: nil,
	}
}

func SmallButtonf(format string, args ...interface{}) *SmallButtonWidget {
	return SmallButton(fmt.Sprintf(format, args...))
}

func (sb *SmallButtonWidget) Build() {
	if imgui.SmallButton(GenAutoID(sb.id)) && sb.onClick != nil {
		sb.onClick()
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

func (ib *InvisibleButtonWidget) Build() {
	if imgui.InvisibleButton(tStr(ib.id), imgui.Vec2{X: ib.width, Y: ib.height}) && ib.onClick != nil {
		ib.onClick()
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

func (i *ImageButtonWidget) Build() {
	if i.texture != nil && i.texture.id != 0 {
		if imgui.ImageButtonV(i.texture.id, imgui.Vec2{X: i.width, Y: i.height}, ToVec2(i.uv0), ToVec2(i.uv1), i.framePadding, ToVec4Color(i.bgColor), ToVec4Color(i.tintColor)) && i.onClick != nil {
			i.onClick()
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

func (b *ImageButtonWithRgbaWidget) BgColor(bgColor color.RGBA) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.BgColor(bgColor)
	return b
}

func (b *ImageButtonWithRgbaWidget) TintColor(tintColor color.RGBA) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.TintColor(tintColor)
	return b
}

func (b *ImageButtonWithRgbaWidget) FramePadding(padding int) *ImageButtonWithRgbaWidget {
	b.ImageButtonWidget.FramePadding(padding)
	return b
}

func (i *ImageButtonWithRgbaWidget) Build() {
	state := Context.GetState(i.id)

	if state == nil {
		Context.SetState(i.id, &ImageState{})

		go func() {
			texture, err := NewTextureFromRgba(i.rgba)
			if err == nil {
				Context.SetState(i.id, &ImageState{texture: texture})
			}
		}()
	} else {
		imgState := state.(*ImageState)
		i.ImageButtonWidget.texture = imgState.texture
	}

	i.ImageButtonWidget.Build()
}

type CheckboxWidget struct {
	text     string
	selected *bool
	onChange func()
}

func (c *CheckboxWidget) Build() {
	if imgui.Checkbox(GenAutoID(c.text), c.selected) && c.onChange != nil {
		c.onChange()
	}
}

func (c *CheckboxWidget) OnChange(onChange func()) *CheckboxWidget {
	c.onChange = onChange
	return c
}

func Checkbox(text string, selected *bool) *CheckboxWidget {
	return &CheckboxWidget{
		text:     tStr(text),
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
	if imgui.RadioButton(GenAutoID(r.text), r.active) && r.onChange != nil {
		r.onChange()
	}
}

func (r *RadioButtonWidget) OnChange(onChange func()) *RadioButtonWidget {
	r.onChange = onChange
	return r
}

func RadioButton(text string, active bool) *RadioButtonWidget {
	return &RadioButtonWidget{
		text:     tStr(text),
		active:   active,
		onChange: nil,
	}
}

type ChildWidget struct {
	width  float32
	height float32
	border bool
	flags  WindowFlags
	layout Layout
}

func (c *ChildWidget) Build() {
	showed := imgui.BeginChildV(GenAutoID("Child"), imgui.Vec2{X: c.width, Y: c.height}, c.border, int(c.flags))
	if showed && c.layout != nil {
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

func (c *ChildWidget) Layout(widgets ...Widget) *ChildWidget {
	c.layout = Layout(widgets)
	return c
}

func Child() *ChildWidget {
	return &ChildWidget{
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
		label:        tStr(label),
		previewValue: tStr(previewValue),
		width:        0,
		flags:        0,
		layout:       nil,
	}
}

func (cc *ComboCustomWidget) Layout(widgets ...Widget) *ComboCustomWidget {
	cc.layout = Layout(widgets)
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

	if imgui.BeginComboV(GenAutoID(cc.label), cc.previewValue, int(cc.flags)) {
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
	for _, item := range items {
		tStr(item)
	}

	return &ComboWidget{
		label:        tStr(label),
		previewValue: tStr(previewValue),
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

	if imgui.BeginComboV(GenAutoID(c.label), c.previewValue, int(c.flags)) {
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
	id          string
	mouseButton MouseButton
	layout      Layout
}

func ContextMenu() *ContextMenuWidget {
	return &ContextMenuWidget{
		mouseButton: MouseButtonRight,
		layout:      nil,
	}
}

func (c *ContextMenuWidget) Layout(widgets ...Widget) *ContextMenuWidget {
	c.layout = Layout(widgets)
	return c
}

func (c *ContextMenuWidget) MouseButton(mouseButton MouseButton) *ContextMenuWidget {
	c.mouseButton = mouseButton
	return c
}

func (c *ContextMenuWidget) ID(id string) *ContextMenuWidget {
	c.id = id
	return c
}

func (c *ContextMenuWidget) Build() {
	if len(c.id) == 0 {
		c.id = GenAutoID("ContextMenu")
	}

	if imgui.BeginPopupContextItemV(c.id, int(c.mouseButton)) {
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
		label:  tStr(label),
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
	imgui.DragIntV(GenAutoID(d.label), d.value, d.speed, d.min, d.max, d.format)
}

type ColumnWidget struct {
	widgets Layout
}

// Column layout will place all widgets one by one vertically.
func Column(widgets ...Widget) *ColumnWidget {
	return &ColumnWidget{
		widgets: widgets,
	}
}

func (g *ColumnWidget) Build() {
	imgui.BeginGroup()

	g.widgets.Build()

	imgui.EndGroup()
}

type ImageWidget struct {
	texture                *Texture
	width                  float32
	height                 float32
	uv0, uv1               image.Point
	tintColor, borderColor color.RGBA
	onClick                func()
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

func (i *ImageWidget) OnClick(cb func()) *ImageWidget {
	i.onClick = cb
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
		// trick: detect click event
		if i.onClick != nil && IsMouseClicked(MouseButtonLeft) {
			cursorPos := GetCursorScreenPos()
			mousePos := GetMousePos()
			mousePos.Add(cursorPos)
			if cursorPos.X <= mousePos.X && cursorPos.Y <= mousePos.Y &&
				cursorPos.X+int(i.width) >= mousePos.X && cursorPos.Y+int(i.height) >= mousePos.Y {
				i.onClick()
			}
		}

		imgui.ImageV(i.texture.id, size, ToVec2(i.uv0), ToVec2(i.uv1), ToVec4Color(i.tintColor), ToVec4Color(i.borderColor))
	} else {
		Dummy(i.width, i.height).Build()
	}
}

type ImageState struct {
	loading bool
	failure bool
	cancel  ctx.CancelFunc
	texture *Texture
}

func (is *ImageState) Dispose() {
	is.texture = nil
	// Cancel ongoing image downloaidng
	if is.loading && is.cancel != nil {
		is.cancel()
	}
}

type ImageWithRgbaWidget struct {
	id      string
	width   float32
	height  float32
	rgba    image.Image
	onClick func()
}

func ImageWithRgba(rgba image.Image) *ImageWithRgbaWidget {
	return &ImageWithRgbaWidget{
		id:     GenAutoID("ImageWithRgba"),
		width:  100,
		height: 100,
		rgba:   rgba,
	}
}

func (i *ImageWithRgbaWidget) Size(width, height float32) *ImageWithRgbaWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithRgbaWidget) OnClick(cb func()) *ImageWithRgbaWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithRgbaWidget) Build() {
	widget := Image(nil).Size(i.width, i.height).OnClick(i.onClick)

	if i.rgba != nil {
		state := Context.GetState(i.id)

		if state == nil {
			Context.SetState(i.id, &ImageState{})

			go func() {
				texture, err := NewTextureFromRgba(i.rgba)
				if err == nil {
					Context.SetState(i.id, &ImageState{texture: texture})
				}
			}()
		} else {
			imgState := state.(*ImageState)
			widget.texture = imgState.texture
		}
	}

	widget.Build()
}

type ImageWithFileWidget struct {
	id      string
	width   float32
	height  float32
	imgPath string
	onClick func()
}

func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		id:      fmt.Sprintf("ImageWithFile_%s", imgPath),
		width:   100,
		height:  100,
		imgPath: imgPath,
	}
}

func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithFileWidget) OnClick(cb func()) *ImageWithFileWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithFileWidget) Build() {
	state := Context.GetState(i.id)

	widget := Image(nil).OnClick(i.onClick).Size(i.width, i.height)

	if state == nil {
		// Prevent multiple invocation to LoadImage.
		Context.SetState(i.id, &ImageState{})

		img, err := LoadImage(i.imgPath)
		if err == nil {
			go func() {
				texture, err := NewTextureFromRgba(img)
				if err == nil {
					Context.SetState(i.id, &ImageState{texture: texture})
				}
			}()
		}
	} else {
		imgState := state.(*ImageState)
		widget.texture = imgState.texture
	}

	widget.Build()
}

type ImageWithUrlWidget struct {
	id              string
	imgUrl          string
	downloadTimeout time.Duration
	width           float32
	height          float32
	whenLoading     Layout
	whenFailure     Layout
	onReady         func()
	onFailure       func(error)
	onClick         func()
}

func ImageWithUrl(url string) *ImageWithUrlWidget {
	return &ImageWithUrlWidget{
		id:              fmt.Sprintf("ImageWithUrl_%s", url),
		imgUrl:          url,
		downloadTimeout: 10 * time.Second,
		width:           100,
		height:          100,
		whenLoading:     Layout{Dummy(100, 100)},
		whenFailure:     Layout{Dummy(100, 100)},
	}
}

// Event trigger when image is downloaded and ready to display.
func (i *ImageWithUrlWidget) OnReady(onReady func()) *ImageWithUrlWidget {
	i.onReady = onReady
	return i
}

func (i *ImageWithUrlWidget) OnFailure(onFailure func(error)) *ImageWithUrlWidget {
	i.onFailure = onFailure
	return i
}

func (i *ImageWithUrlWidget) OnClick(cb func()) *ImageWithUrlWidget {
	i.onClick = cb
	return i
}

func (i *ImageWithUrlWidget) Timeout(downloadTimeout time.Duration) *ImageWithUrlWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

func (i *ImageWithUrlWidget) Size(width, height float32) *ImageWithUrlWidget {
	i.width, i.height = width, height
	return i
}

func (i *ImageWithUrlWidget) LayoutForLoading(widgets ...Widget) *ImageWithUrlWidget {
	i.whenLoading = Layout(widgets)
	return i
}

func (i *ImageWithUrlWidget) LayoutForFailure(widgets ...Widget) *ImageWithUrlWidget {
	i.whenFailure = Layout(widgets)
	return i
}

func (i *ImageWithUrlWidget) Build() {
	state := Context.GetState(i.id)

	widget := Image(nil).OnClick(i.onClick).Size(i.width, i.height)

	if state == nil {
		// Prevent multiple invocation to download image.
		downloadContext, cancalFunc := ctx.WithCancel(ctx.Background())
		Context.SetState(i.id, &ImageState{loading: true, cancel: cancalFunc})

		go func() {
			// Load image from url
			client := resty.New()
			client.SetTimeout(i.downloadTimeout)
			resp, err := client.R().SetContext(downloadContext).Get(i.imgUrl)
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}

			img, _, err := image.Decode(bytes.NewReader(resp.Body()))
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}

			rgba := image.NewRGBA(img.Bounds())
			draw.Draw(rgba, img.Bounds(), img, image.Point{}, draw.Src)

			texture, err := NewTextureFromRgba(rgba)
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}
				return
			}
			Context.SetState(i.id, &ImageState{loading: false, texture: texture})

			// Trigger onReady event
			if i.onReady != nil {
				i.onReady()
			}
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

		widget.texture = imgState.texture
	}

	widget.Build()
}

type InputTextWidget struct {
	label      string
	hint       string
	value      *string
	width      float32
	candidates []string
	flags      InputTextFlags
	cb         imgui.InputTextCallback
	onChange   func()
}

type inputTextState struct {
	autoCompleteCandidates fuzzy.Matches
}

func (s *inputTextState) Dispose() {
	s.autoCompleteCandidates = nil
}

func InputText(value *string) *InputTextWidget {
	return &InputTextWidget{
		hint:     "",
		value:    value,
		width:    0,
		flags:    0,
		cb:       nil,
		onChange: nil,
	}
}

func (i *InputTextWidget) Label(label string) *InputTextWidget {
	i.label = tStr(label)
	return i
}

// Enable auto complete popup by using fuzzy search of current value agains candidates
// Press enter to confirm the first candidate
func (i *InputTextWidget) AutoComplete(candidates []string) *InputTextWidget {
	i.candidates = candidates
	return i
}

func (i *InputTextWidget) Hint(hint string) *InputTextWidget {
	i.hint = tStr(hint)
	return i
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
	if len(i.label) == 0 {
		i.label = GenAutoID(i.label)
	}

	// Get state
	var state *inputTextState
	if s := Context.GetState(i.label); s == nil {
		state = &inputTextState{}
		Context.SetState(i.label, state)
	} else {
		state = s.(*inputTextState)
	}

	if i.width != 0 {
		PushItemWidth(i.width)
	}

	isChanged := imgui.InputTextWithHint(i.label, i.hint, tStrPtr(i.value), int(i.flags), i.cb)

	if i.width != 0 {
		PopItemWidth()
	}

	if isChanged && i.onChange != nil {
		i.onChange()
	}

	if isChanged {
		// Enable auto complete
		if len(i.candidates) > 0 {
			matches := fuzzy.Find(*i.value, i.candidates)
			if matches.Len() > 0 {
				size := int(math.Min(5, float64(matches.Len())))
				matches = matches[:size]

				state.autoCompleteCandidates = matches
			}
		}
	}

	// Draw autocomplete list
	if len(state.autoCompleteCandidates) > 0 {
		labels := make(Layout, len(state.autoCompleteCandidates))
		for i, m := range state.autoCompleteCandidates {
			labels[i] = Label(m.Str)
		}

		SetNextWindowPos(imgui.GetItemRectMin().X, imgui.GetItemRectMax().Y)
		imgui.BeginTooltip()
		labels.Build()
		imgui.EndTooltip()

		// Press enter will replace value string with first match candidate
		if IsKeyPressed(KeyEnter) {
			*i.value = state.autoCompleteCandidates[0].Str
			state.autoCompleteCandidates = nil
		}
	}
}

type InputIntWidget struct {
	label    string
	value    *int32
	width    float32
	flags    InputTextFlags
	onChange func()
}

func InputInt(value *int32) *InputIntWidget {
	return &InputIntWidget{
		value:    value,
		width:    0,
		flags:    0,
		onChange: nil,
	}
}

func (i *InputIntWidget) Label(label string) *InputIntWidget {
	i.label = tStr(label)
	return i
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
	if len(i.label) == 0 {
		i.label = GenAutoID(i.label)
	}

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

func InputFloat(label string, value *float32) *InputFloatWidget {
	return &InputFloatWidget{
		width:    0,
		value:    value,
		format:   "%.3f",
		flags:    0,
		onChange: nil,
	}
}

func (i *InputFloatWidget) Label(label string) *InputFloatWidget {
	i.label = tStr(label)
	return i
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
	if len(i.label) == 0 {
		i.label = GenAutoID(i.label)
	}

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
	label    string
	fontInfo *FontInfo
	wrapped  bool
}

func Label(label string) *LabelWidget {
	return &LabelWidget{
		label:   tStr(label),
		wrapped: false,
	}
}

func Labelf(format string, args ...interface{}) *LabelWidget {
	return Label(fmt.Sprintf(format, args...))
}

func (l *LabelWidget) Wrapped(wrapped bool) *LabelWidget {
	l.wrapped = wrapped
	return l
}

func (l *LabelWidget) Font(font *FontInfo) *LabelWidget {
	l.fontInfo = font
	return l
}

func (l *LabelWidget) Build() {
	if l.wrapped {
		PushTextWrapPos()
	}

	shouldPopFont := false
	if l.fontInfo != nil {
		shouldPopFont = PushFont(l.fontInfo)
	}

	imgui.Text(l.label)

	if shouldPopFont {
		PopFont()
	}

	if l.wrapped {
		PopTextWrapPos()
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

func (m *MainMenuBarWidget) Layout(widgets ...Widget) *MainMenuBarWidget {
	m.layout = Layout(widgets)
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

func (m *MenuBarWidget) Layout(widgets ...Widget) *MenuBarWidget {
	m.layout = Layout(widgets)
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
		label:    tStr(label),
		selected: false,
		enabled:  true,
		onClick:  nil,
	}
}

func MenuItemf(format string, args ...interface{}) *MenuItemWidget {
	return MenuItem(fmt.Sprintf(format, args...))
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
	if imgui.MenuItemV(GenAutoID(m.label), "", m.selected, m.enabled) && m.onClick != nil {
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
		label:   tStr(label),
		enabled: true,
		layout:  nil,
	}
}

func Menuf(format string, args ...interface{}) *MenuWidget {
	return Menu(fmt.Sprintf(format, args...))
}

func (m *MenuWidget) Enabled(e bool) *MenuWidget {
	m.enabled = e
	return m
}

func (m *MenuWidget) Layout(widgets ...Widget) *MenuWidget {
	m.layout = Layout(widgets)
	return m
}

func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(GenAutoID(m.label), m.enabled) {
		if m.layout != nil {
			m.layout.Build()
		}
		imgui.EndMenu()
	}
}

type PopupWidget struct {
	name   string
	flags  WindowFlags
	layout Layout
}

func Popup(name string) *PopupWidget {
	return &PopupWidget{
		name:   tStr(name),
		flags:  0,
		layout: nil,
	}
}

func (p *PopupWidget) Flags(flags WindowFlags) *PopupWidget {
	p.flags = flags
	return p
}

func (p *PopupWidget) Layout(widgets ...Widget) *PopupWidget {
	p.layout = Layout(widgets)
	return p
}

func (p *PopupWidget) Build() {
	if imgui.BeginPopup(p.name, int(p.flags)) {
		if p.layout != nil {
			p.layout.Build()
		}
		imgui.EndPopup()
	}
}

type PopupModalWidget struct {
	name   string
	open   *bool
	flags  WindowFlags
	layout Layout
}

func PopupModal(name string) *PopupModalWidget {
	return &PopupModalWidget{
		name:   tStr(name),
		open:   nil,
		flags:  WindowFlagsNoResize,
		layout: nil,
	}
}

func (p *PopupModalWidget) IsOpen(open *bool) *PopupModalWidget {
	p.open = open
	return p
}

func (p *PopupModalWidget) Flags(flags WindowFlags) *PopupModalWidget {
	p.flags = flags
	return p
}

func (p *PopupModalWidget) Layout(widgets ...Widget) *PopupModalWidget {
	p.layout = Layout(widgets)
	return p
}

func (p *PopupModalWidget) Build() {
	if imgui.BeginPopupModalV(p.name, p.open, int(p.flags)) {
		if p.layout != nil {
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
	p.overlay = tStr(overlay)
	return p
}

func (p *ProgressBarWidget) Overlayf(format string, args ...interface{}) *ProgressBarWidget {
	p.overlay = tStr(fmt.Sprintf(p.overlay))
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
	onDClick func()
}

func Selectable(label string) *SelectableWidget {
	return &SelectableWidget{
		label:    tStr(label),
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
	scale := Context.platform.GetContentScale()
	s.width, s.height = width*scale, height*scale
	return s
}

func (s *SelectableWidget) OnClick(onClick func()) *SelectableWidget {
	s.onClick = onClick
	return s
}

// Handle mouse left button's double click event.
// SelectableFlagsAllowDoubleClick will set once tonDClick callback is notnull
func (s *SelectableWidget) OnDClick(onDClick func()) *SelectableWidget {
	s.onDClick = onDClick
	return s
}

func (s *SelectableWidget) Build() {
	// If onDClick is set, check flags and set related flag when necessary
	if s.onDClick != nil && s.flags&SelectableFlagsAllowDoubleClick != 0 {
		s.flags |= SelectableFlagsAllowDoubleClick
	}

	if imgui.SelectableV(GenAutoID(s.label), s.selected, int(s.flags), imgui.Vec2{X: s.width, Y: s.height}) && s.onClick != nil {
		s.onClick()
	}

	if s.onDClick != nil && IsItemActive() && IsMouseDoubleClicked(MouseButtonLeft) {
		s.onDClick()
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
	label    string
	value    *int32
	min      int32
	max      int32
	format   string
	width    float32
	onChange func()
}

func SliderInt(label string, value *int32, min, max int32) *SliderIntWidget {
	return &SliderIntWidget{
		label:    tStr(label),
		value:    value,
		min:      min,
		max:      max,
		format:   "%d",
		width:    0,
		onChange: nil,
	}
}

func (s *SliderIntWidget) Format(format string) *SliderIntWidget {
	s.format = format
	return s
}

func (s *SliderIntWidget) Size(width float32) *SliderIntWidget {
	s.width = width * Context.platform.GetContentScale()
	return s
}

func (s *SliderIntWidget) OnChange(onChange func()) *SliderIntWidget {
	s.onChange = onChange

	return s
}

func (s *SliderIntWidget) Build() {
	if s.width != 0 {
		PushItemWidth(s.width)
	}

	if imgui.SliderIntV(GenAutoID(s.label), s.value, s.min, s.max, s.format) && s.onChange != nil {
		s.onChange()
	}

	if s.width != 0 {
		PopItemWidth()
	}
}

type VSliderIntWidget struct {
	label    string
	width    float32
	height   float32
	value    *int32
	min      int32
	max      int32
	format   string
	flags    SliderFlags
	onChange func()
}

func VSliderInt(label string, value *int32, min, max int32) *VSliderIntWidget {
	return &VSliderIntWidget{
		label:  tStr(label),
		width:  18,
		height: 60,
		value:  value,
		min:    min,
		max:    max,
		format: "%d",
		flags:  SliderFlagsNone,
	}
}

func (vs *VSliderIntWidget) Size(width, height float32) *VSliderIntWidget {
	vs.width, vs.height = width, height
	return vs
}

func (vs *VSliderIntWidget) Flags(flags SliderFlags) *VSliderIntWidget {
	vs.flags = flags
	return vs
}

func (vs *VSliderIntWidget) Format(format string) *VSliderIntWidget {
	vs.format = format
	return vs
}

func (vs *VSliderIntWidget) OnChange(onChange func()) *VSliderIntWidget {
	vs.onChange = onChange
	return vs
}

func (vs *VSliderIntWidget) Build() {
	if imgui.VSliderIntV(
		GenAutoID(vs.label),
		imgui.Vec2{X: vs.width, Y: vs.height},
		vs.value,
		vs.min,
		vs.max,
		vs.format,
		int(vs.flags),
	) && vs.onChange != nil {
		vs.onChange()
	}
}

type SliderFloatWidget struct {
	label    string
	value    *float32
	min      float32
	max      float32
	format   string
	width    float32
	onChange func()
}

func SliderFloat(label string, value *float32, min, max float32) *SliderFloatWidget {
	return &SliderFloatWidget{
		label:    tStr(label),
		value:    value,
		min:      min,
		max:      max,
		format:   "%.3f",
		width:    0,
		onChange: nil,
	}
}

func (s *SliderFloatWidget) Format(format string) *SliderFloatWidget {
	s.format = format
	return s
}

func (sf *SliderFloatWidget) OnChange(onChange func()) *SliderFloatWidget {
	sf.onChange = onChange

	return sf
}

func (sf *SliderFloatWidget) Size(width float32) *SliderFloatWidget {
	sf.width = width * Context.platform.GetContentScale()
	return sf
}

func (sf *SliderFloatWidget) Build() {
	if sf.width != 0 {
		PushItemWidth(sf.width)
	}

	if imgui.SliderFloatV(GenAutoID(sf.label), sf.value, sf.min, sf.max, sf.format, 1.0) && sf.onChange != nil {
		sf.onChange()
	}

	if sf.width != 0 {
		PopItemWidth()
	}
}

type DummyWidget struct {
	width  float32
	height float32
}

func (d *DummyWidget) Build() {
	w, h := GetAvailableRegion()

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

func HSplitter(delta *float32) *HSplitterWidget {
	return &HSplitterWidget{
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (h *HSplitterWidget) Size(width, height float32) *HSplitterWidget {
	scale := Context.platform.GetContentScale()
	aw, ah := GetAvailableRegion()

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
	h.id = GenAutoID("HSplitter")

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

func VSplitter(delta *float32) *VSplitterWidget {
	return &VSplitterWidget{
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (v *VSplitterWidget) Size(width, height float32) *VSplitterWidget {
	aw, ah := GetAvailableRegion()
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
	v.id = GenAutoID("VSplitter")

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

func TabItem(label string) *TabItemWidget {
	return &TabItemWidget{
		label:  tStr(label),
		open:   nil,
		flags:  0,
		layout: nil,
	}
}

func TabItemf(format string, args ...interface{}) *TabItemWidget {
	return TabItem(fmt.Sprintf(format, args...))
}

func (t *TabItemWidget) IsOpen(open *bool) *TabItemWidget {
	t.open = open
	return t
}

func (t *TabItemWidget) Flags(flags TabItemFlags) *TabItemWidget {
	t.flags = flags
	return t
}

func (t *TabItemWidget) Layout(widgets ...Widget) *TabItemWidget {
	t.layout = Layout(widgets)
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
	id       string
	flags    TabBarFlags
	tabItems []*TabItemWidget
}

func TabBar() *TabBarWidget {
	return &TabBarWidget{
		flags: 0,
	}
}

func (t *TabBarWidget) Flags(flags TabBarFlags) *TabBarWidget {
	t.flags = flags
	return t
}

func (t *TabBarWidget) ID(id string) *TabBarWidget {
	t.id = id
	return t
}

func (t *TabBarWidget) TabItems(items ...*TabItemWidget) *TabBarWidget {
	t.tabItems = items
	return t
}

func (t *TabBarWidget) Build() {
	buildingId := t.id
	if len(buildingId) == 0 {
		buildingId = GenAutoID("TabBar")
	}
	if imgui.BeginTabBarV(buildingId, int(t.flags)) {
		for _, ti := range t.tabItems {
			ti.Build()
		}
		imgui.EndTabBar()
	}
}

type TableRowWidget struct {
	flags        TableRowFlags
	minRowHeight float64
	layout       Layout
	bgColor      *color.RGBA
}

func TableRow(widgets ...Widget) *TableRowWidget {
	return &TableRowWidget{
		flags:        0,
		minRowHeight: 0,
		layout:       widgets,
		bgColor:      nil,
	}
}

func (r *TableRowWidget) BgColor(c *color.RGBA) *TableRowWidget {
	r.bgColor = c
	return r
}

func (r *TableRowWidget) Flags(flags TableRowFlags) *TableRowWidget {
	r.flags = flags
	return r
}

func (r *TableRowWidget) MinHeight(height float64) *TableRowWidget {
	r.minRowHeight = height
	return r
}

func (r *TableRowWidget) Build() {
	imgui.TableNextRow(imgui.TableRowFlags(r.flags), r.minRowHeight)

	for _, w := range r.layout {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopup := w.(*PopupModalWidget)

		if !isTooltip && !isContextMenu && !isPopup {
			imgui.TableNextColumn()
		}

		w.Build()
	}

	if r.bgColor != nil {
		imgui.TableSetBgColor(imgui.TableBgTarget_RowBg0, uint32(imgui.GetColorU32(ToVec4Color(*(r.bgColor)))), -1)
	}
}

type TableColumnWidget struct {
	label              string
	flags              TableColumnFlags
	innerWidthOrWeight float32
	userId             uint32
}

func TableColumn(label string) *TableColumnWidget {
	return &TableColumnWidget{
		label:              tStr(label),
		flags:              0,
		innerWidthOrWeight: 0,
		userId:             0,
	}
}

func (c *TableColumnWidget) Flags(flags TableColumnFlags) *TableColumnWidget {
	c.flags = flags
	return c
}

func (c *TableColumnWidget) InnerWidthOrWeight(w float32) *TableColumnWidget {
	c.innerWidthOrWeight = w
	return c
}

func (c *TableColumnWidget) UserId(id uint32) *TableColumnWidget {
	c.userId = id
	return c
}

func (c *TableColumnWidget) Build() {
	imgui.TableSetupColumn(c.label, imgui.TableColumnFlags(c.flags), c.innerWidthOrWeight, c.userId)
}

type TableWidget struct {
	flags        TableFlags
	size         imgui.Vec2
	innerWidth   float64
	rows         []*TableRowWidget
	columns      []*TableColumnWidget
	fastMode     bool
	freezeRow    int
	freezeColumn int
}

func Table() *TableWidget {
	return &TableWidget{
		flags:        TableFlagsResizable | TableFlagsBorders | TableFlagsScrollY,
		rows:         nil,
		columns:      nil,
		fastMode:     false,
		freezeRow:    -1,
		freezeColumn: -1,
	}
}

// Display visible rows only to boost performance.
func (t *TableWidget) FastMode(b bool) *TableWidget {
	t.fastMode = b
	return t
}

// Freeze columns/rows so they stay visible when scrolled.
func (t *TableWidget) Freeze(col, row int) *TableWidget {
	t.freezeColumn = col
	t.freezeRow = row
	return t
}

func (t *TableWidget) Columns(cols ...*TableColumnWidget) *TableWidget {
	t.columns = cols
	return t
}

func (t *TableWidget) Rows(rows ...*TableRowWidget) *TableWidget {
	t.rows = rows
	return t
}

func (t *TableWidget) Size(width, height float32) *TableWidget {
	t.size = imgui.Vec2{X: width, Y: height}
	return t
}

func (t *TableWidget) InnerWidth(width float64) *TableWidget {
	t.innerWidth = width
	return t
}

func (t *TableWidget) Flags(flags TableFlags) *TableWidget {
	t.flags = flags
	return t
}

func (t *TableWidget) Build() {
	if len(t.rows) == 0 {
		return
	}

	colCount := len(t.columns)
	if colCount == 0 {
		colCount = len(t.rows[0].layout)
	}

	if imgui.BeginTable(GenAutoID("Table"), colCount, imgui.TableFlags(t.flags), t.size, t.innerWidth) {
		if t.freezeColumn >= 0 && t.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(t.freezeColumn, t.freezeRow)
		}

		if len(t.columns) > 0 {
			for _, col := range t.columns {
				col.Build()
			}
			imgui.TableHeadersRow()
		}

		if t.fastMode {
			var clipper imgui.ListClipper
			clipper.Begin(len(t.rows))

			for clipper.Step() {
				for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++ {
					row := t.rows[i]
					row.Build()
				}
			}

			clipper.End()
		} else {
			for _, row := range t.rows {
				row.Build()
			}
		}

		imgui.EndTable()
	}
}

type TreeTableRowWidget struct {
	label    string
	flags    TreeNodeFlags
	layout   Layout
	children []*TreeTableRowWidget
}

func TreeTableRow(label string, widgets ...Widget) *TreeTableRowWidget {
	return &TreeTableRowWidget{
		label:  label,
		layout: widgets,
	}
}

func (ttr *TreeTableRowWidget) Children(rows ...*TreeTableRowWidget) *TreeTableRowWidget {
	ttr.children = rows
	return ttr
}

func (ttr *TreeTableRowWidget) Flags(flags TreeNodeFlags) *TreeTableRowWidget {
	ttr.flags = flags
	return ttr
}

func (ttr *TreeTableRowWidget) Build() {
	imgui.TableNextRow(0, 0)
	imgui.TableNextColumn()

	open := false
	if len(ttr.children) > 0 {
		open = imgui.TreeNodeV(GenAutoID(ttr.label), int(ttr.flags))
	} else {
		ttr.flags |= TreeNodeFlagsLeaf | TreeNodeFlagsNoTreePushOnOpen
		imgui.TreeNodeV(GenAutoID(ttr.label), int(ttr.flags))
	}

	for _, w := range ttr.layout {
		_, isTooltip := w.(*TooltipWidget)
		_, isContextMenu := w.(*ContextMenuWidget)
		_, isPopup := w.(*PopupModalWidget)

		if !isTooltip && !isContextMenu && !isPopup {
			imgui.TableNextColumn()
		}

		w.Build()
	}

	if len(ttr.children) > 0 && open {
		for _, c := range ttr.children {
			c.Build()
		}

		imgui.TreePop()
	}
}

type TreeTableWidget struct {
	flags        TableFlags
	size         imgui.Vec2
	columns      []*TableColumnWidget
	rows         []*TreeTableRowWidget
	freezeRow    int
	freezeColumn int
}

func TreeTable() *TreeTableWidget {
	return &TreeTableWidget{
		flags:   TableFlagsBordersV | TableFlagsBordersOuterH | TableFlagsResizable | TableFlagsRowBg | TableFlagsNoBordersInBody,
		rows:    nil,
		columns: nil,
	}
}

// Freeze columns/rows so they stay visible when scrolled.
func (tt *TreeTableWidget) Freeze(col, row int) *TreeTableWidget {
	tt.freezeColumn = col
	tt.freezeRow = row
	return tt
}

func (tt *TreeTableWidget) Size(width, height float32) *TreeTableWidget {
	tt.size = imgui.Vec2{X: width, Y: height}
	return tt
}

func (tt *TreeTableWidget) Flags(flags TableFlags) *TreeTableWidget {
	tt.flags = flags
	return tt
}

func (tt *TreeTableWidget) Columns(cols ...*TableColumnWidget) *TreeTableWidget {
	tt.columns = cols
	return tt
}

func (tt *TreeTableWidget) Rows(rows ...*TreeTableRowWidget) *TreeTableWidget {
	tt.rows = rows
	return tt
}

func (tt *TreeTableWidget) Build() {
	if len(tt.rows) == 0 {
		return
	}

	colCount := len(tt.columns)
	if colCount == 0 {
		colCount = len(tt.rows[0].layout) + 1
	}

	if imgui.BeginTable(GenAutoID("TreeTable"), colCount, imgui.TableFlags(tt.flags), tt.size, 0) {
		if tt.freezeColumn >= 0 && tt.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(tt.freezeColumn, tt.freezeRow)
		}

		if len(tt.columns) > 0 {
			for _, col := range tt.columns {
				col.Build()
			}
			imgui.TableHeadersRow()
		}

		for _, row := range tt.rows {
			row.Build()
		}

		imgui.EndTable()
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
		tip:    tStr(tip),
		layout: nil,
	}
}

func Tooltipf(format string, args ...interface{}) *TooltipWidget {
	return Tooltip(fmt.Sprintf(format, args...))
}

func (t *TooltipWidget) Layout(widgets ...Widget) *TooltipWidget {
	t.layout = Layout(widgets)
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
		label:        tStr(label),
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

// Create TreeNode with eventHandler
// You could detect events (e.g. IsItemClicked IsMouseDoubleClicked etc...) and handle them for TreeNode inside eventHandler
func (t *TreeNodeWidget) Event(handler func()) *TreeNodeWidget {
	t.eventHandler = handler
	return t
}

func (t *TreeNodeWidget) Layout(widgets ...Widget) *TreeNodeWidget {
	t.layout = Layout(widgets)
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

	child := Child().Border(l.border).Size(l.width, l.height).Layout(Layout{
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
						ContextMenu().Layout(menus).Build()
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
		onChange: func() {}, // small hack - prevent giu from setting nil cb (skip nil check later)
	}
}

func (d *DatePickerWidget) Size(width float32) *DatePickerWidget {
	d.width = width * Context.platform.GetContentScale()
	return d
}

func (d *DatePickerWidget) OnChange(onChange func()) *DatePickerWidget {
	if onChange != nil {
		d.onChange = onChange
	}
	return d
}

func (d *DatePickerWidget) Build() {
	if d.date == nil {
		return
	}

	imgui.PushID(d.id)
	defer imgui.PopID()

	if d.width > 0 {
		PushItemWidth(d.width)
	}

	if imgui.BeginComboV(d.id, d.date.Format("2006-01-02"), imgui.ComboFlagHeightLargest) {
		// Build year widget
		imgui.AlignTextToFramePadding()
		imgui.Text(tStr(" Year"))
		imgui.SameLine()
		imgui.Text(tStr(fmt.Sprintf("%14d", d.date.Year())))
		imgui.SameLine()

		const yearButtonSize = 25
		Button(tStr("-##year")).OnClick(func() {
			*d.date = d.date.AddDate(-1, 0, 0)
			d.onChange()
		}).Size(yearButtonSize, yearButtonSize).Build()
		imgui.SameLine()
		Button(tStr("+##year")).OnClick(func() {
			*d.date = d.date.AddDate(1, 0, 0)
			d.onChange()
		}).Size(yearButtonSize, yearButtonSize).Build()

		// Build month widgets
		imgui.Text(tStr("Month"))
		imgui.SameLine()
		imgui.Text(tStr(fmt.Sprintf("%10s(%02d)", d.date.Month().String(), d.date.Month())))
		imgui.SameLine()
		Button(tStr("-##month")).OnClick(func() {
			*d.date = d.date.AddDate(0, -1, 0)
			d.onChange()
		}).Size(yearButtonSize, yearButtonSize).Build()
		imgui.SameLine()
		Button(tStr("+##month")).OnClick(func() {
			*d.date = d.date.AddDate(0, 1, 0)
			d.onChange()
		}).Size(yearButtonSize, yearButtonSize).Build()

		// Build day widgets
		firstDay := time.Date(d.date.Year(), d.date.Month(), 1, 0, 0, 0, 0, time.Local)
		lastDay := firstDay.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

		// store month days sorted in weeks
		var days [][]int

		// Build first row
		days = append(days, []int{})

		monthDay := 1
		for i := 0; i < 7; i++ {
			// check for the first month weekday
			if i < int(firstDay.Weekday()) {
				days[0] = append(days[0], 0)
				continue
			}

			days[0] = append(days[0], monthDay)
			monthDay++
		}

		// Build rest rows
		for ; monthDay <= lastDay.Day(); monthDay++ {
			if len(days[len(days)-1]) == 7 {
				days = append(days, []int{})
			}

			days[len(days)-1] = append(days[len(days)-1], monthDay)
		}

		// Pad last row
		lastRowLen := len(days[len(days)-1])
		if lastRowLen < 7 {
			for i := lastRowLen; i < 7; i++ {
				days[len(days)-1] = append(days[len(days)-1], 0)
			}
		}

		columns := []*TableColumnWidget{
			TableColumn("S"),
			TableColumn("M"),
			TableColumn("T"),
			TableColumn("W"),
			TableColumn("T"),
			TableColumn("F"),
			TableColumn("S"),
		}

		// Build day widgets
		var rows []*TableRowWidget

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

								d.onChange()
							}).Build()

							if d.date.Year() == today.Year() && d.date.Month() == today.Month() && day == today.Day() {
								imgui.PopStyleColor()
							}
						}),
					)
				}
			}

			rows = append(rows, TableRow(row...))
		}

		Table().Flags(TableFlagsBorders | TableFlagsSizingStretchSame).Columns(columns...).Rows(rows...).Build()

		imgui.EndCombo()
	}

	if d.width > 0 {
		PopItemWidth()
	}
}

type ColorEditWidget struct {
	label    string
	color    *color.RGBA
	flags    ColorEditFlags
	width    float32
	onChange func()
}

func ColorEdit(label string, color *color.RGBA) *ColorEditWidget {
	return &ColorEditWidget{
		label: tStr(label),
		color: color,
		flags: ColorEditFlagsNone,
	}
}

func (ce *ColorEditWidget) OnChange(cb func()) *ColorEditWidget {
	ce.onChange = cb
	return ce
}

func (ce *ColorEditWidget) Flags(f ColorEditFlags) *ColorEditWidget {
	ce.flags = f
	return ce
}

func (ce *ColorEditWidget) Size(width float32) *ColorEditWidget {
	ce.width = width * Context.platform.GetContentScale()
	return ce
}

func (ce *ColorEditWidget) Build() {
	c := ToVec4Color(*ce.color)
	col := [4]float32{
		c.X,
		c.Y,
		c.Z,
		c.W,
	}

	if ce.width > 0 {
		imgui.PushItemWidth(ce.width)
	}

	if imgui.ColorEdit4V(
		GenAutoID(ce.label),
		&col,
		int(ce.flags),
	) {
		*ce.color = Vec4ToRGBA(imgui.Vec4{
			X: col[0],
			Y: col[1],
			Z: col[2],
			W: col[3],
		})
		if ce.onChange != nil {
			ce.onChange()
		}
	}

	if ce.width > 0 {
		imgui.PopItemWidth()
	}
}
