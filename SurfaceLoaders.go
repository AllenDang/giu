package giu

import (
	go_ctx "context"
	"image"
	"image/color"
	"image/draw"
	"net/http"
	"time"
)

// SurfaceLoader interface
type SurfaceLoader interface {
	ServeRGBA() (*image.RGBA, error)
}
type SurfaceLoaderFunc func() (*image.RGBA, error)

func (i *ReflectiveBoundTexture) LoadSurfaceFunc(fn SurfaceLoaderFunc, commit bool) error {
	img, err := fn()
	if err != nil {
		return err
	}
	return i.SetSurfaceFromRGBA(img, commit)
}

func (i *ReflectiveBoundTexture) LoadSurface(loader SurfaceLoader, commit bool) error {
	img, err := loader.ServeRGBA()
	if err != nil {
		return err
	}
	return i.SetSurfaceFromRGBA(img, commit)
}

// FileLoader

type fileLoader struct {
	path string
}

func (f *fileLoader) ServeRGBA() (*image.RGBA, error) {
	img, err := LoadImage(f.path)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func FileLoader(path string) SurfaceLoader {
	return &fileLoader{
		path: path,
	}
}

func (i *ReflectiveBoundTexture) SetSurfaceFromFile(path string, commit bool) error {
	return i.LoadSurface(FileLoader(path), commit)
}

// UrlLoader

type urlLoader struct {
	url     string
	timeout time.Duration
}

func (u *urlLoader) ServeRGBA() (*image.RGBA, error) {
	client := &http.Client{Timeout: u.timeout}
	req, err := http.NewRequestWithContext(go_ctx.Background(), "GET", u.url, http.NoBody)
	if err != nil {
		//errorFn(err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		//errorFn(err)
		return nil, err
	}
	defer resp.Body.Close()
	/*defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			errorFn(closeErr)
		}
	}()*/
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		//errorFn(err)
		return nil, err
	}
	return ImageToRgba(img), nil
}

func URLLoader(url string, timeout time.Duration) SurfaceLoader {
	return &urlLoader{
		url:     url,
		timeout: timeout,
	}
}

func (i *ReflectiveBoundTexture) SetSurfaceFromURL(url string, timeout time.Duration, commit bool) error {
	return i.LoadSurface(URLLoader(url, timeout), commit)
}

// UniformLoader
type uniformLoader struct {
	width, height int
	color         color.Color
}

func (u *uniformLoader) ServeRGBA() (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, u.width, u.height))
	draw.Draw(img, img.Bounds(), &image.Uniform{u.color}, image.ZP, draw.Src)
	return img, nil
}

func UniformLoader(width, height int, color color.Color) SurfaceLoader {
	return &uniformLoader{
		width:  width,
		height: height,
		color:  color,
	}
}

func (i *ReflectiveBoundTexture) SetSurfaceUniform(width int, height int, c color.Color, commit bool) error {
	return i.LoadSurface(UniformLoader(width, height, c), commit)
}
