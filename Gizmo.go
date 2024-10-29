package giu

import "github.com/AllenDang/cimgui-go/imguizmo"

// GizmoI should be implemented by every sub-element of GizmoWidget.
type GizmoI interface {
	Gizmo()
}

var _ Widget = &GizmoWidget{}

// GizmoWidget implement ImGuizmo features.
// It is designed just like PlotWidget.
// This structure provides an "area" where you can put Gizmos (see (*GizmoWidget).Gizmos).
type GizmoWidget struct {
	gizmos []GizmoI
}

// Gizmo creates a new GizmoWidget.
func Gizmo() *GizmoWidget {
	return &GizmoWidget{
		gizmos: []GizmoI{},
	}
}

func (g *GizmoWidget) Gizmos(gizmos ...GizmoI) *GizmoWidget {
	g.gizmos = append(g.gizmos, gizmos...)
	return g
}

func (g *GizmoWidget) build() {
	for _, gizmo := range g.gizmos {
		gizmo.Gizmo()
	}
}

// Build implements Widget interface.
func (g *GizmoWidget) Build() {
	imguizmo.SetDrawlist()
	g.build()
}

// Global works like Build() but does not attach the gizmo to the current window.
func (g *GizmoWidget) Global() {
	g.build()
}
