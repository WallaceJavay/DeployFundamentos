// Golang source

package main

import "fmt"

func main() {

	var peso float64
	var altura float64

	fmt.Print("Digite seu peso (ex: 60.5):")
	fmt.Scanln(&peso)
	fmt.Print("Digite sua altura (ex: 1.70):")
	fmt.Scanln(&altura)

	var imc float64 = peso / (altura * altura)

	fmt.Printf("IMC: %f", imc)

	if imc < 18 {
		fmt.Println("\nAbaixo do peso.")
	} else if imc < 25 {
		fmt.Println("\nPeso normal.")
	} else if imc < 30 {
		fmt.Println("\nSobrepeso.")
	} else if imc < 40 {
		fmt.Println("\nObesidade.")
	}

}
