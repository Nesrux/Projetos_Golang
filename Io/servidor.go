package main

import (
	"fmt"
	"net/http"
)

func Servidor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
