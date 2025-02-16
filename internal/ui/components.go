package ui

import (
	"fyne.io/fyne/v2/widget"
)

// CreateButton создает кастомную кнопку
func CreateButton(label string, onClick func()) *widget.Button {
	return widget.NewButton(label, onClick)
}
