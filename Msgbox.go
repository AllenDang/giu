package giu

import "fmt"

type DialogResult bool

const (
	DialogResultOK     DialogResult = true
	DialogResultCancel DialogResult = false

	DialogResultYes = DialogResultOK
	DialogResultNo  = DialogResultCancel
)

type MsgboxButtons uint8

const (
	MsgboxButtonsYesNo MsgboxButtons = 1 << iota
	MsgboxButtonsOkCancel
	MsgboxButtonsOk
)

type DialogResultCallback func(DialogResult)

type MsgboxState struct {
	title          string
	content        string
	resultCallback DialogResultCallback
	buttons        MsgboxButtons
	open           bool
}

func (ms *MsgboxState) Dispose() {
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

const msgboxId string = "###Msgbox"

// Embed various Msgboxs to layout. Invoke this function in the same layout level where you call g.Msgbox.
func PrepareMsgbox() Layout {
	return Layout{
		Custom(func() {
			var state *MsgboxState

			// Register state.
			stateRaw := Context.GetState(msgboxId)

			if stateRaw == nil {
				state = &MsgboxState{title: "Info", content: "Content", buttons: MsgboxButtonsOk, resultCallback: nil, open: false}
				Context.SetState(msgboxId, state)
			} else {
				state = stateRaw.(*MsgboxState)
			}

			if state.open {
				OpenPopup(msgboxId)
				state.open = false
			}
			SetNextWindowSize(300, 0)
			PopupModal(fmt.Sprintf("%s%s", state.title, msgboxId)).Layout(
				Custom(func() {
					// Ensure the state is valid.
					Context.GetState(msgboxId)
				}),
				Label(state.content).Wrapped(true),
				buildMsgboxButtons(state.buttons, state.resultCallback),
			).Build()
		}),
	}
}

type MsgboxWidget struct{}

func (m *MsgboxWidget) getState() *MsgboxState {
	stateRaw := Context.GetState(msgboxId)
	if stateRaw == nil {
		panic("Msgbox is not prepared. Invoke giu.PrepareMsgbox in the end of the layout.")
	}

	return stateRaw.(*MsgboxState)
}

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

func (m *MsgboxWidget) Buttons(buttons MsgboxButtons) *MsgboxWidget {
	s := m.getState()
	s.buttons = buttons
	return m
}

func (m *MsgboxWidget) ResultCallback(cb DialogResultCallback) *MsgboxWidget {
	s := m.getState()
	s.resultCallback = cb
	return m
}
