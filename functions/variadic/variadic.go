package main

import "fmt"

func main() {
	generalFunc("hi", 6, true)
}
func generalFunc(args ...interface{}) {
	for _, arg := range args {
		switch f := arg.(type) {

		case bool:
			val := arg.(bool)
			fmt.Printf("[%t] %t", f, val)

		case string:
			val := arg.(string)
			fmt.Printf("[%s] %s", f, val)

		case int:
			val := arg.(int)
			fmt.Printf("[%c] %c", f, val)

		case int8:
			val := arg.(int8)
			fmt.Printf("[%c] %c", f, val)

		case int16:
			val := arg.(int16)
			fmt.Printf("[%c] %c", f, val)

		case int32:
			val := arg.(int32)
			fmt.Printf("[%c] %c", f, val)

		case int64:
			val := arg.(int64)
			fmt.Printf("[%c] %c", f, val)

		case float32:
			val := arg.(float32)
			fmt.Printf("[%f] %f", f, val)

		case float64:
			val := arg.(float64)
			fmt.Printf("[%f] %f", f, val)
			/*
				다른 타입들에 대해서 위와 같은 로직
			*/
		}
	}
}
