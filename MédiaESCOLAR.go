// Golang source

package main

import "fmt"

func main() {

	var nota1 float64
	var nota2 float64
	var nota3 float64
	var media float64

	// Recebendo dados
	fmt.Print("1° TRIMESTRE:")
	fmt.Scanln(&nota1)
	fmt.Print("2° TRIMESTRE:")
	fmt.Scanln(&nota2)
	fmt.Print("3° TRIMESTRE:")
	fmt.Scanln(&nota3)

	// Calculando dados
	media = (nota1 + nota2 + nota3) / 3

	// Mostrando dados
	if media >= 6.00 {
		fmt.Printf("Aprovado. Sua média foi %0.2f", media)
	} else {
		fmt.Printf("Reprovado. Sua média foi %0.2f", media)
	}

}
