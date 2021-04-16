package imgui

// #cgo CXXFLAGS: -std=c++11
// #include "Table.h"
import "C"

type TableFlags int

const (
	// Features
	TableFlags_None              TableFlags = 0
	TableFlags_Resizable         TableFlags = 1 << 0 // Enable resizing columns.
	TableFlags_Reorderable       TableFlags = 1 << 1 // Enable reordering columns in header row (need calling TableSetupColumn() + TableHeadersRow() to display headers)
	TableFlags_Hideable          TableFlags = 1 << 2 // Enable hiding/disabling columns in context menu.
	TableFlags_Sortable          TableFlags = 1 << 3 // Enable sorting. Call TableGetSortSpecs() to obtain sort specs. Also see TableFlags_SortMulti and TableFlags_SortTristate.
	TableFlags_NoSavedSettings   TableFlags = 1 << 4 // Disable persisting columns order, width and sort settings in the .ini file.
	TableFlags_ContextMenuInBody TableFlags = 1 << 5 // Right-click on columns body/contents will display table context menu. By default it is available in TableHeadersRow().
	// Decorations
	TableFlags_RowBg                                TableFlags = 1 << 6                                              // Set each RowBg color with Col_TableRowBg or Col_TableRowBgAlt (equivalent of calling TableSetBgColor with TableBgFlags_RowBg0 on each row manually)
	TableFlags_BordersInnerH                        TableFlags = 1 << 7                                              // Draw horizontal borders between rows.
	TableFlags_BordersOuterH                        TableFlags = 1 << 8                                              // Draw horizontal borders at the top and bottom.
	TableFlags_BordersInnerV                        TableFlags = 1 << 9                                              // Draw vertical borders between columns.
	TableFlags_BordersOuterV                        TableFlags = 1 << 10                                             // Draw vertical borders on the left and right sides.
	TableFlags_BordersH                             TableFlags = TableFlags_BordersInnerH | TableFlags_BordersOuterH // Draw horizontal borders.
	TableFlags_BordersV                             TableFlags = TableFlags_BordersInnerV | TableFlags_BordersOuterV // Draw vertical borders.
	TableFlags_BordersInner                         TableFlags = TableFlags_BordersInnerV | TableFlags_BordersInnerH // Draw inner borders.
	TableFlags_BordersOuter                         TableFlags = TableFlags_BordersOuterV | TableFlags_BordersOuterH // Draw outer borders.
	TableFlags_Borders                              TableFlags = TableFlags_BordersInner | TableFlags_BordersOuter   // Draw all borders.
	TableFlags_NoBordersInBody                      TableFlags = 1 << 11                                             // [ALPHA] Disable vertical borders in columns Body (borders will always appears in Headers). -> May move to style
	TableFlags_NoBordersInBodyUntilResizeTableFlags            = 1 << 12                                             // [ALPHA] Disable vertical borders in columns Body until hovered for resize (borders will always appears in Headers). -> May move to style
	// Sizing Policy (read above for defaults)TableFlags
	TableFlags_SizingFixedFit    TableFlags = 1 << 13 // Columns default to _WidthFixed or _WidthAuto (if resizable or not resizable), matching contents width.
	TableFlags_SizingFixedSame   TableFlags = 2 << 13 // Columns default to _WidthFixed or _WidthAuto (if resizable or not resizable), matching the maximum contents width of all columns. Implicitly enable TableFlags_NoKeepColumnsVisible.
	TableFlags_SizingStretchProp TableFlags = 3 << 13 // Columns default to _WidthStretch with default weights proportional to each columns contents widths.
	TableFlags_SizingStretchSame TableFlags = 4 << 13 // Columns default to _WidthStretch with default weights all equal, unless overriden by TableSetupColumn().
	// Sizing Extra Options
	TableFlags_NoHostExtendX        TableFlags = 1 << 16 // Make outer width auto-fit to columns, overriding outer_size.x value. Only available when ScrollX/ScrollY are disabled and Stretch columns are not used.
	TableFlags_NoHostExtendY        TableFlags = 1 << 17 // Make outer height stop exactly at outer_size.y (prevent auto-extending table past the limit). Only available when ScrollX/ScrollY are disabled. Data below the limit will be clipped and not visible.
	TableFlags_NoKeepColumnsVisible TableFlags = 1 << 18 // Disable keeping column always minimally visible when ScrollX is off and table gets too small. Not recommended if columns are resizable.
	TableFlags_PreciseWidths        TableFlags = 1 << 19 // Disable distributing remainder width to stretched columns (width allocation on a 100-wide table with 3 columns: Without this flag: 33,33,34. With this flag: 33,33,33). With larger number of columns, resizing will appear to be less smooth.
	// Clipping
	TableFlags_NoClip TableFlags = 1 << 20 // Disable clipping rectangle for every individual columns (reduce draw command count, items will be able to overflow into other columns). Generally incompatible with TableSetupScrollFreeze().
	// Padding
	TableFlags_PadOuterX   TableFlags = 1 << 21 // Default if BordersOuterV is on. Enable outer-most padding. Generally desirable if you have headers.
	TableFlags_NoPadOuterX TableFlags = 1 << 22 // Default if BordersOuterV is off. Disable outer-most padding.
	TableFlags_NoPadInnerX TableFlags = 1 << 23 // Disable inner padding between columns (double inner padding if BordersOuterV is on, single inner padding if BordersOuterV is off).
	// Scrolling
	TableFlags_ScrollX TableFlags = 1 << 24 // Enable horizontal scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size. Changes default sizing policy. Because this create a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlags_ScrollY TableFlags = 1 << 25 // Enable vertical scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size.
	// Sorting
	TableFlags_SortMulti    TableFlags = 1 << 26 // Hold shift when clicking headers to sort on multiple column. TableGetSortSpecs() may return specs where (SpecsCount > 1).
	TableFlags_SortTristate TableFlags = 1 << 27 // Allow no sorting, disable default sorting. TableGetSortSpecs() may return specs where (SpecsCount == 0).

	// [Internal] Combinations and masks
	TableFlags_SizingMask_ = TableFlags_SizingFixedFit | TableFlags_SizingFixedSame | TableFlags_SizingStretchProp | TableFlags_SizingStretchSame
)

type TableColumnFlags int

const (
	// Input configuration flags
	TableColumnFlags_None                 TableColumnFlags = 0
	TableColumnFlags_DefaultHide          TableColumnFlags = 1 << 0  // Default as a hidden/disabled column.
	TableColumnFlags_DefaultSort          TableColumnFlags = 1 << 1  // Default as a sorting column.
	TableColumnFlags_WidthStretch         TableColumnFlags = 1 << 2  // Column will stretch. Preferable with horizontal scrolling disabled (default if table sizing policy is _SizingStretchSame or _SizingStretchProp).
	TableColumnFlags_WidthFixed           TableColumnFlags = 1 << 3  // Column will not stretch. Preferable with horizontal scrolling enabled (default if table sizing policy is _SizingFixedFit and table is resizable).
	TableColumnFlags_NoResize             TableColumnFlags = 1 << 4  // Disable manual resizing.
	TableColumnFlags_NoReorder            TableColumnFlags = 1 << 5  // Disable manual reordering this column, this will also prevent other columns from crossing over this column.
	TableColumnFlags_NoHide               TableColumnFlags = 1 << 6  // Disable ability to hide/disable this column.
	TableColumnFlags_NoClip               TableColumnFlags = 1 << 7  // Disable clipping for this column (all NoClip columns will render in a same draw command).
	TableColumnFlags_NoSort               TableColumnFlags = 1 << 8  // Disable ability to sort on this field (even if TableFlags_Sortable is set on the table).
	TableColumnFlags_NoSortAscending      TableColumnFlags = 1 << 9  // Disable ability to sort in the ascending direction.
	TableColumnFlags_NoSortDescending     TableColumnFlags = 1 << 10 // Disable ability to sort in the descending direction.
	TableColumnFlags_NoHeaderWidth        TableColumnFlags = 1 << 11 // Disable header text width contribution to automatic column width.
	TableColumnFlags_PreferSortAscending  TableColumnFlags = 1 << 12 // Make the initial sort direction Ascending when first sorting on this column (default).
	TableColumnFlags_PreferSortDescending TableColumnFlags = 1 << 13 // Make the initial sort direction Descending when first sorting on this column.
	TableColumnFlags_IndentEnable         TableColumnFlags = 1 << 14 // Use current Indent value when entering cell (default for column 0).
	TableColumnFlags_IndentDisable        TableColumnFlags = 1 << 15 // Ignore current Indent value when entering cell (default for columns > 0). Indentation changes _within_ the cell will still be honored.

	// Output status flags read-only via TableGetColumnFlags()
	TableColumnFlags_IsEnabled TableColumnFlags = 1 << 20 // Status: is enabled == not hidden by user/api (referred to as "Hide" in _DefaultHide and _NoHide) flags.
	TableColumnFlags_IsVisible TableColumnFlags = 1 << 21 // Status: is visible == is enabled AND not clipped by scrolling.
	TableColumnFlags_IsSorted  TableColumnFlags = 1 << 22 // Status: is currently part of the sort specs
	TableColumnFlags_IsHovered TableColumnFlags = 1 << 23 // Status: is hovered by mouse

	// [Internal] Combinations and masks
	TableColumnFlags_WidthMask_      TableColumnFlags = TableColumnFlags_WidthStretch | TableColumnFlags_WidthFixed
	TableColumnFlags_IndentMask_     TableColumnFlags = TableColumnFlags_IndentEnable | TableColumnFlags_IndentDisable
	TableColumnFlags_StatusMask_     TableColumnFlags = TableColumnFlags_IsEnabled | TableColumnFlags_IsVisible | TableColumnFlags_IsSorted | TableColumnFlags_IsHovered
	TableColumnFlags_NoDirectResize_ TableColumnFlags = 1 << 30 // [Internal] Disable user resizing this column directly (it may however we resized indirectly from its left edge)
)

type TableRowFlags int

const (
	TableRowFlags_None    TableRowFlags = 0
	TableRowFlags_Headers TableRowFlags = 1 << 0 // Identify header row (set default background color + width of its contents accounted different for auto column width)
)

type TableBgTarget int

const (
	TableBgTarget_None   TableBgTarget = 0
	TableBgTarget_RowBg0 TableBgTarget = 1 // Set row background color 0 (generally used for background, automatically set when TableFlags_RowBg is used)
	TableBgTarget_RowBg1 TableBgTarget = 2 // Set row background color 1 (generally used for selection marking)
	TableBgTarget_CellBg TableBgTarget = 3 // Set cell background color (top-most color)
)

func BeginTable(id string, column int, flags TableFlags, outerSize Vec2, innerWidth float64) bool {
	idArg, idDeleter := wrapString(id)
	defer idDeleter()

	outerSizeArg, _ := outerSize.wrapped()

	return C.iggBeginTable(idArg, C.int(column), C.int(flags), outerSizeArg, C.float(innerWidth)) != 0
}

func EndTable() {
	C.iggEndTable()
}

func TableNextRow(rowFlags TableRowFlags, minRowHeight float64) {
	C.iggTableNextRow(C.int(rowFlags), C.float(minRowHeight))
}

func TableNextColumn() {
	C.iggTableNextColumn()
}

func TableSetColumnIndex(columnN int) bool {
	return C.iggTableSetColumnIndex(C.int(columnN)) != 0
}

func TableSetupColumn(label string, flags TableColumnFlags, initWidthOrWeight float32, userId uint32) {
	labelArg, labelDeleter := wrapString(label)
	defer labelDeleter()

	C.iggTableSetupColumn(labelArg, C.int(flags), C.float(initWidthOrWeight), C.uint(userId))
}

func TableSetupScrollFreeze(cols, rows int) {
	C.iggTableSetupScrollFreeze(C.int(cols), C.int(rows))
}

func TableHeadersRow() {
	C.iggTableHeadersRow()
}

func TableHeader(label string) {
	labelArg, labelDeleter := wrapString(label)
	defer labelDeleter()

	C.iggTableHeader(labelArg)
}

func TableGetColumnCount() int {
	return int(C.iggTableGetColumnCount())
}

func TableGetColumnIndex() int {
	return int(C.iggTableGetColumnIndex())
}

func TableGetRowIndex() int {
	return int(C.iggTableGetRowIndex())
}

func TableGetColumnName(columnN int) string {
	return C.GoString(C.iggTableGetColumnName(C.int(columnN)))
}

func TableSetBgColor(target TableBgTarget, color uint32, columnN int) {
	C.iggTableSetBgColor(C.int(target), C.uint(color), C.int(columnN))
}

func TableGetColumnFlags(columnN int) TableColumnFlags {
	return TableColumnFlags(C.iggTableGetColumnFlags(C.int(columnN)))
}
