package giu

import (
	"errors"
	"fmt"
)

var ErrNeedReset = errors.New("cannot load surface without a reset. Should call (*StatefulReflectiveBoundTexture).ResetState()")
var ErrIsLoading = errors.New("cannot reset surface state while loading")

type SurfaceState int

//go:generate stringer -type=SurfaceState
const (
	SsNone SurfaceState = iota
	SsLoading
	SsFailure
	SsSuccess
)

type StatefulReflectiveBoundTexture struct {
	ReflectiveBoundTexture
	state     SurfaceState
	lastError error
	onReset   func()
	onLoading func()
	onSuccess func()
	onFailure func(error)
}

func (s *StatefulReflectiveBoundTexture) GetState() SurfaceState {
	return s.state
}

func (s *StatefulReflectiveBoundTexture) GetLastError() error {
	return s.lastError
}

func (s *StatefulReflectiveBoundTexture) OnReset(fn func()) *StatefulReflectiveBoundTexture {
	s.onReset = fn
	return s
}

func (s *StatefulReflectiveBoundTexture) OnLoading(fn func()) *StatefulReflectiveBoundTexture {
	s.onLoading = fn
	return s
}

func (s *StatefulReflectiveBoundTexture) OnSuccess(fn func()) *StatefulReflectiveBoundTexture {
	s.onSuccess = fn
	return s
}

func (s *StatefulReflectiveBoundTexture) OnFailure(fn func(error)) *StatefulReflectiveBoundTexture {
	s.onFailure = fn
	return s
}

func (s *StatefulReflectiveBoundTexture) ResetState() error {
	switch s.state {
	case SsNone:
		return nil
	case SsLoading:
		return ErrIsLoading
	default:
		s.state = SsNone
		s.lastError = nil

		if s.onReset != nil {
			go s.onReset()
		}
	}

	return nil
}

func (s *StatefulReflectiveBoundTexture) LoadSurfaceAsync(loader SurfaceLoader, commit bool) error {
	if s.state != SsNone {
		return ErrNeedReset
	}

	s.state = SsLoading
	if s.onLoading != nil {
		go s.onLoading()
	}

	go func() {
		img, err := loader.ServeRGBA()
		if err != nil {
			s.state = SsFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after loader.ServeRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		e := s.SetSurfaceFromRGBA(img, commit)
		if e != nil {
			s.state = SsFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after SetSurfaceFromRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		s.state = SsSuccess

		if s.onSuccess != nil {
			go s.onSuccess()
		}
	}()

	return nil
}
