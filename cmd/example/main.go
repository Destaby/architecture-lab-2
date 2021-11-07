package main

import (
  "flag"
  "fmt"
  "os"
  "io"
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

  // TODO

  var out io.Writer

  // var outF *os.File

  // if *outputFile != "" {
    out = os.Stdout
  // } else {
  //   outF, err = os.OpenFile(*outputFile, os.O_WRONLY|os.O_CREATE, 0600)
  //   out = outF.Writer
  //   defer outF.Close()
  //   // var _, err = os.Stat(*outputFile)

  //   // if os.IsNotExist(err) {
  //   //   out, err = os.Create(*outputFile)
  //   //   defer out.Close()
  //   // } else {
  //   //   out, err = os.Open(*outputFile)
  //   //   defer out.Close()
  //   // }
  //   if err != nil {
  //     panic(err)
  //   }
  // }

  // END TODO
 
  handler := &lab2.ComputeHandler{fileOrExpression, out}
  err = handler.Compute()

  fmt.Println(err)
}