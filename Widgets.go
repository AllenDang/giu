package giu

import (
	"fmt"
	"image/color"

	"github.com/AllenDang/imgui-go"
)

// GenAutoID automatically generates fidget's id.
func GenAutoID(id string) string {
	return fmt.Sprintf("%s##%d", id, Context.GetWidgetIndex())
}

var _ Widget = &RowWidget{}

// RowWidget joins a layout into one line
// calls imgui.SameLine().
type RowWidget struct {
	widgets Layout
}

// Row creates RowWidget.
func Row(widgets ...Widget) *RowWidget {
	return &RowWidget{
		widgets: widgets,
	}
}

// Build implements Widget interface.
func (l *RowWidget) Build() {
	isFirst := true
	l.widgets.Range(func(w Widget) {
		switch w.(type) {
		case *TooltipWidget,
			*ContextMenuWidget, *PopupModalWidget,
			*PopupWidget:
			// noop
		default:
			if _, isLabel := w.(*LabelWidget); isLabel {
				AlignTextToFramePadding()
			}

			if !isFirst {
				imgui.SameLine()
			} else {
				isFirst = false
			}
		}

		w.Build()
	})
}

// SameLine wrapps imgui.SomeLine
// Don't use if you don't have to (use RowWidget instead).
func SameLine() {
	imgui.SameLine()
}

var _ Widget = &ChildWidget{}

type ChildWidget struct {
	id     string
	width  float32
	height float32
	border bool
	flags  WindowFlags
	layout Layout
}

// Build implements Widget interface.
func (c *ChildWidget) Build() {
	if imgui.BeginChildV(c.id, imgui.Vec2{X: c.width, Y: c.height}, c.border, int(c.flags)) {
		c.layout.Build()
	}

	imgui.EndChild()
}

func (c *ChildWidget) Border(border bool) *ChildWidget {
	c.border = border
	return c
}

func (c *ChildWidget) Size(width, height float32) *ChildWidget {
	c.width, c.height = width, height
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

// ID sets the interval id of child widgets.
func (c *ChildWidget) ID(id string) *ChildWidget {
	c.id = id
	return c
}

func Child() *ChildWidget {
	return &ChildWidget{
		id:     GenAutoID("Child"),
		width:  0,
		height: 0,
		border: true,
		flags:  0,
		layout: nil,
	}
}

var _ Widget = &ComboCustomWidget{}

// ComboCustomWidget represents a combo with custom layout when opened.
type ComboCustomWidget struct {
	label        string
	previewValue string
	width        float32
	flags        ComboFlags
	layout       Layout
}

// ComboCustom creates a new combo custom widget.
func ComboCustom(label, previewValue string) *ComboCustomWidget {
	return &ComboCustomWidget{
		label:        GenAutoID(label),
		previewValue: tStr(previewValue),
		width:        0,
		flags:        0,
		layout:       nil,
	}
}

// Layout add combo's layout.
func (cc *ComboCustomWidget) Layout(widgets ...Widget) *ComboCustomWidget {
	cc.layout = Layout(widgets)
	return cc
}

// Flags allows to set combo flags (see Flags.go).
func (cc *ComboCustomWidget) Flags(flags ComboFlags) *ComboCustomWidget {
	cc.flags = flags
	return cc
}

// Size sets combo preiview width.
func (cc *ComboCustomWidget) Size(width float32) *ComboCustomWidget {
	cc.width = width
	return cc
}

// Build implements Widget interface.
func (cc *ComboCustomWidget) Build() {
	if cc.width > 0 {
		imgui.PushItemWidth(cc.width)
		defer imgui.PopItemWidth()
	}

	if imgui.BeginComboV(tStr(cc.label), cc.previewValue, int(cc.flags)) {
		cc.layout.Build()
		imgui.EndCombo()
	}
}

var _ Widget = &ComboWidget{}

// ComboWidget is a wrapper of ComboCustomWidget.
// It creates a combo of selectables. (it is the most frequently used).
type ComboWidget struct {
	label        string
	previewValue string
	items        []string
	selected     *int32
	width        float32
	flags        ComboFlags
	onChange     func()
}

// Combo creates a new ComboWidget.
func Combo(label, previewValue string, items []string, selected *int32) *ComboWidget {
	return &ComboWidget{
		label:        GenAutoID(label),
		previewValue: tStr(previewValue),
		items:        tStrSlice(items),
		selected:     selected,
		flags:        0,
		width:        0,
		onChange:     nil,
	}
}

// Build implements Widget interface.
func (c *ComboWidget) Build() {
	if c.width > 0 {
		imgui.PushItemWidth(c.width)
		defer imgui.PopItemWidth()
	}

	if imgui.BeginComboV(tStr(c.label), c.previewValue, int(c.flags)) {
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
}

// Flags allows to set combo flags (see Flags.go).
func (c *ComboWidget) Flags(flags ComboFlags) *ComboWidget {
	c.flags = flags
	return c
}

// Size sets combo's width.
func (c *ComboWidget) Size(width float32) *ComboWidget {
	c.width = width
	return c
}

// OnChange sets callback when combo value gets changed.
func (c *ComboWidget) OnChange(onChange func()) *ComboWidget {
	c.onChange = onChange
	return c
}

var _ Widget = &ContextMenuWidget{}

type ContextMenuWidget struct {
	id          string
	mouseButton MouseButton
	layout      Layout
}

func ContextMenu() *ContextMenuWidget {
	return &ContextMenuWidget{
		mouseButton: MouseButtonRight,
		layout:      nil,
		id:          GenAutoID("ContextMenu"),
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

// Build implements Widget interface.
func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.id, int(c.mouseButton)) {
		c.layout.Build()
		imgui.EndPopup()
	}
}

var _ Widget = &DragIntWidget{}

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
		label:  GenAutoID(label),
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

// Build implements Widget interface.
func (d *DragIntWidget) Build() {
	imgui.DragIntV(tStr(d.label), d.value, d.speed, d.min, d.max, d.format)
}

var _ Widget = &ColumnWidget{}

// ColumnWidget will place all widgets one by one vertically.
type ColumnWidget struct {
	widgets Layout
}

// Column creates a new ColumnWidget.
func Column(widgets ...Widget) *ColumnWidget {
	return &ColumnWidget{
		widgets: widgets,
	}
}

// Build implements Widget interface.
func (g *ColumnWidget) Build() {
	imgui.BeginGroup()

	g.widgets.Build()

	imgui.EndGroup()
}

var _ Widget = &MainMenuBarWidget{}

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

// Build implements Widget interface.
func (m *MainMenuBarWidget) Build() {
	if imgui.BeginMainMenuBar() {
		m.layout.Build()
		imgui.EndMainMenuBar()
	}
}

var _ Widget = &MenuBarWidget{}

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

// Build implements Widget interface.
func (m *MenuBarWidget) Build() {
	if imgui.BeginMenuBar() {
		m.layout.Build()
		imgui.EndMenuBar()
	}
}

var _ Widget = &MenuItemWidget{}

type MenuItemWidget struct {
	label    string
	selected bool
	enabled  bool
	onClick  func()
}

func MenuItem(label string) *MenuItemWidget {
	return &MenuItemWidget{
		label:    GenAutoID(label),
		selected: false,
		enabled:  true,
		onClick:  nil,
	}
}

func MenuItemf(format string, args ...any) *MenuItemWidget {
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

// Build implements Widget interface.
func (m *MenuItemWidget) Build() {
	if imgui.MenuItemV(tStr(m.label), "", m.selected, m.enabled) && m.onClick != nil {
		m.onClick()
	}
}

var _ Widget = &MenuWidget{}

type MenuWidget struct {
	label   string
	enabled bool
	layout  Layout
}

func Menu(label string) *MenuWidget {
	return &MenuWidget{
		label:   GenAutoID(label),
		enabled: true,
		layout:  nil,
	}
}

func Menuf(format string, args ...any) *MenuWidget {
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

// Build implements Widget interface.
func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(tStr(m.label), m.enabled) {
		m.layout.Build()
		imgui.EndMenu()
	}
}

var _ Widget = &ProgressBarWidget{}

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
	p.width, p.height = width, height
	return p
}

func (p *ProgressBarWidget) Overlay(overlay string) *ProgressBarWidget {
	p.overlay = tStr(overlay)
	return p
}

func (p *ProgressBarWidget) Overlayf(format string, args ...any) *ProgressBarWidget {
	return p.Overlay(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, imgui.Vec2{X: p.width, Y: p.height}, p.overlay)
}

var _ Widget = &SeparatorWidget{}

type SeparatorWidget struct{}

// Build implements Widget interface.
func (s *SeparatorWidget) Build() {
	imgui.Separator()
}

func Separator() *SeparatorWidget {
	return &SeparatorWidget{}
}

var _ Widget = &DummyWidget{}

type DummyWidget struct {
	width  float32
	height float32
}

// Build implements Widget interface.
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
		width:  width,
		height: height,
	}
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

func TabItemf(format string, args ...any) *TabItemWidget {
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

// BuildTabItem executes tab item build steps.
func (t *TabItemWidget) BuildTabItem() {
	if imgui.BeginTabItemV(t.label, t.open, int(t.flags)) {
		t.layout.Build()
		imgui.EndTabItem()
	}
}

var _ Widget = &TabBarWidget{}

type TabBarWidget struct {
	id       string
	flags    TabBarFlags
	tabItems []*TabItemWidget
}

func TabBar() *TabBarWidget {
	return &TabBarWidget{
		id:    GenAutoID("TabBar"),
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

// Build implements Widget interface.
func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id, int(t.flags)) {
		for _, ti := range t.tabItems {
			ti.BuildTabItem()
		}
		imgui.EndTabBar()
	}
}

var _ Widget = &TooltipWidget{}

type TooltipWidget struct {
	tip    string
	layout Layout
}

// Build implements Widget interface.
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

func Tooltipf(format string, args ...any) *TooltipWidget {
	return Tooltip(fmt.Sprintf(format, args...))
}

func (t *TooltipWidget) Layout(widgets ...Widget) *TooltipWidget {
	t.layout = Layout(widgets)
	return t
}

var _ Widget = &SpacingWidget{}

type SpacingWidget struct{}

// Build implements Widget interface.
func (s *SpacingWidget) Build() {
	imgui.Spacing()
}

func Spacing() *SpacingWidget {
	return &SpacingWidget{}
}

var _ Widget = &ColorEditWidget{}

type ColorEditWidget struct {
	label    string
	color    *color.RGBA
	flags    ColorEditFlags
	width    float32
	onChange func()
}

func ColorEdit(label string, c *color.RGBA) *ColorEditWidget {
	return &ColorEditWidget{
		label: GenAutoID(label),
		color: c,
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
	ce.width = width
	return ce
}

// Build implements Widget interface.
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
		tStr(ce.label),
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
