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
		g.Row(
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

## Quick intruduction

### What is immediate mode GUI?

Immediate mode GUI system means the UI control doesn't retain it's state and value. For example, call `giu.InputText("ID", &str)` will display a input text box on screen, and the user entered value will be stored in `&str`, input text box doesn't know anything about it. 

And the `loop` method in the *Hello world* example is in charge of **drawing** all widgets based on the parameters passed into them. This method will be invoked 30 times per second to reflect interactive states (like clicked, hovered, value-changed etc...). It will be the place you define the UI structure.

### The layout and sizing system

By default, any widget is placed inside a container's `Layout` will be place vertically.

To create a row of widgets (aka place widgets one by one horizontally), use `Row()` method. For example `giu.Row(Label(...), Button(...))` will create a Label next to a Button.

To creata a column of widgets (aka place widgets one by one vertically) inside a row, use `Column()` method.

Any widget which has a `Size()` method, could set it's size explicitly. Note you could pass negative value to `Size()`, it means avaiable remain width/height - value. For example, `InputText(...).Size(-1)` will create a input text box with longest width it's container has lefted.

### Containers

#### MasterWindow

A `MasterWindow` means the platform native window implemented by OS. All sub window and widgets will be placed inside it.

#### Window

A `Window` is a container with a title bar, and could be collapsed. `SingleWindow` is a special kind of window who will occupy all avaialbe space of `MasterWindow`.

#### Child

A `Child` is like a panel in other GUI framework, it could have a background color and border.

### Widgets

Check `examples/widgets` for all kinds of widgets.

## Install

The backend of giu depends on OpenGL 3.3, make sure your environment supports it (so far as I known some Virual Machine like VirualBox doesn`t support it).

### MacOS

``` sh
xcode-select --install
go get github.com/AllenDang/giu@master
```

### Windows

1. Install mingw [download here](https://github.com/brechtsanders/winlibs_mingw/releases/tag/10.2.0-11.0.0-8.0.0-r8). Thanks @alchem1ster!
2. Add the binaries folder of mingw to the path (usually is *\mingw64\bin*).
3. go get github.com/AllenDang/giu@master.

### Linux

First you need to install libraries
```bash
# apt install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libglx-dev libgl1-mesa-dev libxxf86vm-dev
```
Then simple `go build` will work.

Cross-compiling is a bit more complicated. Let's say that you want to build for arm64. That's what you would need to do:

```bash
# dpkg --add-architecture arm64
# apt update
# apt install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu \
    libx11-dev:arm64 libxcursor-dev:arm64 libxrandr-dev:arm64 libxinerama-dev:arm64 libxi-dev:arm64 libglx-dev:arm64 libgl1-mesa-dev:arm64 libxxf86vm-dev:arm64
$ GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-g++ HOST=aarch64-linux-gnu go build -v
```

## Deploy

### Build MacOS version on MacOS.

``` sh
go build -ldflags "-s -w" .
```

### Build Windows version on Windows.

``` sh
go build -ldflags "-s -w -H=windowsgui -extldflags=-static" .
```

### Build Windows version on MacOS.

1. Install mingw-64.
``` sh
brew install mingw-w64
```

2. Prepare and embed application icon to executable and build.

``` sh
cat > YourExeName.rc << EOL
id ICON "./res/app_win.ico"
GLFW_ICON ICON "./res/app_win.ico"
EOL

x86_64-w64-mingw32-windres YourExeName.rc -O coff -o YourExeName.syso
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -p 4 -v -o YourExeName.exe

rm YourExeName.syso
rm YourExeName.rc
```

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
