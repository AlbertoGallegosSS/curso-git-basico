package main

import "fmt"

func mainflujo() {
	fmt.Println("Control de flujo")

	// if / else corto
	if n := 12; n%2 == 0 {
		fmt.Printf("%d es par\n", n)
	} else {
		fmt.Printf("%d es impar\n", n)
	}

	// switch con expresión
	dia := 3
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	default:
		fmt.Println("Otro día")
	}

	// switch sin expresión
	nota := 87
	switch {
	case nota >= 90:
		fmt.Println("A")
	case nota >= 80:
		fmt.Println("B")
	case nota >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D/F")
	}

	// for clásico
	fmt.Print("for clásico: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for estilo while
	fmt.Print("while (for) sumando: ")
	suma := 0
	for suma < 10 {
		suma += 3
		fmt.Printf("%d ", suma)
	}
	fmt.Println()

	// for range en slice
	slice := []int{10, 20, 30}
	fmt.Print("range slice: ")
	for i, v := range slice {
		fmt.Printf("[i=%d v=%d] ", i, v)
	}
	fmt.Println()

	// for range en map
	mapa := map[string]int{"Ana": 20, "Luis": 25}
	fmt.Print("range map: ")
	for k, v := range mapa {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()

	// break / continue
	fmt.Print("continue solo pares: ")
	for i := 0; i < 7; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
