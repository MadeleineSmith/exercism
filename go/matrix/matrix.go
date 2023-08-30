package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	split := strings.Split(s, "\n")
	theMatrix := make(Matrix, len(split))

	rowLength := 0

	for i, row := range split {
		rowParts := strings.Fields(row)

		newRowLen := len(rowParts)
		if i > 0 && rowLength != newRowLen {
			return nil, errors.New("rows must be equal in length")
		}
		rowLength = newRowLen

		for j, num := range rowParts {
			if j == 0 {
				// could specify len here
				theMatrix[i] = []int{}
			}

			actualNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}

			theMatrix[i] = append(theMatrix[i], actualNum)
		}
	}

	return theMatrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
