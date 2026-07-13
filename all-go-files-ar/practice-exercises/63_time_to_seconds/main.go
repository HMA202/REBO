package main

import "fmt"

func main() {
	fmt.Println(TimeToSeconds("01:02:03"))
	fmt.Println(TimeToSeconds("10:00:00"))
	fmt.Println(TimeToSeconds("bad"))
}
