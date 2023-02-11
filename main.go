package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

const separator = "\t"

func main() {
	var fn, sn, cn string
	var rn int64
	flag.StringVar(&fn, "f", "", "file path. Supported files are XLAM / XLSM / XLSX / XLTM / XLTX.\n"+
		"If not specified, read REXCEL_FILE from environment variables")
	flag.StringVar(&sn, "s", "", "sheet name.\n"+
		"If not specified, read REXCEL_SHEET from environment variables")
	flag.StringVar(&cn, "c", "", "cell. ex: A1")
	flag.Int64Var(&rn, "r", 0, "row no.")
	flag.Parse()
	if fn == "" {
		fn = os.Getenv("REXCEL_FILE")
	}
	if sn == "" {
		sn = os.Getenv("REXCEL_SHEET")
	}
	if cn == "" && rn < 1 {
		fmt.Fprintln(os.Stderr, "row no should be more than 1.")
		return
	}
	f, err := excelize.OpenFile(fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()
	if cn != "" {
		printCell(f, sn, cn)
	} else {
		printRow(err, f, sn, rn)
	}
}

func printCell(f *excelize.File, sn string, cn string) {
	cell, err := f.GetCellValue(sn, cn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(cell)
}

func printRow(err error, f *excelize.File, sn string, rn int64) {
	rows, err := f.GetRows(sn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	row := rows[rn-1]
	rowStr := ""
	for _, cell := range row {
		rowStr = rowStr + cell + separator
	}
	fmt.Println(strings.Trim(rowStr, separator))
}
