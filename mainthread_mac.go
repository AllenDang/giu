//go:build darwin
// +build darwin

package giu

import "golang.design/x/hotkey/mainthread"

func mainthreadCall(c func()) {
	mainthread.init(c)
}
