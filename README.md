# giu

[![Go Report Card](https://goreportcard.com/badge/github.com/AllenDang/giu)](https://goreportcard.com/report/github.com/AllenDang/giu)
![Build Status](https://github.com/AllenDang/giu/actions/workflows/build.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/AllenDang/giu.svg)](https://pkg.go.dev/github.com/AllenDang/giu)
[![Discord Shield](https://discord.com/api/guilds/1306199225616306248/widget.png?style=shield)](https://discord.gg/Tt7eq6YKQS)

A rapid cross-platform GUI framework for Go based on [Dear ImGui](https://github.com/ocornut/imgui) and the great Go binding [cimgui-go](https://github.com/inkyblackness/cimgui-go).

Any contribution (features, widgets, tutorials, documents, etc...) is appreciated!

## Sponsor

(This library is available under a free and permissive license, but needs financial support to sustain its continued improvements. In addition to maintenance and stability there are many desirable features yet to be added. If you are using giu, please consider reaching out.)

Businesses: support continued development and maintenance via invoiced technical support, maintenance, sponsoring contracts:

E-mail: <allengnr@gmail.com>

Individuals: support continued development and maintenance [here](https://patreon.com/AllenDang).

## Documentation

For documentation refer to [our wiki](https://github.com/AllenDang/giu/wiki),
[examples](./examples), [GoDoc](https://pkg.go.dev/github.com/AllenDang/giu),
or just take a look at comments in code.

## Supported Platforms

giu is built upon GLFW v3.3, so ideally giu could support all platforms that GLFW v3.3 supports.
It is also restricted by cimgui-go (which at the moment builds only for linux, windows and mac).

- Windows (Windows 10 x64 and Windows 11 x64)
- macOS (macOS v10.15 and macOS Big Sur)
- Linux (thanks remeh for testing it)
- Raspberry Pi 3B (thanks sndvaps for testing it)

> [!note]
> Because giu relays on C++ code, you need to have C/C++ compiler set up.

## Features

Compared to other Dear ImGui golang bindings, giu has the following features:

- Small executable file size (<3MB after UPX compression for the example/helloworld demo).
- Live-updating during the resizing of the OS window (implemented on GLFW 3.3 and OpenGL 3.2).
- Support for displaying various languages without any font setting. Giu will rebuild font atlas incrementally according to texts in UI between frames.
- Redraws only when user event occurs. Costs only 0.5% CPU usage with 60FPS.
- Declarative UI (see examples for more details).
- DPI awareness (auto scaling font and UI to adapt to high DPI monitors).
- Drop in usage; no need to implement render and platform.
- OS clipboard support.

<table border=0>
<tr><td>

![Screenshot](https://github.com/AllenDang/giu/raw/master/examples/imguidemo/screenshot.png)

<td>

![Screenshot1](https://github.com/AllenDang/giu/blob/master/screenshots/SqlPower.png)

<td>

![Screenshot2](https://github.com/AllenDang/giu/blob/master/screenshots/Chart.png)

</table>

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
        g.SingleWindow().Layout(
                g.Label("Hello world from giu"),
                g.Row(
                        g.Button("Click Me").OnClick(onClickMe),
                        g.Button("I'm so cute").OnClick(onImSoCute),
                ),
        )
}

func main() {
        wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable)
        wnd.Run(loop)
}
```

Here is the result:

![Helloworld](https://github.com/AllenDang/giu/raw/master/examples/helloworld/helloworld.png)

<details><summary><h2>Quick introduction</h2></summary>

### What is immediate mode GUI?

Immediate mode GUI system means the UI control doesn't retain its state and value. For example, calling `giu.InputText(&str)` will display a input text box on screen, and the user entered value will be stored in `&str`. Input text box doesn't know anything about it.

And the `loop` method in the _Hello world_ example is in charge of **drawing** all widgets based on the parameters passed into them. This method will be invoked 30 times per second to reflect interactive states (like clicked, hovered, value-changed, etc.). It will be the place you define the UI structure.

### The layout and sizing system

By default, any widget placed inside a container's `Layout` will be placed vertically.

To create a row of widgets (i.e. place widgets one by one horizontally), use the `Row()` method. For example `giu.Row(Label(...), Button(...))` will create a Label next to a Button.

To create a column of widgets (i.e. place widgets one by one vertically) inside a row, use the `Column()` method.

Any widget that has a `Size()` method, can set its size explicitly. Note that you can pass a negative value to `Size()`, which will fill the remaining width/height value. For example, `InputText(...).Size(giu.Auto)` will create an input text box with the longest width that its container has left.

### Containers

#### MasterWindow

A `MasterWindow` means the platform native window implemented by the OS. All subwindows and widgets will be placed inside it.

#### Window

A `Window` is a container with a title bar, and can be collapsed. `SingleWindow` is a special kind of window that will occupy all the available space of `MasterWindow`.

#### Child

A `Child` is like a panel in other GUI frameworks - it can have a background color and border.

### Widgets

Check `examples/widgets` for all kinds of widgets.

</details>

## Install

The backend of giu depends on OpenGL 3.3, make sure your environment supports it (as far as I know, some Virtual Machines like VirtualBox doesn't support it).

### MacOS

```sh
xcode-select --install
go get github.com/AllenDang/giu
```

### Windows

1. Install mingw [download here](https://github.com/brechtsanders/winlibs_mingw/releases/latest). Thanks @alchem1ster!
2. Add the binaries folder of mingw to the path (usually is _\mingw64\bin_).
3. `go get github.com/AllenDang/giu` in your project

### Linux

First you need to install the required dependencies:

<table>

<tr>
<td>
Debian/Ubuntu
<td>

```bash
sudo apt install libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libglx-dev libgl1-mesa-dev libxxf86vm-dev
```

<tr>
<td>
Fedora/Red Hat/CentOS
<td>

```bash
sudo dnf install libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libGL-devel libXxf86vm-devel
```

<tr>
<td>
Arch Linux
<td>

```bash
sudo pacman -Sy glfw
```

</table>

you may also need to install C/C++ compiler (like g++) if it isn't already installed. Follow go compiler prompts.

Then, a simple `go build` will work.

Cross-compiling is a bit more complicated. Let's say that you want to build for arm64. This is what you would need to do:

```bash
sudo dpkg --add-architecture arm64
sudo apt update
sudo apt install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu \
    libx11-dev:arm64 libxcursor-dev:arm64 libxrandr-dev:arm64 libxinerama-dev:arm64 libxi-dev:arm64 libglx-dev:arm64 libgl1-mesa-dev:arm64 libxxf86vm-dev:arm64
GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-g++ HOST=aarch64-linux-gnu go build -v
```

## Deploying

### MacOS -> MacOS

```sh
go build -ldflags "-s -w" .
```

### Windows -> Windows

```sh
go build -ldflags "-s -w -H=windowsgui -extldflags=-static" .
```

### MacOS/Linux -> Windows

#### Install mingw-64.

<table>
<tr>
<td>
Mac
<td>

```sh
brew install mingw-w64
```

<tr><td>
Fedora/RHEL/CentOS
<td>

```sh
sudo dnf install mingw64-gcc mingw64-gcc-c++ mingw64-winpthreads-static
```

<tr>
<td>
Arch Linux

<td>

```bash
pacman -Sy mingw-w64-gcc
```

</table>

2. Prepare and embed the application icon into the executable and build.

```sh
cat > YourExeName.rc << EOL
id ICON "./res/app_win.ico"
GLFW_ICON ICON "./res/app_win.ico"
EOL

x86_64-w64-mingw32-windres YourExeName.rc -O coff -o YourExeName.syso
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -p 4 -v -o YourExeName.exe

rm YourExeName.syso
rm YourExeName.rc
```

## Contribution

All kinds of pull requests (document, demo, screenshots, code, etc.) are more than welcome!

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=AllenDang/giu&type=Date)](https://star-history.com/#AllenDang/giu&Date)
