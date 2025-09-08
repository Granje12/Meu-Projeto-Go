package fibonacci

// Fibonacci retorna o n-ésimo número da sequência de Fibonacci.
// Fibonacci(0) = 0, Fibonacci(1) = 1
func Fibonacci(n int) int {
	if n < 0 {
		return -1 // valor inválido
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
