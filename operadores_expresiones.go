package main

import "fmt"

func mainoperadores() {
	fmt.Println("== 2.2 Operadores y expresiones ==")

	a, b := 7, 3

	// Aritméticos
	fmt.Printf("a=%d, b=%d\n", a, b)
	fmt.Printf("a+b=%d, a-b=%d, a*b=%d, a/b=%d, a%%b=%d\n", a+b, a-b, a*b, a/b, a%b)

	// Asignación compuesta
	a += 5 // a = a + 5
	b *= 2 // b = b * 2
	fmt.Printf("a+=5 = %d, b*=2 = %d\n", a, b)

	// Comparación
	fmt.Printf("a==b:%t, a!=b:%t, a>b:%t, a<=b:%t\n", a == b, a != b, a > b, a <= b)

	// Lógicos
	x, y := true, false
	fmt.Printf("x&&y:%t, x||y:%t, !x:%t\n", x && y, x || y, !x)

	// Operadores sobre cadenas
	s1, s2 := "Hola", " Go"
	fmt.Println("concatenación:", s1+s2)

}
