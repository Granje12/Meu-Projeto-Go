package main

import (
	"fmt"

	"github.com/Granje12/Meu-Projeto-Go/internal/fibonacci"
	"github.com/Granje12/Meu-Projeto-Go/internal/hello"
)

func main() {
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamar o hello
	hello.SayHello()

	// Chamar a função Fibonacci
	n := 10
	result := fibonacci.Fibonacci(n)
	fmt.Printf("O %dº número da sequência de Fibonacci é: %d\n", n, result)
}
