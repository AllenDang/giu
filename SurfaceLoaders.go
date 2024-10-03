package giu

import (
	go_ctx "context"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io/fs"
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

var _ SurfaceLoader = &FileLoader{}

// FileLoader is a struct that implements the SurfaceLoader interface for loading images from a file.
type FileLoader struct {
	path string
}

// ServeRGBA loads an RGBA image from the file specified by the path in fileLoader.
//
// Returns:
//   - *image.RGBA: The loaded RGBA image.
//   - error: An error if the image could not be loaded.
func (f *FileLoader) ServeRGBA() (*image.RGBA, error) {
	img, err := LoadImage(f.path)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// NewFileLoader creates a new SurfaceLoader that loads images from the specified file path.
//
// Parameters:
//   - path: The path to the file to load the image from.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for loading images from the specified file path.
func NewFileLoader(path string) *FileLoader {
	return &FileLoader{
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
	return i.LoadSurface(NewFileLoader(path), commit)
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
	return s.LoadSurface(NewFileLoader(path), commit)
}

var _ SurfaceLoader = &FsFileLoader{}

// FsFileLoader is a struct that implements the SurfaceLoader interface for loading images from a file and embedded fs.
type FsFileLoader struct {
	file fs.File
}

// ServeRGBA loads an RGBA image from the file specified by the path in fileLoader.
//
// Returns:
//   - *image.RGBA: The loaded RGBA image.
//   - error: An error if the image could not be loaded.
func (f *FsFileLoader) ServeRGBA() (*image.RGBA, error) {
	img, err := PNGToRgba(f.file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// NewFsFileLoader creates a new SurfaceLoader that loads images from the specified file interface.
//
// Parameters:
//   - file: the file interface representing the file
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for loading images from the specified file path.
func NewFsFileLoader(file fs.File) *FsFileLoader {
	return &FsFileLoader{
		file: file,
	}
}

// SetSurfaceFromFsFile loads an image from the specified file interface and sets it as the surface of the ReflectiveBoundTexture.
//
// Parameters:
//   - file: the file interface representing the file
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (i *ReflectiveBoundTexture) SetSurfaceFromFsFile(file fs.File, commit bool) error {
	return i.LoadSurface(NewFsFileLoader(file), commit)
}

// SetSurfaceFromFsFile loads an image from the specified file interface and sets it as the surface of the StatefulReflectiveBoundTexture.
//
// Parameters:
//   - file: the file interface representing the file
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the image could not be loaded or set as the surface.
func (s *StatefulReflectiveBoundTexture) SetSurfaceFromFsFile(file fs.File, commit bool) error {
	return s.LoadSurface(NewFsFileLoader(file), commit)
}

var _ SurfaceLoader = &URLLoader{}

// URLLoader is a SurfaceLoader that loads images from a specified URL.
type URLLoader struct {
	url     string
	timeout time.Duration
	httpdir string
}

// ServeRGBA loads an image from the URL and returns it as an RGBA image.
//
// Returns:
//   - *image.RGBA: The loaded RGBA image.
//   - error: An error if the image could not be loaded.
func (u *URLLoader) ServeRGBA() (*image.RGBA, error) {
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(u.httpdir)))

	client := &http.Client{
		Transport: t,
		Timeout:   u.timeout,
	}

	req, err := http.NewRequestWithContext(go_ctx.Background(), http.MethodGet, u.url, http.NoBody)
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

// NewURLLoader creates a new SurfaceLoader that loads images from the specified URL.
//
// Parameters:
//   - url: The URL to load the image from.
//   - httpdir: The root directory for file:// URLs.
//   - timeout: The timeout duration for the HTTP request.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for loading images from the specified URL.
func NewURLLoader(url, httpdir string, timeout time.Duration) *URLLoader {
	return &URLLoader{
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
	return i.LoadSurface(NewURLLoader(url, i.fsroot, timeout), commit)
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
	return s.LoadSurface(NewURLLoader(url, s.fsroot, timeout), commit)
}

var _ SurfaceLoader = &UniformLoader{}

// UniformLoader is a SurfaceLoader that creates a uniform color image.
type UniformLoader struct {
	width, height int
	color         color.Color
}

// ServeRGBA creates a uniform color image and returns it as an RGBA image.
//
// Returns:
//   - *image.RGBA: The created RGBA image.
//   - error: An error if the image could not be created.
func (u *UniformLoader) ServeRGBA() (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, u.width, u.height))
	draw.Draw(img, img.Bounds(), &image.Uniform{u.color}, image.Point{}, draw.Src)

	return img, nil
}

// NewUniformLoader creates a new SurfaceLoader that creates a uniform color image.
//
// Parameters:
//   - width: The width of the image.
//   - height: The height of the image.
//   - c: The color of the image.
//
// Returns:
//   - SurfaceLoader: A new SurfaceLoader for creating a uniform color image.
func NewUniformLoader(width, height int, c color.Color) *UniformLoader {
	return &UniformLoader{
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
	return i.LoadSurface(NewUniformLoader(width, height, c), commit)
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
	return s.LoadSurface(NewUniformLoader(width, height, c), commit)
}
