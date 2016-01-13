package main

import (
    "github.com/gonum/matrix/mat64"
    "github.com/rodrigolece/transpose/matutil"
)

type dense struct {
    *mat64.Dense
}

func transpose(A *mat64.Dense) *mat64.Dense {
    r, s := A.Dims()
    var data []float64

    for  j := 0 ; j < s ; j++ {
        for i := 0 ; i < r ; i++ {
            data = append(data, A.At(i,j))
        }
    }

    return mat64.NewDense(s, r, data)
}

func (A *dense) transpose() {
    // A.Copy(transpose(A))
    r, s := A.Dims()
    var data []float64

    for  j := 0 ; j < s ; j++ {
        for i := 0 ; i < r ; i++ {
            data = append(data, A.At(i,j))
        }
    }

    *A = dense{mat64.NewDense(s, r, data)}
}


func main() {
    A := mat64.NewDense(2,3, []float64{0,1,0,0,2,3})
    // matutil.PrintMat(A)

    B := dense{A} // no es necesario &dense{A}
    matutil.PrintMat(B)

    B.transpose()
    matutil.PrintMat(B)
}
