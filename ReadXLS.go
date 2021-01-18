package main

import (
	"flag"
	"regexp"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	var in string
	flag.StringVar(&in, "in", "Book1.xlsx", "the input xls")

	var out string
	flag.StringVar(&out, "out", "Book1.md", "the output md file.")

	var sheet string
	flag.StringVar(&sheet, "sheet", "Sheet1", "the input sheet name.")

	flag.Parse()

	f, err := excelize.OpenFile(in)

	if err != nil {
		println(err.Error())
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheet)
	for _, row := range rows {
		for _, colCell := range row {
			print("|")
			re := regexp.MustCompile(`\r?\n`)
			colCell = re.ReplaceAllString(colCell, "<br>")
			print(colCell, "\t")
		}
		println("|")
	}
}
