package main

import "fmt"

//modelar un mensaje de chatbot
type Autor struct {
	ID    string
	EsBot bool
}

type Mensaje struct {
	Autor     Autor
	Texto     string
	Canal     string
	Prioridad int
}

// "Constructor" idiomático
func NuevoMensaje(autorID, texto, canal string, prioridad int, esBot bool) Mensaje {
	return Mensaje{
		Autor:     Autor{ID: autorID, EsBot: esBot},
		Texto:     texto,
		Canal:     canal,
		Prioridad: prioridad,
	}
}

func mainstructurs() {
	fmt.Println("Structs")

	// Literal con nombres de campos
	m1 := Mensaje{
		Autor:     Autor{ID: "u123", EsBot: false},
		Texto:     "Hola, ¿en qué te ayudo?",
		Canal:     "web",
		Prioridad: 1,
	}
	// Constructor
	m2 := NuevoMensaje("bot01", "Mensaje automático", "web", 2, true)

	// Cero-valor y luego asignación
	var m3 Mensaje
	m3.Texto = "vacío con valores por defecto"

	fmt.Printf("m1: %+v\nm2: %+v\nm3: %+v\n", m1, m2, m3)
}
