package main // This line declares the package name as main, which means this file will serve as the entry point for the executable program.

import (
	"fmt" //Provides functions for formatted I/O operations, including printing to the console.
	"os"  //Offers functions for interacting with the operating system, such as accessing command-line arguments.
)

const size = 9 // Defines a constant size set to 9, representing the standard size of a Sudoku grid.

// Declares two global variables:
var FaultFound bool        // mA boolean flag to indicate if any error occurred during the Sudoku board creation or solving process.
var PossibleSolutionNr int // An integer to keep track of the number of possible solutions found.

// Starts the definition of a function named CreateSudokuBoard.
// This function returns a 2D slice of integers ([][]int), representing the Sudoku board.
func CreateSudokuBoard() [][]int {
	argsInArray := os.Args          //argsInArray := os.Args: Stores command-line arguments in argsInArray.
	board := make([][]int, size)    //board := make([][]int, size): Initializes a 2D slice board with dimensions equal to size.
	if len(argsInArray) != size+1 { // Checks if the number of command-line arguments plus the program name equals size+1.
		FaultFound = true //If not, it indicates an error, sets FaultFound to true, and returns the empty board.
		return board
	}
	for i := range board { // Iterates over each line of the board, initializing each line as a slice of integers with a length of size.
		board[i] = make([]int, size)
		for j, value := range argsInArray[i+1] { //Iterates over each character in the i+1th command-line argument.
			if len(argsInArray[i+1]) != size { //Checks if the length of the i+1th argument matches size.
				return board //If not, it indicates an error, sets FaultFound to true, and returns the empty board.
				FaultFound = true
			}
			if value == '.' { // If the character is '.', it replaces it with 0.
				board[i][j] = 0
			} else if value >= '1' && value <= '9' { ////If the character is between '1' and '9',
				board[i][j] = int(value - '0') //it converts it to its numeric ASCII value minus '0' to get the actual digit.
			} else { //// Otherwise, it indicates an error, sets FaultFound to true, and returns the empty board.
				FaultFound = true
				return board
			}
		}
	}
	// checking for duplicate numbers and calling a TheRightDigit func for that
	for line := 0; line < size; line++ { //Nested loops iterate over each cell of the board to check for duplicates.
		for column := 0; column < size; column++ {
			if board[line][column] != 0 { //If the current cell is not empty (!= 0),
				temp_box := board[line][column] //it temporarily stores its value in temp_box and clears the cell.
				board[line][column] = 0
				if !TheRightDigit(board, line, column, temp_box) { //Calls TheRightDigit to check if the temporary value is the right digit for the current position.
					aultFound = true // If not, it indicates an error, sets FaultFound to true,
					return board     //and returns the empty board.
				}
				board[line][column] = temp_box //Then, restores the original value of the cell.
			}
		}
	}
	return board //Returns the initialized and validated Sudoku board.
}

// Recursive function which attepmts to fill in Sudoku board.
// This function attempts to fill in the Sudoku board. It takes two parameters:
//the Sudoku board and a boolean findMultiple indicating whether to find multiple solutions or just one.

func FillinSudoku(board [][]int, findMultiple bool) bool {

	line, column := EmptyCell(board)
	if line == -1 { //If no empty cells are found (line == -1),
		PossibleSolutionNr++ //it increments PossibleSolutionNr
		return true          //and returns true, indicating a successful completion of the puzzle-solving process.
	}
	for num := 1; num <= size; num++ { //Loops through numbers from 1 to size (9 in a standard Sudoku).
		if TheRightDigit(board, line, column, num) { //Checks if placing the current number in the current cell leads to a valid Sudoku solution by calling TheRightDigit.
			board[line][column] = num              //Places the number in the current cell if it passes the validation.
			if FillinSudoku(board, findMultiple) { //Recursively calls FillinSudoku to continue filling the board.
				if !findMultiple || PossibleSolutionNr > 1 { //If findMultiple is false or more than one solution has been found,
					return true //it returns true, indicating success.
				}
			}
			board[line][column] = 0 //Restores the cell to 0 after trying all numbers.
		}
	}
	return false //Returns false if no valid number was found for the current cell, indicating failure to find a solution from the current path.
}

// Finds the next empty cell (=='0') on the board where a number have to be placed.
func EmptyCell(board [][]int) (int, int) { //Finds the next empty cell on the board.
	for i := 0; i < size; i++ { //Iterates over every cell of the board.
		for j := 0; j < size; j++ {
			if board[i][j] == 0 { //Returns the coordinates of the first empty cell found.
				return i, j
			}
		}
	}
	return -1, -1 // Returns the coordinates of the first empty cell found.
}

// Checking each number from 1 to 9, if placing that number in the current cell would solve Sudoku according to Sudoku rules.
// Checking for duplicates in the same line, column, or current 3x3 box.

func TheRightDigit(board [][]int, line, column, num int) bool { //Checks if placing a number in a specific cell would result in a valid Sudoku solution.
	for i := 0; i < size; i++ { //Loops through each row.
		if board[line][i] == num || board[i][column] == num { // Checks if the number already exists in the same row or column.
			return false // If so, returns false.
		}
	}
	startline, startcolumn := (line/3)*3, (column/3)*3 // Calculates the top-left cell of the current 3x3 box.
	for i := startline; i < startline+3; i++ {         // Iterates over the cells in the current 3x3 box.
		for j := startcolumn; j < startcolumn+3; j++ {
			if board[i][j] == num { // Checks if the number exists in the current 3x3 box. If so, returns false.
				return false
			}
		}
	}
	return true // Returns true if no conflicts were found, indicating the number can be placed in the current cell.
}

// Function prints board passed as an argument using space as a delimeter of cells
func printBoard(board [][]int) { //Prints the Sudoku board to the console.
	for i := 0; i < size; i++ { //Iterates over each cell of the board.
		for j := 0; j < size; j++ {
			fmt.Printf("%v", board[i][j]) //Prints the value of the current cell.
			if j < size-1 {               // Adds a space after each cell except the last one in a row.
				fmt.Printf(" ")
			}
		}
		fmt.Println() //Moves to the next line after printing each row.
	}
}

// Main (executive) program
func main() { // Main function where the program execution begins.
	FaultFound = false           //Initializes FaultFound to false.
	board := CreateSudokuBoard() // Creates the Sudoku board.
	if FaultFound {              //Checks if any errors occurred during board creation. If so, prints "Error" and exits.
		fmt.Println("Error")
		return
	}
	// First, count the number of solutions.
	PossibleSolutionNr = 0 //Counts the number of solutions by setting findMultiple to true.
	FillinSudoku(board, true)
	if PossibleSolutionNr == 0 { // Depending on the number of solutions found,
		fmt.Println("Error") //either prints an error message or fills in the board with the solution and prints it.
	} else if PossibleSolutionNr == 1 { // in case only one solution is found fill in board with solution and print it
		FillinSudoku(board, false)
		printBoard(board)
	} else { // in case more than one solution is found return error message
		fmt.Println("Error")
	}
}
