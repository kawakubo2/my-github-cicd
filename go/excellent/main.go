package main

// import (
// 	"fmt"
// )
// func main() {
// 	var num int
// 	fmt.Print("整数値: ")
// 	fmt.Scan(&num)
// 	fmt.Printf("%v is %v\n", num, EvenOrOdd(num))
// }

func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}