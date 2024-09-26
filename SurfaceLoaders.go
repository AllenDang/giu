package giu

import (
	go_ctx "context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"net/http"
	"time"
)

// SurfaceLoader is an interface that defines a method to serve an RGBA image.
type SurfaceLoader interface {
	// ServeRGBA serves an RGBA image.
	//
	// Returns:
	//   - *image.RGBA: The RGBA image.
	//   - error: An error if the image could not be served.
	ServeRGBA() (*image.RGBA, error)
}

// SurfaceLoaderFunc is a function type that serves an RGBA image.
type SurfaceLoaderFunc func() (*image.RGBA, error)

// LoadSurfaceFunc loads a surface using a SurfaceLoaderFunc.
//
// Parameters:
//   - fn: The SurfaceLoaderFunc to use for loading the surface.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the surface could not be loaded.
func (i *ReflectiveBoundTexture) LoadSurfaceFunc(fn SurfaceLoaderFunc, commit bool) error {
	img, err := fn()
	if err != nil {
		return err
	}

	return i.SetSurfaceFromRGBA(img, commit)
}

// LoadSurface loads a surface using a SurfaceLoader.
//
// Parameters:
//   - loader: The SurfaceLoader to use for loading the surface.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the surface could not be loaded.
func (i *ReflectiveBoundTexture) LoadSurface(loader SurfaceLoader, commit bool) error {
	img, err := loader.ServeRGBA()
	if err != nil {
		return fmt.Errorf("in ReflectiveBoundTexture LoadSurface after loader.ServeRGBA: %w", err)
	}

	return i.SetSurfaceFromRGBA(img, commit)
}

// LoadSurface loads a surface asynchronously using a SurfaceLoader.
//
// Parameters:
//   - loader: The SurfaceLoader to use for loading the surface.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the surface could not be loaded.
func (s *StatefulReflectiveBoundTexture) LoadSurface(loader SurfaceLoader, commit bool) error {
	return s.LoadSurfaceAsync(loader, commit)
}

// fileLoader is a struct that implements the SurfaceLoader interface for loading images from a file.
type fileLoader struct {
	path string
}

// ServeRGBA loads an RGBA image from the file specified by the path in fileLoader.
//
// Returns:
//   - *image.RGBA: The loaded RGBA image.
//   - error: An error if the image could not be loaded.
func (f *fileLoader) ServeRGBA() (*image.RGBA, error) {
	img, err := LoadImage(f.path)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// FileLoader creates a new SurfaceLoader that loads images from the specified file path.
//
// Parameters:
//   - path: The path to the file to load the image from.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for loading images from the specified file path.
func FileLoader(path string) SurfaceLoader {
	return &fileLoader{
		path: path,
	}
}

// SetSurfaceFromFile loads an image from the specified file path and sets it as the surface of the ReflectiveBoundTexture.
//
// Parameters:
//   - path: The path to the file to load the image from.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (i *ReflectiveBoundTexture) SetSurfaceFromFile(path string, commit bool) error {
	return i.LoadSurface(FileLoader(path), commit)
}

// SetSurfaceFromFile loads an image from the specified file path and sets it as the surface of the StatefulReflectiveBoundTexture.
//
// Parameters:
//   - path: The path to the file to load the image from.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (s *StatefulReflectiveBoundTexture) SetSurfaceFromFile(path string, commit bool) error {
	return s.LoadSurface(FileLoader(path), commit)
}

// urlLoader is a SurfaceLoader that loads images from a specified URL.
type urlLoader struct {
	url     string
	timeout time.Duration
	httpdir string
}

// ServeRGBA loads an image from the URL and returns it as an RGBA image.
//
// Returns:
//   - *image.RGBA: The loaded RGBA image.
//   - error: An error if the image could not be loaded.
func (u *urlLoader) ServeRGBA() (*image.RGBA, error) {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(u.httpdir)))

	client := &http.Client{
		Transport: t,
		Timeout:   u.timeout}

	req, err := http.NewRequestWithContext(go_ctx.Background(), "GET", u.url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("urlLoader serveRGBA after http.NewRequestWithContext: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("urlLoader serveRGBA after client.Do: %w", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("urlLoader serveRGBA after image.Decode: %w", err)
	}

	return ImageToRgba(img), nil
}

// URLLoader creates a new SurfaceLoader that loads images from the specified URL.
//
// Parameters:
//   - url: The URL to load the image from.
//   - httpdir: The root directory for file:// URLs.
//   - timeout: The timeout duration for the HTTP request.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for loading images from the specified URL.
func URLLoader(url, httpdir string, timeout time.Duration) SurfaceLoader {
	return &urlLoader{
		url:     url,
		timeout: timeout,
		httpdir: httpdir,
	}
}

// SetFSRoot sets the root directory for file:// URLs.
//
// Parameters:
//   - root: The root directory to set.
func (i *ReflectiveBoundTexture) SetFSRoot(root string) {
	i.fsroot = root
}

// GetFSRoot returns the root directory for file:// URLs.
//
// Returns:
//   - string: The root directory.
func (i *ReflectiveBoundTexture) GetFSRoot() string {
	return i.fsroot
}

// SetSurfaceFromURL loads an image from the specified URL and sets it as the surface of the ReflectiveBoundTexture.
//
// Parameters:
//   - url: The URL to load the image from.
//   - timeout: The timeout duration for the HTTP request.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (i *ReflectiveBoundTexture) SetSurfaceFromURL(url string, timeout time.Duration, commit bool) error {
	return i.LoadSurface(URLLoader(url, i.fsroot, timeout), commit)
}

// SetSurfaceFromURL loads an image from the specified URL and sets it as the surface of the StatefulReflectiveBoundTexture.
//
// Parameters:
//   - url: The URL to load the image from.
//   - timeout: The timeout duration for the HTTP request.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (s *StatefulReflectiveBoundTexture) SetSurfaceFromURL(url string, timeout time.Duration, commit bool) error {
	return s.LoadSurface(URLLoader(url, s.fsroot, timeout), commit)
}

// uniformLoader is a SurfaceLoader that creates a uniform color image.
type uniformLoader struct {
	width, height int
	color         color.Color
}

// ServeRGBA creates a uniform color image and returns it as an RGBA image.
//
// Returns:
//   - *image.RGBA: The created RGBA image.
//   - error: An error if the image could not be created.
func (u *uniformLoader) ServeRGBA() (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, u.width, u.height))
	draw.Draw(img, img.Bounds(), &image.Uniform{u.color}, image.Point{}, draw.Src)

	return img, nil
}

// UniformLoader creates a new SurfaceLoader that creates a uniform color image.
//
// Parameters:
//   - width: The width of the image.
//   - height: The height of the image.
//   - c: The color of the image.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for creating a uniform color image.
func UniformLoader(width, height int, c color.Color) SurfaceLoader {
	return &uniformLoader{
		width:  width,
		height: height,
		color:  c,
	}
}

// SetSurfaceUniform creates a uniform color image and sets it as the surface of the ReflectiveBoundTexture.
//
// Parameters:
//   - width: The width of the image.
//   - height: The height of the image.
//   - c: The color of the image.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be created or set as the surface.
func (i *ReflectiveBoundTexture) SetSurfaceUniform(width, height int, c color.Color, commit bool) error {
	return i.LoadSurface(UniformLoader(width, height, c), commit)
}

// SetSurfaceUniform creates a uniform color image and sets it as the surface of the StatefulReflectiveBoundTexture.
//
// Parameters:
//   - width: The width of the image.
//   - height: The height of the image.
//   - c: The color of the image.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be created or set as the surface.
func (s *StatefulReflectiveBoundTexture) SetSurfaceUniform(width, height int, c color.Color, commit bool) error {
	return s.LoadSurface(UniformLoader(width, height, c), commit)
}
