package giu

import (
	"bytes"
	ctx "context"
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/AllenDang/imgui-go"
	resty "github.com/go-resty/resty/v2"
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

func (i *ImageWidget) Uv(uv0, uv1 image.Point) *ImageWidget {
	i.uv0, i.uv1 = uv0, uv1
	return i
}

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
	i.width, i.height = width, height
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

type ImageState struct {
	loading bool
	failure bool
	cancel  ctx.CancelFunc
	texture *Texture
}

func (is *ImageState) Dispose() {
	is.texture = nil
	// Cancel ongoing image downloaidng
	if is.loading && is.cancel != nil {
		is.cancel()
	}
}

var _ Widget = &ImageWithRgbaWidget{}

type ImageWithRgbaWidget struct {
	id   string
	rgba image.Image
	img  *ImageWidget
}

func ImageWithRgba(rgba image.Image) *ImageWithRgbaWidget {
	return &ImageWithRgbaWidget{
		id:   GenAutoID("ImageWithRgba"),
		rgba: rgba,
		img:  Image(nil),
	}
}

func (i *ImageWithRgbaWidget) Size(width, height float32) *ImageWithRgbaWidget {
	i.img.Size(width, height)
	return i
}

func (i *ImageWithRgbaWidget) OnClick(cb func()) *ImageWithRgbaWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithRgbaWidget) Build() {
	if i.rgba != nil {
		var imgState *ImageState
		if state := Context.GetState(i.id); state == nil {
			imgState = &ImageState{}
			Context.SetState(i.id, imgState)

			NewTextureFromRgba(i.rgba, func(tex *Texture) {
				imgState.texture = tex
			})
		} else {
			var isOk bool
			imgState, isOk = state.(*ImageState)
			Assert(isOk, "ImageWithRgbaWidget", "Build", "unexpected type of widget's state recovered")
		}

		i.img.texture = imgState.texture
	}

	i.img.Build()
}

var _ Widget = &ImageWithFileWidget{}

type ImageWithFileWidget struct {
	id      string
	imgPath string
	img     *ImageWidget
}

func ImageWithFile(imgPath string) *ImageWithFileWidget {
	return &ImageWithFileWidget{
		id:      fmt.Sprintf("ImageWithFile_%s", imgPath),
		imgPath: imgPath,
		img:     Image(nil),
	}
}

func (i *ImageWithFileWidget) Size(width, height float32) *ImageWithFileWidget {
	i.img.Size(width, height)
	return i
}

func (i *ImageWithFileWidget) OnClick(cb func()) *ImageWithFileWidget {
	i.img.OnClick(cb)
	return i
}

// Build implements Widget interface.
func (i *ImageWithFileWidget) Build() {
	imgState := &ImageState{}
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
		imgState, isOk = state.(*ImageState)
		Assert(isOk, "ImageWithFileWidget", "Build", "wrong type of widget's state got")
	}

	i.img.texture = imgState.texture
	i.img.Build()
}

var _ Widget = &ImageWithURLWidget{}

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

func (i *ImageWithURLWidget) OnFailure(onFailure func(error)) *ImageWithURLWidget {
	i.onFailure = onFailure
	return i
}

func (i *ImageWithURLWidget) OnClick(cb func()) *ImageWithURLWidget {
	i.img.OnClick(cb)
	return i
}

func (i *ImageWithURLWidget) Timeout(downloadTimeout time.Duration) *ImageWithURLWidget {
	i.downloadTimeout = downloadTimeout
	return i
}

func (i *ImageWithURLWidget) Size(width, height float32) *ImageWithURLWidget {
	i.img.Size(width, height)
	return i
}

func (i *ImageWithURLWidget) LayoutForLoading(widgets ...Widget) *ImageWithURLWidget {
	i.whenLoading = Layout(widgets)
	return i
}

func (i *ImageWithURLWidget) LayoutForFailure(widgets ...Widget) *ImageWithURLWidget {
	i.whenFailure = Layout(widgets)
	return i
}

// Build implements Widget interface.
func (i *ImageWithURLWidget) Build() {
	imgState := &ImageState{}

	if state := Context.GetState(i.id); state == nil {
		Context.SetState(i.id, imgState)

		// Prevent multiple invocation to download image.
		downloadContext, cancalFunc := ctx.WithCancel(ctx.Background())
		Context.SetState(i.id, &ImageState{loading: true, cancel: cancalFunc})

		go func() {
			// Load image from url
			client := resty.New()
			client.SetTimeout(i.downloadTimeout)
			resp, err := client.R().SetContext(downloadContext).Get(i.imgURL)
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}

				return
			}

			img, _, err := image.Decode(bytes.NewReader(resp.Body()))
			if err != nil {
				Context.SetState(i.id, &ImageState{failure: true})

				// Trigger onFailure event
				if i.onFailure != nil {
					i.onFailure(err)
				}

				return
			}

			rgba := ImageToRgba(img)

			NewTextureFromRgba(rgba, func(tex *Texture) {
				Context.SetState(i.id, &ImageState{
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
		imgState, isOk = state.(*ImageState)
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
