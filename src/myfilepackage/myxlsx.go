package myfilepackage

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	Book  string
	Sheet string
}

func (myXlsx Xlsx) GetData(c echo.Context) (res bool) {
	fmt.Println("****************")
	fmt.Println(myXlsx.Book)
	fmt.Println("****************")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	f, err := excelize.OpenFile(os.Getenv("PATH_FILES") + "xlsx/" + myXlsx.Book + ".xlsx")
	if err != nil {
		fmt.Println(err)
		return false
	}

	list := f.GetSheetList()
	for _, valor := range list {

		rows, err := f.GetRows(valor)
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
	}

	return false
}
