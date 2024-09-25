package giu

import (
	"errors"
	"fmt"
)

var errNeedReset = errors.New("cannot load surface without a reset. Should call (*StatefulReflectiveBoundTexture).ResetState()")
var errIsLoading = errors.New("cannot reset surface state while loading")

type surfaceState int

const (
	ssNone surfaceState = iota
	ssLoading
	ssFailure
	ssSuccess
)

type StatefulReflectiveBoundTexture struct {
	ReflectiveBoundTexture
	state     surfaceState
	lastError error
	onReset   func()
	onLoading func()
	onSuccess func()
	onFailure func(error)
}

func (s *StatefulReflectiveBoundTexture) GetState() surfaceState {
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
	case ssNone:
		return nil
	case ssLoading:
		return errIsLoading
	default:
		s.state = ssNone
		s.lastError = nil
		if s.onReset != nil {
			go s.onReset()
		}
	}

	return nil
}

func (s *StatefulReflectiveBoundTexture) LoadSurface(loader SurfaceLoader, commit bool) error {
	if s.state != ssNone {
		return errNeedReset
	}

	s.state = ssLoading
	if s.onLoading != nil {
		go s.onLoading()
	}

	go func() {
		img, err := loader.ServeRGBA()
		if err != nil {
			s.state = ssFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after loader.ServeRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		e := s.SetSurfaceFromRGBA(img, commit)
		if e != nil {
			s.state = ssFailure
			s.lastError = fmt.Errorf("in ReflectiveBoundTexture LoadSurface after SetSurfaceFromRGBA: %w", err)

			if s.onFailure != nil {
				go s.onFailure(s.lastError)
			}

			return
		}

		s.state = ssSuccess

		if s.onSuccess != nil {
			go s.onSuccess()
		}
	}()

	return nil
}
