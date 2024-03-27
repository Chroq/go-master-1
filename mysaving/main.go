package main

import (
	"fmt"

	"github.com/Chroq/mysaving/strategy"
)

const (
	Salary float64 = 4000

	ClassicRate float64 = 0.15
)

func main() {
	simulation := strategy.Simulation{
		Salary: Salary,
		Year:   5,
	}

	savingMethods := []strategy.SavingStrategy{
		strategy.ClassicalStrategy{
			Strategy: strategy.NewStrategy("Classic", simulation),
			Rate:     ClassicRate,
		},
		strategy.PokerStrategy{
			Strategy: strategy.NewStrategy("Poker", simulation),
			MinRate:  strategy.PokerMinRate,
			MaxRate:  strategy.PokerMaxRate,
		},
		&strategy.ProgressiveStrategy{
			Strategy: strategy.NewStrategy("Progressive", simulation),
			Rate:     strategy.ProgressiveRate,
		},
	}

	fmt.Println("Salary: ", simulation.Salary)
	fmt.Println("Year: ", simulation.Year)
	fmt.Println("Result : ")

	for i := range savingMethods {
		fmt.Printf("- %s: %.2f€ \n", savingMethods[i].Name(), savingMethods[i].CalculateSaving())
	}
}
