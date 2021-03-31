#include "imnodes.h"
#include "imnodesWrapper.h"
#include "imguiWrappedHeader.h"
#include "WrapperConverter.h"

void iggImNodesCreateContext()
{
  imnodes::CreateContext();
}

void iggImNodesDestroyContext()
{
  imnodes::DestroyContext();
}

void iggImNodesBeginNodeEditor()
{
  imnodes::BeginNodeEditor();
}

void iggImNodesEndNodeEditor()
{
  imnodes::EndNodeEditor();
}

void iggImNodesBeginNode(int id)
{
  imnodes::BeginNode(id);
}

void iggImNodesEndNode()
{
  imnodes::EndNode();
}

void iggImNodesBeginNodeTitleBar()
{
  imnodes::BeginNodeTitleBar();
}

void iggImNodesEndNodeTitleBar()
{
  imnodes::EndNodeTitleBar();
}

void iggImNodesBeginInputAttribute(int id)
{
  imnodes::BeginInputAttribute(id);
}

void iggImNodesEndInputAttribute()
{
  imnodes::EndInputAttribute();
}

void iggImNodesBeginOutputAttribute(int id)
{
  imnodes::BeginOutputAttribute(id);
}

void iggImNodesEndOutputAttribute()
{
  imnodes::EndOutputAttribute();
}

void iggImNodesLink(int id, int start_attribute_id, int end_attribute_id)
{
  imnodes::Link(id, start_attribute_id, end_attribute_id);
}

IggBool iggImNodesIsLinkCreated(
    int* started_at_node_id,
    int* started_at_attribute_id,
    int* ended_at_node_id,
    int* ended_at_attribute_id,
    IggBool* created_from_snap)
{
  BoolWrapper boolArg(created_from_snap);
  return imnodes::IsLinkCreated(started_at_node_id, started_at_attribute_id, ended_at_node_id, ended_at_attribute_id, boolArg) ? 1 : 0;
}

IggBool iggImNodesIsLinkDestroyed(int* link_id)
{
  return imnodes::IsLinkDestroyed(link_id) ? 1 : 0;
}

void iggImNodesPushAttributeFlag(int flag)
{
  imnodes::PushAttributeFlag(static_cast<imnodes::AttributeFlags>(flag));
}

void iggImNodesPopAttributeFlag()
{
  imnodes::PopAttributeFlag();
}

void iggImNodesEnableDetachWithCtrlClick()
{
  imnodes::IO& io = imnodes::GetIO();
  io.link_detach_with_modifier_click.modifier = &ImGui::GetIO().KeyCtrl;
}

void iggImNodesSetNodeScreenSpacePos(int node_id, const IggVec2 *screen_space_pos)
{
  Vec2Wrapper posArg(screen_space_pos);
  imnodes::SetNodeScreenSpacePos(node_id, *posArg);
}

void iggImNodesSetNodeEditorSpacePos(int node_id, const IggVec2 *editor_space_pos)
{
  Vec2Wrapper posArg(editor_space_pos);
  imnodes::SetNodeEditorSpacePos(node_id, *posArg);
}

void iggImNodesSetNodeGridSpacePos(int node_id, const IggVec2 *grid_pos)
{
  Vec2Wrapper posArg(grid_pos);
  imnodes::SetNodeGridSpacePos(node_id, *posArg);
}

void iggImNodesGetNodeScreenSpacePos(const int node_id, IggVec2 *pos)
{
  exportValue(*pos, imnodes::GetNodeScreenSpacePos(node_id));
}

void iggImNodesGetNodeEditorSpacePos(const int node_id, IggVec2 *pos)
{
  exportValue(*pos, imnodes::GetNodeEditorSpacePos(node_id));
}

void iggImNodesGetNodeGridSpacePos(const int node_id, IggVec2 *pos)
{
  exportValue(*pos, imnodes::GetNodeGridSpacePos(node_id));
}
