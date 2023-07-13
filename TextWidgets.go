package giu

import (
	"fmt"
	"math"

	"golang.org/x/image/colornames"

	"github.com/AllenDang/imgui-go"
	"github.com/sahilm/fuzzy"
)

var _ Widget = &InputTextMultilineWidget{}

// InputTextMultilineWidget is a large (multiline) text input
// see examples/widgets/.
type InputTextMultilineWidget struct {
	label          string
	text           *string
	width, height  float32
	flags          InputTextFlags
	cb             imgui.InputTextCallback
	scrollToBottom bool
	onChange       func()
}

// InputTextMultiline creates InputTextMultilineWidget.
func InputTextMultiline(text *string) *InputTextMultilineWidget {
	return &InputTextMultilineWidget{
		text:     text,
		width:    0,
		height:   0,
		flags:    0,
		cb:       nil,
		onChange: nil,
		label:    GenAutoID("##InputTextMultiline"),
	}
}

// Label sets input field label.
func (i *InputTextMultilineWidget) Label(label string) *InputTextMultilineWidget {
	i.label = label
	return i
}

// Labelf is formatting version of Label.
func (i *InputTextMultilineWidget) Labelf(format string, args ...any) *InputTextMultilineWidget {
	return i.Label(fmt.Sprintf(format, args...))
}

// Flags sets InputTextFlags (see Flags.go).
func (i *InputTextMultilineWidget) Flags(flags InputTextFlags) *InputTextMultilineWidget {
	i.flags = flags
	return i
}

// Callback sets imgui.InputTextCallback.
func (i *InputTextMultilineWidget) Callback(cb imgui.InputTextCallback) *InputTextMultilineWidget {
	i.cb = cb
	return i
}

// OnChange set callback called when user action taken on input text field (when text was changed).
func (i *InputTextMultilineWidget) OnChange(onChange func()) *InputTextMultilineWidget {
	i.onChange = onChange
	return i
}

// Size sets input field size.
func (i *InputTextMultilineWidget) Size(width, height float32) *InputTextMultilineWidget {
	i.width, i.height = width, height
	return i
}

// AutoScrollToBottom Enables/Disables auto scroll to bottom.
func (i *InputTextMultilineWidget) AutoScrollToBottom(b bool) *InputTextMultilineWidget {
	i.scrollToBottom = b
	return i
}

// Build implements Widget interface.
func (i *InputTextMultilineWidget) Build() {
	if imgui.InputTextMultilineV(
		Context.FontAtlas.RegisterString(i.label),
		Context.FontAtlas.RegisterStringPointer(i.text),
		imgui.Vec2{
			X: i.width,
			Y: i.height,
		},
		int(i.flags), i.cb,
	) && i.onChange != nil {
		i.onChange()
	}

	if i.scrollToBottom {
		imgui.BeginChild(i.label)
		imgui.SetScrollHereY(1.0)
		imgui.EndChild()
	}
}

var _ Widget = &BulletWidget{}

// BulletWidget adds a small, white dot (bullet).
// useful in enumerations.
type BulletWidget struct{}

// Bullet creates a bullet widget.
func Bullet() *BulletWidget {
	return &BulletWidget{}
}

// Build implements Widget interface.
func (b *BulletWidget) Build() {
	imgui.Bullet()
}

var _ Widget = &BulletTextWidget{}

// BulletTextWidget does similar to BulletWidget, but allows
// to add a text after a bullet. Very useful to create lists.
type BulletTextWidget struct {
	text string
}

// BulletText creates bulletTextWidget.
func BulletText(text string) *BulletTextWidget {
	return &BulletTextWidget{
		text: Context.FontAtlas.RegisterString(text),
	}
}

// BulletTextf is a formatting version of BulletText.
func BulletTextf(format string, args ...any) *BulletTextWidget {
	return BulletText(fmt.Sprintf(format, args...))
}

// Build implements Widget interface.
func (bt *BulletTextWidget) Build() {
	imgui.BulletText(bt.text)
}

var _ Disposable = &inputTextState{}

type inputTextState struct {
	autoCompleteCandidates fuzzy.Matches
	currentIdx             int
}

// Dispose implements disposable interface.
func (s *inputTextState) Dispose() {
	s.autoCompleteCandidates = nil
	s.currentIdx = 0
}

var _ Widget = &InputTextWidget{}

// InputTextWidget is a single-line text input.
type InputTextWidget struct {
	label      string
	hint       string
	value      *string
	width      float32
	candidates []string
	flags      InputTextFlags
	cb         imgui.InputTextCallback
	onChange   func()
}

// InputText creates new input text widget.
func InputText(value *string) *InputTextWidget {
	return &InputTextWidget{
		label:    GenAutoID("##InputText"),
		hint:     "",
		value:    value,
		width:    0,
		flags:    0,
		cb:       nil,
		onChange: nil,
	}
}

// Label adds label (alternatively you can use it to set widget's id).
func (i *InputTextWidget) Label(label string) *InputTextWidget {
	i.label = Context.FontAtlas.RegisterString(label)
	return i
}

// Labelf adds formatted label.
func (i *InputTextWidget) Labelf(format string, args ...any) *InputTextWidget {
	return i.Label(fmt.Sprintf(format, args...))
}

// AutoComplete enables auto complete popup by using fuzzy search of current value against candidates
// Press enter to confirm the first candidate.
func (i *InputTextWidget) AutoComplete(candidates []string) *InputTextWidget {
	i.candidates = candidates
	return i
}

// Hint sets hint text.
func (i *InputTextWidget) Hint(hint string) *InputTextWidget {
	i.hint = Context.FontAtlas.RegisterString(hint)
	return i
}

// Size sets field's width.
func (i *InputTextWidget) Size(width float32) *InputTextWidget {
	i.width = width
	return i
}

// Flags sets flags.
func (i *InputTextWidget) Flags(flags InputTextFlags) *InputTextWidget {
	i.flags = flags
	return i
}

// Callback sets input text callback.
func (i *InputTextWidget) Callback(cb imgui.InputTextCallback) *InputTextWidget {
	i.cb = cb
	return i
}

// OnChange sets callback when text was changed.
func (i *InputTextWidget) OnChange(onChange func()) *InputTextWidget {
	i.onChange = onChange
	return i
}

// Build implements Widget interface.
func (i *InputTextWidget) Build() {
	// Get state
	var state *inputTextState
	if state = GetState[inputTextState](Context, i.label); state == nil {
		state = &inputTextState{}
		SetState(Context, i.label, state)
	}

	if i.width != 0 {
		PushItemWidth(i.width)

		defer PopItemWidth()
	}

	isChanged := imgui.InputTextWithHint(i.label, i.hint, Context.FontAtlas.RegisterStringPointer(i.value), int(i.flags), i.cb)

	if isChanged && i.onChange != nil {
		i.onChange()
	}

	if isChanged {
		// Enable auto complete
		if len(i.candidates) > 0 {
			matches := fuzzy.Find(*i.value, i.candidates)
			size := int(math.Min(5, float64(matches.Len())))
			matches = matches[:size]

			state.autoCompleteCandidates = matches
		}
	}

	// Draw autocomplete list
	if len(state.autoCompleteCandidates) > 0 && imgui.IsItemFocused() {
		i.handleAutoComplete(state)
	}
}

func (i *InputTextWidget) handleAutoComplete(state *inputTextState) {
	if state.currentIdx >= len(state.autoCompleteCandidates) {
		state.currentIdx = 0
	}

	labels := make(Layout, len(state.autoCompleteCandidates))
	for i, m := range state.autoCompleteCandidates {
		labels[i] = Label(m.Str)
		if i == state.currentIdx {
			labels[i] = Layout{
				Custom(func() { PushStyleColor(StyleColorText, colornames.Blue) }),
				labels[i],
				Custom(func() { PopStyleColor() }),
			}
		}
	}

	SetNextWindowPos(imgui.GetItemRectMin().X, imgui.GetItemRectMax().Y)
	imgui.BeginTooltip()
	labels.Build()
	imgui.EndTooltip()

	// Press enter will replace value string with first match candidate
	switch {
	case IsKeyPressed(KeyEnter) || IsKeyPressed(KeyTab):
		*i.value = state.autoCompleteCandidates[state.currentIdx].Str
		state.autoCompleteCandidates = nil

		if i.onChange != nil {
			i.onChange()
		}
	case IsKeyPressed(KeyDown):
		state.currentIdx++
		if state.currentIdx >= state.autoCompleteCandidates.Len() {
			state.currentIdx = 0
		}
	case IsKeyPressed(KeyUp):
		state.currentIdx--
		if state.currentIdx < 0 {
			state.currentIdx = len(state.autoCompleteCandidates) - 1
		}
	}
}

var _ Widget = &InputIntWidget{}

// InputIntWidget is an input text field accepting integer values only.
type InputIntWidget struct {
	label    string
	value    *int32
	width    float32
	flags    InputTextFlags
	onChange func()
	step     int
	stepFast int
}

// InputInt creates input int widget
// NOTE: value is int32, so its size is up to 10^32-1.
// to process greater values, you need to use InputTextWidget
// with InputTextFlagsCharsDecimal and strconv.ParseInt in OnChange callback.
func InputInt(value *int32) *InputIntWidget {
	return &InputIntWidget{
		label:    GenAutoID("##InputInt"),
		value:    value,
		width:    0,
		flags:    0,
		onChange: nil,
	}
}

// Label sets label (id).
func (i *InputIntWidget) Label(label string) *InputIntWidget {
	i.label = Context.FontAtlas.RegisterString(label)
	return i
}

// Labelf sets formatted label.
func (i *InputIntWidget) Labelf(format string, args ...any) *InputIntWidget {
	return i.Label(fmt.Sprintf(format, args...))
}

// Size sets input's width.
func (i *InputIntWidget) Size(width float32) *InputIntWidget {
	i.width = width
	return i
}

// Flags sets flags.
func (i *InputIntWidget) Flags(flags InputTextFlags) *InputIntWidget {
	i.flags = flags
	return i
}

// StepSize sets the step size.
func (i *InputIntWidget) StepSize(step int) *InputIntWidget {
	i.step = step
	return i
}

// StepSizeFast sets the fast step size.
func (i *InputIntWidget) StepSizeFast(stepFast int) *InputIntWidget {
	i.stepFast = stepFast
	return i
}

// OnChange adds on change callback.
func (i *InputIntWidget) OnChange(onChange func()) *InputIntWidget {
	i.onChange = onChange
	return i
}

// Build implements Widget interface.
func (i *InputIntWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)

		defer PopItemWidth()
	}

	if imgui.InputIntV(i.label, i.value, i.step, i.stepFast, int(i.flags)) && i.onChange != nil {
		i.onChange()
	}
}

var _ Widget = &InputFloatWidget{}

// InputFloatWidget does similar to InputIntWIdget, but accepts float numbers.
type InputFloatWidget struct {
	label    string
	value    *float32
	width    float32
	flags    InputTextFlags
	format   string
	onChange func()
	step     float32
	stepFast float32
}

// InputFloat constructs InputFloatWidget.
func InputFloat(value *float32) *InputFloatWidget {
	return &InputFloatWidget{
		label:    GenAutoID("##InputFloatWidget"),
		width:    0,
		value:    value,
		format:   "%.3f",
		flags:    0,
		onChange: nil,
	}
}

// Label sets label of input field.
func (i *InputFloatWidget) Label(label string) *InputFloatWidget {
	i.label = Context.FontAtlas.RegisterString(label)
	return i
}

// Labelf sets formatted label.
func (i *InputFloatWidget) Labelf(format string, args ...any) *InputFloatWidget {
	return i.Label(fmt.Sprintf(format, args...))
}

// Size sets input field's width.
func (i *InputFloatWidget) Size(width float32) *InputFloatWidget {
	i.width = width
	return i
}

// Flags sets flags.
func (i *InputFloatWidget) Flags(flags InputTextFlags) *InputFloatWidget {
	i.flags = flags
	return i
}

// Format sets data format (e.g. %.3f).
func (i *InputFloatWidget) Format(format string) *InputFloatWidget {
	i.format = format
	return i
}

// OnChange sets callback called when text is changed.
func (i *InputFloatWidget) OnChange(onChange func()) *InputFloatWidget {
	i.onChange = onChange
	return i
}

// StepSize sets the step size.
func (i *InputFloatWidget) StepSize(step float32) *InputFloatWidget {
	i.step = step
	return i
}

// StepSizeFast sets the fast step size.
func (i *InputFloatWidget) StepSizeFast(stepFast float32) *InputFloatWidget {
	i.stepFast = stepFast
	return i
}

// Build implements Widget interface.
func (i *InputFloatWidget) Build() {
	if i.width != 0 {
		PushItemWidth(i.width)

		defer PopItemWidth()
	}

	if imgui.InputFloatV(i.label, i.value, i.step, i.stepFast, i.format, int(i.flags)) && i.onChange != nil {
		i.onChange()
	}
}

var _ Widget = &LabelWidget{}

// LabelWidget is a plain text label.
type LabelWidget struct {
	label    string
	fontInfo *FontInfo
	wrapped  bool
}

// Label constructs label widget.
func Label(label string) *LabelWidget {
	return &LabelWidget{
		label:   Context.FontAtlas.RegisterString(label),
		wrapped: false,
	}
}

// Labelf allows to add formatted label.
func Labelf(format string, args ...any) *LabelWidget {
	return Label(fmt.Sprintf(format, args...))
}

// Wrapped determines if label is wrapped.
func (l *LabelWidget) Wrapped(wrapped bool) *LabelWidget {
	l.wrapped = wrapped
	return l
}

// Font sets specific font (does like Style().SetFont).
func (l *LabelWidget) Font(font *FontInfo) *LabelWidget {
	l.fontInfo = font
	return l
}

// Build implements Widget interface.
func (l *LabelWidget) Build() {
	if l.wrapped {
		PushTextWrapPos()

		defer PopTextWrapPos()
	}

	if l.fontInfo != nil {
		if PushFont(l.fontInfo) {
			defer PopFont()
		}
	}

	imgui.Text(l.label)
}
