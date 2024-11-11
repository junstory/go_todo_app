package main

import (
	"fmt"
	//"github.com/junstory/golangStudy"
)

// func test() {
// 	fmt.Println("test")
// }

func main() {
	defer fmt.Println("프로그램이 종료되었습니다.")
	//fmt.Println(golangStudy.IsPrime(3))
	test1 := "Hello \nWorld"
	test2 := `Hello \nWorld`

	fmt.Println(test1)
	fmt.Println(test2)

	var num1 int
	var num2 int
	var op string
	fmt.Println("Calculator")
	fmt.Println("첫 번째 숫자 입력: ")
	fmt.Scan(&num1)
	fmt.Println("두 번째 숫자 입력: ")
	fmt.Scan(&num2)
	fmt.Println("연산자 입력(+, -, *, /) : ")
	fmt.Scan(&op)
	calc(num1, num2, op)

	fmt.Println("\nEven or Odd")
	fmt.Print("숫자를 입력하세요:")
	fmt.Scan(&num1)
	if num1%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}

	fmt.Println("\n배열 분석")
	arr := []int{3, 5, 1, 2, 0}
	fmt.Print("배열 출력: ")
	for _, i := range arr {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println("배열의 총합: ", sumArray(arr))

	n_max, n_min := findMaxMin(arr)
	fmt.Println("최대값: ", n_max)
	fmt.Println("최소값: ", n_min)

	length := len(arr)
	switch {
	case length < 3:
		fmt.Println("배열의 길이가 짧습니다.")
	case length == 3:
		fmt.Println("배열의 길이가 적당합니다.")
	default:
		fmt.Println("배열의 길이가 깁니다.")
	}
}

func calc(num1 int, num2 int, op string) {
	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Print(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("0으로 나눌 수 없습니다.")
		} else {
			fmt.Print(float32(num1) / float32(num2))
		}
	default:
		fmt.Println("잘못된 연산자")
	}
}

func sumArray(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func findMaxMin(arr []int) (int, int) {
	n_max := arr[0]
	n_min := arr[0]
	for _, i := range arr {
		if i > n_max {
			n_max = i
		}
		if i < n_min {
			n_min = i
		}
	}
	return n_max, n_min
}
