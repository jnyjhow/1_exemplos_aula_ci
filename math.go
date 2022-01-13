package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 5))

}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}

/* 
func Multiplicacao(a int, b int) int {
	return a * b
}

func Divisao(a int, b int) int {
	return a / b
}

func SomaX(a int, b int) int {
	return a + b + a
} */