package main

import "fmt"

func main () {

var estudiante  string
fmt.Print("Escriba el nombre del estudiante: ")
fmt.Scan(&estudiante)

var nota1, nota2, nota3 float64 
var notafinal float64
fmt.Print("Esriba la nota 1: ")
fmt.Scan(&nota1)

fmt.Print("Esriba la nota 2: ")
fmt.Scan(&nota2)

fmt.Print("Esriba la nota 3: ")
fmt.Scan(&nota3)

notafinal =((nota1 + nota2 + nota3)/3)

fmt.Println("su nota final es: ", notafinal)

if (notafinal >= 3.0) {
fmt.Println("el estudiante", estudiante,"aprueba")
}else{
	fmt.Println("el estudiante", estudiante, "no aprueba")
}
}