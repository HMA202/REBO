package main

import "fmt"

func main() {
	maze := []string{
		"S..#",
		".#..",
		"...E",
	}

	fmt.Println(MazeShortestPath(maze))
}
