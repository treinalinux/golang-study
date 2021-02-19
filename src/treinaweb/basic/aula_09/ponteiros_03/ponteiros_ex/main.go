package main

import (
	"fmt"
)

func main() {
	var inteiro = 0
	teste(&inteiro)
	fmt.Println(inteiro)
}

func teste(v *int) {
	*v = 33
}
