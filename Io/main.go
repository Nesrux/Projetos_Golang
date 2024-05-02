package main

import (
	"log"
	"net/http"
)

func main(){
	tratador := http.HandlerFunc(Servidor)
	if err := http.ListenAndServe(":8080", tratador);
	err != nil{
		log.Fatalf("não foi possivel escutar a porta 8080 %v", err)
	}
}