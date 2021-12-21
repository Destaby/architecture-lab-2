package lab2

import (
	"testing"
	"fmt"
)

func BenchmarkCalculatePostfix(b *testing.B) { 
	postfixes := []string{
		"4 2 - 3 * 5 +",
		"9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / +",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 +",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / +",
		"4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 +",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
		"4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 1 2 + 4 * 3 + 4 2 - 3 * 6 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + * 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 4 2 - 3 * 5 + 4 2 - 3 * 5 + 3 4 2 * 1 3 - 2 ^ / + 3 4 2 * 1 3 - 2 ^ / + 9 3 4 + * 6 / 1 - 8 2 ^ 5 1 - + *",
	}
  l := len(postfixes);
	for i := 0; i < l; i++ { 
		b.Run(fmt.Sprintf("len=%d", len(postfixes[i])), func(b *testing.B) { 
			CalculatePostfix(postfixes[i]) 
		}) 
	} 
}
