package controllers

import (
	"api/src/modelos"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	caseTest := []struct {
		UserCase modelos.Usuarios
		Expected string
	}{
		{UserCase: modelos.Usuarios{Id: 0, Nome: "deivid", Nick: "Kob", Email: "de@gmail.com", Senha: "123"}, Expected: "deivid"},
	}

	srv := httptest.NewServer(http.HandlerFunc(CriarUsuario))
	defer srv.Close()

	for _, tc := range caseTest {
		path := fmt.Sprintf("%s/usuarios", srv.URL)

		userJSON, err := json.Marshal(tc.UserCase)
		if err != nil {
			t.Fatalf("Erro ao serializar usuário: %v", err)
		}

		req, err := http.NewRequest("POST", path, bytes.NewBuffer(userJSON))
		if err != nil {
			t.Fatalf("Error in created request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Error in execute request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected %v, got %v", http.StatusOK, resp.StatusCode)
		}

		var user modelos.Usuarios
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			t.Fatalf("Erro ao decodificar resposta: %v", err)
		}

		if user.Nome != tc.Expected {
			t.Errorf("Expected %v, got %v", tc.Expected, user.Nome)
		}
	}
}
