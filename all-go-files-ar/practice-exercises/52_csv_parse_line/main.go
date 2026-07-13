package main

import "fmt"

func main() {
	fmt.Println(CSVParseLine(`Ali,25,Bahrain`))
	fmt.Println(CSVParseLine(`Ali,"Manama, Bahrain",25`))
}
