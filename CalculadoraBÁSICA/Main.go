// Golang main

package main

import "fmt"

func main() {

	var num1 int
	var num2 int

	// Entrada de dados
	fmt.Print("Digite o primeiro número:")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número:")
	fmt.Scanln(&num2)

	// Somando
	somar := FuncSomar(num1, num2)
	fmt.Printf("Soma: %d\n", somar)

	// Subtraindo
	subtrair := FuncSubtrair(num1, num2)
	fmt.Printf("Subtração: %d\n", subtrair)

	// Multiplicar
	multiplicar := FuncMultiplicar(num1, num2)
	fmt.Printf("Multiplicação: %d\n", multiplicar)

	// Dividir
	dividir := FuncDividir(num1, num2)
	fmt.Printf("Divisão: %d\n", dividir)

}
