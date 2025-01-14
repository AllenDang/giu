package giu

import (
	"strings"
)

// Translator should be implemented by type that can translate a string to another string.
type Translator interface {
	// Translate returns tranalted string
	Translate(string) string
	// SetLanguage changes language of the translation.
	SetLanguage(string) error
}

// SetTranslator allows to change the default (&EmptyTranslator{})
// This will raise a panic if t is nil.
// Note that using translator will change labels of widgets,
// so this might affect internal imgui's state.
// See also Context.RegisterString.
func (c *GIUContext) SetTranslator(t Translator) {
	Assert(t != nil, "Context", "SetTranslator", "Translator must not be nil.")
	c.Translator = t
}

var _ Translator = &EmptyTranslator{}

// EmptyTranslator is the default one (to save resources).
// It does nothing.
type EmptyTranslator struct{}

// Translate implements Translator interface.
func (t *EmptyTranslator) Translate(s string) string {
	return s
}

// SetLanguage implements Translator interface.
func (t *EmptyTranslator) SetLanguage(_ string) error {
	return nil
}

var _ Translator = &BasicTranslator{}

// BasicTranslator is a simpliest implementation of translation mechanism.
// This is NOT thread-safe yet. If you need thread-safety, write your own implementation.
//
// It is supposed to be used in the following way:
// - create translator
// - add languages (tip: you can use empty map for default language)
// - set translator in Context
// - set default language
// - write your UI as always.
type BasicTranslator struct {
	// language tag -> key -> value
	source          map[string]map[string]string
	currentLanguage string
}

// NewBasicTranslator creates a new BasicTranslator with the given language tag.
func NewBasicTranslator() *BasicTranslator {
	return &BasicTranslator{}
}

// Translate implements Translator interface.
// It translates s to the current language, under the following conditions:
// - If s is empty, an empty string is returned with no further processing.
// - If t.currentLanguage is empty, a panic will be raised.
// - If t.currentLanguage is not in a.source, BasicTranslator raises panic.
// - If s is not in source[currentLanguage], s is returned as-is.
func (t *BasicTranslator) Translate(s string) string {
	s = strings.Split(s, "##")[0]
	if s == "" {
		return ""
	}

	Assert(t.currentLanguage != "", "BasicTranslator", "Translate", "Current language is not set, so there is no sense in using BasicTranslator.")
	locale, ok := t.source[t.currentLanguage]
	Assert(ok, "BasicTranslator", "Translate", "There is no language tag %s known by the translator. Did you add it?", t.currentLanguage)

	translated, ok := locale[s]
	if !ok {
		return s
	}

	return translated
}

// SetLanguage sets the current language of the translator.
func (t *BasicTranslator) SetLanguage(tag string) error {
	t.currentLanguage = tag
	return nil
}

// AddLanguage adds a new "dictionary" to the translator.
func (t *BasicTranslator) AddLanguage(tag string, source map[string]string) *BasicTranslator {
	if t.source == nil {
		t.source = make(map[string]map[string]string)
	}

	t.source[tag] = source

	return t
}
