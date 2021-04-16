#include "imguiWrappedHeader.h"
#include "DragDropWrapper.h"

void* iggPayloadGetData(IggPayload handle) {
  ImGuiPayload *payload = reinterpret_cast<ImGuiPayload*>(handle);
  return payload->Data;
}

IggBool iggBeginDragDropSource(int flags) {
  return ImGui::BeginDragDropSource(flags) != 0 ? 1 : 0;
}

IggBool iggSetDragDropPayload(const char* type, const void* data, unsigned int sz, int cond) {
  return ImGui::SetDragDropPayload(type, data, sz, cond) != 0 ? 1 : 0;
}

void iggEndDragDropSource() {
  ImGui::EndDragDropSource();
}

IggBool iggBeginDragDropTarget() {
  return ImGui::BeginDragDropTarget() != 0 ? 1 : 0;
}

IggPayload iggAcceptDragDropPayload(const char *type, int flags) {
  const ImGuiPayload *payload = ImGui::AcceptDragDropPayload(type, flags);
  return static_cast<IggPayload>(const_cast<ImGuiPayload*>(payload));
}

void iggEndDragDropTarget() {
  ImGui::EndDragDropTarget();
}
