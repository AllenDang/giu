package giu

import (
	"image/color"

	"github.com/AllenDang/cimgui-go/implot"
)

//go:generate go run github.com/dmarkham/enumer@latest -linecomment -type=PlotProperty .

// PlotProperty represents a property of a single plot on the canvas.
// See PlotSpec.
type PlotProperty byte

const (
	PlotPropertyLineColor        PlotProperty = PlotProperty(implot.PropLineColor)
	PlotPropertyLineColors       PlotProperty = PlotProperty(implot.PropLineColors)
	PlotPropertyLineWeight       PlotProperty = PlotProperty(implot.PropLineWeight)
	PlotPropertyFillColor        PlotProperty = PlotProperty(implot.PropFillColor)
	PlotPropertyFillColors       PlotProperty = PlotProperty(implot.PropFillColors)
	PlotPropertyFillAlpha        PlotProperty = PlotProperty(implot.PropFillAlpha)
	PlotPropertyMarker           PlotProperty = PlotProperty(implot.PropMarker)
	PlotPropertyMarkerSize       PlotProperty = PlotProperty(implot.PropMarkerSize)
	PlotPropertyMarkerSizes      PlotProperty = PlotProperty(implot.PropMarkerSizes)
	PlotPropertyMarkerLineColor  PlotProperty = PlotProperty(implot.PropMarkerLineColor)
	PlotPropertyMarkerLineColors PlotProperty = PlotProperty(implot.PropMarkerLineColors)
	PlotPropertyMarkerFillColor  PlotProperty = PlotProperty(implot.PropMarkerFillColor)
	PlotPropertyMarkerFillColors PlotProperty = PlotProperty(implot.PropMarkerFillColors)
	PlotPropertySize             PlotProperty = PlotProperty(implot.PropSize)
	PlotPropertyOffset           PlotProperty = PlotProperty(implot.PropOffset)
	PlotPropertyStride           PlotProperty = PlotProperty(implot.PropStride)
	PlotPropertyFlags            PlotProperty = PlotProperty(implot.PropFlags)
)

// PlotSpec is a wrapper for implot.Spec. It allows to set various properties for plots.
type PlotSpec struct {
	spec *implot.Spec
}

func NewPlotSpec() *PlotSpec {
	return &PlotSpec{spec: implot.NewSpec()}
}

// SetProperty allows to set choosen property to value. Type of value could be:
// - uint16
// - uint32
// - *uint32
// - uint64
// - uint8
// - vec4
// - float64
// - int16
// - int
// - int64
// TODO: This could be made a generic method as soon as go adds them (probably go 1.27)
func (ps *PlotSpec) SetProperty(property PlotProperty, value any) *PlotSpec {
	switch value.(type) {
	case uint16:
		val, ok := value.(uint16)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type uint16")
		ps.spec.SetPropU16(implot.Prop(property), val)
	case uint32:
		val, ok := value.(uint32)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type uint32")
		ps.spec.SetPropU32(implot.Prop(property), val)
	case *uint32:
		val, ok := value.(*uint32)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type *uint32")
		ps.spec.SetPropU32Ptr(implot.Prop(property), val)
	case uint64:
		val, ok := value.(uint64)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type uint64")
		ps.spec.SetPropU64(implot.Prop(property), val)
	case uint8:
		val, ok := value.(uint8)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type uint8")
		ps.spec.SetPropU8(implot.Prop(property), val)
	case color.Color:
		val, ok := value.(color.Color)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type color.Color")
		ps.spec.SetPropVec4(implot.Prop(property), ToVec4Color(val))
	case float64:
		val, ok := value.(float64)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type float64")
		ps.spec.SetPropdouble(implot.Prop(property), val)
	case int16:
		ps.spec.SetPropS16(implot.Prop(property), value.(int16))
	case int:
		val, ok := value.(int)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type int")
		ps.spec.SetPropS32(implot.Prop(property), val)
	case int64:
		val, ok := value.(int64)
		Assert(ok, "PlotSpec", "SetProperty", "value should be of type int64")
		ps.spec.SetPropS64(implot.Prop(property), val)
	default:
		panic("unsupported type")
	}

	return ps
}

func (ps *PlotSpec) GetSpec() *implot.Spec {
	return ps.spec
}
