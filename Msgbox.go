package giu

import "fmt"

// DialogResult represents dialog result
// dialog resullt is bool. if OK/Yes it is true, else (Cancel/No) - false.
type DialogResult bool

// dialog results.
const (
	DialogResultOK     DialogResult = true
	DialogResultCancel DialogResult = false

	DialogResultYes = DialogResultOK
	DialogResultNo  = DialogResultCancel
)

// MsgboxButtons determines which buttons are in the dialog.
type MsgboxButtons uint8

// button sets.
const (
	// Yes-No question.
	MsgboxButtonsYesNo MsgboxButtons = 1 << iota
	// Ok / Cancel dialog.
	MsgboxButtonsOkCancel
	// info.
	MsgboxButtonsOk
)

// DialogResultCallback is a callback for dialogs.
type DialogResultCallback func(DialogResult)

var _ Disposable = &msgboxState{}

type msgboxState struct {
	title          string
	content        string
	resultCallback DialogResultCallback
	buttons        MsgboxButtons
	open           bool
}

// Dispose implements disposable interface.
func (ms *msgboxState) Dispose() {
	// Nothing to do here.
}

func msgboxInvokeCallback(result DialogResult, callback DialogResultCallback) {
	CloseCurrentPopup()
	if callback != nil {
		callback(result)
	}
}

func buildMsgboxButtons(buttons MsgboxButtons, callback DialogResultCallback) Layout {
	switch buttons {
	case MsgboxButtonsOk:
		return Layout{
			Button(" Ok ").OnClick(func() {
				msgboxInvokeCallback(DialogResultOK, callback)
			}),
		}
	case MsgboxButtonsOkCancel:
		return Layout{
			Row(
				Button("  Ok  ").OnClick(func() {
					msgboxInvokeCallback(DialogResultOK, callback)
				}),
				Button("Cancel").OnClick(func() {
					msgboxInvokeCallback(DialogResultCancel, callback)
				}),
			),
		}
	case MsgboxButtonsYesNo:
		return Layout{
			Row(
				Button(" Yes ").OnClick(func() {
					msgboxInvokeCallback(DialogResultYes, callback)
				}),
				Button("  No  ").OnClick(func() {
					msgboxInvokeCallback(DialogResultNo, callback)
				}),
			),
		}
	default:
		return Layout{
			Button("  Ok  ").OnClick(func() {
				msgboxInvokeCallback(DialogResultOK, callback)
			}),
		}
	}
}

const msgboxID string = "###Msgbox"

// PrepareMsgbox should be invoked in function in the same layout level where you call g.Msgbox.
// BUG: calling this more than 1 time per frame causes unexpected
// merging msgboxes layouts (see https://github.com/AllenDang/giu/issues/290)
func PrepareMsgbox() Layout {
	return Layout{
		Custom(func() {
			var state *msgboxState

			// Register state.
			stateRaw := Context.GetState(msgboxID)

			if stateRaw == nil {
				state = &msgboxState{title: "Info", content: "Content", buttons: MsgboxButtonsOk, resultCallback: nil, open: false}
				Context.SetState(msgboxID, state)
			} else {
				var isOk bool
				state, isOk = stateRaw.(*msgboxState)
				Assert(isOk, "MsgboxWidget", "PrepareMsgbox", "got state of unexpected type")
			}

			if state.open {
				OpenPopup(msgboxID)
				state.open = false
			}
			SetNextWindowSize(300, 0)
			PopupModal(fmt.Sprintf("%s%s", state.title, msgboxID)).Layout(
				Custom(func() {
					// Ensure the state is valid.
					Context.GetState(msgboxID)
				}),
				Label(state.content).Wrapped(true),
				buildMsgboxButtons(state.buttons, state.resultCallback),
			).Build()
		}),
	}
}

// MsgboxWidget represents message dialog.
type MsgboxWidget struct{}

func (m *MsgboxWidget) getState() *msgboxState {
	stateRaw := Context.GetState(msgboxID)
	if stateRaw == nil {
		panic("Msgbox is not prepared. Invoke giu.PrepareMsgbox in the end of the layout.")
	}

	result, isOk := stateRaw.(*msgboxState)
	Assert(isOk, "MsgboxWidget", "getState", "unexpected type of widget's state recovered")

	return result
}

// Msgbox opens message box.
// call it whenever you want to open popup with
// question / info.
func Msgbox(title, content string) *MsgboxWidget {
	result := &MsgboxWidget{}

	state := result.getState()
	state.title = title
	state.content = content

	state.buttons = MsgboxButtonsOk
	state.resultCallback = nil

	state.open = true

	return result
}

// Buttons sets which buttons should be possible.
func (m *MsgboxWidget) Buttons(buttons MsgboxButtons) *MsgboxWidget {
	s := m.getState()
	s.buttons = buttons
	return m
}

// ResultCallback sets result callback.
func (m *MsgboxWidget) ResultCallback(cb DialogResultCallback) *MsgboxWidget {
	s := m.getState()
	s.resultCallback = cb
	return m
}
