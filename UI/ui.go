package UI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func UI_init() fyne.CanvasObject {

	Button1 := widget.NewButtonWithIcon("Start", theme.ComputerIcon(), func() {
		//启动逻辑start.go(待完成)
	})

	Button2 := widget.NewButtonWithIcon("Versions", theme.MenuExpandIcon(), func() {
		//版本任务versions.go(待完成)
	})

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			//设置逻辑setting.go
		}),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			//疑问逻辑help.go
		}),
	)

	buttonCon := container.NewVBox(Button1, Button2)

	tol := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, buttonCon, nil, toolbar), buttonCon, toolbar)

	content := container.NewVBox(tol, buttonCon)

	return content
}

func Ui_Start() {
	//创建一个窗口,名字叫PECL
	a := app.New()
	b := a.NewWindow("PECL")
	// 设置窗口大小
	b.Resize(fyne.NewSize(400, 400))
	// 窗口居中
	b.CenterOnScreen()
	//收集canvas.object
	object := UI_init()
	//展示窗口 并运行程序
	b.SetContent(object)
	b.ShowAndRun()
}
