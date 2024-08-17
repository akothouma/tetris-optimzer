# tetris-optimizer

This  program is writen in Golang

## Prerequisite
go version 1.22.2 or higher

## Objective
 The program receives only one argument, a path to a text file which will contain a list of tetrominoes and assembles them in order to create the smallest square possible.

## Exacution logic
Tetrominoes are read from the sample.txt file

Tetrominoes are validated as correct depending on the connections,number of # and length of each tetromino lines.

Then they are optimized for placing on the board by removing the rows and columns that do not contribute to the tetromino shape.

 A tetromino is placed on the board recurrsively  and using backtracking algorithm for optimization.Should they not fill the board,the board size is increased until a solution is reached.


## Cloning the project
```bash
git clone https://learn.zone01kisumu.ke/git/lakoth/tetris-optimizer
cd tetris-optimizer
```
## Initializing the module
If the go mod version in the project is an issue,delete the go.mod file and initialize it again as
```bash
go mod init tetris
```
## Running the program
```bash
go run . sample.txt
```
You are free to change the contents of the sample.txt file and re-run the program to place the tetrominoes on the board.

## Author
**[Lorna Akoth](https://github.com/akothouma)**

## Contribution
To contribute, make a pull request 
















































































































