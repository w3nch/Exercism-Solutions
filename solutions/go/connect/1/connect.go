package connect

import "strings"

func ResultOf(grid []string) (string, error) {
	h := len(grid)
	if h == 0 {
		return "", nil
	}
	
	flat := make([][]byte, h)
	for i, row := range grid {
		flat[i] = []byte(strings.ReplaceAll(row, " ", ""))
	}
	w := len(flat[0])

	// Check Player O (Top to Bottom)
	for j := 0; j < w; j++ {
		if flat[0][j] == 'O' {
			seen := make([][]bool, h)
			for i := range seen {
				seen[i] = make([]bool, w)
			}
			if walk(flat, 0, j, h, w, 'O', seen) {
				return "O", nil
			}
		}
	}

	// Check Player X (Left to Right)
	for i := 0; i < h; i++ {
		if flat[i][0] == 'X' {
			seen := make([][]bool, h)
			for i := range seen {
				seen[i] = make([]bool, w)
			}
			if walk(flat, i, 0, h, w, 'X', seen) {
				return "X", nil
			}
		}
	}

	return "", nil
}

func walk(data [][]byte, r, c, h, w int, char byte, seen [][]bool) bool {
	if char == 'O' && r == h-1 {
		return true
	}
	if char == 'X' && c == w-1 {
		return true
	}

	seen[r][c] = true

	// Hexagonal adjacent offsets
	coords := [][2]int{
		{r - 1, c}, {r - 1, c + 1},
		{r, c - 1}, {r, c + 1},
		{r + 1, c - 1}, {r + 1, c},
	}

	for _, p := range coords {
		nr, nc := p[0], p[1]
		if nr >= 0 && nr < h && nc >= 0 && nc < w && data[nr][nc] == char && !seen[nr][nc] {
			if walk(data, nr, nc, h, w, char, seen) {
				return true
			}
		}
	}
	return false
}