package main

import (
	"fmt"
)

func add(inputChan, outputChan chan int) {
	num1 := <-inputChan
	num2 := <-inputChan
	outputChan <- num1 + num2
}

func mul(inputChan, outputChan chan int) {
	ret := 1
	for x := range inputChan {
		fmt.Print(x, " ")
		ret *= x
		fmt.Print(ret, " ")
	}
	outputChan <- ret
}

func main() {

	var num1 int
	var num2 int
	//var op string
	inputChannel := make(chan int, 2)
	inputChannel2 := make(chan int, 2)
	//	opChannel := make(chan int, 1)
	outputChannel := make(chan int, 2)
	fmt.Print("첫 번째 정수 : ")
	fmt.Scan(&num1)
	fmt.Print("첫 번째 정수 : ")
	fmt.Scan(&num2)
	go add(inputChannel, outputChannel)
	go mul(inputChannel2, outputChannel)
	//go printer(outputChannel)
	inputChannel <- num1
	inputChannel <- num2
	inputChannel2 <- num1
	inputChannel2 <- num2
	//	opChannel <- op
	// for i := 0; i < 1; i++ {
	// 	fmt.Println("Add : ", <-outputChannel)
	// 	//fmt.Println("Mul : ", <-outputChannel)
	// }
	close(inputChannel)
	close(inputChannel2)
}
