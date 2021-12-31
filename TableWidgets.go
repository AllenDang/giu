package giu

import (
	"image/color"

	"github.com/AllenDang/imgui-go"
)

type TableRowWidget struct {
	flags        TableRowFlags
	minRowHeight float64
	layout       Layout
	bgColor      color.Color
}

func TableRow(widgets ...Widget) *TableRowWidget {
	return &TableRowWidget{
		flags:        0,
		minRowHeight: 0,
		layout:       widgets,
		bgColor:      nil,
	}
}

func (r *TableRowWidget) BgColor(c color.Color) *TableRowWidget {
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

// BuildTableRow executes table row build steps.
func (r *TableRowWidget) BuildTableRow() {
	imgui.TableNextRow(imgui.TableRowFlags(r.flags), r.minRowHeight)

	for _, w := range r.layout {
		switch w.(type) {
		case *TooltipWidget,
			*ContextMenuWidget, *PopupModalWidget:
			// noop
		default:
			imgui.TableNextColumn()
		}

		w.Build()
	}

	if r.bgColor != nil {
		imgui.TableSetBgColor(imgui.TableBgTarget_RowBg0, uint32(imgui.GetColorU32(ToVec4Color(r.bgColor))), -1)
	}
}

type TableColumnWidget struct {
	label              string
	flags              TableColumnFlags
	innerWidthOrWeight float32
	userID             uint32
}

func TableColumn(label string) *TableColumnWidget {
	return &TableColumnWidget{
		label:              tStr(label),
		flags:              0,
		innerWidthOrWeight: 0,
		userID:             0,
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

func (c *TableColumnWidget) UserID(id uint32) *TableColumnWidget {
	c.userID = id
	return c
}

// BuildTableColumn executes table column build steps.
func (c *TableColumnWidget) BuildTableColumn() {
	imgui.TableSetupColumn(c.label, imgui.TableColumnFlags(c.flags), c.innerWidthOrWeight, c.userID)
}

var _ Widget = &TableWidget{}

type TableWidget struct {
	id           string
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
		id:           GenAutoID("Table"),
		flags:        TableFlagsResizable | TableFlagsBorders | TableFlagsScrollY,
		rows:         nil,
		columns:      nil,
		fastMode:     false,
		freezeRow:    -1,
		freezeColumn: -1,
	}
}

// ID sets the internal id of table widget.
func (t *TableWidget) ID(id string) *TableWidget {
	t.id = id
	return t
}

// FastMode Displays visible rows only to boost performance.
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

// Build implements Widget interface.
func (t *TableWidget) Build() {
	if len(t.rows) == 0 {
		return
	}

	colCount := len(t.columns)
	if colCount == 0 {
		colCount = len(t.rows[0].layout)
	}

	if imgui.BeginTable(t.id, colCount, imgui.TableFlags(t.flags), t.size, t.innerWidth) {
		if t.freezeColumn >= 0 && t.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(t.freezeColumn, t.freezeRow)
		}

		if len(t.columns) > 0 {
			for _, col := range t.columns {
				col.BuildTableColumn()
			}
			imgui.TableHeadersRow()
		}

		if t.fastMode {
			clipper := imgui.NewListClipper()
			defer clipper.Delete()

			clipper.Begin(len(t.rows))

			for clipper.Step() {
				for i := clipper.DisplayStart(); i < clipper.DisplayEnd(); i++ {
					row := t.rows[i]
					row.BuildTableRow()
				}
			}

			clipper.End()
		} else {
			for _, row := range t.rows {
				row.BuildTableRow()
			}
		}

		imgui.EndTable()
	}
}
