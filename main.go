package main

import (
	"bk_goui/layout/layoutwidget"
	"fmt"
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()

	// 建立按鈕元件
	var (
		button      widget.Clickable
		inputEditor widget.Editor
	)

	inputEditor.Filter = "0123456789"

	label := layoutwidget.NewLabel(theme, "文字顯示", 16)
	// 按鈕
	buttonWidget := layoutwidget.NewButton(theme, &button, func() {
		// inputContent = inputEditor.Text()

		count, err := strconv.Atoi(inputEditor.Text())
		if err != nil {
			label.AddNewLine("數量輸入錯誤")
		} else {
			label.AddNewLine(fmt.Sprintf("執行帳號數量:%d", count))
		}

		// 在此處理輸入的內容，例如打印到控制台
		// println("本次處理的帳號數量：", inputContent)
	})

	// 輸入框
	inputWidget := layoutwidget.NewInput(theme, &inputEditor, "輸入帳號數量")

	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// 此圖形上下文用於管理渲染狀態。
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical, // 元件垂直排列
			}.Layout(gtx,
				// 輸入框元件
				layout.Rigid(inputWidget.Widget),
				// 按鈕元件
				layout.Rigid(buttonWidget.Widget),
				// 顯示文字區塊
				layout.Rigid(label.Widget),
			)

			// 將繪圖操作傳遞給 GPU。
			e.Frame(gtx.Ops)
		}
	}
}
