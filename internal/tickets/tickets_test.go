package tickets_test

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	destino := "Brazil"
	resultado, err := tickets.GetTotalTickets(destino)
	expected := 45
	if err != nil || resultado != expected {
		fmt.Println(resultado)
		t.Errorf("El resultado calculado %d no coincide con el valor esperado %d", resultado, expected)
	}

	destino1 := "Finland"
	resultado, err = tickets.GetTotalTickets(destino1)
	expected = 8
	if err != nil || resultado != expected {
		fmt.Println(resultado)
		t.Errorf("El resultado calculado %d no coincide con el valor esperado %d", resultado, expected)
	}
}
