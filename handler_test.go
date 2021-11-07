package lab2

import (
	"testing"
	"io"
	"strings"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
  var out io.Writer
  fileOrExpressionReader := strings.NewReader("4 2 - 3 * ? & 5 +")
  handler := &ComputeHandler{fileOrExpressionReader, out}
  err := handler.Compute()
  assert.NotNil(t, err)

  fileOrExpressionReader = strings.NewReader("4 2 - 3 * 5 +")
  buf := new(bytes.Buffer)
  handler = &ComputeHandler{fileOrExpressionReader, buf}
  err = handler.Compute()

  assert.Nil(t, err)

  if assert.Nil(t, err) {
	assert.Equal(t, "11", buf.String())
  }
}
