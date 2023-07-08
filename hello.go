package main

import "fmt"

func main() {
	fmt.Print("Enter number : ")
	var num int
	fmt.Scan(&num)
	for i := 0; i <= 12; i++ {
		fmt.Println(num, "*", i, "=", num*i)
	}
}
