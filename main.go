
package main

import (
	"fmt"
	"math"
)

func f(x float64) float64  { return math.Pow(x-3, 2) }
func df(x float64) float64 { return 2 * (x - 3) }

// Implementa el algoritmo de descenso de gradiente
func decensoGradianes(x0, alpha float64, maxIterations int) (float64, float64) {
	x := x0
	for i := 0; i < maxIterations; i++ {
		x -= alpha * df(x)
	}
	return x, f(x)
}

func main() {
	var x0, alpha float64
	var maxIterations int
	// Entradas
	fmt.Print("Ingrese el punto inicial (x0): ")
	fmt.Scan(&x0)
	fmt.Print("Ingrese el tamaño de paso (alpha): ")
	fmt.Scan(&alpha)
	fmt.Print("Ingrese el número máximo de iteraciones: ")
	fmt.Scan(&maxIterations)

	// Ejecuta el algoritmo de descenso de gradiente
	xMin, fMin := decensoGradianes(x0, alpha, maxIterations)

	// Resultados
	fmt.Printf("Valor de x que minimiza la función: %.4f\n", xMin)
	fmt.Printf("Valor mínimo de la función en ese punto: %.4f\n", fMin)
}
