package main

import (
	"testing"
)

// Tests if addition has been successful
func TestMatrixAddSuccessful(t *testing.T) {
	matA := Matrix{{1, 1}}
	matB := Matrix{{1, 1}}
	matC, err := matA.Add(matB)

	if err != nil {
		t.Fatal("Although calculation should be correct, error appeared")
	}
	if matC == nil {
		t.Fatal("Calculation failed for some reason")
	}
}

// Test if matrix addition has failed as expected
func TestMatrixAddFailed(t *testing.T) {
	// Test ob Addition korrekt funktioniert
	matA := Matrix{{1, 1}, {2, 2}}
	matB := Matrix{{1, 1}}
	matC, err := matA.Add(matB)

	if err == nil {
		t.Fatal("Despite failed calculation, error hasn't been thrown")
	}
	if matC != nil {
		t.Fatal("Calculated matA + matB despite dimension error")
	}
}

// Tests if transposing an matrix is correct
func TestMatrixTranspose(t *testing.T) {
	matA := Matrix{{1, 2}, {3, 4}}
	matATransposed := Matrix{{1, 3}, {2, 4}}
	numRowsTransposedMatA, numColsTransposedMatA := len(matA[0]), len(matA)

	matA.Transpose()

	if matA == nil {
		t.Fatal("Matrix emptied")
	}

	if len(matA) != numRowsTransposedMatA || len(matA[0]) != numColsTransposedMatA {
		t.Fatal("Created matrix with wrong dimensions")
	}

	for i := 0; i < len(matA); i++ {
		for j := 0; j < len(matA[0]); j++ {
			if matA[i][j] != matATransposed[i][j] {
				t.Errorf("Transposing matrix failed")
			}
		}
	}
}

func TestMatrixDet(t *testing.T) {
	// 2x2 Ergebnis = 1
	// 3x3 Ergebnis = 1.283185306
}