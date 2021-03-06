package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Destaby/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File")
	outputFile      = flag.String("o", "", "Output")
)

func main() {
	flag.Parse()

	fmt.Println("inputExpression", *inputExpression)
	fmt.Println("inputFile", *inputFile)
	fmt.Println("outputFile", *outputFile)

	var err error

	var fileOrExpression io.Reader
	if *inputFile != "" {
		fileOrExpression, err = os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
	} else {
		fileOrExpression = strings.NewReader(*inputExpression)
	}

	var out io.Writer
	var outF *os.File

	if *outputFile != "" {
		outF, err = os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE, 0600)
		defer outF.Close()
		out = outF
		if err != nil {
			panic(err)
		}
	} else {
		out = os.Stdout
	}

	handler := &lab2.ComputeHandler{fileOrExpression, out}
	err = handler.Compute()

	if err != nil {
		io.Copy(os.Stderr, strings.NewReader(err.Error()))
	}
}
