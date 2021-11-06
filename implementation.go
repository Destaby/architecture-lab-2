package lab2

import (
	"errors"
    "math"
    "strconv"
)

type Stack []int

func (array *Stack) IsStackEmpty() bool {
    return len(*array) == 0
}

func (array *Stack) Push(data int) {
    *array = append(*array, data)
}

func (array *Stack) Pop() bool {
    if array.IsStackEmpty() {
        return false
    } else {
        index := len(*array) - 1
        *array = (*array)[:index]
        return true
    }
}

func (array *Stack) Top() int {
    if array.IsStackEmpty() {
        return 0
    } else {
        index := len(*array) - 1
        element := (*array)[index]
        return element
    }
}

func CalculatePostfix(postfix string) (int, error) {
    var calcArray Stack
	if len(postfix) == 0 {
		return 0, errors.New("Empty input")
	};
	var atLeastOneNumber bool = false
    for _, char := range postfix {
        opchar := string(char)
		if opchar == " " {
		    continue
		}
        if opchar >= "0" && opchar <= "9" {
            i1, _ := strconv.Atoi(opchar)
            calcArray.Push(i1)
			atLeastOneNumber = true
        } else {
            opr1 := calcArray.Top()
            calcArray.Pop()
            opr2 := calcArray.Top()
            calcArray.Pop()
            switch char {
            case '^':
                x := math.Pow(float64(opr2), float64(opr1))
                calcArray.Push(int(x))
            case '+':
                calcArray.Push(opr2 + opr1)

            case '-':
                calcArray.Push(opr2 - opr1)

            case '*':
                calcArray.Push(opr2 * opr1)

            case '/':
                calcArray.Push(opr2 / opr1)

			default:
				return 0, errors.New("Invalid input")	
            }

        }
    }
	if !atLeastOneNumber {
		return 0, errors.New("Invalid input")
	}
    return calcArray.Top(), nil
}
