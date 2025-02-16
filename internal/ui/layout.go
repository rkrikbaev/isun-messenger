package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// CreateMainLayout создает основной макет приложения
func CreateMainLayout(content fyne.CanvasObject) *fyne.Container {
	return container.NewBorder(nil, nil, nil, nil, content)
}
