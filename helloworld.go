/*package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")

}*/

//Esta es la sintaxis normal o comun
/*package main		En Go, cada programa forma parte de un paquete. Lo definimos con la packagepalabra clave. En este ejemplo, el programa pertenece al mainpaquete.
import ("fmt")		import ("fmt") nos permite importar archivos incluidos en el fmtpaquete.
					Una línea en blanco. Go ignora los espacios en blanco. Tener espacios en blanco en el código facilita su lectura.
func main()			 es una función. Se ejecutará func main() {}cualquier código entre sus llaves .{}
{
  fmt.Println("Hello World!")	fmt.Println() es una función disponible en el fmtpaquete. Se utiliza para imprimir texto. En nuestro ejemplo, mostrará "¡Hola mundo!".
}*/

//Codigo compactado

//package main; import ("fmt"); func main() { fmt.Println("Hello World!");}

//Variables
//var variablename type = value //Primer forma

//variablename := value //Segunda forma

/*package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
} */

package main

import (
	"fmt"
)

func mainhola() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
