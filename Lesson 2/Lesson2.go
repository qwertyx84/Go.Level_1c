package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
	var a, b float32
	fmt.Println("Task 1")
	fmt.Print("Введите первое число a: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число b: ")
	fmt.Scanln(&b)
	fmt.Printf("Площадь прямоугольника со сторонами a и b: %f\n", a*b)

	//Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
	var s float64
	var r float64
	fmt.Println("Task 2")
	fmt.Print("Введите площадь круга S: ")
	fmt.Scanln(&s)

	r = math.Sqrt(s / math.Pi)

	fmt.Printf("Диаметр окружности: %f\n", 2*r)
	fmt.Printf("Длина окружности: %f\n", 2*math.Pi*r)

	//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	var strVal string
	var intVal int
	fmt.Println("Введите трехзначное число: ")
	fmt.Scanln(&intVal)
	strVal = strconv.FormatInt(int64(intVal), 10)
	fmt.Printf("Число содержит %c сотен, %c десятков %c единиц\n", strVal[0], strVal[1], strVal[2])

	//через деление
	var digit1 int = intVal % 10
	var digit2 int = (intVal - digit1) / 10 % 10
	var digit3 int = (intVal - digit2*10 - digit1) / 100
	fmt.Printf("Число содержит %d сотен, %d десятков %d единиц\n", digit3, digit2, digit1)

	fmt.Scanf("")

}
