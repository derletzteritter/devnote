package wnd

import (
	"github.com/AllenDang/giu"
	"github.com/itschip/dev-note/internal/files"
)

const (
	windowTitle = "DevNote"
)

var (
	searchValue string
)

func wndLoop() {
	giu.SingleWindow().Layout(
		giu.Label("Search for files"),
		giu.Spacing(),
		giu.Row(
			giu.InputText(&searchValue),
			giu.Button("New file")),
		giu.Label("All files"), giu.Table().Rows(files.BuildFilesLayout()...))
}

func MasterWindow() {
	wnd := giu.NewMasterWindow(windowTitle, 1000, 700, giu.MasterWindowFlagsNotResizable)

	wnd.Run(wndLoop)
}
