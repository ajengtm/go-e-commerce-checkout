package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader to read input from standard input (console)
	reader := bufio.NewReader(os.Stdin)
	var m, n int

	// Read the width of the grid (m)
	fmt.Fscanf(reader, "%d\n", &m)
	// Read the height of the grid (n)
	fmt.Fscanf(reader, "%d\n", &n)

	// Allocate a 2D slice to store the grid
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		// Read each line of the grid, which is a string of '.' and '#'
		// Slice the line to keep only the first m characters (since Fscanf doesn't consume the newline)
		line, _ := reader.ReadString('\n')
		grid[i] = line[:m]
	}

	// Call the function to count lakes and print the result
	fmt.Println(countLakes(grid, m, n))
}

// Function to count the number of lakes in the grid
func countLakes(grid []string, m, n int) int {
	// Create a 2D slice to keep track of visited cells
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m) // Initialize each row with false
	}

	// Variable to keep track of the number of lakes
	lakes := 0

	// Iterate over each cell in the grid
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// Check if the cell is water and unvisited
			if grid[i][j] == '.' && !visited[i][j] {
				// Start a flood-fill from this cell
				floodFillDFS(grid, visited, i, j, m, n)
				lakes++ // Increment the lake count
			}
		}
	}

	// Return the total number of lakes found
	return lakes
}

// Function to perform a depth-first search (DFS) flood-fill from a given cell
func floodFillDFS(grid []string, visited [][]bool, x, y, m, n int) {
	// Use a stack to facilitate the DFS process
	stack := [][]int{{x, y}} // Start the stack with the initial cell
	visited[x][y] = true     // Mark the starting cell as visited

	// Possible 8 directions (up, down, left, right, and the 4 diagonals)
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	// Continue while there are cells to process in the stack
	for len(stack) > 0 {
		// Pop the current cell from the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Explore all 8 possible directions
		for _, d := range directions {
			newX, newY := current[0]+d[0], current[1]+d[1] // Calculate new coordinates

			// Check if the new coordinates are valid, contain water, and are unvisited
			if isValid(newX, newY, m, n) && grid[newX][newY] == '.' && !visited[newX][newY] {
				visited[newX][newY] = true               // Mark the new cell as visited
				stack = append(stack, []int{newX, newY}) // Add the new cell to the stack
			}
		}
	}
}

// Helper function to check if a cell is within the grid boundaries
func isValid(x, y, m, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
}
