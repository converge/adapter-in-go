package adapters

import (
	"adapter-in-go/domain"
)

type WindowsAdapter struct {
	WindowMachine *domain.Windows
}

func (w *WindowsAdapter) InsertInSquarePort() {
	w.WindowMachine.InterInCirclePort()
}