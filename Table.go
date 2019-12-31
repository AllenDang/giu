package giu

import "github.com/AllenDang/giu/imgui"

type RowWidget struct {
	BaseWidget
	widgets []Widget
}

func Row(widgets ...Widget) *RowWidget {
	return &RowWidget{
		BaseWidget: BaseWidget{width: 0},
		widgets:    widgets,
	}
}

func (r *RowWidget) Build() {
	for _, w := range r.widgets {
		w.Build()
		imgui.NextColumn()
	}
}

func (r *RowWidget) Count() int {
	return len(r.widgets)
}

type TableWidget struct {
	BaseWidget
	label  string
	border bool
	rows   []*RowWidget
}

func Table(label string, border bool, rows ...*RowWidget) *TableWidget {
	return &TableWidget{
		BaseWidget: BaseWidget{width: 0},
		label:      label,
		border:     border,
		rows:       rows,
	}
}

func (t *TableWidget) Build() {
	if len(t.rows) > 0 {
		imgui.ColumnsV(t.rows[0].Count(), t.label, t.border)

		for _, r := range t.rows {
			if t.border {
				imgui.Separator()
			}
			r.Build()
		}

		if t.border {
			imgui.Columns()
			imgui.Separator()
		}
	}
}
