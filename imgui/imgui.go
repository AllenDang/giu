package imgui

// #cgo CXXFLAGS: -std=c++11
// #include "imguiWrapper.h"
import "C"
import (
	"math"
)

// Version returns a version string e.g. "1.23".
func Version() string {
	return C.GoString(C.iggGetVersion())
}

// CurrentIO returns access to the ImGui communication struct for the currently active context.
func CurrentIO() IO {
	return IO{handle: C.iggGetCurrentIO()}
}

// CurrentStyle returns the UI Style for the currently active context.
func CurrentStyle() Style {
	return Style(C.iggGetCurrentStyle())
}

// NewFrame starts a new ImGui frame, you can submit any command from this point until Render()/EndFrame().
func NewFrame() {
	C.iggNewFrame()
}

// Render ends the ImGui frame, finalize the draw data.
// After this method, call RenderedDrawData to retrieve the draw commands and execute them.
func Render() {
	C.iggRender()
}

// EndFrame ends the ImGui frame. Automatically called by Render(), so most likely don't need to ever
// call that yourself directly. If you don't need to render you may call EndFrame() but you'll have
// wasted CPU already. If you don't need to render, better to not create any imgui windows instead!
func EndFrame() {
	C.iggEndFrame()
}

func GetEventWaitingTime() float64 {
	return float64(C.iggGetEventWaitingTime())
}

// RenderedDrawData returns the created draw commands, which are valid after Render() and
// until the next call to NewFrame(). This is what you have to render.
func RenderedDrawData() DrawData {
	return DrawData(C.iggGetDrawData())
}

// ShowDemoWindow creates a demo/test window. Demonstrates most ImGui features.
// Call this to learn about the library! Try to make it always available in your application!
func ShowDemoWindow(open *bool) {
	openArg, openFin := wrapBool(open)
	defer openFin()
	C.iggShowDemoWindow(openArg)
}

// ShowUserGuide adds basic help/info block (not a window): how to manipulate ImGui as a end-user (mouse/keyboard controls).
func ShowUserGuide() {
	C.iggShowUserGuide()
}

// BeginV pushes a new window to the stack and start appending to it.
// You may append multiple times to the same window during the same frame.
// If the open argument is provided, the window can be closed, in which case the value will be false after the call.
//
// Returns false if the window is currently not visible.
// Regardless of the return value, End() must be called for each call to Begin().
func BeginV(id string, open *bool, flags int) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBegin(idArg, openArg, C.int(flags)) != 0
}

// Begin calls BeginV(id, nil, 0).
func Begin(id string) bool {
	return BeginV(id, nil, 0)
}

// End closes the scope for the previously opened window.
// Every call to Begin() must be matched with a call to End().
func End() {
	C.iggEnd()
}

// BeginChildV pushes a new child to the stack and starts appending to it.
// flags are the WindowFlags to apply.
func BeginChildV(id string, size Vec2, border bool, flags int) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggBeginChild(idArg, sizeArg, castBool(border), C.int(flags)) != 0
}

// BeginChild calls BeginChildV(id, Vec2{0,0}, false, 0).
func BeginChild(id string) bool {
	return BeginChildV(id, Vec2{}, false, 0)
}

// EndChild closes the scope for the previously opened child.
// Every call to BeginChild() must be matched with a call to EndChild().
func EndChild() {
	C.iggEndChild()
}

// WindowPos returns the current window position in screen space.
// This is useful if you want to do your own drawing via the DrawList API.
func WindowPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowPos(valueArg)
	valueFin()
	return value
}

// WindowSize returns the size of the current window.
func WindowSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowSize(valueArg)
	valueFin()
	return value
}

// WindowWidth returns the width of the current window.
func WindowWidth() float32 {
	return float32(C.iggWindowWidth())
}

// WindowHeight returns the height of the current window.
func WindowHeight() float32 {
	return float32(C.iggWindowHeight())
}

// ContentRegionAvail returns the size of the content region that is available (based on the current cursor position).
func ContentRegionAvail() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggContentRegionAvail(valueArg)
	valueFin()
	return value
}

func IsWindowAppearing() bool {
	return C.iggIsWindowAppearing() != 0
}

// SetNextWindowPosV sets next window position.
// Call before Begin(). Use pivot=(0.5,0.5) to center on given point, etc.
func SetNextWindowPosV(pos Vec2, cond Condition, pivot Vec2) {
	posArg, _ := pos.wrapped()
	pivotArg, _ := pivot.wrapped()
	C.iggSetNextWindowPos(posArg, C.int(cond), pivotArg)
}

// SetNextWindowPos calls SetNextWindowPosV(pos, 0, Vec{0,0})
func SetNextWindowPos(pos Vec2) {
	SetNextWindowPosV(pos, 0, Vec2{})
}

// SetNextWindowSizeV sets next window size.
// Set axis to 0.0 to force an auto-fit on this axis. Call before Begin().
func SetNextWindowSizeV(size Vec2, cond Condition) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowSize(sizeArg, C.int(cond))
}

// SetNextWindowSize calls SetNextWindowSizeV(size, 0)
func SetNextWindowSize(size Vec2) {
	SetNextWindowSizeV(size, 0)
}

// SetNextWindowContentSize sets next window content size (~ enforce the range of scrollbars).
// Does not include window decorations (title bar, menu bar, etc.).
// Set one axis to 0.0 to leave it automatic. This function must be called before Begin() to take effect.
func SetNextWindowContentSize(size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowContentSize(sizeArg)
}

// SetNextWindowFocus sets next window to be focused / front-most. Call before Begin().
func SetNextWindowFocus() {
	C.iggSetNextWindowFocus()
}

// SetNextWindowBgAlpha sets next window background color alpha.
// Helper to easily modify ImGuiCol_WindowBg/ChildBg/PopupBg.
func SetNextWindowBgAlpha(value float32) {
	C.iggSetNextWindowBgAlpha(C.float(value))
}

// PushFont adds the given font on the stack. Use DefaultFont to refer to the default font.
func PushFont(font Font) {
	C.iggPushFont(font.handle())
}

// PopFont removes the previously pushed font from the stack.
func PopFont() {
	C.iggPopFont()
}

// PushStyleColor pushes the current style color for given ID on a stack and sets the given one.
// To revert to the previous color, call PopStyleColor().
func PushStyleColor(id StyleColorID, color Vec4) {
	colorArg, _ := color.wrapped()
	C.iggPushStyleColor(C.int(id), colorArg)
}

// PopStyleColorV reverts the given amount of style color changes.
func PopStyleColorV(count int) {
	C.iggPopStyleColor(C.int(count))
}

// PopStyleColor calls PopStyleColorV(1).
func PopStyleColor() {
	PopStyleColorV(1)
}

// PushStyleVarFloat pushes a float value on the stack to temporarily modify a style variable.
func PushStyleVarFloat(id StyleVarID, value float32) {
	C.iggPushStyleVarFloat(C.int(id), C.float(value))
}

// PushStyleVarVec2 pushes a Vec2 value on the stack to temporarily modify a style variable.
func PushStyleVarVec2(id StyleVarID, value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggPushStyleVarVec2(C.int(id), valueArg)
}

// PopStyleVarV reverts the given amount of style variable changes.
func PopStyleVarV(count int) {
	C.iggPopStyleVar(C.int(count))
}

// PopStyleVar calls PopStyleVarV(1).
func PopStyleVar() {
	PopStyleVarV(1)
}

// FontSize returns the current font size (= height in pixels) of the current font with the current scale applied.
func FontSize() float32 {
	return float32(C.iggGetFontSize())
}

// CalcTextSize calculate the size of the text
func CalcTextSize(text string, hideTextAfterDoubleHash bool, wrapWidth float32) Vec2 {
	CString := newStringBuffer(text)
	defer CString.free()

	var vec2 Vec2
	valueArg, returnFunc := vec2.wrapped()

	C.iggCalcTextSize((*C.char)(CString.ptr), C.int(CString.size), castBool(hideTextAfterDoubleHash), C.float(wrapWidth), valueArg)
	returnFunc()

	return vec2
}

// Get ColorU32 from Vec4
func GetColorU32(col Vec4) uint {
	valueArg, _ := col.wrapped()
	return uint(C.iggGetColorU32(*valueArg))
}

func StyleColorsDark() {
	C.iggStyleColorsDark()
}

func StyleColorsClassic() {
	C.iggStyleColorsClassic()
}

func StyleColorsLight() {
	C.iggStyleColorsLight()
}

// PushItemWidth sets width of items for the common item+label case, in pixels.
// 0.0f = default to ~2/3 of windows width, >0.0f: width in pixels,
// <0.0f align xx pixels to the right of window (so -1.0f always align width to the right side).
func PushItemWidth(width float32) {
	C.iggPushItemWidth(C.float(width))
}

// PopItemWidth must be called for each call to PushItemWidth().
func PopItemWidth() {
	C.iggPopItemWidth()
}

// CalcItemWidth returns the width of items given pushed settings and current cursor position.
func CalcItemWidth() float32 {
	return float32(C.iggCalcItemWidth())
}

// PushTextWrapPosV defines word-wrapping for Text() commands.
// < 0.0f: no wrapping; 0.0f: wrap to end of window (or column); > 0.0f: wrap at 'wrapPosX' position in window local space.
// Requires a matching call to PopTextWrapPos().
func PushTextWrapPosV(wrapPosX float32) {
	C.iggPushTextWrapPos(C.float(wrapPosX))
}

// PushTextWrapPos calls PushTextWrapPosV(0.0).
func PushTextWrapPos() {
	PushTextWrapPosV(0.0)
}

// PopTextWrapPos resets the last pushed position.
func PopTextWrapPos() {
	C.iggPopTextWrapPos()
}

// PushID pushes the given identifier into the ID stack. IDs are hash of the entire stack!
func PushID(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggPushID(idArg)
}

// PopID removes the last pushed identifier from the ID stack.
func PopID() {
	C.iggPopID()
}

// Text adds formatted text. See PushTextWrapPosV() or PushStyleColorV() for modifying the output.
// Without any modified style stack, the text is unformatted.
func Text(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	// Internally we use ImGui::TextUnformatted, for the most direct call.
	C.iggTextUnformatted(textArg)
}

// LabelText adds text+label aligned the same way as value+label widgets.
func LabelText(label, text string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggLabelText(labelArg, textArg)
}

// ButtonV returning true if it is pressed.
func ButtonV(id string, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggButton(idArg, sizeArg) != 0
}

// Button calls ButtonV(id, Vec2{0,0}).
func Button(id string) bool {
	return ButtonV(id, Vec2{})
}

func InvisibleButton(id string, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggInvisibleButton(idArg, sizeArg) != 0
}

// ImageV adds an image based on given texture ID.
// Refer to TextureID what this represents and how it is drawn.
func ImageV(id TextureID, size Vec2, uv0, uv1 Vec2, tintCol, borderCol Vec4) {
	sizeArg, _ := size.wrapped()
	uv0Arg, _ := uv0.wrapped()
	uv1Arg, _ := uv1.wrapped()
	tintColArg, _ := tintCol.wrapped()
	borderColArg, _ := borderCol.wrapped()
	C.iggImage(id.handle(), sizeArg, uv0Arg, uv1Arg, tintColArg, borderColArg)
}

// Image calls ImageV(id, size, Vec2{0,0}, Vec2{1,1}, Vec4{1,1,1,1}, Vec4{0,0,0,0}).
func Image(id TextureID, size Vec2) {
	ImageV(id, size, Vec2{X: 0, Y: 0}, Vec2{X: 1, Y: 1}, Vec4{X: 1, Y: 1, Z: 1, W: 1}, Vec4{X: 0, Y: 0, Z: 0, W: 0})
}

// ImageButtonV adds a button with an image, based on given texture ID.
// Refer to TextureID what this represents and how it is drawn.
// <0 framePadding uses default frame padding settings. 0 for no padding.
func ImageButtonV(id TextureID, size Vec2, uv0, uv1 Vec2, framePadding int, bgCol Vec4, tintCol Vec4) bool {
	sizeArg, _ := size.wrapped()
	uv0Arg, _ := uv0.wrapped()
	uv1Arg, _ := uv1.wrapped()
	bgColArg, _ := bgCol.wrapped()
	tintColArg, _ := tintCol.wrapped()
	return C.iggImageButton(id.handle(), sizeArg, uv0Arg, uv1Arg, C.int(framePadding), bgColArg, tintColArg) != 0
}

// ImageButton calls ImageButtonV(id, size, Vec2{0,0}, Vec2{1,1}, -1, Vec4{0,0,0,0}, Vec4{1,1,1,1}).
func ImageButton(id TextureID, size Vec2) bool {
	return ImageButtonV(id, size, Vec2{X: 0, Y: 0}, Vec2{X: 1, Y: 1}, -1, Vec4{X: 0, Y: 0, Z: 0, W: 0}, Vec4{X: 1, Y: 1, Z: 1, W: 1})
}

// Checkbox creates a checkbox in the selected state.
// The return value indicates if the selected state has changed.
func Checkbox(id string, selected *bool) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	selectedArg, selectedFin := wrapBool(selected)
	defer selectedFin()
	return C.iggCheckbox(idArg, selectedArg) != 0
}

func RadioButton(label string, active bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggRadioButton(labelArg, castBool(active)) != 0
}

// ProgressBarV creates a progress bar.
// size (for each axis) < 0.0f: align to end, 0.0f: auto, > 0.0f: specified size
func ProgressBarV(fraction float32, size Vec2, overlay string) {
	sizeArg, _ := size.wrapped()
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()
	C.iggProgressBar(C.float(fraction), sizeArg, overlayArg)
}

// ProgressBar calls ProgressBarV(fraction, Vec2{X: -1, Y: 0}, "").
func ProgressBar(fraction float32) {
	ProgressBarV(fraction, Vec2{X: -1, Y: 0}, "")
}

// BeginComboV creates a combo box with complete control over the content to the user.
// Call EndCombo() if this function returns true.
// flags are the ComboFlags to apply.
func BeginComboV(label, previewValue string, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	previewValueArg, previewValueFin := wrapString(previewValue)
	defer previewValueFin()
	return C.iggBeginCombo(labelArg, previewValueArg, C.int(flags)) != 0
}

// BeginCombo calls BeginComboV(label, previewValue, 0).
func BeginCombo(label, previewValue string) bool {
	return BeginComboV(label, previewValue, 0)
}

// EndCombo must be called if BeginComboV() returned true.
func EndCombo() {
	C.iggEndCombo()
}

// DragFloatV creates a draggable slider for floats.
func DragFloatV(label string, value *float32, speed, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragFloat(labelArg, valueArg, C.float(speed), C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// DragFloat calls DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", 1.0).
func DragFloat(label string, value *float32) bool {
	return DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", 1.0)
}

// DragIntV creates a draggable slider for integers.
func DragIntV(label string, value *int32, speed float32, min, max int32, format string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragInt(labelArg, valueArg, C.float(speed), C.int(min), C.int(max), formatArg) != 0
}

// DragInt calls DragIntV(label, value, 1.0, 0, 0, "%d").
func DragInt(label string, value *int32) bool {
	return DragIntV(label, value, 1.0, 0, 0, "%d")
}

// SliderFloatV creates a slider for floats.
func SliderFloatV(label string, value *float32, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderFloat(labelArg, valueArg, C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// SliderFloat calls SliderIntV(label, value, min, max, "%.3f", 1.0).
func SliderFloat(label string, value *float32, min, max float32) bool {
	return SliderFloatV(label, value, min, max, "%.3f", 1.0)
}

// SliderFloat3V creates slider for a 3D vector.
func SliderFloat3V(label string, values *[3]float32, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggSliderFloatN(labelArg, cvalues, 3, C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// SliderFloat3 calls SliderFloat3V(label, values, min, max, "%.3f", 1,0).
func SliderFloat3(label string, values *[3]float32, min, max float32) bool {
	return SliderFloat3V(label, values, min, max, "%.3f", 1.0)
}

// SliderIntV creates a slider for integers.
func SliderIntV(label string, value *int32, min, max int32, format string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderInt(labelArg, valueArg, C.int(min), C.int(max), formatArg) != 0
}

// SliderInt calls SliderIntV(label, value, min, max, "%d").
func SliderInt(label string, value *int32, min, max int32) bool {
	return SliderIntV(label, value, min, max, "%d")
}

// InputTextV creates a text field for dynamic text input.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextV(label string, text *string, flags int, cb InputTextCallback) bool {
	if text == nil {
		panic("text can't be nil")
	}
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	state := newInputTextState(*text, cb)
	defer func() {
		*text = state.buf.toGo()
		state.release()
	}()

	return C.iggInputText(labelArg, (*C.char)(state.buf.ptr), C.uint(state.buf.size),
		C.int(flags|inputTextFlagsCallbackResize), state.key) != 0
}

// InputText calls InputTextV(label, string, 0, nil)
func InputText(label string, text *string) bool {
	return InputTextV(label, text, 0, nil)
}

// InputTextMultilineV provides a field for dynamic text input of multiple lines.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextMultilineV(label string, text *string, size Vec2, flags int, cb InputTextCallback) bool {
	if text == nil {
		panic("text can't be nil")
	}
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	state := newInputTextState(*text, cb)
	defer func() {
		*text = state.buf.toGo()
		state.release()
	}()

	return C.iggInputTextMultiline(labelArg, (*C.char)(state.buf.ptr), C.uint(state.buf.size), sizeArg,
		C.int(flags|inputTextFlagsCallbackResize), state.key) != 0
}

// InputTextMultiline calls InputTextMultilineV(label, text, Vec2{0,0}, 0, nil)
func InputTextMultiline(label string, text *string) bool {
	return InputTextMultilineV(label, text, Vec2{}, 0, nil)
}

func InputInt(label string, value *int32) bool {
	return InputIntV(label, value, 0, 100, 0)
}

func InputIntV(label string, value *int32, step, step_fast int, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valueArg, valueFin := wrapInt32(value)
	defer valueFin()

	return C.iggInputInt(labelArg, valueArg, C.int(0), C.int(100), C.int(flags)) != 0
}

func InputFloat(label string, value *float32) bool {
	return InputFloatV(label, value, 0.0, 0.0, "%.3f", 0)
}

func InputFloatV(label string, value *float32, step, step_fast float32, format string, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()

	return C.iggInputFloat(labelArg, valueArg, C.float(step), C.float(step_fast), formatArg, C.int(flags)) != 0
}

// ColorEdit3 calls ColorEdit3V(label, col, 0)
func ColorEdit3(label string, col *[3]float32) bool {
	return ColorEdit3V(label, col, 0)
}

// ColorEdit3V will show a clickable little square which will open a color picker window for 3D vector (rgb format).
func ColorEdit3V(label string, col *[3]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit3(labelArg, ccol, C.int(flags)) != 0
}

// ColorEdit4 calls ColorEdit4V(label, col, 0)
func ColorEdit4(label string, col *[4]float32) bool {
	return ColorEdit4V(label, col, 0)
}

// ColorEdit4V will show a clickable little square which will open a color picker window for 4D vector (rgba format).
func ColorEdit4V(label string, col *[4]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit4(labelArg, ccol, C.int(flags)) != 0
}

// ColorPicker3 calls ColorPicker3(label, col, 0)
func ColorPicker3(label string, col *[3]float32, flags int) bool {
	return ColorPicker3V(label, col, 0)
}

// ColorPicker3V will show directly a color picker control for editing a color in 3D vector (rgb format).
func ColorPicker3V(label string, col *[3]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker3(labelArg, ccol, C.int(flags)) != 0
}

// ColorPicker4 calls ColorPicker4(label, col, 0)
func ColorPicker4(label string, col *[4]float32) bool {
	return ColorPicker4V(label, col, 0)
}

// ColorPicker4V will show directly a color picker control for editing a color in 4D vector (rgba format).
func ColorPicker4V(label string, col *[4]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker4(labelArg, ccol, C.int(flags)) != 0
}

// Separator is generally horizontal. Inside a menu bar or in horizontal layout mode, this becomes a vertical separator.
func Separator() {
	C.iggSeparator()
}

// SameLineV is between widgets or groups to layout them horizontally.
func SameLineV(posX float32, spacingW float32) {
	C.iggSameLine(C.float(posX), C.float(spacingW))
}

// SameLine calls SameLineV(0, -1).
func SameLine() {
	SameLineV(0, -1)
}

// Spacing adds vertical spacing.
func Spacing() {
	C.iggSpacing()
}

// Dummy adds a dummy item of given size.
func Dummy(size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggDummy(sizeArg)
}

// BeginGroup locks horizontal starting position + capture group bounding box into one "item"
// (so you can use IsItemHovered() or layout primitives such as SameLine() on whole group, etc.)
func BeginGroup() {
	C.iggBeginGroup()
}

// EndGroup must be called for each call to BeginGroup().
func EndGroup() {
	C.iggEndGroup()
}

// CursorPos returns the cursor position in window coordinates (relative to window position).
func CursorPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorPos(valueArg)
	valueFin()
	return value
}

// CursorPosX returns the x-coordinate of the cursor position in window coordinates.
func CursorPosX() float32 {
	return float32(C.iggCursorPosX())
}

// CursorPosY returns the y-coordinate of the cursor position in window coordinates.
func CursorPosY() float32 {
	return float32(C.iggCursorPosY())
}

// CursorStartPos returns the initial cursor position in window coordinates.
func CursorStartPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorStartPos(valueArg)
	valueFin()
	return value
}

// CursorScreenPos returns the cursor position in absolute screen coordinates.
func CursorScreenPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorScreenPos(valueArg)
	valueFin()
	return value
}

// SetCursorPos sets the cursor relative to the current window.
func SetCursorPos(localPos Vec2) {
	localPosArg, _ := localPos.wrapped()
	C.iggSetCursorPos(localPosArg)
}

// SetCursorScreenPos sets the cursor position in absolute screen coordinates.
func SetCursorScreenPos(absPos Vec2) {
	absPosArg, _ := absPos.wrapped()
	C.iggSetCursorScreenPos(absPosArg)
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	C.iggAlignTextToFramePadding()
}

// TextLineHeight returns ~ FontSize.
func TextLineHeight() float32 {
	return float32(C.iggGetTextLineHeight())
}

// TextLineHeightWithSpacing returns ~ FontSize + style.ItemSpacing.y (distance in pixels between 2 consecutive lines of text).
func TextLineHeightWithSpacing() float32 {
	return float32(C.iggGetTextLineHeightWithSpacing())
}

// TreeNodeV returns true if the tree branch is to be rendered. Call TreePop() in this case.
func TreeNodeV(label string, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggTreeNode(labelArg, C.int(flags)) != 0
}

// TreeNode calls TreeNodeV(label, 0).
func TreeNode(label string) bool {
	return TreeNodeV(label, 0)
}

// TreePop finishes a tree branch. This has to be called for a matching TreeNodeV call returning true.
func TreePop() {
	C.iggTreePop()
}

// SetNextItemOpen sets the open/collapsed state of the following tree node.
func SetNextItemOpen(open bool, cond Condition) {
	C.iggSetNextItemOpen(castBool(open), C.int(cond))
}

// TreeNodeToLabelSpacing returns the horizontal distance preceding label for a regular unframed TreeNode.
func TreeNodeToLabelSpacing() float32 {
	return float32(C.iggGetTreeNodeToLabelSpacing())
}

// SelectableV returns true if the user clicked it, so you can modify your selection state.
// flags are the SelectableFlags to apply.
// size.x==0.0: use remaining width, size.x>0.0: specify width.
// size.y==0.0: use label height, size.y>0.0: specify height
func SelectableV(label string, selected bool, flags int, size Vec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	return C.iggSelectable(labelArg, castBool(selected), C.int(flags), sizeArg) != 0
}

// Selectable calls SelectableV(label, false, 0, Vec2{0, 0})
func Selectable(label string) bool {
	return SelectableV(label, false, 0, Vec2{})
}

// ListBoxV creates a list of selectables of given items with equal height, enclosed with header and footer.
// This version accepts a custom item height.
// The function returns true if the selection was changed. The value of currentItem will indicate the new selected item.
func ListBoxV(label string, currentItem *int32, items []string, heightItems int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valueArg, valueFin := wrapInt32(currentItem)
	defer valueFin()

	itemsCount := len(items)

	argv := make([]*C.char, itemsCount)
	for i, item := range items {
		itemArg, itemDeleter := wrapString(item)
		defer itemDeleter()
		argv[i] = itemArg
	}

	return C.iggListBoxV(labelArg, valueArg, &argv[0], C.int(itemsCount), C.int(heightItems)) != 0
}

// ListBox calls ListBoxV(label, currentItem, items, -1)
// The function returns true if the selection was changed. The value of currentItem will indicate the new selected item.
func ListBox(label string, currentItem *int32, items []string) bool {
	return ListBoxV(label, currentItem, items, -1)
}

// PlotLines draws an array of floats as a line graph.
// It calls PlotLinesV using no overlay text and automatically calculated scale and graph size.
func PlotLines(label string, values []float32) {
	PlotLinesV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, Vec2{})
}

// PlotLinesV draws an array of floats as a line graph with additional options.
// valuesOffset specifies an offset into the values array at which to start drawing, wrapping around when the end of the values array is reached.
// overlayText specifies a string to print on top of the graph.
// scaleMin and scaleMax define the scale of the y axis, if either is math.MaxFloat32 that value is calculated from the input data.
// graphSize defines the size of the graph, if either coordinate is zero the default size for that direction is used.
func PlotLinesV(label string, values []float32, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valuesCount := len(values)
	valuesArray := make([]C.float, valuesCount)
	for i, value := range values {
		valuesArray[i] = C.float(value)
	}

	var overlayTextArg *C.char
	if overlayText != "" {
		var overlayTextFinisher func()
		overlayTextArg, overlayTextFinisher = wrapString(overlayText)
		defer overlayTextFinisher()
	}

	graphSizeArg, _ := graphSize.wrapped()

	C.iggPlotLines(labelArg, &valuesArray[0], C.int(valuesCount), C.int(valuesOffset), overlayTextArg, C.float(scaleMin), C.float(scaleMax), graphSizeArg)
}

// PlotHistogram draws an array of floats as a bar graph.
// It calls PlotHistogramV using no overlay text and automatically calculated scale and graph size.
func PlotHistogram(label string, values []float32) {
	PlotHistogramV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, Vec2{})
}

// PlotHistogramV draws an array of floats as a bar graph with additional options.
// valuesOffset specifies an offset into the values array at which to start drawing, wrapping around when the end of the values array is reached.
// overlayText specifies a string to print on top of the graph.
// scaleMin and scaleMax define the scale of the y axis, if either is math.MaxFloat32 that value is calculated from the input data.
// graphSize defines the size of the graph, if either coordinate is zero the default size for that direction is used.
func PlotHistogramV(label string, values []float32, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valuesCount := len(values)
	valuesArray := make([]C.float, valuesCount)
	for i, value := range values {
		valuesArray[i] = C.float(value)
	}

	var overlayTextArg *C.char
	if overlayText != "" {
		var overlayTextFinisher func()
		overlayTextArg, overlayTextFinisher = wrapString(overlayText)
		defer overlayTextFinisher()
	}

	graphSizeArg, _ := graphSize.wrapped()

	C.iggPlotHistogram(labelArg, &valuesArray[0], C.int(valuesCount), C.int(valuesOffset), overlayTextArg, C.float(scaleMin), C.float(scaleMax), graphSizeArg)
}

// SetTooltip sets a text tooltip under the mouse-cursor, typically use with IsItemHovered().
// Overrides any previous call to SetTooltip().
func SetTooltip(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggSetTooltip(textArg)
}

// BeginTooltip begins/appends to a tooltip window. Used to create full-featured tooltip (with any kind of contents).
// Requires a call to EndTooltip().
func BeginTooltip() {
	C.iggBeginTooltip()
}

// EndTooltip closes the previously started tooltip window.
func EndTooltip() {
	C.iggEndTooltip()
}

// BeginMainMenuBar creates and appends to a full screen menu-bar.
// If the return value is true, then EndMainMenuBar() must be called!
func BeginMainMenuBar() bool {
	return C.iggBeginMainMenuBar() != 0
}

// EndMainMenuBar finishes a main menu bar.
// Only call EndMainMenuBar() if BeginMainMenuBar() returns true!
func EndMainMenuBar() {
	C.iggEndMainMenuBar()
}

// BeginMenuBar appends to menu-bar of current window.
// This requires WindowFlagsMenuBar flag set on parent window.
// If the return value is true, then EndMenuBar() must be called!
func BeginMenuBar() bool {
	return C.iggBeginMenuBar() != 0
}

// EndMenuBar finishes a menu bar.
// Only call EndMenuBar() if BeginMenuBar() returns true!
func EndMenuBar() {
	C.iggEndMenuBar()
}

// BeginMenuV creates a sub-menu entry.
// If the return value is true, then EndMenu() must be called!
func BeginMenuV(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggBeginMenu(labelArg, castBool(enabled)) != 0
}

// BeginMenu calls BeginMenuV(label, true).
func BeginMenu(label string) bool {
	return BeginMenuV(label, true)
}

// EndMenu finishes a sub-menu entry.
// Only call EndMenu() if BeginMenu() returns true!
func EndMenu() {
	C.iggEndMenu()
}

// MenuItemV adds a menu item with given label.
// Returns true if the item is selected.
// If selected is not nil, it will be toggled when true is returned.
// Shortcuts are displayed for convenience but not processed by ImGui at the moment.
func MenuItemV(label string, shortcut string, selected bool, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	shortcutArg, shortcutFin := wrapString(shortcut)
	defer shortcutFin()
	return C.iggMenuItem(labelArg, shortcutArg, castBool(selected), castBool(enabled)) != 0
}

// MenuItem calls MenuItemV(label, "", false, true).
func MenuItem(label string) bool {
	return MenuItemV(label, "", false, true)
}

// OpenPopup marks popup as open (don't call every frame!).
// Popups are closed when user click outside, or if CloseCurrentPopup() is called within a BeginPopup()/EndPopup() block.
// By default, Selectable()/MenuItem() are calling CloseCurrentPopup().
// Popup identifiers are relative to the current ID-stack (so OpenPopup and BeginPopup needs to be at the same level).
func OpenPopup(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggOpenPopup(idArg)
}

// BeginPopupModalV creates modal dialog (regular window with title bar, block interactions behind the modal window,
// can't close the modal window by clicking outside).
func BeginPopupModalV(name string, open *bool, flags int) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBeginPopupModal(nameArg, openArg, C.int(flags)) != 0
}

// BeginPopupModal calls BeginPopupModalV(name, nil, 0)
func BeginPopupModal(name string) bool {
	return BeginPopupModalV(name, nil, 0)
}

// BeginPopupContextItemV returns true if the identified mouse button was pressed
// while hovering over the last item.
func BeginPopupContextItemV(label string, mouseButton int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggBeginPopupContextItem(labelArg, C.int(mouseButton)) != 0
}

// BeginPopupContextItem calls BeginPopupContextItemV("", 1)
func BeginPopupContextItem() bool {
	return BeginPopupContextItemV("", 1)
}

// EndPopup finshes a popup. Only call EndPopup() if BeginPopupXXX() returns true!
func EndPopup() {
	C.iggEndPopup()
}

// CloseCurrentPopup closes the popup we have begin-ed into.
// Clicking on a MenuItem or Selectable automatically close the current popup.
func CloseCurrentPopup() {
	C.iggCloseCurrentPopup()
}

// IsItemHoveredV returns true if the last item is hovered.
// (and usable, aka not blocked by a popup, etc.). See HoveredFlags for more options.
func IsItemHoveredV(flags int) bool {
	return C.iggIsItemHovered(C.int(flags)) != 0
}

// IsItemHovered calls IsItemHoveredV(HoveredFlagsNone)
func IsItemHovered() bool {
	return IsItemHoveredV(HoveredFlagsNone)
}

func IsItemActive() bool {
	return C.iggIsItemActive() != 0
}

// IsAnyItemActive returns true if the any item is active.
func IsAnyItemActive() bool {
	return C.iggIsAnyItemActive() != 0
}

// IsKeyDown returns true if the corresponding key is currently being held down.
func IsKeyDown(key int) bool {
	return C.iggIsKeyDown(C.int(key)) != 0
}

// IsKeyPressedV returns true if the corresponding key was pressed (went from !Down to Down).
// If repeat=true and the key is being held down then the press is repeated using io.KeyRepeatDelay and KeyRepeatRate
func IsKeyPressedV(key int, repeat bool) bool {
	return C.iggIsKeyPressed(C.int(key), castBool(repeat)) != 0
}

// IsKeyPressed calls IsKeyPressedV(key, true).
func IsKeyPressed(key int) bool {
	return IsKeyPressedV(key, true)
}

// IsKeyReleased returns true if the corresponding key was released (went from Down to !Down).
func IsKeyReleased(key int) bool {
	return C.iggIsKeyReleased(C.int(key)) != 0
}

// IsMouseDown returns true if the corresponding mouse button is currently being held down.
func IsMouseDown(button int) bool {
	return C.iggIsMouseDown(C.int(button)) != 0
}

// IsAnyMouseDown returns true if any mouse button is currently being held down.
func IsAnyMouseDown() bool {
	return C.iggIsAnyMouseDown() != 0
}

// IsMouseClickedV returns true if the mouse button was clicked (0=left, 1=right, 2=middle)
// If repeat=true and the mouse button is being held down then the click is repeated using io.KeyRepeatDelay and KeyRepeatRate
func IsMouseClickedV(button int, repeat bool) bool {
	return C.iggIsMouseClicked(C.int(button), castBool(repeat)) != 0
}

// IsMouseClicked calls IsMouseClickedV(key, false).
func IsMouseClicked(button int) bool {
	return IsMouseClickedV(button, false)
}

// IsMouseReleased returns true if the mouse button was released (went from Down to !Down).
func IsMouseReleased(button int) bool {
	return C.iggIsMouseReleased(C.int(button)) != 0
}

// IsMouseDoubleClicked returns true if the mouse button was double-clicked (0=left, 1=right, 2=middle).
func IsMouseDoubleClicked(button int) bool {
	return C.iggIsMouseDoubleClicked(C.int(button)) != 0
}

// Columns calls ColumnsV(1, "", false).
func Columns() {
	ColumnsV(1, "", false)
}

// ColumnsV creates a column layout of the specified number of columns.
// The brittle columns API will be superseded by an upcoming 'table' API.
func ColumnsV(count int, label string, border bool) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggColumns(C.int(count), labelArg, castBool(border))
}

// NextColumn next column, defaults to current row or next row if the current row is finished.
func NextColumn() {
	C.iggNextColumn()
}

// ColumnIndex get current column index.
func ColumnIndex() int {
	return int(C.iggGetColumnIndex())
}

// ColumnWidth calls ColumnWidthV(-1).
func ColumnWidth() int {
	return ColumnWidthV(-1)
}

// ColumnWidthV get column width (in pixels). pass -1 to use current column.
func ColumnWidthV(index int) int {
	return int(C.iggGetColumnWidth(C.int(index)))
}

// SetColumnWidth sets column width (in pixels). pass -1 to use current column.
func SetColumnWidth(index int, width float32) {
	C.iggSetColumnWidth(C.int(index), C.float(width))
}

// ColumnOffset calls ColumnOffsetV(-1)
func ColumnOffset() float32 {
	return ColumnOffsetV(-1)
}

// ColumnOffsetV get position of column line (in pixels, from the left side of the contents region). pass -1 to use
// current column, otherwise 0..GetColumnsCount() inclusive. column 0 is typically 0.0.
func ColumnOffsetV(index int) float32 {
	return float32(C.iggGetColumnOffset(C.int(index)))
}

// SetColumnOffset set position of column line (in pixels, from the left side of the contents region). pass -1 to use
// current column.
func SetColumnOffset(index int, offsetX float32) {
	C.iggSetColumnOffset(C.int(index), C.float(offsetX))
}

// ColumnsCount returns number of current columns.
func ColumnsCount() int {
	return int(C.iggGetColumnsCount())
}

// BeginTabBarV create and append into a TabBar
func BeginTabBarV(strID string, flags int) bool {
	idArg, idFin := wrapString(strID)
	defer idFin()

	return C.iggBeginTabBar(idArg, C.int(flags)) != 0
}

// BeginTabBar calls BeginTabBarV(strId, 0)
func BeginTabBar(strID string) bool {
	return BeginTabBarV(strID, 0)
}

// EndTabBar only call EndTabBar() if BeginTabBar() returns true!
func EndTabBar() {
	C.iggEndTabBar()
}

// BeginTabItemV create a Tab. Returns true if the Tab is selected.
func BeginTabItemV(label string, open *bool, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	openArg, openFin := wrapBool(open)
	defer openFin()

	return C.iggBeginTabItem(labelArg, openArg, C.int(flags)) != 0
}

// BeginTabItem calls BeginTabItemV(label, nil, 0)
func BeginTabItem(label string) bool {
	return BeginTabItemV(label, nil, 0)
}

// EndTabItem Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
func EndTabItem() {
	C.iggEndTabItem()
}

// SetTabItemClosed notify TabBar or Docking system of a closed tab/window ahead
// (useful to reduce visual flicker on reorderable tab bars). For tab-bar: call
// after BeginTabBar() and before Tab submissions. Otherwise call with a window name.
func SetTabItemClosed(tabOrDockedWindowLabel string) {
	labelArg, labelFin := wrapString(tabOrDockedWindowLabel)
	defer labelFin()
	C.iggSetTabItemClosed(labelArg)
}

// SetScrollHereY adjusts scrolling amount to make current cursor position visible.
// ratio=0.0: top, 0.5: center, 1.0: bottom.
// When using to make a "default/current item" visible, consider using SetItemDefaultFocus() instead.
func SetScrollHereY(ratio float32) {
	C.iggSetScrollHereY(C.float(ratio))
}

// SetItemDefaultFocus makes the last item the default focused item of a window.
func SetItemDefaultFocus() {
	C.iggSetItemDefaultFocus()
}

// IsItemFocused returns true if the last item is focused.
func IsItemFocused() bool {
	return C.iggIsItemFocused() != 0
}

// IsAnyItemFocused returns true if any item is focused.
func IsAnyItemFocused() bool {
	return C.iggIsAnyItemFocused() != 0
}

// MouseCursor returns desired cursor type, reset in imgui.NewFrame(), this is updated during the frame.
// Valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you.
func MouseCursor() int {
	return int(C.iggGetMouseCursor())
}

// SetMouseCursor sets desired cursor type.
func SetMouseCursor(cursor int) {
	C.iggSetMouseCursor(C.int(cursor))
}

// SetKeyboardFocusHere calls SetKeyboardFocusHereV(0)
func SetKeyboardFocusHere() {
	C.iggSetKeyboardFocusHere(0)
}

// SetKeyboardFocusHereV gives keyboard focus to next item
func SetKeyboardFocusHereV(offset int) {
	C.iggSetKeyboardFocusHere(C.int(offset))
}

func GetWindowDrawList() DrawList {
	return DrawList(C.iggGetWindowDrawList())
}

func GetItemRectMin() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetItemRectMin(valueArg)
	valueFin()
	return value
}

func GetItemRectMax() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetItemRectMax(valueArg)
	valueFin()
	return value
}

func GetItemRectSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetItemRectSize(valueArg)
	valueFin()
	return value
}
