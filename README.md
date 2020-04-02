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
    g.SingleWindow("hello world", g.Layout{
        g.Label("Hello world from giu"),
        g.Line(
            g.Button("Click Me", onClickMe),
            g.Button("I'm so cute", onImSoCute)),
        })
}

func main() {
    wnd := g.NewMasterWindow("Hello world", 400, 200, g.MasterWindowFlagsNotResizable, nil)
    wnd.Main(loop)
}
```

Here is result.

![Helloworld](https://github.com/AllenDang/giu/raw/master/examples/helloworld/helloworld.png)

## Document

Check [Wiki](https://github.com/AllenDang/giu/wiki)

## Embed Lua as script language to create UI

This is a very interesting use case and it is incredibly easy.

```go
package main

import (
	g "github.com/AllenDang/giu"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// Define a simple plugin struct
type LuaPlugin struct {
	Name   string
	Layout g.Layout
}

// Genreate a string pointer for lua
func GStrPtr() *string {
	var str string
	return &str
}

// Receive string value from pointer
func ToStr(str *string) string {
	return *str
}

var luaPlugin LuaPlugin

func onRunScript() {
	luaPlugin.Name = ""
	luaPlugin.Layout = g.Layout{}

	luaState := lua.NewState()
	defer luaState.Close()

	// Pass luaPlugin into lua VM.
	luaState.SetGlobal("luaPlugin", luar.New(luaState, &luaPlugin))

	// Register some method (giu widget creator)
	luaState.SetGlobal("GStrPtr", luar.New(luaState, GStrPtr))
	luaState.SetGlobal("ToStr", luar.New(luaState, ToStr))

	luaState.SetGlobal("Label", luar.New(luaState, g.Label))
	luaState.SetGlobal("Button", luar.New(luaState, g.Button))
	luaState.SetGlobal("InputText", luar.New(luaState, g.InputText))

	// Simple lua code
	luaCode := `
    luaPlugin.Name = "test"

    name = GStrPtr()
    
    function onGreeting()
	  print(string.format("Greeting %s", ToStr(name)))
    end
    
    luaPlugin.Layout = {
      Label("Label from lua, tell me your name"),
      InputText("##name", 200, name),
      Button("Greeting", onGreeting),
    }
  `

	// Run lua script
	if err := luaState.DoString(luaCode); err != nil {
		panic(err)
	}
}

func loop() {
	g.SingleWindow("Lua test", g.Layout{
		g.Button("Load from lua", onRunScript),
		luaPlugin.Layout,
	})
}

func main() {
	wnd := g.NewMasterWindow("Lua test", 400, 300, 0, nil)
	wnd.Main(loop)
}


```

## Contribution

All kinds of pull request (document, demo, screenshots, code, etc...) are more then welcome!

## Projects using giu

### [PipeIt](https://github.com/AllenDang/PipeIt)

PipeIt is a text transformation, conversion, cleansing and extraction tool.

![PipeIt Demo](https://github.com/AllenDang/PipeIt/raw/master/screenshot/findimageurl.gif)
