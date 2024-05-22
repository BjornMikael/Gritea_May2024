package main

import (
	"os"
)

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isValidInt(s string) bool {
	if s == "" {
		return false
	}
	for i, c := range s {
		if !(isDigit(c) || (i == 0 && (c == '-' || c == '+'))) {
			return false
		}
	}
	return true
}

func atoi(s string) (int64, bool) {
	if !isValidInt(s) {
		return 0, false
	}
	var res int64
	var sign int64 = 1
	for i, c := range s {
		if c == '-' && i == 0 {
			sign = -1
		} else if c == '+' && i == 0 {
			sign = 1
		} else if isDigit(c) {
			res = res*10 + int64(c-'0')
		}
	}
	return sign * res, true
}

func printInt64(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	var buf [20]byte
	var i int
	var negative bool

	if n < 0 {
		negative = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		i++
		n /= 10
	}

	if negative {
		buf[i] = '-'
		i++
	}

	for j := i - 1; j >= 0; j-- {
		os.Stdout.Write([]byte{buf[j]})
	}
	os.Stdout.Write([]byte("\n"))
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	val1, ok1 := atoi(os.Args[1])
	operator := os.Args[2]
	val2, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	switch operator {
	case "+":
		result := val1 + val2
		if (val1 > 0 && val2 > 0 && result < 0) || (val1 < 0 && val2 < 0 && result > 0) {
			return // Overflow
		}
		printInt64(result)
	case "-":
		result := val1 - val2
		if (val1 > 0 && val2 < 0 && result < 0) || (val1 < 0 && val2 > 0 && result > 0) {
			return // Overflow
		}
		printInt64(result)
	case "*":
		result := val1 * val2
		if val1 != 0 && result/val1 != val2 {
			return // Overflow
		}
		printInt64(result)
	case "/":
		if val2 == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
		} else {
			printInt64(val1 / val2)
		}
	case "%":
		if val2 == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
		} else {
			printInt64(val1 % val2)
		}
	default:
		return
	}
}
