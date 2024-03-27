package main

import (
	"fmt"
	"math"
	"math/rand"
)

type SavingInterface interface {
	CalculateSaving() float64
	Name() string
}

type Simulation struct {
	Salary float64
	Year   int
}

type SavingMethod struct {
	name       string
	simulation Simulation
}

func NewSavingMethod(name string, simulation Simulation) SavingMethod {
	return SavingMethod{
		name:       name,
		simulation: simulation,
	}
}

func (m SavingMethod) Name() string {
	return m.name
}

type ClassicalSavingMethod struct {
	SavingMethod
	Rate float64
}

type ProgressiveSavingMethod struct {
	SavingMethod
	Rate float64
}

type PokerSavingMethod struct {
	SavingMethod
	MinRate float64
	MaxRate float64
}

func (m ClassicalSavingMethod) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		sum += m.simulation.Salary * m.Rate
	}

	return sum
}

func (m PokerSavingMethod) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		sum += math.RoundToEven(m.MinRate + rand.Float64()*(m.MaxRate-m.MinRate)*m.simulation.Salary)
	}

	return sum
}

func (m *ProgressiveSavingMethod) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		m.Rate += ProgressiveIncrease

		sum += math.RoundToEven(m.simulation.Salary * m.Rate)
	}

	return sum
}

const (
	Salary float64 = 4000

	ClassicRate float64 = 0.15

	ProgressiveRate     float64 = 0.05
	ProgressiveIncrease float64 = 0.005

	PokerMinRate float64 = 0.05
	PokerMaxRate float64 = 0.25
)

func main() {
	simulation := Simulation{
		Salary: Salary,
		Year:   5,
	}

	savingMethods := []SavingInterface{
		ClassicalSavingMethod{
			SavingMethod: NewSavingMethod("Classic", simulation),
			Rate:         ClassicRate,
		},
		PokerSavingMethod{
			SavingMethod: NewSavingMethod("Poker", simulation),
			MinRate:      PokerMinRate,
			MaxRate:      PokerMaxRate,
		},
		&ProgressiveSavingMethod{
			SavingMethod: NewSavingMethod("Progressive", simulation),
			Rate:         ProgressiveRate,
		},
	}

	fmt.Println("Salary: ", simulation.Salary)
	fmt.Println("Year: ", simulation.Year)
	fmt.Println("Result : ")

	for i := range savingMethods {
		fmt.Printf("- %s: %.2f€ \n", savingMethods[i].Name(), savingMethods[i].CalculateSaving())
	}
}
