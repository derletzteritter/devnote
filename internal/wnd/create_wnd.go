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
	showFileDialog bool
	newFileName string
	fileContent string
	noteName string
	showFileEditor bool
	editor *giu.CodeEditorWidget
)

func onShowFileDialog() {
	showFileDialog = true
}

func onHideFileDialog()  {
	showFileDialog = false
}

func onShowFileEditor() {
	showFileEditor = true
}

func onHideFileEditor() {
	showFileEditor = false
}

func wndLoop() {

	giu.SingleWindow().Layout(
		giu.Label("Search for files"),
		giu.Spacing(),
		giu.Row(
			giu.InputText(&searchValue),
			giu.Button("New file").OnClick(onShowFileDialog)),
		giu.Label("All files"), giu.Table().Rows(files.BuildFilesLayout(func(file string, filename string) {
			fileContent = file
			noteName = filename

			editor = giu.CodeEditor().ShowWhitespaces(true).Text(fileContent).Border(true)

			onShowFileEditor()
		})...))

	if showFileDialog {
		giu.Window("New file").IsOpen(&showFileDialog).Size(300, 300).
			Layout(
				giu.Label("Enter a file name"),
				giu.InputText(&newFileName),
				giu.Button("Create file").OnClick(func() {
					files.CreateNewFile(newFileName)
					onHideFileDialog()
				}),
				giu.Button("Cancel").OnClick(onHideFileDialog))
	}

	if showFileEditor {
		if editor.IsTextChanged() {
			fileContent = editor.GetText()
		}

		giu.Window("Editor").IsOpen(&showFileEditor).Size(800, 650).
			Layout(
				giu.SplitLayout(giu.DirectionHorizontal, 200, giu.Markdown(&fileContent), editor),
				giu.Button("Close").OnClick(onHideFileEditor),
				giu.Button("Save").OnClick(func() {
					content := []byte(fileContent)

					files.SaveNote(noteName, content)
				}))
	}
}

func MasterWindow() {

	wnd := giu.NewMasterWindow(windowTitle, 1000, 700, giu.MasterWindowFlagsNotResizable)

	wnd.Run(wndLoop)
}
