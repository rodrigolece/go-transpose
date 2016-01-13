/* La idea que tenía para la implemetación era tener dos funciones. Si A es una
matriz mat64.Matrix:
- transpose(A) que no modifica a A
- A.transpose() que modifica a A

Como no puede definir métodos sobre tipos importados, definimos dense que tiene
un tipo embedded *mat64.Dense. Usamos Dense y no Matrix, porque Dense tiene
definido el tipo Copy que modifica a su argumento: A.Copy(B) copia el contenido
de B dentro de A si A es de tipo Dense.

*/

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

func transpose(A *dense) *dense{
    // Esta sería la parte padre, como ya sabemos invertir una matriz Dense
    // la transponemos con la función que definimos arriba
    return &dense{transpose(mat64.Dense(A))}
    // El problema es que Go no sabe convertir A de tipo dense a mat64.Dense    
}

func (A *dense) transpose() {
    A.Copy(transpose(A))
}


func main() {
    A := mat64.NewDense(2,3, []float64{0,1,0,0,2,3})
    matutil.PrintMat(transpose(A)) // A no se modifica

    B := dense{A}
    matutil.PrintMat(B) // imprime la misma matriz que con A

    B.transpose() // B se modifica
    matutil.PrintMat(B) // Ahora imprime la transpuesta
}
