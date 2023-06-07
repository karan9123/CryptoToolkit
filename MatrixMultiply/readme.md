# Matrix Multiplication

This program multiplies two matrices entered by the user. It validates the matrix sizes and performs the multiplication if they are compatible.

## Prerequisites

- Go programming language (Golang) installed

## Usage

1. Clone the repository or download the source code.

2. Navigate to the project directory.

3. Run the following command to compile and execute the program:

   ```shell
   go run matrixMultiply.go
   ```
   
Follow the instructions provided by the program to enter the matrix dimensions and values.

The program will multiply the matrices and display the resultant matrix.

Functionality
The program ensures that the entered matrix dimensions are positive integers and performs input validation.

It checks if the number of columns in the first matrix is equal to the number of rows in the second matrix, as this is a requirement for matrix multiplication. If the sizes are incompatible, the program will display an error message and terminate.

The program prompts the user to enter the values for each element of the matrices.

After multiplying the matrices, the program displays the resultant matrix.

## Example:
```
Enter the number of rows for the first matrix:
2
Enter the number of columns for the first matrix:
3
Enter the number of rows for the second matrix:
3
Enter the number of columns for the second matrix:
2

Enter the value for element [0][0]: 1
Enter the value for element [0][1]: 2
Enter the value for element [0][2]: 3
Enter the value for element [1][0]: 4
Enter the value for element [1][1]: 5
Enter the value for element [1][2]: 6

The matrix you entered is:
1 2 3
4 5 6

Enter the value for element [0][0]: 7
Enter the value for element [0][1]: 8
Enter the value for element [1][0]: 9
Enter the value for element [1][1]: 10
Enter the value for element [2][0]: 11
Enter the value for element [2][1]: 12

The matrix you entered is:
7 8
9 10
11 12

Resultant Matrix:

58 64
139 154
```