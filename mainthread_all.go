//go:build !darwin && !windows
// +build !darwin,!windows

package giu

import "github.com/faiface/mainthread"

func mainthreadCallPlatform(c func()) {
	mainthread.Run(c)
}
