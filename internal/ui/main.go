package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// InitUI инициализирует и запускает графический интерфейс
func InitUI() {
	application := app.New()
	window := application.NewWindow("Messenger UI")

	// Основной контейнер
	mainContainer := container.NewVBox(
		widget.NewLabel("Welcome to Messenger UI"),
		widget.NewButton("Exit", func() {
			application.Quit()
		}),
	)

	window.SetContent(mainContainer)
	window.ShowAndRun()
}
