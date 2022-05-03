package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	var uint64Res uint64 = 1

	fmt.Print("Введите арифметическую операцию (+, -, *, /,!,^): ")
	fmt.Scanln(&op)

	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	if op == "!" {
		if a-float64(uint64(a)) > 0 {
			fmt.Print("Операция вычисляется только для целого числа!")
			os.Exit(1)
		} else if a < 0 {
			fmt.Print("Операция вычисляется только для положительных чисел!")
			os.Exit(1)
		}

	} else {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("На 0 делить нельзя!")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "!":
		for i := 1; i <= int(a); i++ {
			uint64Res *= uint64(i)
		}
		res = float64(uint64Res)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
	fmt.Scanln()
}
