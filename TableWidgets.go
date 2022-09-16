package giu

import (
	"image/color"

	imgui "github.com/AllenDang/cimgui-go"
)

type TableRowWidget struct {
	flags        imgui.ImGuiTableRowFlags
	minRowHeight float32
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

func (r *TableRowWidget) Flags(flags imgui.ImGuiTableRowFlags) *TableRowWidget {
	r.flags = flags
	return r
}

func (r *TableRowWidget) MinHeight(height float32) *TableRowWidget {
	r.minRowHeight = height
	return r
}

// BuildTableRow executes table row build steps.
func (r *TableRowWidget) BuildTableRow() {
	imgui.TableNextRowV(r.flags, r.minRowHeight)

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
		imgui.TableSetBgColorV(imgui.ImGuiTableBgTarget_RowBg0, uint32(imgui.GetColorU32_Vec4(ToVec4Color(r.bgColor))), -1)
	}
}

type TableColumnWidget struct {
	label              string
	flags              imgui.ImGuiTableColumnFlags
	innerWidthOrWeight float32
	userID             imgui.ImGuiID
}

func TableColumn(label string) *TableColumnWidget {
	return &TableColumnWidget{
		label:              Context.FontAtlas.RegisterString(label),
		flags:              0,
		innerWidthOrWeight: 0,
		userID:             0,
	}
}

func (c *TableColumnWidget) Flags(flags imgui.ImGuiTableColumnFlags) *TableColumnWidget {
	c.flags = flags
	return c
}

func (c *TableColumnWidget) InnerWidthOrWeight(w float32) *TableColumnWidget {
	c.innerWidthOrWeight = w
	return c
}

func (c *TableColumnWidget) UserID(id uint32) *TableColumnWidget {
	c.userID = imgui.ImGuiID(id)
	return c
}

// BuildTableColumn executes table column build steps.
func (c *TableColumnWidget) BuildTableColumn() {
	imgui.TableSetupColumnV(c.label, imgui.ImGuiTableColumnFlags(c.flags), c.innerWidthOrWeight, c.userID)
}

var _ Widget = &TableWidget{}

type TableWidget struct {
	id           string
	flags        imgui.ImGuiTableFlags
	size         imgui.ImVec2
	innerWidth   float32
	rows         []*TableRowWidget
	columns      []*TableColumnWidget
	fastMode     bool
	freezeRow    int32
	freezeColumn int32
}

func Table() *TableWidget {
	return &TableWidget{
		id:           GenAutoID("Table"),
		flags:        imgui.ImGuiTableFlags_Resizable | imgui.ImGuiTableFlags_Borders | imgui.ImGuiTableFlags_ScrollY,
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
func (t *TableWidget) Freeze(col, row int32) *TableWidget {
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
	t.size = imgui.ImVec2{X: width, Y: height}
	return t
}

func (t *TableWidget) InnerWidth(width float32) *TableWidget {
	t.innerWidth = width
	return t
}

func (t *TableWidget) Flags(flags imgui.ImGuiTableFlags) *TableWidget {
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

	if imgui.BeginTableV(t.id, int32(colCount), t.flags, t.size, t.innerWidth) {
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
			clipper := imgui.NewImGuiListClipper()
			defer clipper.Destroy()

			clipper.Begin(int32(len(t.rows)))

			for clipper.Step() {
				for i := clipper.GetDisplayStart(); i < clipper.GetDisplayEnd(); i++ {
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
