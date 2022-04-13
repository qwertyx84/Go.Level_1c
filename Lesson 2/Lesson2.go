package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
	var a, b float32
	fmt.Print("Task 1")
	fmt.Print("Введите первое число a: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число b: ")
	fmt.Scanln(&b)
	fmt.Printf("Площадь прямоугольника со сторонами a и b: %f\n", a*b)

	//Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
	var s float64
	var r float64
	fmt.Print("Task 2")
	fmt.Print("Введите площадь круга S: ")
	fmt.Scanln(&s)

	r = math.Sqrt(s / math.Pi)

	fmt.Printf("Диаметр окружности: %f\n", 2*r)
	fmt.Printf("Длина окружности: %f\n", 2*math.Pi*r)

	//С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	var strVal string
	var intVal int64
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&intVal)
	strVal = strconv.FormatInt(intVal, 10)
	fmt.Printf("Число содержит %c сотен, %c десятков %c единиц\n", strVal[0], strVal[1], strVal[2])

	fmt.Scanf(" ")

}
