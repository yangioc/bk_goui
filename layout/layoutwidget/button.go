package layoutwidget

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// theme 場景
//
// triger 觸發物件
//
// event 觸發事件處理
func NewButton(theme *material.Theme, triger *widget.Clickable, event func()) *Button {
	return &Button{
		theme:  theme,
		triger: triger,
		event:  event,
	}
}

type Button struct {
	theme  *material.Theme
	triger *widget.Clickable
	event  func()
}

func (self Button) Widget(gtx layout.Context) layout.Dimensions {
	btn := material.Button(self.theme, self.triger, "提交")
	if self.triger.Clicked(gtx) {
		if self.event != nil {
			self.event()
		}
	}
	return btn.Layout(gtx)
}
