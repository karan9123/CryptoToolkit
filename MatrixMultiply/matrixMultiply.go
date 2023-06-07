package main

import (
	"fmt"
	"strconv"
)

// takePositiveInt reads an integer input from the user and validates if it is a positive number.
// It returns the validated positive integer.
func takePositiveInt() int {
	var input string
	var number int
	var err error

	for {
		_, errored := fmt.Scanln(&input)

		if errored != nil {
			panic("Something went wrong while reading input. Please run the program again.")
		}

		if number, err = strconv.Atoi(input); err != nil {
			fmt.Println("Invalid input. Please enter a positive integer.")
			continue
		}

		if number <= 0 {
			fmt.Println("Invalid input. Please enter a positive integer.")
			continue
		}

		break
	}

	fmt.Println("You entered:", number)

	return number
}

// createMatrix creates a matrix with the specified number of rows and columns.
// It reads input from the user to populate the matrix elements and returns the created matrix.
func createMatrix(rows, columns int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("Enter the value for element [%d][%d]: ", i, j)
			_, errored := fmt.Scanln(&matrix[i][j])

			if errored != nil {
				panic("Something went wrong while reading input. Please run the program again.")
			}
		}
	}

	fmt.Println("The matrix you entered is:")
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	return matrix
}

// multiplyMatrices multiplies two matrices and returns the resulting matrix.
// It validates the sizes of the matrices and panics if they cannot be multiplied.
func multiplyMatrices() [][]int {
	fmt.Println("Enter the number of rows for the first matrix:")
	rowsA := takePositiveInt()
	fmt.Println("Enter the number of columns for the first matrix:")
	colsA := takePositiveInt()
	fmt.Println("Enter the number of rows for the second matrix:")
	rowsB := takePositiveInt()
	fmt.Println("Enter the number of columns for the second matrix:")
	colsB := takePositiveInt()

	if colsA != rowsB {
		panic("Matrices cannot be multiplied. Number of columns in the first matrix must be equal to the number of rows in the second matrix.")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	matrixA := createMatrix(rowsA, colsA)
	matrixB := createMatrix(rowsB, colsB)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

func main() {
	resultantMatrix := multiplyMatrices()

	fmt.Println("\nResultant Matrix:")
	fmt.Println()
	for i := 0; i < len(resultantMatrix); i++ {
		for j := 0; j < len(resultantMatrix[0]); j++ {
			fmt.Printf("%d ", resultantMatrix[i][j])
		}
		fmt.Println()
	}
}
