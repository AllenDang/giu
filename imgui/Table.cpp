#include "imguiWrappedHeader.h"
#include "Table.h"
#include "WrapperConverter.h"

IggBool iggBeginTable(const char *str_id, int column, int flags,
                      IggVec2 const *outer_size, float inner_width) {
  Vec2Wrapper outerSizeArg(outer_size);
  return ImGui::BeginTable(str_id, column, flags, *outerSizeArg, inner_width) ? 1 : 0;
}

void iggEndTable() { ImGui::EndTable(); }

void iggTableNextRow(int row_flags, float min_row_height) {
  ImGui::TableNextRow(row_flags, min_row_height);
}

IggBool iggTableNextColumn() { return ImGui::TableNextColumn() ? 1 : 0; }

IggBool iggTableSetColumnIndex(int column_n) {
  return ImGui::TableSetColumnIndex(column_n) ? 1 : 0;
}

void iggTableSetupColumn(const char *label, int flags,
                         float init_width_or_weight, unsigned int user_id) {
  ImGui::TableSetupColumn(label, flags, init_width_or_weight, user_id);
}

void iggTableSetupScrollFreeze(int cols, int rows) {
  ImGui::TableSetupScrollFreeze(cols, rows);
}

void iggTableHeadersRow() { ImGui::TableHeadersRow(); }

void iggTableHeader(const char *label) { ImGui::TableHeader(label); }

IggImGuiTableSortSpecs *iggTableGetSortSpecs() {
  ImGuiTableSortSpecs *specs = ImGui::TableGetSortSpecs();
  return reinterpret_cast<IggImGuiTableSortSpecs *>(specs);
}

int iggTableGetColumnCount() { return ImGui::TableGetColumnCount(); }

int iggTableGetColumnIndex() { return ImGui::TableGetColumnIndex(); }

int iggTableGetRowIndex() { return ImGui::TableGetRowIndex(); }

const char *iggTableGetColumnName(int column_n) {
  return ImGui::TableGetColumnName(column_n);
}

void iggTableSetBgColor(int target, unsigned int color, int column_n) {
  ImGui::TableSetBgColor(target, color, column_n);
}

int iggTableGetColumnFlags(int column_n) {
  return ImGui::TableGetColumnFlags(column_n);
}
