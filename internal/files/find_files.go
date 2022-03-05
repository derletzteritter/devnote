package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/AllenDang/giu"
)

func findFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir("../../notes")

	if err != nil {
		fmt.Println(err.Error())
	}

	return files
}

func BuildFilesLayout(searchValue *string, openFile func(file string, noteName string)) []*giu.TableRowWidget {
	files := findFiles()

	rows := make([]*giu.TableRowWidget, len(files))

	for i := range rows {
		rows[i] = giu.TableRow(
			giu.Label(
				files[i].Name()),
			giu.Row(
				giu.Button("Open").OnClick(func() {
					file := OpenNote(files[i].Name())
					openFile(string(file), files[i].Name())
				}),
				giu.Button("Delete").OnClick(func() {
					DeleteNote(files[i].Name())
				})))
	}

	return rows
}
