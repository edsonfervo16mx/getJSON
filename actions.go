package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xuri/excelize"
)

func Index(w http.ResponseWriter, r *http.Request) {
	testReadTxt()
	testReadCsv()
	f, err := excelize.OpenFile("./test_files/nombres.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Hoja1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Hoja1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
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

func testReadCsv() {
	file, err := ioutil.ReadFile("./test_files/lista.csv")
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
}
