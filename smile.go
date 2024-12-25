package main

import "fmt"

func main() {
	var mobileNum interface{} = int64(8073215402)

	if num, ok := mobileNum.(int64); ok {
		fmt.Println("Mobile number is valid:", num)
	} else {
		fmt.Println("Mobile number is not valid")
	}
}
