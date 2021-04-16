package imgui

// #include "imguiWrapperTypes.h"
import "C"

// Vec2 represents a two-dimensional vector.
type Vec2 struct {
	X float32
	Y float32
}

func (vec *Vec2) wrapped() (out *C.IggVec2, finisher func()) {
	if vec != nil {
		out = &C.IggVec2{
			x: C.float(vec.X),
			y: C.float(vec.Y),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}

// Set sets values of the vec with args.
func (vec *Vec2) Set(x, y float32) {
	vec.X = x
	vec.Y = y
}

// Plus returns vec + other.
func (vec Vec2) Plus(other Vec2) Vec2 {
	return Vec2{
		X: vec.X + other.X,
		Y: vec.Y + other.Y,
	}
}

// Minus returns vec - other.
func (vec Vec2) Minus(other Vec2) Vec2 {
	return Vec2{
		X: vec.X - other.X,
		Y: vec.Y - other.Y,
	}
}

// Times returns vec * value.
func (vec Vec2) Times(value float32) Vec2 {
	return Vec2{
		X: vec.X * value,
		Y: vec.Y * value,
	}
}

// Vec4 represents a four-dimensional vector.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (vec *Vec4) wrapped() (out *C.IggVec4, finisher func()) {
	if vec != nil {
		out = &C.IggVec4{
			x: C.float(vec.X),
			y: C.float(vec.Y),
			z: C.float(vec.Z),
			w: C.float(vec.W),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
			vec.Z = float32(out.z) // nolint: gotype
			vec.W = float32(out.w) // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}

// Set sets values of the vec with args.
func (vec *Vec4) Set(x, y, z, w float32) {
	vec.X = x
	vec.Y = y
	vec.Z = z
	vec.W = w
}

// Plus returns vec + other.
func (vec Vec4) Plus(other Vec4) Vec4 {
	return Vec4{
		X: vec.X + other.X,
		Y: vec.Y + other.Y,
		Z: vec.Z + other.Z,
		W: vec.W + other.W,
	}
}

// Minus returns vec - other.
func (vec Vec4) Minus(other Vec4) Vec4 {
	return Vec4{
		X: vec.X - other.X,
		Y: vec.Y - other.Y,
		Z: vec.Z - other.Z,
		W: vec.W - other.W,
	}
}

// Times returns vec * value.
func (vec Vec4) Times(value float32) Vec4 {
	return Vec4{
		X: vec.X * value,
		Y: vec.Y * value,
		Z: vec.Z * value,
		W: vec.W * value,
	}
}
