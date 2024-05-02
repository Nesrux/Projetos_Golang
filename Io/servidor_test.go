package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test_obterJogador(t *testing.T){
	t.Run("retorna resuldado de maria", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/jogadores/maria", nil)
		resp := httptest.NewRecorder()

		Servidor(resp, req)

		recebido := resp.Body.String()
		esperado := "20"

		if recebido != esperado{
			t.Errorf("recebido %s, esperado %s", recebido, esperado)
		}
	})
}