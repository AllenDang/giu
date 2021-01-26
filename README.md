# giu

[![Join the chat at https://gitter.im/AllenDang-giu/community](https://badges.gitter.im/AllenDang-giu/community.svg)](https://gitter.im/AllenDang-giu/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Go Report Card](https://goreportcard.com/badge/github.com/AllenDang/giu)](https://goreportcard.com/report/github.com/AllenDang/giu) [![Build Status](https://travis-ci.org/AllenDang/giu.svg?branch=master)](https://travis-ci.org/AllenDang/giu) [![Godoc Card](https://camo.githubusercontent.com/fd3cd5d5f44237541b35fcfdcba2fd4466a60c12/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476f646f632d7265666572656e63652d626c75652e737667)](https://pkg.go.dev/github.com/AllenDang/giu?tab=doc)

Cross platform rapid GUI framework for golang based on [Dear ImGui](https://github.com/ocornut/imgui) and the great golang binding [imgui-go](https://github.com/inkyblackness/imgui-go).

Any contribution (features, widgets, tutorials, documents and etc...) is appreciated!

## Supported Platforms

giu is built upon GLFW v3.3, so idealy giu could support all platforms that GLFW v3.3 supports.

- Windows (only tested on Windows 10 x64)
- MacOS (only tested on MacOS v10.15)
- Linux (thanks remeh to test it)
- Raspberry pi 3b (thanks sndvaps to test it)

## Features

Compare to other Dear ImGui golang bindings, giu has following features:

- Small executable file size (<3mb after upx compression for the example/helloworld demo).
- Live-update during the resizing of OS window (implemented on GLFW 3.3 and OpenGL 3.2).
- Redraw only when user event occurred. Costs only 0.5% CPU usage with 60FPS.
- Declarative UI (see examples for more detail).
- DPI awareness (auto scale font and UI to adapte high DPI monitor).
- Drop in usage, no need to implement render and platform.
- Freetype font rendering support.
- OS clipboard support.

![Screenshot](https://github.com/AllenDang/giu/raw/master/examples/imguidemo/screenshot.png)
![Screenshot1](https://github.com/AllenDang/giu/blob/master/screenshots/SqlPower.png)
![Screenshot2](https://github.com/AllenDang/giu/blob/master/screenshots/Chart.png)

## Hello world

```go
package main

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop() {
	g.SingleWindow("hello world").Layout(
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me").OnClick(onClickMe),
			g.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
	wnd.Run(loop)
}
```

Here is result.

![Helloworld](https://github.com/AllenDang/giu/raw/master/examples/helloworld/helloworld.png)

## Document

Check [Wiki](https://github.com/AllenDang/giu/wiki)

## Contribution

All kinds of pull request (document, demo, screenshots, code, etc...) are more then welcome!

## Projects using giu

### [PipeIt](https://github.com/AllenDang/PipeIt)

PipeIt is a text transformation, conversion, cleansing and extraction tool.

![PipeIt Demo](https://github.com/AllenDang/PipeIt/raw/master/screenshot/findimageurl.gif)

### [NVTool](https://github.com/Nicify/nvtool)

NVTool is a video encoding tool based on NVEncC.

![NVTool Screenshots](https://images-cdn.shimo.im/dLiWypVO9fbgAXPb__original.png)
