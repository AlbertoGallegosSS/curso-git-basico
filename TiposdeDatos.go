package main

import "fmt"

// Struct
type Personatype struct {
	Nombre  string
	Edad    int
	Alumno  bool
	AlturaM float64
}

const (
	Pi              = 3.14159
	SegundosPorHora = 60 * 60
)

func main() {

	var entero int = 10
	real := 3.5
	texto := "Hola Go"
	var bandera bool
	r := 'G'
	by := byte(255)

	var ceroInt int
	var ceroStr string

	fmt.Println("Tipos de datos y variables")
	fmt.Printf("entero=%d, real=%.2f, texto=%q, bandera=%t, rune=%c, byte=%d\n",
		entero, real, texto, bandera, r, by)
	fmt.Printf("constantes  Pi=%.5f, SegundosPorHora=%d\n", Pi, SegundosPorHora)
	fmt.Printf("valores cero  int:%d, string:%q\n\n", ceroInt, ceroStr)

	//Tipos compuestos
	arreglo := [3]int{1, 2, 3}
	slice := []int{10, 20, 30, 40}
	mapa := map[string]int{"Ana": 20, "Luis": 25}
	p := Personatype{"Mía", 19, true, 1.67}

	fmt.Println("array:", arreglo)
	fmt.Println("slice:", slice)
	fmt.Println("map:", mapa)
	fmt.Printf("struct: %+v\n\n", p)

	// Conversión de tiOS
	var x int = 7
	y := float64(x)
	fmt.Printf("conversión int->float64: %d -> %.1f\n", x, y)
}
