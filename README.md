# giu

[![Join the chat at https://gitter.im/AllenDang-giu/community](https://badges.gitter.im/AllenDang-giu/community.svg)](https://gitter.im/AllenDang-giu/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Cross platform rapid GUI framework for golang based on [Dear ImGui](https://github.com/ocornut/imgui) and the great golang binding [imgui-go](https://github.com/inkyblackness/imgui-go).

[Document](https://github.com/AllenDang/giu/wiki)

## Supported Platforms

The supported platforms depends on GLFW v3.3, so idealy giu could support all platforms that GLFW v3.3 supports.

- Windows (only tested on Windows 10 x64)
- MacOS (only tested on MacOS v10.15)
- Linux (thanks remeh to test it)
- Raspberry pi 3b (thanks sndvaps to test it)

## Features

Compare to other Dear ImGui golang bindings, giu has following features:

- Live-update during the resizing of OS window (implemented on GLFW 3.3 and OpenGL 3.2).
- Redraw only when user event occurred. Costs only 0.5% CPU usage with 60FPS.
- Declarative UI (see examples for more detail).
- DPI awareness (auto scale font and UI to adapte high DPI monitor).
- Drop in usage, no need to implement render and platform.
- Freetype font rendering support.
- OS clipboard support.

![Screenshot](https://github.com/AllenDang/giu/raw/master/examples/widgets/screenshot.png)

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
    g.SingleWindow("hello world", g.Layout{
        g.Label("Hello world from giu"),
        g.Line(
            g.Button("Click Me", onClickMe),
            g.Button("I'm so cute", onImSoCute)),
        })
}

func main() {
    wnd := g.NewMasterWindow("Hello world", 400, 200, false, nil)
    wnd.Main(loop)
}
```

Here is result.

![Helloworld](https://github.com/AllenDang/giu/raw/master/examples/helloworld/helloworld.png)

## Projects using giu

### [PipeIt](https://github.com/AllenDang/PipeIt)

PipeIt is a text transformation, conversion, cleansing and extraction tool.

![PipeIt Demo](https://github.com/AllenDang/PipeIt/raw/master/screenshot/findimageurl.gif)

### [EbookDownloader (Chinese language only)](https://github.com/sndnvaps/ebookdownloader)

Ebook downloader supports novel downloading from various web sites.
