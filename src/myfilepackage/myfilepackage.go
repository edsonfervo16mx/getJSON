package myfilepackage

import "fmt"

type File struct {
	Name      string
	Extension string
	Size      float32
}

func (myFile File) GetMetaData() {
	fmt.Println(myFile.Name, "archivo")
}
