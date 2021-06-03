package main

import (
	"fmt"
	"strings"

	g "github.com/AllenDang/giu"
)

var (
	dropInFiles string
)

func loop() {
	g.SingleWindow("On Drop Demo").Layout(
		g.Label("Drop file to this window"),
		g.InputTextMultiline("#DroppedFiles", &dropInFiles).Size(-1, -1).Flags(g.InputTextFlagsReadOnly),
	)
}

func onDrop(names []string) {
	var sb strings.Builder
	for _, n := range names {
		sb.WriteString(fmt.Sprintf("%s\n", n))
	}

	dropInFiles = sb.String()
	g.Update()
}

func main() {
	wnd := g.NewMasterWindow("On Drop Demo", 600, 400, g.MasterWindowFlagsNotResizable)
	wnd.SetDropCallback(onDrop)
	wnd.Run(loop)
}
