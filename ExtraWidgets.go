package giu

import (
	"fmt"
	"image"
	"time"

	"github.com/AllenDang/imgui-go"
)

var _ Widget = &HSplitterWidget{}

type HSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func HSplitter(delta *float32) *HSplitterWidget {
	return &HSplitterWidget{
		id:     GenAutoID("HSplitter"),
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (h *HSplitterWidget) Size(width, height float32) *HSplitterWidget {
	aw, ah := GetAvailableRegion()

	if width == 0 {
		h.width = aw
	} else {
		h.width = width
	}

	if height == 0 {
		h.height = ah
	} else {
		h.height = height
	}

	return h
}

func (h *HSplitterWidget) ID(id string) *HSplitterWidget {
	h.id = id
	return h
}

// Build implements Widget interface
// nolint:dupl // will fix later
func (h *HSplitterWidget) Build() {
	// Calc line position.
	width := 40
	height := 2

	pt := GetCursorScreenPos()

	centerX := int(h.width / 2)
	centerY := int(h.height / 2)

	ptMin := image.Pt(centerX-width/2, centerY-height/2)
	ptMax := image.Pt(centerX+width/2, centerY+height/2)

	style := imgui.CurrentStyle()
	c := Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButton(h.id, imgui.Vec2{X: h.width, Y: h.height})
	if imgui.IsItemActive() {
		*(h.delta) = imgui.CurrentIO().GetMouseDelta().Y
	} else {
		*(h.delta) = 0
	}
	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeNS)
		c = Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrabActive))
	}

	// Draw a line in the very center
	canvas := GetCanvas()
	canvas.AddRectFilled(pt.Add(ptMin), pt.Add(ptMax), c, 0, 0)
}

var _ Widget = &VSplitterWidget{}

type VSplitterWidget struct {
	id     string
	width  float32
	height float32
	delta  *float32
}

func VSplitter(delta *float32) *VSplitterWidget {
	return &VSplitterWidget{
		id:     GenAutoID("VSplitter"),
		width:  0,
		height: 0,
		delta:  delta,
	}
}

func (v *VSplitterWidget) Size(width, height float32) *VSplitterWidget {
	aw, ah := GetAvailableRegion()

	if width == 0 {
		v.width = aw
	} else {
		v.width = width
	}

	if height == 0 {
		v.height = ah
	} else {
		v.height = height
	}

	return v
}

func (v *VSplitterWidget) ID(id string) *VSplitterWidget {
	v.id = id
	return v
}

// Build implements Widget interface
// nolint:dupl // will fix later
func (v *VSplitterWidget) Build() {
	// Calc line position.
	width := 2
	height := 40

	pt := GetCursorScreenPos()

	centerX := int(v.width / 2)
	centerY := int(v.height / 2)

	ptMin := image.Pt(centerX-width/2, centerY-height/2)
	ptMax := image.Pt(centerX+width/2, centerY+height/2)

	style := imgui.CurrentStyle()
	c := Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrab))

	// Place a invisible button to capture event.
	imgui.InvisibleButton(v.id, imgui.Vec2{X: v.width, Y: v.height})
	if imgui.IsItemActive() {
		*(v.delta) = imgui.CurrentIO().GetMouseDelta().X
	} else {
		*(v.delta) = 0
	}

	if imgui.IsItemHovered() {
		imgui.SetMouseCursor(imgui.MouseCursorResizeEW)
		c = Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrabActive))
	}

	// Draw a line in the very center
	canvas := GetCanvas()
	canvas.AddRectFilled(pt.Add(ptMin), pt.Add(ptMax), c, 0, 0)
}

type TreeTableRowWidget struct {
	label    string
	flags    TreeNodeFlags
	layout   Layout
	children []*TreeTableRowWidget
}

func TreeTableRow(label string, widgets ...Widget) *TreeTableRowWidget {
	return &TreeTableRowWidget{
		label:  GenAutoID(label),
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

// BuildTreeTableRow executes table row building steps.
func (ttr *TreeTableRowWidget) BuildTreeTableRow() {
	imgui.TableNextRow(0, 0)
	imgui.TableNextColumn()

	open := false
	if len(ttr.children) > 0 {
		open = imgui.TreeNodeV(tStr(ttr.label), int(ttr.flags))
	} else {
		ttr.flags |= TreeNodeFlagsLeaf | TreeNodeFlagsNoTreePushOnOpen
		imgui.TreeNodeV(tStr(ttr.label), int(ttr.flags))
	}

	for _, w := range ttr.layout {
		switch w.(type) {
		case *TooltipWidget,
			*ContextMenuWidget, *PopupModalWidget:
			// noop
		default:
			imgui.TableNextColumn()
		}

		w.Build()
	}

	if len(ttr.children) > 0 && open {
		for _, c := range ttr.children {
			c.BuildTreeTableRow()
		}

		imgui.TreePop()
	}
}

var _ Widget = &TreeTableWidget{}

type TreeTableWidget struct {
	id           string
	flags        TableFlags
	size         imgui.Vec2
	columns      []*TableColumnWidget
	rows         []*TreeTableRowWidget
	freezeRow    int
	freezeColumn int
}

func TreeTable() *TreeTableWidget {
	return &TreeTableWidget{
		id:      GenAutoID("TreeTable"),
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

// Build implements Widget interface.
func (tt *TreeTableWidget) Build() {
	if len(tt.rows) == 0 {
		return
	}

	colCount := len(tt.columns)
	if colCount == 0 {
		colCount = len(tt.rows[0].layout) + 1
	}

	if imgui.BeginTable(tt.id, colCount, imgui.TableFlags(tt.flags), tt.size, 0) {
		if tt.freezeColumn >= 0 && tt.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(tt.freezeColumn, tt.freezeRow)
		}

		if len(tt.columns) > 0 {
			for _, col := range tt.columns {
				col.BuildTableColumn()
			}
			imgui.TableHeadersRow()
		}

		for _, row := range tt.rows {
			row.BuildTreeTableRow()
		}

		imgui.EndTable()
	}
}

var _ Widget = &CustomWidget{}

type CustomWidget struct {
	builder func()
}

// Build implements Widget interface.
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

var _ Widget = &ConditionWidget{}

type ConditionWidget struct {
	cond       bool
	layoutIf   Layout
	layoutElse Layout
}

func Condition(cond bool, layoutIf, layoutElse Layout) *ConditionWidget {
	return &ConditionWidget{
		cond:       cond,
		layoutIf:   layoutIf,
		layoutElse: layoutElse,
	}
}

// Build implements Widget interface.
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

// RangeBuilder batch create widgets and render only which is visible.
func RangeBuilder(id string, values []any, builder func(int, any) Widget) Layout {
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

var _ Widget = &ListBoxWidget{}

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
	l.width, l.height = width, height
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

// Build implements Widget interface
// nolint:gocognit // will fix later
func (l *ListBoxWidget) Build() {
	var state *ListBoxState
	if s := Context.GetState(l.id); s == nil {
		state = &ListBoxState{selectedIndex: 0}
		Context.SetState(l.id, state)
	} else {
		var isOk bool
		state, isOk = s.(*ListBoxState)
		Assert(isOk, "ListBoxWidget", "Build", "wrong state type recovered")
	}

	child := Child().Border(l.border).Size(l.width, l.height).Layout(Layout{
		Custom(func() {
			clipper := imgui.NewListClipper()
			defer clipper.Delete()

			clipper.Begin(len(l.items))

			for clipper.Step() {
				for i := clipper.DisplayStart(); i < clipper.DisplayEnd(); i++ {
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

var _ Widget = &DatePickerWidget{}

type DatePickerWidget struct {
	id          string
	date        *time.Time
	width       float32
	onChange    func()
	format      string
	startOfWeek time.Weekday
}

func DatePicker(id string, date *time.Time) *DatePickerWidget {
	return &DatePickerWidget{
		id:          GenAutoID(id),
		date:        date,
		width:       100,
		startOfWeek: time.Sunday,
		onChange:    func() {}, // small hack - prevent giu from setting nil cb (skip nil check later)
	}
}

func (d *DatePickerWidget) Size(width float32) *DatePickerWidget {
	d.width = width
	return d
}

func (d *DatePickerWidget) OnChange(onChange func()) *DatePickerWidget {
	if onChange != nil {
		d.onChange = onChange
	}
	return d
}

func (d *DatePickerWidget) Format(format string) *DatePickerWidget {
	d.format = format
	return d
}

func (d *DatePickerWidget) StartOfWeek(weekday time.Weekday) *DatePickerWidget {
	d.startOfWeek = weekday
	return d
}

func (d *DatePickerWidget) getFormat() string {
	if d.format == "" {
		return "2006-01-02" // default
	}
	return d.format
}

func (d *DatePickerWidget) offsetDay(offset int) time.Weekday {
	day := (int(d.startOfWeek) + offset) % 7
	// offset may be negative, thus day can be negative
	day = (day + 7) % 7
	return time.Weekday(day)
}

// Build implements Widget interface.
func (d *DatePickerWidget) Build() {
	if d.date == nil {
		return
	}

	imgui.PushID(d.id)
	defer imgui.PopID()

	if d.width > 0 {
		PushItemWidth(d.width)
		defer PopItemWidth()
	}

	if imgui.BeginComboV(d.id+"##Combo", d.date.Format(d.getFormat()), imgui.ComboFlagsHeightLargest) {
		// --- [Build year widget] ---
		imgui.AlignTextToFramePadding()

		const yearButtonSize = 25

		Row(
			Label(tStr(" Year")),
			Labelf("%14d", d.date.Year()),
			Button("-##"+d.id+"year").OnClick(func() {
				*d.date = d.date.AddDate(-1, 0, 0)
				d.onChange()
			}).Size(yearButtonSize, yearButtonSize),
			Button("+##"+d.id+"year").OnClick(func() {
				*d.date = d.date.AddDate(1, 0, 0)
				d.onChange()
			}).Size(yearButtonSize, yearButtonSize),
		).Build()

		// --- [Build month widgets] ---
		Row(
			Label("Month"),
			Labelf("%10s(%02d)", d.date.Month().String(), d.date.Month()),
			Button("-##"+d.id+"month").OnClick(func() {
				*d.date = d.date.AddDate(0, -1, 0)
				d.onChange()
			}).Size(yearButtonSize, yearButtonSize),
			Button("+##"+d.id+"month").OnClick(func() {
				*d.date = d.date.AddDate(0, 1, 0)
				d.onChange()
			}).Size(yearButtonSize, yearButtonSize),
		).Build()

		// --- [Build day widgets] ---
		days := d.getDaysGroups()

		// Create calendar (widget)
		columns := make([]*TableColumnWidget, 7)

		for i := 0; i < 7; i++ {
			firstChar := d.offsetDay(i).String()[0:1]
			columns[i] = TableColumn(firstChar)
		}

		// Build day widgets
		var rows []*TableRowWidget

		for _, week := range days {
			var row []Widget

			for _, day := range week {
				day := day // hack for golang ranges
				if day == 0 {
					row = append(row, Label(" "))
					continue
				}

				row = append(row, d.calendarField(day))
			}

			rows = append(rows, TableRow(row...))
		}

		Table().Flags(TableFlagsBorders | TableFlagsSizingStretchSame).Columns(columns...).Rows(rows...).Build()

		imgui.EndCombo()
	}
}

// store month days sorted in weeks.
func (d *DatePickerWidget) getDaysGroups() (days [][]int) {
	firstDay := time.Date(d.date.Year(), d.date.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDay := firstDay.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	// calculate first week
	days = append(days, make([]int, 7))

	monthDay := 1
	emptyDaysInFirstWeek := (int(firstDay.Weekday()) - int(d.startOfWeek) + 7) % 7
	for i := emptyDaysInFirstWeek; i < 7; i++ {
		days[0][i] = monthDay
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

	return days
}

func (d *DatePickerWidget) calendarField(day int) Widget {
	today := time.Now()
	highlightColor := imgui.CurrentStyle().GetColor(imgui.StyleColorPlotHistogram)
	return Custom(func() {
		isToday := d.date.Year() == today.Year() && d.date.Month() == today.Month() && day == today.Day()
		if isToday {
			imgui.PushStyleColor(imgui.StyleColorText, highlightColor)
		}

		Selectable(fmt.Sprintf("%02d", day)).Selected(isToday).OnClick(func() {
			*d.date = time.Date(
				d.date.Year(), d.date.Month(), day,
				0, 0, 0, 0,
				d.date.Location())
			d.onChange()
		}).Build()

		if isToday {
			imgui.PopStyleColor()
		}
	})
}
