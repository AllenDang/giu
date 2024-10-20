//go:build windows
// +build windows

package giu

import "runtime"

// according to:
// https://github.com/AllenDang/giu/issues/881
// this init solves some windows problems.
// However I'm not sure about implications of this. Will turn out later.
func init() {
	runtime.LockOSThread()
}

// I have no working mainthread library for windows.
// - this one for macOS crashes app immediately
// - this for linux (and everything else) freezes after a few seconds
//
// this seems to solve an issue: https://github.com/AllenDang/giu/issues/881
func mainthreadCallPlatform(c func()) {
	c()
}
