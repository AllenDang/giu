//go:build darwin
// +build darwin

package giu

import "golang.design/x/hotkey/mainthread"

func mainthreadCallPlatform(c func()) {
	mainthread.Call(c)
}
