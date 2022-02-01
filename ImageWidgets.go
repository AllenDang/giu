package giu

import (
	ctx "context"
	"fmt"
	"image"
	"image/color"
	"net/http"
	"time"

	"github.com/AllenDang/imgui-go"
)

var _ Widget = &ImageWidget{}

// ImageWidget adds an image.
// NOTE: ImageWidget is going to be deprecated. ImageWithRGBAWidget
// should be used instead, however, because it is a native
// imgui's solution it is still there.
type ImageWidget struct {
	texture                *Texture
	width                  float32
	height                 float32
	uv0, uv1               image.Point
	tintColor, borderColor color.Color
	onClick                func()
}

// Image adds an image from giu.Texture.
func Image(texture *Texture) *ImageWidget {
	return &ImageWidget{
		texture:     texture,
		width:       100,
		height:      100,
		uv0:         image.Point{X: 0, Y: 0},
		uv1:         image.Point{X: 1, Y: 1},
		tintColor:   color.RGBA{255, 255, 255, 255},
		borderColor: color.RGBA{0, 0, 0, 0},
	}
}

// Uv allows to specify uv parameters.
func (i *ImageWidget) Uv(uv0, uv1 image.Point) *ImageWidget {
	i.uv0, i.uv1 = uv0, uv1
	return i
}

// TintColor sets image's tint color.
func (i *ImageWidget) TintColor(tintColor color.Color) *ImageWidget {
	i.tintColor = tintColor
	return i
}

// BorderCol sets color of the border.
func (i *ImageWidget) BorderCol(borderColor color.Color) *ImageWidget {
	i.borderColor = borderColor
	return i
}

// OnClick adds on-click-callback.
func (i *ImageWidget) OnClick(cb func()) *ImageWidget {
	i.onClick = cb
	return i
}

// Size sets image size.
func (i *ImageWidget) Size(width, height float32) *ImageWidget {
	// Size image with DPI scaling
	factor := Context.GetPlatform().GetContentScale()
	i.width, i.height = width*factor, height*factor
	return i
}

// Build implements Widget interface.
func (i *ImageWidget) Build() {
	size := imgui.Vec2{X: i.width, Y: i.height}
	rect := imgui.ContentRegionAvail()
	if size.X == -1 {
		size.X = rect.X
	}
	if size.Y == -1 {
		size.Y = rect.Y
	}

	if i.texture == nil || i.texture.id == 0 {
		Dummy(size.X, size.Y).Build()
		return
	}

	// trick: detect click event
	if i.onClick != nil && IsMouseClicked(MouseButtonLeft) {
		cursorPos := GetCursorScreenPos()
		mousePos := GetMousePos()
		mousePos.Add(cursorPos)
		if cursorPos.X <= mousePos.X && cursorPos.Y <= mousePos.Y &&
			cursorPos.X+int(i.width) >= mousePos.X && cursorPos.Y+int(i.height) >= mousePos.Y {
			i.onClick()
		}
	}

	imgui.ImageV(i.texture.id, size, ToVec2(i.uv0), ToVec2(i.uv1), ToVec4Color(i.tintColor), ToVec4Color(i.borderColor))
}

type imageState struct {
	loading bool
	failure bool
	cancel  ctx.CancelFunc
	texture *Texture
}

// Dispose cleans imageState (implements Disposable interface).
func (is *imageState) Dispose() {
	is.texture = nil
	// Cancel ongoing image downloaidng
	if is.loading && is.cancel != nil {
		is.cancel()
	}
}

var _ Widget = &ImageWithRgbaWidget{}

// ImageWithRgbaWidget wrapps ImageWidget.
// It is more useful because it doesn't make you to care about
// imgui textures. You can just pass golang-native image.Image and
// display it in giu.
type ImageWithRgbaWidget struct {
	id   string
	rgba image.Image
	img  *ImageWidget
}

// ImageWithRgba creates ImageWithRgbaWidget.
func ImageWithRgba(rgba image.Image) *ImageWithRgbaWidget {
	return &ImageWithRgbaWidget{
		id:   GenAutoID("ImageWithRgba"),
		rgba: rgba,
		img:  Image(nil),
	}
}

// Size sets image's size.
func (i *ImageWithRgbaWidget) Size(width, height float32) *ImageWithRgbaWidget {
	i.img.Size(width, height)
	return i
}

// OnClick sets click callback.
func (i *ImageWithRgbaWidget) OnClick(cb func()) *ImageWithRgbaWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithRgbaWidget) Build() {
	if i.rgba != nil {
		var imgState *imageState
		if state := Context.GetState(i.id); state == nil {
			imgState = &imageState{}
			Context.SetState(i.id, imgState)

			NewTextureFromRgba(i.rgba, func(tex *Texture) {
				imgState.texture = tex
			})
		} else {
			var isOk bool
			imgState, isOk = state.(*imageState)
			Assert(isOk, "ImageWithRgbaWidget", "Build", "unexpected type of widget's state recovered")
		}

		i.img.texture = imgState.texture
	}

	i.img.Build()
}

var _ Widget = &ImageWithFileWidget{}

// ImageWithFileWidget allows to display an image directly
// from .png file.
// NOTE: Be aware that project using this solution may not be portable
// because files are not included in executable binaries!
// You may want to use "embed" package and ImageWithRgba instead.
type ImageWithFileWidget struct {
	id      string
	imgPath string
	img     *ImageWidget
}

// ImageWithFile constructs a new ImageWithFileWidget.
func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		id:      fmt.Sprintf("ImageWithFile_%s", imgPath),
		imgPath: imgPath,
		img:     Image(nil),
	}
}

// Size sets image's size.
func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.img.Size(width, height)
	return i
}

// OnClick sets click callback.
func (i *ImageWithFileWidget) OnClick(cb func()) *ImageWithFileWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithFileWidget) Build() {
	imgState := &imageState{}
	if state := Context.GetState(i.id); state == nil {
		// Prevent multiple invocation to LoadImage.
		Context.SetState(i.id, imgState)

		img, err := LoadImage(i.imgPath)
		if err == nil {
			NewTextureFromRgba(img, func(tex *Texture) {
				imgState.texture = tex
			})
		}
	} else {
		var isOk bool
		imgState, isOk = state.(*imageState)
		Assert(isOk, "ImageWithFileWidget", "Build", "wrong type of widget's state got")
	}

	i.img.texture = imgState.texture
	i.img.Build()
}

var _ Widget = &ImageWithURLWidget{}

// ImageWithURLWidget allows to display an image using
// an URL as image source.
type ImageWithURLWidget struct {
	id              string
	imgURL          string
	downloadTimeout time.Duration
	whenLoading     Layout
	whenFailure     Layout
	onReady         func()
	onFailure       func(error)
	img             *ImageWidget
}

// ImageWithURL creates ImageWithURLWidget.
func ImageWithURL(url string) *ImageWithURLWidget {
	return &ImageWithURLWidget{
		id:              fmt.Sprintf("ImageWithURL_%s", url),
		imgURL:          url,
		downloadTimeout: 10 * time.Second,
		whenLoading:     Layout{Dummy(100, 100)},
		whenFailure:     Layout{Dummy(100, 100)},
		img:             Image(nil),
	}
}

// OnReady sets event trigger when image is downloaded and ready to display.
func (i *ImageWithURLWidget) OnReady(onReady func()) *ImageWithURLWidget {
	i.onReady = onReady
	return i
}

// OnFailure sets event trigger when image failed to download/load.
func (i *ImageWithURLWidget) OnFailure(onFailure func(error)) *ImageWithURLWidget {
	i.onFailure = onFailure
	return i
}

// OnClick sets click callback.
func (i *ImageWithURLWidget) OnClick(cb func()) *ImageWithURLWidget {
	i.img.OnClick(cb)
	return i
}

// Timeout sets download timeout.
func (i *ImageWithURLWidget) Timeout(downloadTimeout time.Duration) *ImageWithURLWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

// Size sets image's size.
func (i *ImageWithURLWidget) Size(width, height float32) *ImageWithURLWidget {
	i.img.Size(width, height)
	return i
}

// LayoutForLoading allows to set layout rendered while loading an image.
func (i *ImageWithURLWidget) LayoutForLoading(widgets ...Widget) *ImageWithURLWidget {
	i.whenLoading = Layout(widgets)
	return i
}

// LayoutForFailure allows to specify layout when image failed to download.
func (i *ImageWithURLWidget) LayoutForFailure(widgets ...Widget) *ImageWithURLWidget {
	i.whenFailure = Layout(widgets)
	return i
}

// Build implements Widget interface.
func (i *ImageWithURLWidget) Build() {
	imgState := &imageState{}

	if state := Context.GetState(i.id); state == nil {
		Context.SetState(i.id, imgState)

		// Prevent multiple invocation to download image.
		downloadContext, cancalFunc := ctx.WithCancel(ctx.Background())
		Context.SetState(i.id, &imageState{loading: true, cancel: cancalFunc})

		errorFn := func(err error) {
			Context.SetState(i.id, &imageState{failure: true})

			// Trigger onFailure event
			if i.onFailure != nil {
				i.onFailure(err)
			}
		}

		go func() {
			// Load image from url
			client := &http.Client{Timeout: i.downloadTimeout}
			req, err := http.NewRequestWithContext(downloadContext, "GET", i.imgURL, http.NoBody)
			if err != nil {
				errorFn(err)
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				errorFn(err)
				return
			}

			defer func() {
				if closeErr := resp.Body.Close(); closeErr != nil {
					errorFn(closeErr)
				}
			}()

			img, _, err := image.Decode(resp.Body)
			if err != nil {
				errorFn(err)
				return
			}

			rgba := ImageToRgba(img)

			NewTextureFromRgba(rgba, func(tex *Texture) {
				Context.SetState(i.id, &imageState{
					loading: false,
					failure: false,
					texture: tex,
				})
			})

			// Trigger onReady event
			if i.onReady != nil {
				i.onReady()
			}
		}()
	} else {
		var isOk bool
		imgState, isOk = state.(*imageState)
		Assert(isOk, "ImageWithURLWidget", "Build", "wrong type of widget's state recovered.")
	}

	switch {
	case imgState.failure:
		i.whenFailure.Build()
	case imgState.loading:
		i.whenLoading.Build()
	default:
		i.img.texture = imgState.texture
		i.img.Build()
	}
}
