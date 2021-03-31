#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggImNodesCreateContext();
extern void iggImNodesDestroyContext();

extern void iggImNodesBeginNodeEditor();
extern void iggImNodesEndNodeEditor();

extern void iggImNodesBeginNode(int id);
extern void iggImNodesEndNode();

extern void iggImNodesBeginNodeTitleBar();
extern void iggImNodesEndNodeTitleBar();

extern void iggImNodesBeginInputAttribute(int id);
extern void iggImNodesEndInputAttribute();

extern void iggImNodesBeginOutputAttribute(int id);
extern void iggImNodesEndOutputAttribute();

extern void iggImNodesLink(int id, int start_attribute_id, int end_attribute_id);

extern IggBool iggImNodesIsLinkCreated(
    int* started_at_node_id,
    int* started_at_attribute_id,
    int* ended_at_node_id,
    int* ended_at_attribute_id,
    IggBool* created_from_snap);

extern IggBool iggImNodesIsLinkDestroyed(int* link_id);

extern void iggImNodesPushAttributeFlag(int flag);
extern void iggImNodesPopAttributeFlag();

extern void iggImNodesEnableDetachWithCtrlClick();

extern void iggImNodesSetNodeScreenSpacePos(int node_id, const IggVec2 *screen_space_pos);
extern void iggImNodesSetNodeEditorSpacePos(int node_id, const IggVec2 *editor_space_pos);
extern void iggImNodesSetNodeGridSpacePos(int node_id, const IggVec2 *grid_pos);

extern void iggImNodesGetNodeScreenSpacePos(const int node_id, IggVec2 *pos);
extern void iggImNodesGetNodeEditorSpacePos(const int node_id, IggVec2 *pos);
extern void iggImNodesGetNodeGridSpacePos(const int node_id, IggVec2 *pos);

#ifdef __cplusplus
}
#endif
