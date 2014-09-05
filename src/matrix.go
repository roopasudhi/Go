package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type MatrixOp interface {
	Make([][]string) Matrix
	Inverse(Matrix) Matrix
	Multiply(Matrix, Matrix) Matrix
	String() string
}

type Matrix struct {
	elements   []Cell
	rows, cols int
}

type Cell struct {
	rows  int
	cols  int
	value float64
}

func NewMatrix(l int) *Matrix {
	return &Matrix{elements: make([]Cell, l)}
}
func (m Matrix) transpose() *Matrix {
	matB := NewMatrix(len(m.elements))
	for i := 0; i < len(m.elements); i++ {
		matB.elements[i] = Cell{m.elements[i].cols, m.elements[i].rows, m.elements[i].value}
	}
	matB.rows = m.rows
	matB.cols = m.cols
	return matB
}
func (m Matrix) get(r int, c int) float64 {
	for i := range m.elements {
		if m.elements[i].cols == c && m.elements[i].rows == r {
			return m.elements[i].value
		}
	}
	return 0
}
func (m *Matrix) print() {
	for i := 0; i <= m.rows; i++ {
		for j := 0; j <= m.cols; j++ {
			fmt.Print(m.get(i, j), " ")
		}
		fmt.Println()
	}
}
func (m Matrix) Multiply(b Matrix) *Matrix {
	result := NewMatrix(len(m.elements))
	result.rows = m.rows
	result.cols = m.cols
	for i := 0; i <= m.rows; i++ {
		sum := make([]float64, b.cols+1)

		for k := 0; k <= m.cols; k++ {
			for j := 0; j <= b.cols; j++ {

				sum[j] += m.get(i, k) * b.get(k, j)
			}
		}

		for j := 0; j <= b.cols; j++ {
			index := i*(m.cols+1) + j
			result.elements[index] = Cell{i, j, sum[j]}

		}

	}
	return result
}
func Build(record [][]string) *Matrix {
	l := len(record)
	var rows, cols int
	matrix := NewMatrix(l)

	for i := 0; i < l; i++ {
		r, _ := strconv.Atoi(record[i][0])
		c, _ := strconv.Atoi(record[i][1])
		v, _ := strconv.ParseFloat(record[i][2], 64)
		matrix.elements[i] = Cell{r, c, v}
		if c > cols {
			cols = c
		}
		if r > rows {
			rows = r
		}
	}
	matrix.rows = rows
	matrix.cols = cols
	return matrix
}

func main() {

	f, _ := os.Open("matin.csv")
	defer f.Close()
	csvreader := csv.NewReader(f)
	record, _ := csvreader.ReadAll()
	matA := Build(record)
	matA.print()
	matB := matA.transpose()
	matB.print()
	matC := matA.Multiply(*matB)
	matC.print()
}
