package myfilepackage

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	Sheet string
}

func (myXlsx Xlsx) TestRead() bool {
	fmt.Println(myXlsx.Sheet, "Hoja1")
	f, err := excelize.OpenFile("./test_files/nombres.xlsx")
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Hoja1", "B2")
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Hoja1")
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	return false
}
