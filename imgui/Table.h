#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C" {
#endif

IggBool iggBeginTable(const char *str_id, int column, int flags,
                      IggVec2 const *outer_size, float inner_width);
void iggEndTable(); // only call EndTable() if BeginTable() returns true!
void iggTableNextRow(
    int row_flags,
    float min_row_height);    // append into the first cell of a new row.
IggBool iggTableNextColumn(); // append into the next column (or first column of
                              // next row if currently in last column). Return
                              // true when column is visible.
IggBool
iggTableSetColumnIndex(int column_n); // append into the specified column.
                                      // Return true when column is visible.

void iggTableSetupColumn(const char *label, int flags,
                         float init_width_or_weight, unsigned int user_id);
void iggTableSetupScrollFreeze(
    int cols,
    int rows); // lock columns/rows so they stay visible when scrolled.
void iggTableHeadersRow(); // submit all headers cells based on data provided to
                           // TableSetupColumn() + submit context menu
void iggTableHeader(
    const char *label); // submit one header cell manually (rarely used)

typedef void *IggImGuiTableSortSpecs;

IggImGuiTableSortSpecs *iggTableGetSortSpecs(); // get latest sort specs for the
                                                // table (NULL if not sorting).

int iggTableGetColumnCount(); // return number of columns (value passed to
                              // BeginTable)
int iggTableGetColumnIndex(); // return current column index.
int iggTableGetRowIndex();    // return current row index.
const char *iggTableGetColumnName(
    int column_n); // return "" if column didn't have a name declared by
                   // TableSetupColumn(). Pass -1 to use current column.
void iggTableSetBgColor(
    int target, unsigned int color,
    int column_n); // change the color of a cell, row, or column. See
                   // ImGuiTableBgTarget_ flags for details.
int iggTableGetColumnFlags(
    int column_n); // return column flags so you can query their
                   // Enabled/Visible/Sorted/Hovered status flags. Pass -1 to
                   // use current column.

#ifdef __cplusplus
}
#endif
