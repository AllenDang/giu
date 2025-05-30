package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/imgui"
)

// SortDirection tells how the data are sorted (ascending/descending).
type SortDirection byte

// Possible sort directions.
const (
	SortAscending  SortDirection = 1
	SortDescending SortDirection = 2
)

// TableRowWidget represents a row in a table.
type TableRowWidget struct {
	flags        TableRowFlags
	minRowHeight float64
	layout       Layout
	bgColor      color.Color
}

// TableRow creates a TbleRowWidget.
// Each widget will be rendered in a separated column.
// NOTE: if you want to put multiple widgets in one cell, enclose them in Layout{}.
func TableRow(widgets ...Widget) *TableRowWidget {
	return &TableRowWidget{
		flags:        0,
		minRowHeight: 0,
		layout:       widgets,
		bgColor:      nil,
	}
}

// BgColor sets the background color of the row.
func (r *TableRowWidget) BgColor(c color.Color) *TableRowWidget {
	r.bgColor = c
	return r
}

// Flags sets the flags of the row.
func (r *TableRowWidget) Flags(flags TableRowFlags) *TableRowWidget {
	r.flags = flags
	return r
}

// MinHeight sets the minimum height of the row.
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

// TableColumnWidget allows to configure table columns headers.
type TableColumnWidget struct {
	label              string
	flags              TableColumnFlags
	innerWidthOrWeight float32
	userID             uint32
	sortFn             func(SortDirection)
}

// TableColumn creates a new TableColumnWidget.
func TableColumn(label string) *TableColumnWidget {
	return &TableColumnWidget{
		label:              Context.FontAtlas.RegisterString(label),
		flags:              0,
		innerWidthOrWeight: 0,
		userID:             0,
	}
}

// Flags sets the flags of the column.
func (c *TableColumnWidget) Flags(flags TableColumnFlags) *TableColumnWidget {
	c.flags = flags
	return c
}

// InnerWidthOrWeight sets the inner width or weight of the column.
func (c *TableColumnWidget) InnerWidthOrWeight(w float32) *TableColumnWidget {
	c.innerWidthOrWeight = w
	return c
}

// UserID sets the user id of the column.
func (c *TableColumnWidget) UserID(id uint32) *TableColumnWidget {
	c.userID = id
	return c
}

// Sort allows you to set Sort function for that column. I talso could be used to detect click event.
func (c *TableColumnWidget) Sort(s func(SortDirection)) *TableColumnWidget {
	c.sortFn = s
	return c
}

// BuildTableColumn executes table column build steps.
func (c *TableColumnWidget) BuildTableColumn() {
	imgui.TableSetupColumnV(c.label, imgui.TableColumnFlags(c.flags), c.innerWidthOrWeight, imgui.ID(c.userID))
}

var _ Widget = &TableWidget{}

// TableWidget is a table widget.
// - Call Table to create new
// - Then use Rows method to add content
// - Use Columns method to configure columns (optional).
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
	noHeader     bool
}

// Table creates new TableWidget.
func Table() *TableWidget {
	return &TableWidget{
		id:           GenAutoID("Table"),
		flags:        TableFlagsResizable | TableFlagsBorders | TableFlagsScrollY,
		rows:         nil,
		columns:      nil,
		fastMode:     false,
		freezeRow:    -1,
		freezeColumn: -1,
		noHeader:     false,
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

// NoHeader indicates that the column header should not be shown. This allows
// the use of the Columns() function to configure table columns (eg. column
// width) but without showing the table header.
func (t *TableWidget) NoHeader(b bool) *TableWidget {
	t.noHeader = b
	return t
}

// Freeze columns/rows so they stay visible when scrolled.
func (t *TableWidget) Freeze(col, row int) *TableWidget {
	t.freezeColumn = col
	t.freezeRow = row

	return t
}

// Columns adds a list of column widgets to be used in the table. Columns added
// with this function will cause the table header to be shown. If the table
// header is not required then the NoHeader() function can be used.
func (t *TableWidget) Columns(cols ...*TableColumnWidget) *TableWidget {
	t.columns = cols
	return t
}

// Rows sets the rows of the table.
func (t *TableWidget) Rows(rows ...*TableRowWidget) *TableWidget {
	t.rows = rows
	return t
}

// Size sets the size of the table.
func (t *TableWidget) Size(width, height float32) *TableWidget {
	t.size = imgui.Vec2{X: width, Y: height}
	return t
}

// InnerWidth sets the inner width of the table.
func (t *TableWidget) InnerWidth(width float64) *TableWidget {
	t.innerWidth = width
	return t
}

// Flags sets the flags of the table.
func (t *TableWidget) Flags(flags TableFlags) *TableWidget {
	t.flags = flags
	return t
}

// helper function to find out number of columns in the table.
func (t *TableWidget) colCount() int {
	colCount := len(t.columns)

	if colCount == 0 {
		if len(t.rows) > 0 {
			return len(t.rows[0].layout)
		}

		// No rows or columns, pass a single column to BeginTable
		return 1
	}

	return colCount
}

func (t *TableWidget) handleSort() {
	if specs := imgui.TableGetSortSpecs(); specs != nil {
		if specs.SpecsDirty() {
			// Evil bithack - we assume that array==pointer, so specs.Specs() points to the first element of that array.
			cs := specs.Specs() // this in fact is []TableColumnSortSpecs but should be also (*TableColumnSortSpecs)
			colIdx := cs.ColumnIndex()
			sortDir := cs.SortDirection()

			if col := t.columns[colIdx]; col.sortFn != nil {
				col.sortFn(SortDirection(sortDir))
			}

			specs.SetSpecsDirty(false)
		}
	}
}

// Build implements Widget interface.
func (t *TableWidget) Build() {
	if imgui.BeginTableV(t.id.String(), int32(t.colCount()), imgui.TableFlags(t.flags), t.size, float32(t.innerWidth)) {
		if t.freezeColumn >= 0 && t.freezeRow >= 0 {
			imgui.TableSetupScrollFreeze(int32(t.freezeColumn), int32(t.freezeRow))
		}

		if len(t.columns) > 0 {
			for _, col := range t.columns {
				col.BuildTableColumn()
			}

			if !t.noHeader {
				imgui.TableHeadersRow()
			}

			if t.flags&TableFlags(imgui.TableFlagsSortable) != 0 {
				t.handleSort()
			}
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
