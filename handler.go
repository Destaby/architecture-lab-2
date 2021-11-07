package lab2

import (
	"io"
	"strings"
	"bytes"
	"strconv"
)

type ComputeHandler struct { 
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(ch.Input)

	res, err := CalculatePostfix(buf.String())
	if (err != nil) {
		return err
	}

	resString := strconv.Itoa(res)
	_, err = io.Copy(ch.Output, strings.NewReader(resString))
	if (err != nil) { 
		return err
	}
	return nil
}
