package files

import (
	"fmt"
	"github.com/AllenDang/giu"
	"io/fs"
	"io/ioutil"
)

func findFiles() []fs.FileInfo {
	files, err := ioutil.ReadDir("notes")

	if err != nil {
		fmt.Println("Failed to find or load files")
	}

	return files
}

func BuildFilesLayout() []*giu.TableRowWidget {
	files := findFiles()

	rows := make([]*giu.TableRowWidget, len(files))

	for i := range rows {
		rows[i] = giu.TableRow(giu.Label(files[i].Name()), giu.Row(giu.Button("Open"), giu.Button("Delete")))
	}

	return rows
}