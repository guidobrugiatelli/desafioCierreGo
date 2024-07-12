package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	total, err := tickets.GetTotalTickets("Colombia")
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
	ticketsTurn := tickets.CountTicketsByTurn()
	fmt.Printf("Tickets por turno:\n")
	for turn, count := range ticketsTurn {
		fmt.Printf("%s: %d\n", turn, count)
	}
	percentage := tickets.AverageDestination("Colombia")
	fmt.Printf("Porcentaje de tickets con destino a Brazil: %.2f%%\n", percentage)
}
