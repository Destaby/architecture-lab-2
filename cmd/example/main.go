package main

import (
	"flag"
	"fmt"
	//lab2 "https://github.com/Destaby/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File")
	outputFile      = flag.String("o", "", "Output")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	//res, _ := lab2.CalculatePostfix("+ 2 2")
	//fmt.Println(res)
	fmt.Println("inputExpression", *inputExpression)
	fmt.Println("inputFile", *inputFile)
	fmt.Println("outputFile", *outputFile)
}
