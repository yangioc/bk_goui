package layoutwidget

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// theme 場景
//
// input 輸入物件
//
// tipStr 提示文字
func NewInput(theme *material.Theme, input *widget.Editor, tipStr string) *Input {
	return &Input{
		theme:      theme,
		input:      input,
		contextStr: tipStr,
	}
}

type Input struct {
	theme      *material.Theme
	input      *widget.Editor
	contextStr string
}

func (self Input) Widget(gtx layout.Context) layout.Dimensions {
	editor := material.Editor(self.theme, self.input, self.contextStr)
	return editor.Layout(gtx)
}
