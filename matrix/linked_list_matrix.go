package matrix

import (
	"errors"

	"gcoll/list"
)

var (
	ERR_INDEX_OUT_OF_BOUND error = errors.New("Index out of bound")
)

type LinkedListMatrix struct {
	rows    *list.LinkedList
	numRows int
	numCols int
}

func (this *LinkedListMatrix) Get(x int, y int) any {
	if x+1 > this.numRows || x < 0 || y+1 > this.numCols || y < 0 {
		return nil
	}
	row := this.rows.Get(x).(*list.LinkedList)
	return row.Get(y)
}

func (this *LinkedListMatrix) Set(x int, y int, value any) error {
	if x+1 > this.numRows || x < 0 || y+1 > this.numCols || y < 0 {
		return ERR_INDEX_OUT_OF_BOUND
	}
	row := this.rows.Get(x).(*list.LinkedList)
	row.Set(y, value)
	return nil
}

func (this *LinkedListMatrix) Size() (int, int) {
	return this.numRows, this.numCols
}

func (this *LinkedListMatrix) Resize(numRows int, numCols int, initVal any) error {
	if numRows < 0 || numCols < 0 {
		return ERR_INDEX_OUT_OF_BOUND
	}
	// handle rows
	if numRows < this.numRows {
		// be careful that when resizing to a smaller matrix, some data will be lost
		for i := numRows; i < this.numRows; i++ {
			this.rows.RemoveAt(numRows)
		}
	} else if numRows > this.numRows {
		for i := 0; i < numRows-this.numRows; i++ {
			row := list.NewLinkedList()
			for j := 0; j < this.numCols; j++ {
				row.Add(initVal)
			}
			this.rows.Add(row)
		}
	}
	// handle columns
	for i := 0; i < numRows; i++ {
		row := this.rows.Get(i).(*list.LinkedList)
		if numCols < this.numCols {
			// be careful that when resizing to a smaller matrix, some data will be lost
			for j := numCols; j < this.numCols; j++ {
				row.RemoveAt(numCols)
			}
		} else if numCols > this.numCols {
			for j := 0; j < numCols-this.numCols; j++ {
				row.Add(initVal)
			}
		}
	}
	this.numRows = numRows
	this.numCols = numCols
	return nil
}

func (this *LinkedListMatrix) RemoveAt(index int) error {
	if index >= this.numRows {
		return ERR_INDEX_OUT_OF_BOUND
	}
	this.rows.RemoveAt(index)
	for i := 0; i < this.rows.Size(); i++ {
		row := this.rows.Get(i).(*list.LinkedList)
		row.RemoveAt(index)
	}
	return nil
}

func NewLinkedMatrix(rows int, columns int, initVal any) *LinkedListMatrix {
	matrix := &LinkedListMatrix{rows: list.NewLinkedList()}
	matrix.Resize(rows, columns, initVal)
	return matrix
}
