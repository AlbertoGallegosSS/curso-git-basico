package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func duplicarPorValor(n int) {
	n *= 2
	fmt.Println("valor   =>", n)
}

func duplicarPorReferencia(n *int) {
	*n *= 2
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func cumplirAnios(p *Persona) {
	p.Edad++
}

func mainpunteros() {
	fmt.Println("Punteros ")

	x := 5
	px := &x // dirección de x

	fmt.Printf("x=%d, &x=%p, px=%p, *px=%d\n", x, &x, px, *px)

	duplicarPorValor(x)
	fmt.Println("afuera (valor)   =>", x)

	duplicarPorReferencia(&x)
	fmt.Println("afuera (puntero) =>", x)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println("swap => a, b =", a, b)

	p := Persona{"Mía", 19}
	cumplirAnios(&p)
	fmt.Printf("struct via puntero => %+v\n", p)
}
