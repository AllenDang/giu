// Package imgui contains all the functions to create an immediate mode graphical user interface based on Dear ImGui.
//
// Setup
//
// For integration, please refer to the dedicated repository
// https://github.com/inkyblackness/imgui-go-examples ,
// which contains ported examples of the C++ version, available to Go.
//
// Conventions
//
// The exported functions and constants are named closely to that of the wrapped library.
// If a function has optional parameters, it will be available in two versions:
// A verbose one, which has all optional parameters listed, and a terse one, with only the mandatory parameters in its signature.
// The verbose variant will have the suffix V in its name. For example, there are
//         func Button(id string) bool
// and
//         func ButtonV(id string, size Vec2) bool
// The terse variant will list the default parameters it uses to call the verbose variant.
//
// There are several types which are based on uintptr. These are references to the wrapped instances in C++.
// You will always get to such a reference via some function - you never "instantiate" such an instance on your own.
package imgui
