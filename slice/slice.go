package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	idx := 1

	fmt.Println(slice1)

	// 1)
	// slice1 = append(slice1, 0)

	// for i := len(slice1) - 2; i >= idx; i-- {
	// 	slice1[i+1] = slice1[i]
	// }

	// slice1[idx] = 200

	// fmt.Println("After ---------------")

	// 2)
	// temp := slice1[idx:]
	// slice1 = append(append(slice1[:idx:1], 200), temp...)

	temp := make([]int, idx)
	copy(temp, slice1[:idx])
	temp = append(temp, 200)
	slice1 = append(temp, slice1[idx:]...)

	fmt.Println(slice1)
}
