package main

import (
	"log"
	"net/http"
)

func main() {
	servidor := &ServidorJogador{}

	if err := http.ListenAndServe(":8080", servidor); err != nil {
		log.Fatalf("não foi possivel escutar a porta 8080 %v", err)
	}
}
