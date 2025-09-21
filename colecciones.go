package main

import "fmt"

func main() {
	fmt.Println(" Arreglos, slices y mapas")

	//Arreglos (tamaño fijo, se copian por valor)
	arr := [4]int{1, 2, 3, 4}
	arrCopia := arr
	arrCopia[0] = 99
	fmt.Println("array      =", arr)
	fmt.Println("arrayCopia =", arrCopia)

	//Slices
	s := []int{10, 20, 30}
	fmt.Printf("slice=%v len=%d cap=%d\n", s, len(s), cap(s))
	s = append(s, 40, 50)
	fmt.Printf("append => %v len=%d cap=%d\n", s, len(s), cap(s))

	sub := s[1:3] // comparte memoria con s
	sub[0] = 999  // modifica s también
	fmt.Println("sub-slice =", sub, " | slice =", s)

	//make y copy
	a := make([]int, 3, 5)
	copy(a, s) // copia hasta len(a)
	fmt.Println("make/copy =>", a)

	//Mapas (clave-valor, referencia)
	m := map[string]int{"Ana": 20}
	m["Luis"] = 25
	edad, ok := m["Eva"] // ok=false si no existe
	fmt.Printf("map=%v | Eva:%d ok=%t\n", m, edad, ok)
	delete(m, "Ana")
	for nombre, e := range m {
		fmt.Printf("%s:%d ", nombre, e)
	}
	fmt.Println()
}
