package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {

	xlFile := getExcelFile("TheProfessionals.xlsx")

	worksheets := xlFile.GetSheetList()

	for i := range worksheets {
		createCSVFile(xlFile, worksheets[i])
	}
}

func getExcelFile(fileName string) *excelize.File {

	xlFile, xlErr := excelize.OpenFile(fileName)
	if xlErr != nil {
		panic(xlErr)
	}
	defer func() {
		if xlErr := xlFile.Close(); xlErr != nil {
			panic(xlErr)
		}
	}()

	return xlFile
}

func createCSVFile(xlFile *excelize.File, worksheet string) {

	allRows, arErr := xlFile.GetRows(worksheet)
	if arErr != nil {
		panic(arErr)
	}

	csvFile, csvErr := os.Create(worksheet + ".csv")
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
