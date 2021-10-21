package giu

import (
	"github.com/AllenDang/imgui-go"
)

type memoryEditorState struct {
	editor imgui.MemoryEditor
}

// Dispose implements Disposable interface.
func (s *memoryEditorState) Dispose() {
	// noop
}

// MemoryEditorWidget - Mini memory editor for Dear ImGui
// (to embed in your game/tools)
//
// Right-click anywhere to access the Options menu!
// You can adjust the keyboard repeat delay/rate in ImGuiIO.
// The code assume a mono-space font for simplicity!
// If you don't use the default font, use ImGui::PushFont()/PopFont() to switch to a mono-space font before calling this.
type MemoryEditorWidget struct {
	id       string
	contents []byte
}

// MemoryEditor creates nwe memory editor widget.
func MemoryEditor() *MemoryEditorWidget {
	return &MemoryEditorWidget{
		id: GenAutoID("memoryEditor"),
	}
}

// Contents sets editor's conents.
func (me *MemoryEditorWidget) Contents(contents []byte) *MemoryEditorWidget {
	me.contents = contents
	return me
}

// Build implements widget inetrface.
func (me *MemoryEditorWidget) Build() {
	me.getState().editor.DrawContents(me.contents)
}

func (me *MemoryEditorWidget) getState() (state *memoryEditorState) {
	if s := Context.GetState(me.id); s == nil {
		state = &memoryEditorState{
			editor: imgui.NewMemoryEditor(),
		}

		Context.SetState(me.id, state)
	} else {
		var ok bool
		state, ok = s.(*memoryEditorState)
		Assert(ok, "MemoryEditorWidget", "getState", "incorrect state type recovered.")
	}

	return state
}
