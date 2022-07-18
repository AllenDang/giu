package giu

// Code in this file allows to apply style using a CSS stylesheet
// The intention of the code looks as follows:
// - The CSS stylesheet is parsed and the resulting rules are stored in giu context
//   NOTE: more than one CSS stylesheets can be parsed, however the app could panic if the same tag is present twice
// - CSSTagWidget allows to apply style to a specified layout
// - main tag allows to apply style to the whole application

// ParseCSSStyleSheet parses CSS stylesheet and stores the rules in giu context
func ParseCSSStyleSheet(css string) {

}

// cssStylesheet is a map tag:StyleSetter
type cssStylesheet map[string]StyleSetter

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
	Assert(exists, "CSSTagWidget", "Build", "CSS stylesheet doesn't contain tag: ", c.tag)

	style.To(c.layout).Build()
}
