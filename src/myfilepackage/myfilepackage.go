package myfilepackage

import "fmt"

type File struct {
	Name      string
	Extension string
	Size      float32
}

//Obtener datos del archivo
func (myFile File) GetMetadata() bool {
	fmt.Println(myFile.Name, "archivo")
	return true
}
