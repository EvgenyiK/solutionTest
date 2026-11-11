package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Println("Введите количество чисел N: ")
	if _,err:= fmt.Scan(&n); err !=nil || n<=0{
		fmt.Println("Ошибка: введите целое положительное число.")
		return
	}

	found:= 0

	for num := 0; found < n; num++ {
		digits := len(strconv.Itoa(num))

		if num%digits == 0 {
			fmt.Printf("%d ", num)
			found++
		}
	}
	fmt.Println()

}