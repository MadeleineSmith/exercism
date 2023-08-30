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
	new3dSlice := make([][]int, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			new3dSlice[i] = append(new3dSlice[i], m[j][i])
		}
	}

	return new3dSlice
}

func (m Matrix) Rows() [][]int {
	new3dSlice := make([][]int, len(m))

	for i, row := range m {
		for j, num := range row {
			if j == 0 {
				// could specify len here
				new3dSlice[i] = []int{}
			}

			new3dSlice[i] = append(new3dSlice[i], num)
		}
	}

	return new3dSlice
}

func (m Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
