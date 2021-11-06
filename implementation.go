package lab2

import (
    "fmt"
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

func CalculatePostfix(postfix string) int {
    var calcArray Stack
    for _, char := range postfix {
        opchar := string(char)
				if opchar == " " {
					continue
				}
        if opchar >= "0" && opchar <= "9" {
            i1, _ := strconv.Atoi(opchar)
            calcArray.Push(i1)
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
            }
        }
    }
    return calcArray.Top()
}