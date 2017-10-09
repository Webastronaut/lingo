package main

import (
	"fmt"
)

type Value float64
type Matrix [][]Value

const errResult Value  = -9999999999

// TODO wieder die ganzen Error returns einbauen --> best practice in Go
func dimensionError(method string) error {
	return fmt.Errorf("lingo: Dimension error occured with (one of) the parameter(s)" +
		" in method \"%s\"", method)
}

func (matA Matrix) isNil() {
	if matA == nil {
		panic("Matrix is empty")
	}
}

// Print pretty prints a given matrix
func (matA Matrix) Print() {
	matA.isNil()

	for i := 0; i < len(matA); i++ {
		fmt.Print("|")

		for j := 0; j < len(matA[0]); j++ {
			fmt.Print(" ", matA[i][j], " ")
		}

		fmt.Println("|")
	}

	fmt.Println("")
}

// Add adds a given matrix matB to a baked in matrix matA
func (matA Matrix) Add(matB Matrix) (Matrix, error) {
	// Allocate space for result matrix
	var matC Matrix = make(Matrix, len(matA))

	matB.isNil()

	// Dimensions of matA and matB must fit
	if len(matA) == len(matB) && len(matA[0]) == len(matB[0]) {

		for i := 0; i < len(matA); i++ {
			// Create new columns in current row
			matC[i] = make([]Value, len(matA[0]))
			for j := 0; j < len(matA[0]); j++ {
				matC[i][j] = matA[i][j] + matB[i][j]
			}
		}

		return matC, nil

	} else {

		var matC Matrix = Matrix{}
		return matC, dimensionError("Add")

	}
}

// TODO Check if Add and Sub can be combined
func (matA Matrix) Sub(matB Matrix) Matrix {
	// Allocate space for result matrix
	var matC Matrix = make(Matrix, len(matA))

	matB.isNil()

	// Dimensions of matA and matB must fit
	if len(matA) == len(matB) && len(matA[0]) == len(matB[0]) {

		for i := 0; i < len(matA); i++ {
			// Create new columns in current row
			matC[i] = make([]Value, len(matA[0]))
			for j := 0; j < len(matA[0]); j++ {
				matC[i][j] = matA[i][j] - matB[i][j]
			}
		}

	} else {

		panic("Dimension error")

	}

	return matC
}

// Transpose transposes any given matrix
func (matA *Matrix) Transpose() {
	matA.isNil()

	// Allocate space for transposed matrix
	matATransposed := make(Matrix, len((*matA)[0]))

	// Create correct number of columns in each row of transposed matrix
	for i := 0; i < len((*matA)[0]); i++ {
		matATransposed[i] = make([]Value, len(*matA))
	}

	// Fill transposed matrix with values of matA
	for i := 0; i < len(*matA); i++ {
		for j := 0; j < len((*matA)[0]); j++ {

			matATransposed[j][i] = (*matA)[i][j]

		}
	}

	*matA = matATransposed
}

func (matA Matrix) IsQuadratic() bool {
	matA.isNil()

	if len(matA) == 0 || len(matA) / len(matA[0]) == 1 {
		return true
	} else {
		return false
	}
}

func (matA *Matrix) ScalarMult(s Value) error {
	matA.isNil()

	if len(*matA) == 0 {
		return dimensionError("ScalarMult")
	}

	for i := 0; i < len(*matA); i++ {
		for j := 0; j < len((*matA)[0]); j++ {
			(*matA)[i][j] *= s
		}
	}

	return nil
}

//TODO Testing
func (matA Matrix) IsTriangle() (bool, error) {
	matA.isNil()

	if !matA.IsQuadratic() || len(matA) == 0 {
		return false, dimensionError("IsTriangle")
	}

	isTriangle := false

	for i := 1; i < len(matA); i++ {
		for j := 0; j < len(matA); j++ {
			if j < i {
				if matA[i][j] == 0 {
					isTriangle = true
				} else {
					isTriangle = false
					return isTriangle, nil
				}
			}
		}
	}

	return isTriangle, nil
}

// https://math.stackexchange.com/questions/383219/divide-and-conquer-matrices-to-calculate-determinant#383249
// TODO Testing
// TODO: Divide&Conquer, see: https://math.stackexchange.com/questions/383219/divide-and-conquer-matrices-to-calculate-determinant#383249
func (matA Matrix) Det() (Value, error) {
	var result Value = errResult

	if _, err := matA.IsTriangle(); err != nil {
		return result, dimensionError("Det")
	}

	// Fehler
	if isTriangle, err := matA.IsTriangle(); err == nil {
		if isTriangle {
			for i := 0; i < len(matA); i++ {
				if i == 0 {
					result = matA[i][i]
				} else {
					result *= matA[i][i]
				}
			}
			fmt.Print("Matrix is for some reason accepted as triangle")

			return result, nil
		}
	} else {
		return result, dimensionError("Det")
	}

	switch len(matA) {
	case 1:
		result = matA[0][0]
	case 2:
		result = matA[0][0] * matA[1][1] - matA[0][1] * matA[1][0]
	default:
		// Dogson Condensation --> D&C determinant calculation
		matA.Print()
		var matA11, matA12, matA21, matA22, matAMiddle Matrix =
			(matA[0:len(matA) - 1])[0:len(matA[0]) - 1],
			matA[0:len(matA) - 1][1:len(matA[0])],
			matA[1:][0:len(matA[0]) - 1],
		matA,
			//matA[1:][1:len(matA[0])],
			matA[1:len(matA) - 1][1:len(matA[0]) - 1]

		matA11.Print()
		matA12.Print()
		matA21.Print()
		matA22.Print()
		matAMiddle.Print()
		//matA11, matA12, matA21, matA22, matAMiddle := matA.sliceMatDet()

		matA11Det, err := matA11.Det()
		if err != nil { return result, err }

		matA12Det, err := matA12.Det()
		if err != nil { return result, err }

		matA21Det, err := matA21.Det()
		if err != nil { return result, err }

		matA22Det, err := matA22.Det()
		if err != nil { return result, err }

		matAMiddleDet, err := matAMiddle.Det()
		if err != nil { return result, err }

		result = (matA11Det * matA22Det - matA12Det * matA21Det) / matAMiddleDet
	}

	return result, nil
}

// Calculates the trace of a matrix
func (matA Matrix) Trace() (Value, error) {
	matA.isNil()

	var result Value = errResult

	if !matA.IsQuadratic() || len(matA) == 0 {
		return result, dimensionError("Trace")
	} else {
		result = 0
	}

	for i := 0; i < len(matA); i++ {
		result += matA[i][i]
	}

	return result, nil
}

func main() {
	fmt.Println("+++++++++++++ + Addition ++++++++++++")
	matA := Matrix{{1, 2}, {3, 4}, {5, 6}}
	matB := Matrix{{1, 1}, {1, 1}, {1, 1}}
	//matB := Matrix{}

	matC, err := matA.Add(matB)
	if err == nil {
		matC.Print()
	} else {
		fmt.Print(err)
	}

	fmt.Println("+++++++++++++ = Addition ++++++++++++\n")

	fmt.Println("+++++++++++++ + Transpose ++++++++++++")
	matA.Print()


	matA.Transpose()
	matA.Print()
	matA.Transpose()
	matA.Print()

	fmt.Println("+++++++++++++ = Transpose ++++++++++++\n")

	/*fmt.Println("+++++++++++++ + Determinant ++++++++++++")
	matD := Matrix{{2, 5}, {1, 3}}
	detMatD := matD.Det()
	fmt.Println(detMatD)

	//matE := Matrix{{2, 5, 4}, {1, 3.141592653, 7}, {2, 5, 5}}
	//detMatE := matE.Det()
	//fmt.Println(detMatE)

	fmt.Println("+++++++++++++ = Determinant ++++++++++++\n")*/

	fmt.Println("+++++++++++++ + IsTriangle ++++++++++++\n")

	matF := Matrix{{1, 2}, {0, 3}}

	matF.Print()
	res, err := matF.IsTriangle()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(res)

	fmt.Println("+++++++++++++ = IsTriangle ++++++++++++\n")

	fmt.Println("+++++++++++++ + ScalarMult ++++++++++++\n")

	matG := Matrix{{1, 2}, {0, 3}}
	err = matG.ScalarMult(2)
	if err != nil {
		fmt.Print(err)
		return
	}
	matG.Print()

	fmt.Println("+++++++++++++ = ScalarMult ++++++++++++\n")

	fmt.Println("+++++++++++++ + Trace ++++++++++++\n")

	matH := Matrix{
		{1, 2, 3},
		{0, 1, 7},
		{0, 0, 1}}

	fmt.Println(matH.Trace())

	fmt.Println("+++++++++++++ = Trace ++++++++++++\n")

	fmt.Println("+++++++++++++ + Det ++++++++++++\n")
	matJ := Matrix{
		{1333, 42, 3, 432},
		{0, 56, 7, 8},
		{3, 0, 11, 12},
		{234.23423, 0, 0, 16}}

	fmt.Println(matJ.Det())

	fmt.Println("+++++++++++++ = Det ++++++++++++\n")
}