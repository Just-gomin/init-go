package main

import (
	"fmt"
)

/* -------------------------------- */
const consoleRed string = "\033[31m"
const consoleGreen string = "\033[32m"
const consoleYellow string = "\033[33m"
const consoleBlue string = "\033[34m"
const consolePurple string = "\033[35m"
const consoleCyan string = "\033[36m"
const consoleWhite string = "\033[37m"

/* -------------------------------- */

type Container struct {
	amount   int // 내장된 물품의 수량
	capacity int // 해당 컨테이너의 최대 수용량
}

// 컨테이너의 잔여 수용량
func (container *Container) remain() int {
	return container.capacity - container.amount
}

// 물량 입고 포인터 메소드
func (container *Container) packageInPointer(num int) {
	if num > container.remain() {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Exceeded the capacity.")
	} else {
		container.amount += num
	}
}

// 물량 출고 포인터 메소드
func (container *Container) packageOutPointer(num int) {
	if num > container.amount {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Out of stock.")
	} else {
		container.amount -= num
	}
}

// 물량 입고 값 메소드
func (container Container) packageInValue(num int) {
	if num > container.remain() {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Exceeded the capacity.")
	} else {
		container.amount += num
	}
}

// 물량 출고 값 메소드
func (container Container) packageOutValue(num int) {
	if num > container.amount {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Out of stock.")
	} else {
		container.amount -= num
	}
}

// 물량 입고 값 타입 메서드
func (container Container) packageInReturn(num int) Container {
	if num > container.remain() {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Exceeded the capacity.")
	} else {
		container.amount += num
	}

	return container
}

// 물량 출고 값 타입 메서드
func (container Container) packageOutReturn(num int) Container {
	if num > container.amount {
		fmt.Println(consoleRed, "[!]", consoleWhite, "Out of stock.")
	} else {
		container.amount -= num
	}

	return container
}

func main() {
	myContainer := Container{0, 100}

	fmt.Println("[myContainerInfo]", myContainer.amount, myContainer.remain())

	// 1) PackageIn Test
	myContainer.packageInPointer(10)
	fmt.Println(consoleBlue, "[packageInPointer]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer.packageInValue(10)
	fmt.Println(consoleGreen, "[packageInValue]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer.packageInReturn(10)
	fmt.Println(consolePurple, "[packageInReturn]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer = myContainer.packageInReturn(10)
	fmt.Println(consolePurple, "[packageInReturn]", consoleWhite, myContainer.amount, myContainer.remain())

	// 2) PackageOut Test
	myContainer.packageOutPointer(10)
	fmt.Println(consoleBlue, "[packageOutPointer]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer.packageOutValue(10)
	fmt.Println(consoleGreen, "[packageOutValue]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer.packageOutReturn(10)
	fmt.Println(consolePurple, "[packageOutReturn]", consoleWhite, myContainer.amount, myContainer.remain())

	myContainer = myContainer.packageOutReturn(10)
	fmt.Println(consolePurple, "[packageOutReturn]", consoleWhite, myContainer.amount, myContainer.remain())

	// 3) out of range test
	myContainer.packageInPointer(110)
	myContainer.packageOutPointer(110)
}
