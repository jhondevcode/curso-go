package main

import (
	"errors"
	"fmt"
	"log"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Cannot be divided by 0")
	}
	return dividendo / divisor, nil
}

func main() {
	val, err := division(4, 2)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println(val)
	}
}
