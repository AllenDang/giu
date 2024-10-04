package giu

import (
	"errors"
	"fmt"
)

// ErrNeedReset is an error indicating that the surface cannot be loaded without a reset.
// The method (*StatefulReflectiveBoundTexture).ResetState() should be called.
var ErrNeedReset = errors.New("cannot load surface without a reset. Should call (*StatefulReflectiveBoundTexture).ResetState()")

// ErrIsLoading is an error indicating that the surface state cannot be reset while loading.
var ErrIsLoading = errors.New("cannot reset surface state while loading")

// SurfaceState represents the state of the surface.
type SurfaceState int

//go:generate stringer -type=SurfaceState
const (
	// SurfaceStateNone indicates that the surface state is none.
	SurfaceStateNone SurfaceState = iota
	// SurfaceStateLoading indicates that the surface is currently loading.
	SurfaceStateLoading
	// SurfaceStateFailure indicates that the surface loading has failed.
	SurfaceStateFailure
	// SurfaceStateSuccess indicates that the surface loading was successful.
	SurfaceStateSuccess
)

// StatefulReflectiveBoundTexture is a ReflectiveBoundTexture with added async, states, and event callbacks.
type StatefulReflectiveBoundTexture struct {
	ReflectiveBoundTexture
	state     SurfaceState
	lastError error
	onReset   func()
	onLoading func()
	onSuccess func()
	onFailure func(error)
}

// GetState returns the current state of the surface.
//
// Returns:
//   - SurfaceState: The current state of the surface.
func (s *StatefulReflectiveBoundTexture) GetState() SurfaceState {
	return s.state
}

// GetLastError returns the last error that occurred during surface loading.
//
// Returns:
//   - error: The last error that occurred, or nil if no error occurred.
func (s *StatefulReflectiveBoundTexture) GetLastError() error {
	return s.lastError
}

// OnReset sets the callback function to be called when the surface state is reset.
//
// Parameters:
//   - fn: The callback function to be called on reset.
//
// Returns:
//   - *StatefulReflectiveBoundTexture: The current instance of StatefulReflectiveBoundTexture.
func (s *StatefulReflectiveBoundTexture) OnReset(fn func()) *StatefulReflectiveBoundTexture {
	s.onReset = fn
	return s
}

// OnLoading sets the callback function to be called when the surface is loading.
//
// Parameters:
//   - fn: The callback function to be called on loading.
//
// Returns:
//   - *StatefulReflectiveBoundTexture: The current instance of StatefulReflectiveBoundTexture.
func (s *StatefulReflectiveBoundTexture) OnLoading(fn func()) *StatefulReflectiveBoundTexture {
	s.onLoading = fn
	return s
}

// OnSuccess sets the callback function to be called when the surface loading is successful.
//
// Parameters:
//   - fn: The callback function to be called on success.
//
// Returns:
//   - *StatefulReflectiveBoundTexture: The current instance of StatefulReflectiveBoundTexture.
func (s *StatefulReflectiveBoundTexture) OnSuccess(fn func()) *StatefulReflectiveBoundTexture {
	s.onSuccess = fn
	return s
}

// OnFailure sets the callback function to be called when the surface loading fails.
//
// Parameters:
//   - fn: The callback function to be called on failure, with the error as a parameter.
//
// Returns:
//   - *StatefulReflectiveBoundTexture: The current instance of StatefulReflectiveBoundTexture.
func (s *StatefulReflectiveBoundTexture) OnFailure(fn func(error)) *StatefulReflectiveBoundTexture {
	s.onFailure = fn
	return s
}

// ResetState resets the state of the StatefulReflectiveBoundTexture.
//
// Returns:
//   - error: An error if the state is currently loading, otherwise nil.
func (s *StatefulReflectiveBoundTexture) ResetState() error {
	switch s.state {
	case SurfaceStateNone:
		return nil
	case SurfaceStateLoading:
		return ErrIsLoading
	default:
		s.state = SurfaceStateNone
		s.lastError = nil

		if s.onReset != nil {
			go s.onReset()
		}
	}

	return nil
}

// LoadSurfaceAsync loads the surface asynchronously using the provided SurfaceLoader.
// It sets the state to loading, and upon completion, updates the state to success or failure
// based on the result. It also triggers the appropriate callback functions.
//
// Parameters:
//   - loader: The SurfaceLoader to use for loading the surface.
//   - commit: A boolean flag indicating whether to commit the changes.
//
// Returns:
//   - error: An error if the state is not SsNone, otherwise nil.
func (s *StatefulReflectiveBoundTexture) LoadSurfaceAsync(loader SurfaceLoader, commit bool) error {
	if s.state != SurfaceStateNone {
		return ErrNeedReset
	}

	s.state = SurfaceStateLoading
	if s.onLoading != nil {
		go s.onLoading()
	}

	go func() {
		img, err := loader.ServeRGBA()
		if err != nil {
			s.state = SurfaceStateFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after loader.ServeRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		e := s.SetSurfaceFromRGBA(img, commit)
		if e != nil {
			s.state = SurfaceStateFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after SetSurfaceFromRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		s.state = SurfaceStateSuccess

		if s.onSuccess != nil {
			go s.onSuccess()
		}
	}()

	return nil
}
