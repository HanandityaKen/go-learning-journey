package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    // Read input
    var dimensions string
    var targetStr string
    fmt.Scanln(&dimensions)
    fmt.Scanln(&targetStr)

    // Parse dimensions
    dimParts := strings.Split(dimensions, ",")
    rows, _ := strconv.Atoi(dimParts[0])
    cols, _ := strconv.Atoi(dimParts[1])

    // Parse target number
    target, _ := strconv.Atoi(targetStr)

    // Predefined grid
    grid := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
    }

    // Flag untuk mengecek apakah target ditemukan
    found := false

    // Labeled break
		searchLoop:
				for i := 0; i < rows; i++ {
						for j := 0; j < cols; j++ {
								// Cek agar tidak keluar dari batas grid
								if i >= len(grid) || j >= len(grid[i]) {
										continue
								}

								if grid[i][j] == target {
										fmt.Printf("Found %d at position (%d, %d)\n", target, i, j)
										found = true
										break searchLoop
								}
						}
				}

    if !found {
        fmt.Printf("Target %d not found\n", target)
    }
}
