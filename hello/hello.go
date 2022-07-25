package main

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("[1]")

	slice1 := arr[2:6]
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))

	slice2 := arr[:5]
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice3 := arr[3:]
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))

	arr[3] = 100
	fmt.Println("[2]")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))

	slice2 = append(slice2, 200, 300, 400)

	fmt.Println("[3]")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))

	slice4 := slice2[1:3:4]
	slice5 := slice2[1:4:4]

	fmt.Println("[4]")
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
	fmt.Println("slice5:", slice5, len(slice5), cap(slice5))

	slice4 = append(slice4, 500)

	fmt.Println("[5]")
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
	fmt.Println("slice5:", slice5, len(slice5), cap(slice5))

}
