package main

import (
	"errors"
	"fmt"
	"log"
)

func factorial(num int) (error, int) {
	if num >= 0 {
		if num == 0 {
			return nil, 1
		} else {
			_, n := factorial(num - 1)
			return nil, num * n
		}
	} else {
		return errors.New("Invalid natural number"), 0
	}
}

func main() {
	const num int = 6
	err, fact := factorial(num)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("El factorial de %d es %d\n", num, fact)
	}
}
