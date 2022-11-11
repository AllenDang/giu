package giu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mazznoer/csscolorparser"
	"github.com/napsy/go-css"
)

// Code in this file allows to apply style using a CSS stylesheet
// The intention of the code looks as follows:
// - The CSS stylesheet is parsed and the resulting rules are stored in giu context
//   NOTE: more than one CSS stylesheets can be parsed, however the app could panic if the same tag is present twice
// - CSSTagWidget allows to apply style to a specified layout
// - main tag allows to apply style to the whole application
//
// tools:
// css parser - for now github.com/napsy/go-css. it is a bit poor, but I don't think we need anything more
// css colors - github.com/mazznoer/csscolorparser

// ParseCSSStyleSheet parses CSS stylesheet and stores the rules in giu context
func ParseCSSStyleSheet(data []byte) error {
	stylesheet, err := css.Unmarshal(data)
	if err != nil {
		return err
	}

	for rule, style := range stylesheet {
		setter := Style()
		for styleVarName, styleVarValue := range style {
			// convert style variable name to giu style variable name
			var styleVarID StyleVarID
			err := panicToErr(func() {
				styleVarID = StyleVarIDFromString(styleVarName)
			})

			if err == nil {
				// the style is StyleVarID - set it
				f, err := strconv.ParseFloat(styleVarValue, 32)
				if err == nil {
					setter.SetStyleFloat(styleVarID, float32(f))
				}

				// so maybe it is a vec2 value:
				// var-name: x, y;
				vec2 := strings.Split(styleVarValue, ",")
				if len(vec2) != 2 {
					return fmt.Errorf("unable to parse value %v is not float nor vec2: %w", styleVarValue, err)
				}

				for i, v := range vec2 {
					vec2[i] = strings.ReplaceAll(v, " ", "")
				}

				x, err := strconv.ParseFloat(vec2[0], 32)
				if err != nil {
					return fmt.Errorf("unable to parse value %v is not float: %w", vec2[0], err)
				}

				y, err := strconv.ParseFloat(vec2[1], 32)
				if err != nil {
					return fmt.Errorf("unable to parse value %v is not float: %w", vec2[1], err)
				}

				setter.SetStyle(styleVarID, float32(x), float32(y))

				continue
			}

			var styleColorID StyleColorID
			err = panicToErr(func() {
				styleColorID = StyleColorIDFromString(styleVarName)
			})

			if err != nil {
				return fmt.Errorf("cannot parse style variable ID: %v", styleVarName)
			}

			col, err := csscolorparser.Parse(styleVarValue)
			if err != nil {
				return fmt.Errorf("cannot parse color %v: %w", styleVarValue, err)
			}

			setter.SetColor(styleColorID, col)
		}

		Context.cssStylesheet[string(rule)] = setter
	}

	return nil
}

func panicToErr(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	f()
	return err
}

// cssStylesheet is a map tag:StyleSetter
type cssStylesheet map[string]*StyleSetter

var _ Widget = &CSSTagWidget{}

// CSSTagWidget is a widget that allows to apply CSS style to a specified layout
type CSSTagWidget struct {
	tag    string
	layout Layout
}

// CSSTag creates CSSTagWidget
func CSSTag(tag string) *CSSTagWidget {
	return &CSSTagWidget{tag: tag}
}

// To specifies a layout to which the style will be applied
func (c *CSSTagWidget) To(layout ...Widget) *CSSTagWidget {
	c.layout = layout
	return c
}

// Build implements Widget interface
func (c *CSSTagWidget) Build() {
	// get style from context.
	// if it doesn't exist Assert.
	style, exists := Context.cssStylesheet[c.tag]
	Assert(exists, "CSSTagWidget", "Build", "CSS stylesheet doesn't contain tag: %s", c.tag)

	style.To(c.layout).Build()
}
