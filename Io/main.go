package main

import (
	"log"
	"net/http"
)

type ArmazenamentoJogadorEmMemoria struct {}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoJogador(nome string) int {
	return 123
}

func main() {
	servidor := &ServidorJogador{&ArmazenamentoJogadorEmMemoria{}}

	if err := http.ListenAndServe(":8080", servidor); err != nil {
		log.Fatalf("não foi possivel escutar a porta 8080 %v", err)
	}
}
