# giu
Cross platform rapid GUI framework for golang based on [imgui](https://github.com/ocornut/imgui) and the great golang binding [imgui-go](https://github.com/inkyblackness/imgui-go).

## Features

Compare to other imgui golang bindings, giu has following features:

- Live-update during the resizing of OS window (implemented on GLFW 3.3 and OpenGL 3.2).
- Redraw only when user event occurred. Costs only 0.5% CPU usage with 60FPS.
- Declarative UI (see examples for more detail).
- Drop in usage, no need to implement render and platform.
- Freetype font rendering support.
- OS clipboard support.

![Screenshot](https://github.com/AllenDang/giu/raw/master/examples/widgets/screenshot.png)

## Projects using giu

### [PipeIt](https://github.com/AllenDang/PipeIt)

PipeIt is a text transformation, conversion, cleansing and extraction tool.

![PipeIt Demo](https://github.com/AllenDang/PipeIt/raw/master/screenshot/findimageurl.gif)

