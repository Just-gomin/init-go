package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	slice := arr[2:5]

	fmt.Println(slice)

	slice = arr[3:]

	fmt.Println(slice)
}
