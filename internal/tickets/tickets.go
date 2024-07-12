package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Ticket struct {
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	totalTickets := 0
	file, err := os.Open("/Users/gbrugiatelli/go/src/CierreBasesGo/desafio-go-bases/tickets.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return totalTickets, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return totalTickets, err
	}

	for _, record := range records {
		if record[3] == destination {
			totalTickets++
		}
	}
	return totalTickets, nil
}

// ejemplo 2
func CountTicketsByTurn() map[string]int {
	turnCounts := map[string]int{
		"Madrugada": 0,
		"Mañana":    0,
		"Tarde":     0,
		"Noche":     0,
	}

	file, err := os.Open("tickets.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return turnCounts
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return turnCounts
	}

	for _, record := range records {
		flightTimeHour, error1 := strconv.Atoi(record[4][:2]) // Se extraen los primeros 2 caracteres.
		if err != nil {
			fmt.Println("Error al convertir la hora:", error1)
			continue
		}

		switch {
		case flightTimeHour >= 0 && flightTimeHour <= 6:
			turnCounts["Madrugada"]++
		case flightTimeHour >= 7 && flightTimeHour <= 12:
			turnCounts["Mañana"]++
		case flightTimeHour >= 13 && flightTimeHour <= 19:
			turnCounts["Tarde"]++
		case flightTimeHour >= 20 && flightTimeHour <= 23:
			turnCounts["Noche"]++
		}
	}
	return turnCounts
}

// ejemplo 3
func AverageDestination(destination string) float64 {
	totalTickets := 0
	destinationCount := 0

	file, err := os.Open("tickets.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return 0.0
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return 0.0
	}

	for _, record := range records {
		totalTickets++

		if record[3] == destination {
			destinationCount++
		}
	}

	percentage := float64(destinationCount) / float64(totalTickets) * 100
	return percentage
}
