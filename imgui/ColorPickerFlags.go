package imgui

const (
	// ColorPickerFlagsNone default = 0
	ColorPickerFlagsNone = 0
	// ColorPickerFlagsNoAlpha ignoreс Alpha component (read 3 components from the input pointer).
	ColorPickerFlagsNoAlpha = 1 << iota
	// ColorPickerFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorPickerFlagsNoSmallPreview = 1 << 4
	// ColorPickerFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorPickerFlagsNoInputs = 1 << 5
	// ColorPickerFlagsNoTooltip disables tooltip when hovering the preview.
	ColorPickerFlagsNoTooltip = 1 << 6
	// ColorPickerFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorPickerFlagsNoLabel = 1 << 7
	// ColorPickerFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorPickerFlagsNoSidePreview = 1 << 8

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions(). The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorPickerOptions() during startup.

	// ColorPickerFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorPickerFlagsAlphaBar = 1 << 16
	// ColorPickerFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreview = 1 << 17
	// ColorPickerFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreviewHalf = 1 << 18
	// ColorPickerFlagsRGB sets the format as RGB
	ColorPickerFlagsRGB = 1 << 20
	// ColorPickerFlagsHSV sets the format as HSV
	ColorPickerFlagsHSV = 1 << 21
	// ColorPickerFlagsHEX sets the format as HEX
	ColorPickerFlagsHEX = 1 << 22
	// ColorPickerFlagsUint8 _display_ values formatted as 0..255.
	ColorPickerFlagsUint8 = 1 << 23
	// ColorPickerFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorPickerFlagsFloat = 1 << 24
	// ColorPickerFlagsPickerHueBar bar for Hue, rectangle for Sat/Value.
	ColorPickerFlagsPickerHueBar = 1 << 25
	// ColorPickerFlagsPickerHueWheel wheel for Hue, triangle for Sat/Value.
	ColorPickerFlagsPickerHueWheel = 1 << 26
)
