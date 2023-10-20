//go:build !darwin
// +build !darwin

package giu

import "github.com/faiface/mainthread"

func mainthreadCallPlatform(c func()) {
	mainthread.Run(c)
}
