package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

const separator = "\t"

var (
	version    string
	fn, sn, cn string
	v          bool
	rn         int64
)

func main() {
	parseFlag()
	printVersion()
	checkFlag()
	printCellOrRow()
}

func parseFlag() {
	flag.StringVar(&fn, "f", "", "file path. Supported files are XLAM / XLSM / XLSX / XLTM / XLTX.\n"+
		"If not specified, read REXCEL_FILE from environment variables")
	flag.StringVar(&sn, "s", "", "sheet name.\n"+
		"If not specified, read REXCEL_SHEET from environment variables")
	flag.StringVar(&cn, "c", "", "cell. ex: A1")
	flag.Int64Var(&rn, "r", -1, "row no.")
	flag.BoolVar(&v, "v", false, "print version")
	flag.Parse()
	if fn == "" {
		fn = os.Getenv("REXCEL_FILE")
	}
	if sn == "" {
		sn = os.Getenv("REXCEL_SHEET")
	}
}

func printVersion() {
	if v {
		fmt.Println(version)
		os.Exit(0)
	}
}

func checkFlag() {
	if fn == "" {
		exitError("Specify file path.", 1)
	}
	if sn == "" {
		exitError("Specify sheet name.", 1)
	}
	if cn == "" && rn == -1 {
		exitError("Specify cell or row.", 1)
	}
	if cn == "" && rn < 1 {
		exitError("row no should be more than 1.", 1)
	}
}

func printCellOrRow() {
	f, err := excelize.OpenFile(fn)
	if err != nil {
		exitError(err, 1)
	}
	defer func() {
		if err := f.Close(); err != nil {
			exitError(err, 1)
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
		exitError(err, 1)
	}
	fmt.Println(cell)
}

func printRow(err error, f *excelize.File, sn string, rn int64) {
	rows, err := f.GetRows(sn)
	if err != nil {
		exitError(err, 1)
	}
	row := rows[rn-1]
	rowStr := ""
	for _, cell := range row {
		rowStr = rowStr + cell + separator
	}
	fmt.Println(strings.Trim(rowStr, separator))
}

func exitError[T any](e T, code int) {
	fmt.Fprintln(os.Stderr, e)
	os.Exit(code)
}
