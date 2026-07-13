package main

type point struct {
	row  int
	col  int
	dist int
}

func MazeShortestPath(maze []string) int {
	if len(maze) == 0 {
		return -1
	}

	rows := len(maze)
	cols := len(maze[0])

	start := point{-1, -1, 0}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if maze[r][c] == 'S' {
				start = point{r, c, 0}
			}
		}
	}

	if start.row == -1 {
		return -1
	}

	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	queue := []point{start}
	visited[start.row][start.col] = true

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if maze[current.row][current.col] == 'E' {
			return current.dist
		}

		for _, d := range directions {
			nr := current.row + d[0]
			nc := current.col + d[1]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if !visited[nr][nc] && maze[nr][nc] != '#' {
					visited[nr][nc] = true
					queue = append(queue, point{nr, nc, current.dist + 1})
				}
			}
		}
	}

	return -1
}
