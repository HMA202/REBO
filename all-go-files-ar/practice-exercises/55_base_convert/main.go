package main

import "fmt"

func main() {
	fmt.Println(BaseConvert("1010", 2, 10))
	fmt.Println(BaseConvert("255", 10, 16))
	fmt.Println(BaseConvert("FF", 16, 2))
}
