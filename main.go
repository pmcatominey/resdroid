package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	drawableReportFile = flag.String("drawable", "", "file to save drawable report to")
)

func main() {
	flag.Parse()
	resPath := flag.Arg(0)

	res, err := NewResDirectory(resPath)
	handleError(err)

	if len(*drawableReportFile) > 0 {
		file, err := os.Create(*drawableReportFile)
		handleError(err)

		err = GenerateDrawableReport(res, file)
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
}
