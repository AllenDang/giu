package giu

import (
	"fmt"
	"image"
	"time"

	"github.com/AllenDang/imgui-go"
)

var _ Widget = &SplitterWidget{}

// SplitterWidget is a line (vertical or horizontal) that splits layout (child)
// Int two pieces. It has a tiny button in the middle of that line and its creator
// takes float pointer so that you can read user's movement of this rect.
// Generally used by SplitLayoutWidget.
type SplitterWidget struct {
	id        string
	width     float32
	height    float32
	delta     *float32
	direction SplitDirection
}

// Splitter creates new SplitterWidget.
func Splitter(direction SplitDirection, delta *float32) *SplitterWidget {
	return &SplitterWidget{
		id:        GenAutoID("Splitter"),
		width:     0,
		height:    0,
		delta:     delta,
		direction: direction,
	}
}

// Size sets size of the button aray.
func (h *SplitterWidget) Size(width, height float32) *SplitterWidget {
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

// ID allows to set widget's ID manually.
func (h *SplitterWidget) ID(id string) *SplitterWidget {
	h.id = id
	return h
}

// Build implements Widget interface.
func (h *SplitterWidget) Build() {
	// Calc line position.
	var width, height int

	switch h.direction {
	case DirectionHorizontal:
		width = 40
		height = 2
	case DirectionVertical:
		width = 2
		height = 40
	}

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
		switch h.direction {
		case DirectionHorizontal:
			*(h.delta) = imgui.CurrentIO().GetMouseDelta().Y
		case DirectionVertical:
			*(h.delta) = imgui.CurrentIO().GetMouseDelta().X
		}
	} else {
		*(h.delta) = 0
	}

	if imgui.IsItemHovered() {
		switch h.direction {
		case DirectionHorizontal:
			imgui.SetMouseCursor(imgui.MouseCursorResizeNS)
		case DirectionVertical:
			imgui.SetMouseCursor(imgui.MouseCursorResizeEW)
		}

		c = Vec4ToRGBA(style.GetColor(imgui.StyleColorScrollbarGrabActive))
	}

	// Draw a line in the very center
	canvas := GetCanvas()
	canvas.AddRectFilled(pt.Add(ptMin), pt.Add(ptMax), c, 0, 0)
}

var _ Widget = &CustomWidget{}

// CustomWidget allows you to do whatever you want.
// This includes:
// - using functions from upstream imgui instead of thes from giu
// - build widgets in loop (see also RangeBuilder)
// - do any calculations needed in this part of rendering.
type CustomWidget struct {
	builder func()
}

// Custom creates a new custom widget.
func Custom(builder func()) *CustomWidget {
	return &CustomWidget{
		builder: builder,
	}
}

// Build implements Widget interface.
func (c *CustomWidget) Build() {
	if c.builder != nil {
		c.builder()
	}
}

// Plot implements Plot interface.
func (c *CustomWidget) Plot() {
	c.Build()
}

var _ Widget = &ConditionWidget{}

// ConditionWidget allows to build if a condition is met
// it is like:
//
//	if condition {
//	   layoutIf.Build()
//	} else {
//
//	   layoutElse.Build()
//	}
type ConditionWidget struct {
	cond       bool
	layoutIf   Layout
	layoutElse Layout
}

// Condition creates new COnditionWidget.
func Condition(cond bool, layoutIf, layoutElse Layout) *ConditionWidget {
	return &ConditionWidget{
		cond:       cond,
		layoutIf:   layoutIf,
		layoutElse: layoutElse,
	}
}

// Range implements extra abilities (see Splittablle).
func (c *ConditionWidget) Range(rangeFunc func(w Widget)) {
	var l Layout
	if c.cond {
		l = c.layoutIf
	} else {
		l = c.layoutElse
	}

	if l != nil {
		l.Range(rangeFunc)
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

type listBoxState struct {
	selectedIndex int32
}

func (s *listBoxState) Dispose() {
	// Nothing to do here.
}

var _ Widget = &ListBoxWidget{}

// ListBoxWidget is a field with selectable items (Child with Selectables).
type ListBoxWidget struct {
	selectedIndex *int32
	id            string
	width         float32
	height        float32
	border        bool
	items         []string
	menus         []string
	onChange      func(selectedIndex int)
	onDClick      func(selectedIndex int)
	onMenu        func(selectedIndex int, menu string)
}

// ListBox creates new ListBoxWidget.
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

func (l *ListBoxWidget) SelectedIndex(i *int32) *ListBoxWidget {
	l.selectedIndex = i
	return l
}

// Size sets size of the box.
func (l *ListBoxWidget) Size(width, height float32) *ListBoxWidget {
	l.width, l.height = width, height
	return l
}

// Border sets whether box should have border (see Child().Border(...).
func (l *ListBoxWidget) Border(b bool) *ListBoxWidget {
	l.border = b
	return l
}

// ContextMenu adds item in context menu which is opened when user right-click on item.
func (l *ListBoxWidget) ContextMenu(menuItems []string) *ListBoxWidget {
	l.menus = menuItems
	return l
}

// OnChange sets callback called when user changes their selection.
func (l *ListBoxWidget) OnChange(onChange func(selectedIndex int)) *ListBoxWidget {
	l.onChange = onChange
	return l
}

// OnDClick sets callback on double click.
func (l *ListBoxWidget) OnDClick(onDClick func(selectedIndex int)) *ListBoxWidget {
	l.onDClick = onDClick
	return l
}

// OnMenu sets callback called when context menu item clicked.
func (l *ListBoxWidget) OnMenu(onMenu func(selectedIndex int, menu string)) *ListBoxWidget {
	l.onMenu = onMenu
	return l
}

// Build implements Widget interface
//
//nolint:gocognit // will fix later
func (l *ListBoxWidget) Build() {
	selectedIndex := l.selectedIndex
	if selectedIndex == nil {
		var state *listBoxState
		if state = GetState[listBoxState](Context, l.id); state == nil {
			state = &listBoxState{selectedIndex: 0}
			SetState(Context, l.id, state)
		}

		selectedIndex = &state.selectedIndex
	}

	child := Child().Border(l.border).Size(l.width, l.height).Layout(Layout{
		Custom(func() {
			clipper := imgui.NewListClipper()
			defer clipper.Delete()

			clipper.Begin(len(l.items))

			for clipper.Step() {
				for i := clipper.DisplayStart(); i < clipper.DisplayEnd(); i++ {
					selected := i == int(*selectedIndex)
					item := l.items[i]
					Selectable(item).Selected(selected).Flags(SelectableFlagsAllowDoubleClick).OnClick(func() {
						if *selectedIndex != int32(i) {
							*selectedIndex = int32(i)
							if l.onChange != nil {
								l.onChange(i)
							}
						}
					}).Build()

					if IsItemHovered() && IsMouseDoubleClicked(MouseButtonLeft) && l.onDClick != nil {
						l.onDClick(int(*selectedIndex))
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

// DatePickerWidget is a simple Calender widget.
// It allow user to select a day and convert it to time.Time go type.
// It consists of a Combo widget which (after opening) contains
// a calender-like table.
type DatePickerWidget struct {
	id          string
	date        *time.Time
	width       float32
	onChange    func()
	format      string
	startOfWeek time.Weekday
}

// DatePicker creates new DatePickerWidget.
func DatePicker(id string, date *time.Time) *DatePickerWidget {
	return &DatePickerWidget{
		id:          GenAutoID(id),
		date:        date,
		width:       100,
		startOfWeek: time.Sunday,
		onChange:    func() {}, // small hack - prevent giu from setting nil cb (skip nil check later)
	}
}

// Size sets combo widget's size.
func (d *DatePickerWidget) Size(width float32) *DatePickerWidget {
	d.width = width
	return d
}

// OnChange sets callback called when date is changed.
func (d *DatePickerWidget) OnChange(onChange func()) *DatePickerWidget {
	if onChange != nil {
		d.onChange = onChange
	}

	return d
}

// Format sets date format of displayed (in combo) date.
// Compatible with (time.Time).Format(...)
// Default: "2006-01-02".
func (d *DatePickerWidget) Format(format string) *DatePickerWidget {
	d.format = format
	return d
}

// StartOfWeek sets first day of the week
// Default: Sunday.
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
			Label(Context.FontAtlas.RegisterString(" Year")),
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
