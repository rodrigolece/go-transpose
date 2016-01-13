package main

import (
    "fmt"
    "github.com/gonum/matrix/mat64"
)

func test_tipo(A mat64.Matrix) {
    switch t := A.(type) {
    case mat64.Matrix:
        fmt.Println("interfaz: ", t)
    // impossible type switch case: A (type mat64.Matrix)
    // cannot have dynamic type mat64.Dense (missing At method)
    // case mat64.Dense:
    //     fmt.Println("tipo concreto: ", t)
    default:
        fmt.Println("otro")
    }
}

func main() {
    A := mat64.NewDense(2,3, []float64{0,1,0,0,2,3})

    test_tipo(A)
}
