package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {

	xlFile, xlErr := excelize.OpenFile("Book1.xlsx")
	if xlErr != nil {
		fmt.Println(xlErr)
		return
	}
	defer func() {
		if xlErr := xlFile.Close(); xlErr != nil {
			panic(xlErr)
		}
	}()

	allRows, arErr := xlFile.GetRows("Sheet3")
	if arErr != nil {
		panic(arErr)
	}

	csvFile, csvErr := os.Create("text.csv")
	if csvErr != nil {
		fmt.Println(csvErr)
	}
	defer func() {
		if csvErr := csvFile.Close(); csvErr != nil {
			panic(csvErr)
		}
	}()

	writer := csv.NewWriter(csvFile)

	var writerErr error = writer.WriteAll(allRows)
	if writerErr != nil {
		fmt.Println(writerErr)
	}

}
