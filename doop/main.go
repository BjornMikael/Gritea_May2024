package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	val1, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	operator := os.Args[2]
	val2, err2 := strconv.ParseInt(os.Args[3], 10, 64)

	if err1 != nil || err2 != nil {
		return
	}

	switch operator {
	case "+":
		result := val1 + val2
		if (val1 > 0 && val2 > 0 && result < 0) || (val1 < 0 && val2 < 0 && result > 0) {
			return // Overflow
		}
		fmt.Println(result)
	case "-":
		result := val1 - val2
		if (val1 > 0 && val2 < 0 && result < 0) || (val1 < 0 && val2 > 0 && result > 0) {
			return // Overflow
		}
		fmt.Println(result)
	case "*":
		result := val1 * val2
		if val1 != 0 && result/val1 != val2 {
			return // Overflow
		}
		fmt.Println(result)
	case "/":
		if val2 == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(val1 / val2)
		}
	case "%":
		if val2 == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(val1 % val2)
		}
	default:
		return
	}
}
