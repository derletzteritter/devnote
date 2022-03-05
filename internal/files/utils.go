package files

import "fmt"

const (
	NotesPath = "../../notes"
)

func FormatFilePath(filename string) string {
	return fmt.Sprintf("%v\\%v.md", NotesPath, filename)
}
