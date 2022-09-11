package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic Test Code
func TestSquare1(t *testing.T) {
	fmt.Println("[TestSquare1] Start...")
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) returns %d", rst)
	}
	fmt.Println("[TestSquare1] Finish...")
}

func TestSquare2(t *testing.T) {
	fmt.Println("[TestSquare2] Start...")
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be 9 but square(3) returns %d", rst)
	}
	fmt.Println("[TestSquare2] Finish...")
}

// Test Code with Testify
func TestSquare3(t *testing.T) {
	fmt.Println("[TestSquare3] Start...")
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
	fmt.Println("[TestSquare3] Start...")
}

func TestSquare4(t *testing.T) {
	fmt.Println("[TestSquare4] Start...")
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
	fmt.Println("[TestSquare4] Start...")
}
