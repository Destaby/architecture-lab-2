package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestPostfix(t *testing.T) {
	res, err := CalculatePostfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 11, res)
	}

	res, err = CalculatePostfix("1 2 + 4 * 3 +")
	if assert.Nil(t, err) {
		assert.Equal(t, 15, res)
	}

	res, err = CalculatePostfix("3 4 2 * 1 3 - 2 ^ / +")
	if assert.Nil(t, err) {
		assert.Equal(t, 5, res)
	}

	res, err = CalculatePostfix("9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *")
	if assert.Nil(t, err) {
		assert.Equal(t, 612, res)
	}

	res, err = CalculatePostfix("1 7 - 9 * 3 4 + 7 / 1 4 6 - * + +")
	if assert.Nil(t, err) {
		assert.Equal(t, -55, res)
	}

	res, err = CalculatePostfix("")
	assert.NotNil(t, err)

	res, err = CalculatePostfix(" ")
	assert.NotNil(t, err)

	res, err = CalculatePostfix("test")
	assert.NotNil(t, err)
}

func ExamplePostfix(t *testing.T) {
	res, err := CalculatePostfix("1 2 + 4 * 3 +")
	if assert.Nil(t, err) {
		fmt.Println(res)
	}

	// Output:
    // 15
}
