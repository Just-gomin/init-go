package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)
	slice4 := make([]int, len(slice1))

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)
	cnt3 := copy(slice4, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
	fmt.Println(cnt3, slice4)
}
