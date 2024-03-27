package strategy

import (
	"math"
	"math/rand"
)

const (
	ProgressiveRate     float64 = 0.05
	ProgressiveIncrease float64 = 0.005

	PokerMinRate float64 = 0.05
	PokerMaxRate float64 = 0.25
)

type SavingStrategy interface {
	CalculateSaving() float64
	Name() string
}

type Simulation struct {
	Salary float64
	Year   int
}

type Strategy struct {
	name       string
	simulation Simulation
}

func NewStrategy(name string, simulation Simulation) Strategy {
	return Strategy{
		name:       name,
		simulation: simulation,
	}
}

func (m Strategy) Name() string {
	return m.name
}

type ClassicalStrategy struct {
	Strategy
	Rate float64
}

type ProgressiveStrategy struct {
	Strategy
	Rate float64
}

type PokerStrategy struct {
	Strategy
	MinRate float64
	MaxRate float64
}

func (m ClassicalStrategy) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		sum += m.simulation.Salary * m.Rate
	}

	return sum
}

func (m PokerStrategy) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		sum += math.RoundToEven(m.MinRate + rand.Float64()*(m.MaxRate-m.MinRate)*m.simulation.Salary)
	}

	return sum
}

func (m *ProgressiveStrategy) CalculateSaving() float64 {
	sum := 0.0
	for i := 0; i < m.simulation.Year*12; i++ {
		m.Rate += ProgressiveIncrease

		sum += math.RoundToEven(m.simulation.Salary * m.Rate)
	}

	return sum
}
