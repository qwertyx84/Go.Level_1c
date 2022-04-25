package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// InsertionSort do sort slice []int64
func InsertionSort(s []int64) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func main() {

	var inputNums []int64

	scanner := bufio.NewScanner(os.Stdin)
	// записываем введённые числа в массив
	for i := 0; true; i++ {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				break
			}
			inputNums = append(inputNums, num)
		}
	}

	InsertionSort(inputNums)
	fmt.Println(inputNums)

}
