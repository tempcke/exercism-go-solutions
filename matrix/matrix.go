package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix rows and columns of ints
type Matrix [][]int

func (m Matrix) width() (w int) {
	if m.height() > 0 {
		w = len(m[0])
	}
	return w
}

func (m Matrix) height() int {
	return len(m)
}

// New Matrix
func New(s string) (Matrix, error) {
	// split rows from newline seperated list
	rowStrings := strings.Split(s, "\n")
	m := make([][]int, len(rowStrings))
	var colLen int

	for rowNum, rowString := range rowStrings {
		// split cols from space seperated list
		colStrings := strings.Fields(rowString)

		// ensure each row contains the same number of columns
		cl := len(colStrings)
		if rowNum == 0 {
			colLen = cl
		}
		if cl != colLen {
			return m, errors.New("Error: Uneven rows")
		}

		// build list of ints for row columns
		m[rowNum] = make([]int, cl)
		for colNum, colString := range colStrings {
			n, err := strconv.Atoi(colString)
			if err != nil {
				return m, err
			}
			m[rowNum][colNum] = n
		}
	}
	return m, nil
}

// Rows of matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, m.height())
	for i, row := range m {
		rows[i] = append([]int{}, row...)
	}
	return rows
}

// Cols of matrix
func (m Matrix) Cols() [][]int {
	colCount := m.width()
	cols := make([][]int, colCount)
	for colIndex := 0; colIndex < colCount; colIndex++ {
		cols[colIndex] = make([]int, m.height())
		for i := 0; i < m.height(); i++ {
			cols[colIndex][i] = m[i][colIndex]
		}
	}
	return cols
}

// Set matrix value at index
func (m Matrix) Set(row, col, val int) bool {
	// out of range?
	if row < 0 || col < 0 || row >= m.height() || col >= m.width() {
		return false
	}

	// set value
	m[row][col] = val
	return true
}
