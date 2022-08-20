package main

import "fmt"

func main() {
	generalFunc("hi", 6, true)

	fmt.Println(sum(1, 2, 3, 4, 5))
	// fmt.Println(sum2(1, 2, 3, 4, 5, 6))
	fmt.Println(sum())
	fmt.Println(sum2([]int{1, 2, 3}))
}

func generalFunc(args ...interface{}) {
	fmt.Printf("args Type : %T\n", args)

	for _, arg := range args {
		switch f := arg.(type) {

		case bool:
			val := arg.(bool)
			fmt.Printf("[%T] %t\n", f, val)

		case string:
			val := arg.(string)
			fmt.Printf("[%T] %s\n", f, val)

		case int:
			val := arg.(int)
			fmt.Printf("[%T] %c\n", f, val)

		case int8:
			val := arg.(int8)
			fmt.Printf("[%T] %c\n", f, val)

		case int16:
			val := arg.(int16)
			fmt.Printf("[%T] %c\n", f, val)

		case int32:
			val := arg.(int32)
			fmt.Printf("[%T] %c\n", f, val)

		case int64:
			val := arg.(int64)
			fmt.Printf("[%T] %c\n", f, val)

		case float32:
			val := arg.(float32)
			fmt.Printf("[%T] %f\n", f, val)

		case float64:
			val := arg.(float64)
			fmt.Printf("[%T] %f\n", f, val)
			/*
				다른 타입들에 대해서 위와 같은 로직
			*/
		}
	}
}

func sum(nums ...int) int {
	sum := 0
	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func sum2(nums []int) int {
	sum := 0
	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}
