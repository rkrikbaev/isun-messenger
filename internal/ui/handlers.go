package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateMainUI создает основной интерфейс приложения
func CreateMainUI(w fyne.Window) fyne.CanvasObject {
	// Кнопки управления
	btnRefresh := widget.NewButton("Обновить", func() {
		// Логика обновления данных
	})

	btnSettings := widget.NewButton("Настройки", func() {
		// Открытие окна настроек
	})

	// Контейнер с кнопками
	buttons := container.NewVBox(btnRefresh, btnSettings)

	// Главный контейнер
	content := container.NewBorder(nil, buttons, nil, nil, widget.NewLabel("Главное окно"))

	return content
}
