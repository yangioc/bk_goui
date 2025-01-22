package layoutwidget

import (
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// 創建文字顯示物件
//
// theme 物件所屬場景
//
// tipText 提示文字
//
// textSize 文字大小
func NewLabel(theme *material.Theme, tipText string, textSize float32) *Label {
	editor := new(widget.Editor)
	editor.ReadOnly = true
	editorStyle := material.Editor(theme, editor, tipText)
	editorStyle.TextSize = unit.Sp(textSize)
	return &Label{
		theme:           theme,
		editor:          editor,
		editorStyle:     editorStyle,
		contextTextLine: []string{},
		scrollUp:        true,
	}
}

type Label struct {
	theme                *material.Theme
	editor               *widget.Editor
	editorStyle          material.EditorStyle
	contextTextLine      []string
	contextTextLineRever []string
	scrollUp             bool
}

func (self *Label) Widget(gtx layout.Context) layout.Dimensions {
	return self.editorStyle.Layout(gtx)
}

func (self *Label) SetScroll(isUp bool) {
	self.scrollUp = isUp
}

// 返回當前內容
func (self *Label) Text() string {
	return self.editor.Text()
}

// SetText 重置內容
func (self *Label) SetText(txt string) {
	self.contextTextLine = []string{txt}
	self.editor.SetText(txt)
}

// AddNewLine 加入新內容
func (self *Label) AddNewLine(txt string) {
	self.contextTextLine = append(self.contextTextLine, txt)
	var builder strings.Builder

	if self.scrollUp {
		for i := len(self.contextTextLine) - 1; i >= 0; i-- {
			builder.WriteString(self.contextTextLine[i])
			if i > 0 {
				builder.WriteString("\n") // 添加换行符，除非是最后一个元素
			}
		}

		self.editor.SetText(builder.String())
	} else {
		self.editor.SetText(strings.Join(self.contextTextLine, "\n"))
	}
}
