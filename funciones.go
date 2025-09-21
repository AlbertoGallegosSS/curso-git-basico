package main

import "fmt"

// Función simple
func sumar(a, b int) int {
	return a + b
}

// Múltiples retornos
func dividirSegura(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Variádica
func sumaVariadica(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Retornos con nombre
func promedio(nums []float64) (prom float64) {
	if len(nums) == 0 {
		return 0
	}
	for _, n := range nums {
		prom += n
	}
	prom /= float64(len(nums))
	return
}

func mainfunciones() {
	fmt.Println("Funciones ")

	fmt.Println("sumar(5, 8) =", sumar(5, 8))

	if q, ok := dividirSegura(10, 2); ok {
		fmt.Printf("dividirSegura(10,2) = %.2f\n", q)
	}
	if _, ok := dividirSegura(5, 0); !ok {
		fmt.Println("dividirSegura(5,0): división por cero detectada")
	}

	fmt.Println("sumaVariadica(1,2,3,4) =", sumaVariadica(1, 2, 3, 4))

	notas := []float64{9.5, 8.0, 7.5, 10.0}
	fmt.Printf("promedio(%v) = %.2f\n", notas, promedio(notas))
}
