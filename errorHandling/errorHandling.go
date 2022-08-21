package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ------------------------
//	  File Read & Write
// ------------------------

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

// 오류가 발생하더라도 프로그램이 종료되는 것이 아닌, 후속 조치를 취하고 에러 메세지를 전달 함으로 사용자 경험에 도움을 줍니다.
func fileReadWriteProc() {
	const filename string = "data.txt"

	line, err := ReadFile(filename)

	if err != nil {
		fmt.Println("파일을 생성합니다.")
		err = WriteFile(filename, "This is WriteFile.")

		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}

		line, err = ReadFile(filename)

		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}

		fmt.Println("파일내용:", line)
	} else {
		fmt.Println("파일 읽기에 성공했습니다.")
		fmt.Println("파일내용:", line)
	}
}

// ------------------------
//	    Custom Error
// ------------------------

// 1. Errorf
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func sqrtProc() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}

// 2. Custom Error Type
type PasswordError struct {
	Len        int
	RequireLen int
}

// 문자열을 반환하는 Error 함수를 포함하므로 Error 타입으로 사용할 수 있습니다.
func (err PasswordError) Error() string {
	return "암호의 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func passwordErrorProc() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len: %d Requiredlen: %d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}

// 3. Error Rapping
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader((str)))
	scanner.Split(bufio.ScanWords)

	pos := 0
	res := 1
	for pos < len(str) {
		a, n, err := readNextInt(scanner)
		if err != nil {
			return 0, fmt.Errorf("Failed to readNextInt(), pos: %d err: %w", pos, err)
		}
		res *= a
		pos += n + 1
	}

	return res, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word: %s err: %w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func multipleFromStringProc() {
	readEq("123 3")
	readEq("123 abc")
	readEq("1 2 3 4 5")
}

// ------------------------
//	        Panic
// ------------------------

// 1. Panic
func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func divideProc() {
	divide(9, 3)
	divide(9, 0)
}

// 2. Panic Propagation
func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("계수는 0일 수 없습니다.")
	}
	return a / b
}

func panicProc() {
	f()
	fmt.Println("프로그램이 계속 실행됨")
}

func main() {
	// fileReadWriteProc()

	// sqrtProc()

	// passwordErrorProc()

	// multipleFromStringProc()

	// divideProc()

	panicProc()
}
