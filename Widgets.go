package giu

import (
	"fmt"
	"image/color"
	"math"

	"github.com/AllenDang/cimgui-go/imgui"
)

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

// SameLine wraps imgui.SomeLine
// Don't use if you don't have to (use RowWidget instead).
func SameLine() {
	imgui.SameLine()
}

var _ Widget = &ChildWidget{}

// ChildWidget is a container widget. It will have a separated scroll bar.
// Use Child if you want to create a layout of e specific size.
type ChildWidget struct {
	id     ID
	width  float32
	height float32
	border bool
	flags  WindowFlags
	layout Layout
}

// Child creates a new ChildWidget.
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

// Border sets whether child should have border
// You can use imgui.ChildFlagsBorders as well.
func (c *ChildWidget) Border(border bool) *ChildWidget {
	c.border = border
	return c
}

// Size sets child size.
func (c *ChildWidget) Size(width, height float32) *ChildWidget {
	c.width, c.height = width, height
	return c
}

// Flags allows to specify Child flags.
func (c *ChildWidget) Flags(flags WindowFlags) *ChildWidget {
	c.flags = flags
	return c
}

// Layout sets widgets that will be rendered inside of the Child.
func (c *ChildWidget) Layout(widgets ...Widget) *ChildWidget {
	c.layout = Layout(widgets)
	return c
}

// ID sets the interval id of child widgets.
func (c *ChildWidget) ID(id ID) *ChildWidget {
	c.id = id
	return c
}

// Build makes a Child.
func (c *ChildWidget) Build() {
	if imgui.BeginChildStrV(c.id.String(), imgui.Vec2{X: c.width, Y: c.height}, func() imgui.ChildFlags {
		if c.border {
			return imgui.ChildFlagsBorders
		}

		return 0
	}(), imgui.WindowFlags(c.flags)) {
		c.layout.Build()
	}

	imgui.EndChild()
}

var _ Widget = &ComboCustomWidget{}

// ComboCustomWidget represents a combo with custom layout when opened.
type ComboCustomWidget struct {
	label        ID
	previewValue string
	width        float32
	flags        ComboFlags
	layout       Layout
}

// ComboCustom creates a new combo custom widget.
func ComboCustom(label, previewValue string) *ComboCustomWidget {
	return &ComboCustomWidget{
		label:        GenAutoID(label),
		previewValue: Context.FontAtlas.RegisterString(previewValue),
		width:        0,
		flags:        0,
		layout:       nil,
	}
}

// Layout add combo's layout.
func (cc *ComboCustomWidget) Layout(widgets ...Widget) *ComboCustomWidget {
	cc.layout = widgets
	return cc
}

// Flags allows to set combo flags (see Flags.go).
func (cc *ComboCustomWidget) Flags(flags ComboFlags) *ComboCustomWidget {
	cc.flags = flags
	return cc
}

// Size sets combo preview width.
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

	if imgui.BeginComboV(Context.PrepareString(cc.label.String()), cc.previewValue, imgui.ComboFlags(cc.flags)) {
		cc.layout.Build()
		imgui.EndCombo()
	}
}

var _ Disposable = &comboFilterState{}

type comboFilterState struct {
	filter *imgui.TextFilter
}

func (c *comboFilterState) Dispose() {
	// noop
}

var _ Widget = &ComboWidget{}

// ComboWidget is a wrapper of ComboCustomWidget.
// It creates a combo of selectables. (it is the most frequently used).
type ComboWidget struct {
	label        ID
	previewValue string
	items        []string
	selected     *int32
	width        float32
	flags        ComboFlags
	filter       bool
	filterLabel  ID
	onChange     func()
}

// Combo creates a new ComboWidget.
func Combo(label, previewValue string, items []string, selected *int32) *ComboWidget {
	return &ComboWidget{
		label:        GenAutoID(label),
		previewValue: Context.FontAtlas.RegisterString(previewValue),
		items:        Context.FontAtlas.RegisterStringSlice(items),
		selected:     selected,
		flags:        0,
		width:        0,
		filter:       false,
		filterLabel:  GenAutoID("##Filter"),
		onChange:     nil,
	}
}

// ID sets the interval id of combo. (overrides label).
func (c *ComboWidget) ID(id ID) *ComboWidget {
	c.label = id
	return c
}

// Flags allows to set combo flags (see Flags.go).
func (c *ComboWidget) Flags(flags ComboFlags) *ComboWidget {
	c.flags = flags
	return c
}

// Filter enables/disables the combo filter.
func (c *ComboWidget) Filter(filter bool) *ComboWidget {
	c.filter = filter
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

// Build implements Widget interface.
func (c *ComboWidget) Build() {
	if c.width > 0 {
		imgui.PushItemWidth(c.width)
		defer imgui.PopItemWidth()
	}

	var state *comboFilterState
	if c.filter {
		if state = GetState[comboFilterState](Context, c.label); state == nil {
			state = &comboFilterState{filter: imgui.NewEmptyTextFilter()}
			SetState(Context, c.label, state)
		}
	}

	if imgui.BeginComboV(Context.PrepareString(c.label.String()), c.previewValue, imgui.ComboFlags(c.flags)) {
		if c.filter {
			if imgui.IsWindowAppearing() {
				imgui.SetKeyboardFocusHere()
				state.filter.Clear()
			}

			state.filter.DrawV(Context.PrepareString(c.filterLabel.String()), -math.SmallestNonzeroFloat32)
		}

		for i, item := range c.items {
			if c.filter && !state.filter.PassFilter(item) {
				continue
			}

			if imgui.SelectableBool(fmt.Sprintf("%s##%d", Context.PrepareString(item), i)) {
				*c.selected = int32(i)
				if c.onChange != nil {
					c.onChange()
				}
			}
		}

		imgui.EndCombo()
	}
}

var _ Widget = &ContextMenuWidget{}

// ContextMenuWidget is a context menu on another widget. (e.g. right-click menu on button).
type ContextMenuWidget struct {
	id          ID
	mouseButton MouseButton
	layout      Layout
}

// ContextMenu creates new ContextMenuWidget.
func ContextMenu() *ContextMenuWidget {
	return &ContextMenuWidget{
		mouseButton: MouseButtonRight,
		layout:      nil,
		id:          GenAutoID("ContextMenu"),
	}
}

// Layout sets layout of the context menu.
func (c *ContextMenuWidget) Layout(widgets ...Widget) *ContextMenuWidget {
	c.layout = Layout(widgets)
	return c
}

// MouseButton sets mouse button that will trigger the context menu.
func (c *ContextMenuWidget) MouseButton(mouseButton MouseButton) *ContextMenuWidget {
	c.mouseButton = mouseButton
	return c
}

// ID sets the interval id of context menu.
func (c *ContextMenuWidget) ID(id ID) *ContextMenuWidget {
	c.id = id
	return c
}

// Build implements Widget interface.
func (c *ContextMenuWidget) Build() {
	if imgui.BeginPopupContextItemV(c.id.String(), imgui.PopupFlags(c.mouseButton)) {
		c.layout.Build()
		imgui.EndPopup()
	}
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

// MainMenuBarWidget is a widget that creates a main menu bar.
// Main means that it will be docked to the MasterWindow.
// Do NOT use with SingleWindow (see MenuBarWidget).
type MainMenuBarWidget struct {
	layout Layout
}

// MainMenuBar creates new MainMenuBarWidget.
func MainMenuBar() *MainMenuBarWidget {
	return &MainMenuBarWidget{
		layout: nil,
	}
}

// Layout sets layout of the menu bar. (See MenuWidget).
func (m *MainMenuBarWidget) Layout(widgets ...Widget) *MainMenuBarWidget {
	m.layout = widgets
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

// MenuBarWidget is a widget that creates a menu bar for a window.
// Use it e.g. with SingleWindowWithMenuBar.
type MenuBarWidget struct {
	layout Layout
}

// MenuBar creates new MenuBarWidget.
func MenuBar() *MenuBarWidget {
	return &MenuBarWidget{
		layout: nil,
	}
}

// Layout sets layout of the menu bar. (See MenuWidget).
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

// MenuItemWidget is a menu node. Commonly used inside of MenuWidget.
type MenuItemWidget struct {
	label    ID
	shortcut string
	selected bool
	enabled  bool
	onClick  func()
}

// MenuItem creates new MenuItemWidget.
func MenuItem(label string) *MenuItemWidget {
	return &MenuItemWidget{
		label:    GenAutoID(label),
		shortcut: "",
		selected: false,
		enabled:  true,
		onClick:  nil,
	}
}

// MenuItemf creates MenuItem with formated label.
func MenuItemf(format string, args ...any) *MenuItemWidget {
	return MenuItem(fmt.Sprintf(format, args...))
}

// Shortcut sets shortcut of the item (grayed, right-aligned text). Used for presenting e.g. keyboard shortcuts (e.g. "Ctrl+S")
// NOTE: this is only a visual effect. It has nothing to do with keyboard shortcuts.
func (m *MenuItemWidget) Shortcut(s string) *MenuItemWidget {
	m.shortcut = s
	return m
}

// Selected sets whether the item is selected.
func (m *MenuItemWidget) Selected(s bool) *MenuItemWidget {
	m.selected = s
	return m
}

// Enabled sets whether the item is enabled.
func (m *MenuItemWidget) Enabled(e bool) *MenuItemWidget {
	m.enabled = e
	return m
}

// OnClick sets callback that will be executed when item is clicked.
func (m *MenuItemWidget) OnClick(onClick func()) *MenuItemWidget {
	m.onClick = onClick
	return m
}

// Build implements Widget interface.
func (m *MenuItemWidget) Build() {
	if imgui.MenuItemBoolV(Context.PrepareString(m.label.String()), m.shortcut, m.selected, m.enabled) && m.onClick != nil {
		m.onClick()
	}
}

var _ Widget = &MenuWidget{}

// MenuWidget is a node of (Main)MenuBarWidget.
// See also: MenuItemWidget, MenuBarWidget, MainMenuBarWidget.
type MenuWidget struct {
	label   ID
	enabled bool
	layout  Layout
}

// Menu creates new MenuWidget.
func Menu(label string) *MenuWidget {
	return &MenuWidget{
		label:   GenAutoID(label),
		enabled: true,
		layout:  nil,
	}
}

// Menuf is alias to Menu(fmt.Sprintf(format, args...)).
func Menuf(format string, args ...any) *MenuWidget {
	return Menu(fmt.Sprintf(format, args...))
}

// Enabled sets whether the menu is enabled.
func (m *MenuWidget) Enabled(e bool) *MenuWidget {
	m.enabled = e
	return m
}

// Layout sets layout of the menu. (See MenuItemWidget).
func (m *MenuWidget) Layout(widgets ...Widget) *MenuWidget {
	m.layout = widgets
	return m
}

// Build implements Widget interface.
func (m *MenuWidget) Build() {
	if imgui.BeginMenuV(Context.PrepareString(m.label.String()), m.enabled) {
		m.layout.Build()
		imgui.EndMenu()
	}
}

var _ Widget = &ProgressBarWidget{}

// ProgressBarWidget is a progress bar (like in windows' copy-file dialog).
// It is a perfect solution to indicate percentage progress of some action.
type ProgressBarWidget struct {
	fraction float32
	width    float32
	height   float32
	overlay  string
}

// ProgressBar creates new ProgressBar.
func ProgressBar(fraction float32) *ProgressBarWidget {
	return &ProgressBarWidget{
		fraction: fraction,
		width:    0,
		height:   0,
		overlay:  "",
	}
}

// Size sets size of the bar.
func (p *ProgressBarWidget) Size(width, height float32) *ProgressBarWidget {
	p.width, p.height = width, height
	return p
}

// Overlay sets custom overlay displayed on the bar.
func (p *ProgressBarWidget) Overlay(overlay string) *ProgressBarWidget {
	p.overlay = Context.FontAtlas.RegisterString(overlay)
	return p
}

// Overlayf is alias to Overlay(fmt.Sprintf(format, args...)).
func (p *ProgressBarWidget) Overlayf(format string, args ...any) *ProgressBarWidget {
	return p.Overlay(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (p *ProgressBarWidget) Build() {
	imgui.ProgressBarV(p.fraction, imgui.Vec2{X: p.width, Y: p.height}, p.overlay)
}

var _ Widget = &SeparatorWidget{}

// SeparatorWidget is like <hr> in HTML.
// Creates a layout-wide line.
type SeparatorWidget struct{}

// Separator creates new SeparatorWidget.
func Separator() *SeparatorWidget {
	return &SeparatorWidget{}
}

// Build implements Widget interface.
func (s *SeparatorWidget) Build() {
	imgui.Separator()
}

var _ Widget = &DummyWidget{}

// DummyWidget creates an empty space (moves drawing cursor by width and height).
type DummyWidget struct {
	width  float32
	height float32
}

// Dummy creates new DummyWidget.
func Dummy(width, height float32) *DummyWidget {
	return &DummyWidget{
		width:  width,
		height: height,
	}
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

// TabItemWidget is an item in TabBarWidget.
type TabItemWidget struct {
	label        string
	open         *bool
	flags        TabItemFlags
	layout       Layout
	eventHandler *EventHandler
}

// TabItem creates new TabItem.
func TabItem(label string) *TabItemWidget {
	return &TabItemWidget{
		label:  label,
		open:   nil,
		flags:  0,
		layout: nil,
	}
}

// TabItemf creates tab item with formated label.
func TabItemf(format string, args ...any) *TabItemWidget {
	return TabItem(fmt.Sprintf(format, args...))
}

// IsOpen takes a pointer to a boolean.
// Value of this pointer indicated whether TabItem is currently selected.
// NOTE: The item will NOT be opened/closed if this value is changed.
// It has only one-side effect.
func (t *TabItemWidget) IsOpen(open *bool) *TabItemWidget {
	t.open = open
	return t
}

// Flags allows to set item's flags.
func (t *TabItemWidget) Flags(flags TabItemFlags) *TabItemWidget {
	t.flags = flags
	return t
}

// EventHandler allows to attach a custym EventHandler to the tab item in order to detect events on it.
func (t *TabItemWidget) EventHandler(handler *EventHandler) *TabItemWidget {
	t.eventHandler = handler
	return t
}

// Layout is a layout displayed when item is opened.
func (t *TabItemWidget) Layout(widgets ...Widget) *TabItemWidget {
	t.layout = widgets
	return t
}

// BuildTabItem executes tab item build steps.
func (t *TabItemWidget) BuildTabItem() {
	start := imgui.BeginTabItemV(
		Context.PrepareString(t.label),
		t.open, imgui.TabItemFlags(t.flags),
	)

	if t.eventHandler != nil {
		t.eventHandler.Build()
	}

	if start {
		t.layout.Build()
		imgui.EndTabItem()
	}
}

var _ Widget = &TabBarWidget{}

// TabBarWidget is a bar of TabItemWidgets.
type TabBarWidget struct {
	id       ID
	flags    TabBarFlags
	tabItems []*TabItemWidget
}

// TabBar creates new TabBarWidget.
func TabBar() *TabBarWidget {
	return &TabBarWidget{
		id:    GenAutoID("TabBar"),
		flags: 0,
	}
}

// Flags allows to set TabBArFlags.
func (t *TabBarWidget) Flags(flags TabBarFlags) *TabBarWidget {
	t.flags = flags
	return t
}

// ID manually sets widget's ID.
func (t *TabBarWidget) ID(id ID) *TabBarWidget {
	t.id = id
	return t
}

// TabItems sets list of TabItemWidgets in the bar.
func (t *TabBarWidget) TabItems(items ...*TabItemWidget) *TabBarWidget {
	t.tabItems = items
	return t
}

// Build implements Widget interface.
func (t *TabBarWidget) Build() {
	if imgui.BeginTabBarV(t.id.String(), imgui.TabBarFlags(t.flags)) {
		for _, ti := range t.tabItems {
			ti.BuildTabItem()
		}

		imgui.EndTabBar()
	}
}

var _ Widget = &TooltipWidget{}

// TooltipWidget sets a tooltip on the previous widget.
// The tooltip can be anything.
type TooltipWidget struct {
	tip    string
	layout Layout
	to     Layout
}

// Tooltip creates new tooltip with given label
// NOTE: you can set the empty label and use Layout() method.
func Tooltip(tip string) *TooltipWidget {
	return &TooltipWidget{
		tip:    Context.FontAtlas.RegisterString(tip),
		layout: nil,
	}
}

// Tooltipf sets formated label.
func Tooltipf(format string, args ...any) *TooltipWidget {
	return Tooltip(fmt.Sprintf(format, args...))
}

// Layout sets a custom layout of tooltip.
func (t *TooltipWidget) Layout(widgets ...Widget) *TooltipWidget {
	t.layout = widgets
	return t
}

// To sets layout to which the tooltip should be attached.
// NOTE: This is an optional approach. By default tooltip is attached to the previous widget.
func (t *TooltipWidget) To(layout ...Widget) *TooltipWidget {
	t.to = Layout(layout)
	return t
}

// Build implements Widget interface.
func (t *TooltipWidget) Build() {
	if t.to != nil {
		t.to.Range(func(w Widget) {
			w.Build()
			t.buildTooltip()
		})

		return
	}

	t.buildTooltip()
}

func (t *TooltipWidget) buildTooltip() {
	if imgui.IsItemHovered() {
		if t.layout != nil {
			imgui.BeginTooltip()
			t.layout.Build()
			imgui.EndTooltip()
		} else {
			imgui.SetTooltip(Context.PrepareString(t.tip))
		}
	}
}

var _ Widget = &SpacingWidget{}

// SpacingWidget increases a spacing between two widgets a bit.
type SpacingWidget struct{}

// Spacing creates new SpacingWidget.
func Spacing() *SpacingWidget {
	return &SpacingWidget{}
}

// Build implements Widget interface.
func (s *SpacingWidget) Build() {
	imgui.Spacing()
}

var _ Widget = &ColorEditWidget{}

// ColorEditWidget is a widget that provides a color editor.
type ColorEditWidget struct {
	label    ID
	color    *color.RGBA
	flags    ColorEditFlags
	width    float32
	onChange func()
}

// ColorEdit creates new ColorEditWidget.
func ColorEdit(label string, c *color.RGBA) *ColorEditWidget {
	return &ColorEditWidget{
		label: GenAutoID(label),
		color: c,
		// flags: ColorEditFlagsNone,
	}
}

// OnChange sets callback that will be executed when color is changed.
func (ce *ColorEditWidget) OnChange(cb func()) *ColorEditWidget {
	ce.onChange = cb
	return ce
}

// Flags allows to set ColorEditFlags.
func (ce *ColorEditWidget) Flags(f ColorEditFlags) *ColorEditWidget {
	ce.flags = f
	return ce
}

// Size sets width of the color editor.
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
		Context.PrepareString(ce.label.String()),
		&col,
		imgui.ColorEditFlags(ce.flags),
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
