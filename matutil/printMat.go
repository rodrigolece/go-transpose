package matutil
// Herramientas para trabajar con matrices

import (
    "fmt"
    "github.com/gonum/matrix/mat64"
)

// PrintMat imprime una matriz con un formato entendible
func PrintMat(A mat64.Matrix) {
    r, s := A.Dims()

    fmt.Printf("(")
    for  i := 0 ; i < r ; i++ {
        fmt.Printf("\t")
        for j := 0 ; j < s ; j ++ {
            fmt.Printf("%0.1f\t", A.At(i,j))
        }
        if i != r-1 {
            fmt.Printf("\n")
        }
    }
    fmt.Printf(")\n")
}
