package reports

import (
	"fmt"
	"reflect"

	"github.com/xuri/excelize/v2"
)

func toAlphaString(i int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""
	for i > 0 {
		i--
		result = string(letters[i%26]) + result
		i /= 26
	}
	return result
}

func GenerateXlsxGeneric(nameSheet string, records interface{}, headers []string) {
	f := excelize.NewFile()
	index, _ := f.NewSheet(nameSheet)
	for i, header := range headers {
		cell := toAlphaString(i+1) + "1"
		f.SetCellValue(nameSheet, cell, header)
	}

	reflectValue := reflect.ValueOf(records)

	if reflectValue.Kind() != reflect.Slice {
		fmt.Println("Error: records no es un slice")
		return
	}

	for i := 0; i < reflectValue.Len(); i++ {
		record := reflectValue.Index(i)
		row := i + 2
		for j := 0; j < record.NumField(); j++ {
			cell := fmt.Sprintf("%s%d", toAlphaString(j+1), row)
			f.SetCellValue(nameSheet, cell, fmt.Sprintf("%v", record.Field(j).Interface()))
		}
	}

	f.SetActiveSheet(index)

	name := nameSheet + ".xlsx"

	if err := f.SaveAs(name); err != nil {
		fmt.Println(err)
		return
	}
}
