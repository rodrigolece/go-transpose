package main

import (
    "fmt"
    "github.com/gonum/matrix/mat64"
)

func transpose(A mat64.Matrix) mat64.Matrix{
    r, s := A.Dims()
    var data []float64

    for  j := 0 ; j < s ; j++ {
        for i := 0 ; i < r ; i++ {
            data = append(data, A.At(i,j))
        }
    }

    // Está medio chafa que regrese Dense, cómo hacemos para que regrese
    // el mismo tipo de A?
    return mat64.NewDense(s, r, data)
}

// cannot define new methods on non-local type mat64.Dense
// func (A *mat64.Dense) transpose() {
//     A.Copy(transpose(A))
// }


func main() {
    A := mat64.NewDense(2,3, []float64{0,1,0,0,2,3})
    printMat(A)

    printMat(transpose(A))
}


func printMat(A mat64.Matrix) {
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


/* Dos diferencias con Julia me molestan un poco:
 - No se pueden definir nuevos métodos en tipos importados (no-locales)
 - No sé como regresar una transpuesta que sea del mismo tipo que la matriz
 con la que se llama a transpose. Hace falta algo del tipo de similar()

 Vale la pena buscar el equivalente de show(), que aquí sería cómo extender
 los métodos de print para lo que me sirve a mí. (Valdría la pena asegurarme
 de que los prints son la única manera de interactuar con el mundo exterior.

 Fuera de las desventajas mencionadas, es muy fácil implementar el algoritmo
 de reescribir las columnas de A como renglones, porque basta llenar data en el
 orden correcto. */
