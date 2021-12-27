package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateNewFile(filename string)  {
	path := FormatFilePath(filename)

	err := ioutil.WriteFile(path, []byte(""), 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
}


func DeleteNote(filename string) {
	path := NotesPath + "\\" + filename

	err := os.Remove(path)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func OpenNote(filename string) []byte {
	path := NotesPath + "\\" + filename

	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(f))

	return f
}