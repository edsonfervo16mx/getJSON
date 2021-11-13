package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func Index(w http.ResponseWriter, r *http.Request) {
	testReadTxt()
	fmt.Fprintf(w, "API getJSON")
}

func testReadTxt() {
	file, err := ioutil.ReadFile("./test_files/texto.txt")
	if err != nil {
			log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
}
