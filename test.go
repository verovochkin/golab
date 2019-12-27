package main

import "testing"

func Test_calculateSLE(t *testing.T) {
	var matrix = [][]float64{
		{1, 5, 3, -4, 20},
		{3, 1, -2, 0, 9},
		{5, -7, 0, 10, -9},
		{0, 3, -5, 0, 1}}

	var expectedResult = []float64{2.9999999999999996, 2, 0.9999999999999998, -1.0000000000000004}

	var actualResult = calculateSLE(matrix)

	checkArrayResult(actualResult, expectedResult, t)
}

func Test_formValuesArray(t *testing.T) {
	var matrix = [][]float64{
		{1, 3, -4, 20},
		{3, 1, 0, 9},
		{5, -7, 0, -9}}

	var expectedResult = []float64{20, 9, -9}

	var actualResult = formValuesArray(matrix)

	checkArrayResult(actualResult, expectedResult, t)
}

func Test_formMatrix(t *testing.T) {
	var matrix = [][]float64{
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4, 20}}

	var expectedResult = [][]float64{
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4},
		{1, 3, -4, 20,1, 3, -4, 20,1, 3, -4}}

	var actualResult = formMatrix(matrix)

	checkMatrixResult(actualResult, expectedResult, t)
}

func Test_getDeterminant(t *testing.T) {
	var matrix = [][]float64{
		{3, 8, 9, 4},
		{-5, 0, 5, 3},
		{-8, -2, 6, 7},
		{10, 8, -7, 6}}

	var expectedResult float64 = 1562

	var actualResult = getDeterminant(matrix)

	if expectedResult != actualResult {
		t.Error("Error")
	}
}

func Test_isMinus(t *testing.T) {
	var actualResult = isMinus(0, 1)
	if true != actualResult {
		t.Error("Error")
	}
	actualResult = isMinus(0, 2)
	if true == actualResult {
		t.Error("Error")
	}
}

func Test_calculateDeterminant(t *testing.T) {
	var matrix = [][]float64{
		{3, 8, 9, 4},
		{-5, 0, 5, 3},
		{-8, -2, 6, 7},
		{10, 8, -7, 6}}

	var expectedResult float64 = 781

	var actualResult = calculateDeterminant(0.5, matrix, false)

	if expectedResult != actualResult {
		t.Error("Error")
	}
}

func Test_getBuffMatrix(t *testing.T) {
	var matrix = [][]float64{
		{3, 8, 9, 4},
		{-5, 0, 5, 3},
		{-8, -2, 6, 7},
		{10, 8, -7, 6}}

	var expectedResult = [][]float64{
		{3, 9, 4},
		{-5, 5, 3},
		{10, -7, 6}}

	var actualResult = getBuffMatrix(matrix, 2, 1)

	checkMatrixResult(actualResult, expectedResult, t)
}

func Test_getDeterminantFrom2x2(t *testing.T) {
	var matrix = [][]float64{
		{2, 2},
		{1, 9}}

	var expectedResult float64 = 16

	var actualResult = getDeterminantFrom2x2(matrix)

	if expectedResult != actualResult {
		t.Error("Error")
	}
}

func Test_getRevertMatrix(t *testing.T) {
	var matrix = [][]float64{
		{3, 9, 4},
		{-5, 5, 3},
		{10, -7, 6}}

	var expectedResult = [][]float64{
		{3, -5, 10},
		{9, 5, -7},
		{4, 3, 6}}

	var actualResult = getRevertMatrix(matrix)

	checkMatrixResult(actualResult, expectedResult, t)
}

func Test_getTransposedMatrix(t *testing.T) {
	var matrix = [][]float64{
		{1, 2, 3},
		{-5, 0, 3},
		{10, 8, 6}}

	var expectedResult = [][]float64{
		{1, -0.5, -0.25},
		{-2.5, 1, 0.75},
		{1.6666666666666667, -0.5, -0.4166666666666667}}

	var actualResult = getTransposedMatrix(getRevertMatrix(matrix), getDeterminant(matrix))

	checkMatrixResult(actualResult, expectedResult, t)
}

func Test_getResultArray(t *testing.T) {
	var matrix = [][]float64{
		{1, 2, 3},
		{-5, 0, 3},
		{10, 8, 6}}

	var valuesArray = []float64{3, 3, 6}

	var expectedResult = []float64{27, 3, 90}

	var actualResult = getResultArray(matrix, valuesArray)

	checkArrayResult(actualResult, expectedResult, t)
}

func Test_printMatrix(t *testing.T) {
	var matrix = [][]float64{
		{1, 2},
		{3, 4}}
	var expectedResult = "1.000000\t2.000000\t\n3.000000\t4.000000\t\n"

	var actualResult = printMatrix(matrix)

	if expectedResult != actualResult {
		t.Error("Error")
	}
}

func checkMatrixResult(actualResult [][]float64, expectedResult [][]float64, t *testing.T) {
	if actualResult != nil {
		for i := 0; i < len(actualResult); i++ {
			for j := 0; j < len(actualResult[i]); j++ {
				if expectedResult[i][j] != actualResult[i][j] {
					t.Error("Error")
				}
			}
		}
	}
}

func checkArrayResult(actualResult []float64, expectedResult []float64, t *testing.T) {
	if actualResult != nil {
		for i := 0; i < len(actualResult); i++ {
			if expectedResult[i] != actualResult[i] {
				t.Error("Error")
			}
		}
	}
}
