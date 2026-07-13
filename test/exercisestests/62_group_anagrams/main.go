package main

import "fmt"

func main() {
	groups := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	for _, group := range groups {
		if len(group) > 1 {
			fmt.Println(group)
		}
	}
}
