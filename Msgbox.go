package giu

import "fmt"

type DialogResult uint8

const (
	DialogResultOK DialogResult = 1 << iota
	DialogResultCancel
	DialogResultYes
	DialogResultNo
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
			Button(" Ok ", func() {
				msgboxInvokeCallback(DialogResultOK, callback)
			}),
		}
	case MsgboxButtonsOkCancel:
		return Layout{
			Line(
				Button("  Ok  ", func() {
					msgboxInvokeCallback(DialogResultOK, callback)
				}),
				Button("Cancel", func() {
					msgboxInvokeCallback(DialogResultCancel, callback)
				}),
			),
		}
	case MsgboxButtonsYesNo:
		return Layout{
			Line(
				Button(" Yes ", func() {
					msgboxInvokeCallback(DialogResultYes, callback)
				}),
				Button("  No  ", func() {
					msgboxInvokeCallback(DialogResultNo, callback)
				}),
			),
		}
	default:
		return Layout{
			Button("  Ok  ", func() {
				msgboxInvokeCallback(DialogResultOK, callback)
			}),
		}
	}
}

const msgboxId string = "###Msgbox"

// Embed various Msgboxs to layout. Invoke this function in the same layout level where you call g.Msgbox.
func PrepareMsgbox() Layout {
	var state *MsgboxState
	// Register state.
	stateRaw := Context.GetState(msgboxId)

	if stateRaw == nil {
		state = &MsgboxState{title: "Info", content: "Content", buttons: MsgboxButtonsOk, resultCallback: nil, open: false}
		Context.SetState(msgboxId, state)
	} else {
		state = stateRaw.(*MsgboxState)
	}

	return Layout{
		Custom(func() {
			if state.open {
				OpenPopup(msgboxId)
				state.open = false
			}
			SetNextWindowSize(300, 0)
		}),
		PopupModal(fmt.Sprintf("%s%s", state.title, msgboxId), Layout{
			Custom(func() {
				// Ensure the state is valid.
				Context.GetState(msgboxId)
			}),
			LabelWrapped(state.content),
			buildMsgboxButtons(state.buttons, state.resultCallback),
		}),
	}
}

func Msgbox(title, content string) {
	MsgboxV(title, content, MsgboxButtonsOk, nil)
}

func MsgboxV(title, content string, buttons MsgboxButtons, resultCallback func(DialogResult)) {
	stateRaw := Context.GetState(msgboxId)
	if stateRaw == nil {
		fmt.Println("Msgbox is not prepared. Invoke giu.PrepareMsgbox in the end of the layout.")
		return
	}

	state := stateRaw.(*MsgboxState)
	state.title = title
	state.content = content
	state.buttons = buttons
	state.resultCallback = resultCallback

	state.open = true
}
