package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

// by reference
func changeNum1(num *int) { //* pointer location means address   
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	changeNum(num)
	changeNum1(&num)

	fmt.Println("Memory address", &num)//& memory address show
	fmt.Println("After changeNum in main", num)

} 