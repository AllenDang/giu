package giu

import (
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
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
	imgui.TableNextRowV(imgui.TableRowFlags(r.flags), float32(r.minRowHeight))

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
		imgui.TableSetBgColorV(imgui.TableBgTargetRowBg0, imgui.ColorU32Vec4(ToVec4Color(r.bgColor)), -1)
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
		label:              Context.FontAtlas.RegisterString(label),
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
	imgui.TableSetupColumnV(c.label, imgui.TableColumnFlags(c.flags), c.innerWidthOrWeight, imgui.ID(c.userID))
}

var _ Widget = &TableWidget{}

type TableWidget struct {
	id           ID
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
func (t *TableWidget) ID(id ID) *TableWidget {
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
	colCount := len(t.columns)
	if colCount == 0 {
		if len(t.rows) > 0 {
			colCount = len(t.rows[0].layout)
		} else {
			// No rows or columns, pass a single column to BeginTable
			colCount = 1
		}
	}

	if imgui.BeginTableV(t.id.String(), int32(colCount), imgui.TableFlags(t.flags), t.size, float32(t.innerWidth)) {
		if t.freezeColumn >= 0 && t.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(int32(t.freezeColumn), int32(t.freezeRow))
		}

		if len(t.columns) > 0 {
			for _, col := range t.columns {
				col.BuildTableColumn()
			}

			imgui.TableHeadersRow()
		}

		if t.fastMode {
			clipper := imgui.NewListClipper()
			defer clipper.Destroy()

			clipper.Begin(int32(len(t.rows)))

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

// TreeTableRowWidget is a row in TreeTableWidget.
type TreeTableRowWidget struct {
	label    ID
	flags    TreeNodeFlags
	layout   Layout
	children []*TreeTableRowWidget
}

// TreeTableRow creates new TreeTableRowWidget.
func TreeTableRow(label string, widgets ...Widget) *TreeTableRowWidget {
	return &TreeTableRowWidget{
		label:  GenAutoID(label),
		layout: widgets,
	}
}

// Children sets child rows of tree row.
func (ttr *TreeTableRowWidget) Children(rows ...*TreeTableRowWidget) *TreeTableRowWidget {
	ttr.children = rows
	return ttr
}

// Flags sets row's flags.
func (ttr *TreeTableRowWidget) Flags(flags TreeNodeFlags) *TreeTableRowWidget {
	ttr.flags = flags
	return ttr
}

// BuildTreeTableRow executes table row building steps.
func (ttr *TreeTableRowWidget) BuildTreeTableRow() {
	imgui.TableNextRowV(0, 0)
	imgui.TableNextColumn()

	open := false
	if len(ttr.children) > 0 {
		open = imgui.TreeNodeExStrV(Context.FontAtlas.RegisterString(ttr.label.String()), imgui.TreeNodeFlags(ttr.flags))
	} else {
		ttr.flags |= TreeNodeFlagsLeaf | TreeNodeFlagsNoTreePushOnOpen
		imgui.TreeNodeExStrV(Context.FontAtlas.RegisterString(ttr.label.String()), imgui.TreeNodeFlags(ttr.flags))
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

// TreeTableWidget is a table that consists of TreeNodeWidgets.
type TreeTableWidget struct {
	id           ID
	flags        TableFlags
	size         imgui.Vec2
	columns      []*TableColumnWidget
	rows         []*TreeTableRowWidget
	freezeRow    int
	freezeColumn int
}

// TreeTable creates new TreeTableWidget.
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

// Size sets size of the table.
func (tt *TreeTableWidget) Size(width, height float32) *TreeTableWidget {
	tt.size = imgui.Vec2{X: width, Y: height}
	return tt
}

// Flags sets table flags.
func (tt *TreeTableWidget) Flags(flags TableFlags) *TreeTableWidget {
	tt.flags = flags
	return tt
}

// Columns sets table's columns.
func (tt *TreeTableWidget) Columns(cols ...*TableColumnWidget) *TreeTableWidget {
	tt.columns = cols
	return tt
}

// Rows sets TreeTable rows.
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

	if imgui.BeginTableV(tt.id.String(), int32(colCount), imgui.TableFlags(tt.flags), tt.size, 0) {
		if tt.freezeColumn >= 0 && tt.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(int32(tt.freezeColumn), int32(tt.freezeRow))
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
