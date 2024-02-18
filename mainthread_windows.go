//go:build windows
// +build windows

package giu

// TODO: I have no working mainthread library for windows.
// - this one for macOS crashes app immediately
// - this for linux (and everything else) freezes after a few seconds
// 
// With no mianthread support this at least runs sometimes. Just keep your giu calls in one thread and everything should work.
func mainthreadCallPlatform(c func()) {
	c()
}
