package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- fmt.Sprintf("%s No esta disponible", servidor)
	} else {
		canal <- fmt.Sprintf("%s Esta funcionando", servidor)
	}
}

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"https://www.udemy.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	fin := time.Since(inicio)
	fmt.Println("Tiempo de ejecucion:", fin)
}
