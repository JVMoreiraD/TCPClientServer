package matrix

type Matrix [2][2]int

type Data struct {
	Matrix1 Matrix
	Matrix2 Matrix
}

func Multiply(matrix1 Matrix, matrix2 Matrix) Matrix {
	var result Matrix

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = 0
			for k := 0; k < 2; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}
