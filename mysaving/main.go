package main

import (
	"fmt"
	"math"
	"math/rand"
)

type SavingMethod struct {
	Name   string
	method func(salary float64, year int) float64
}

const (
	Salary      float64 = 4000
	ClassicRate float64 = 0.15

	ProgressiveRate     float64 = 0.05
	ProgressiveIncrease float64 = 0.005

	PokerMinRate float64 = 0.05
	PokerMaxRate float64 = 0.25
)

func main() {
	year := 5

	savingMethods := []SavingMethod{
		{
			Name: "Classic",
			method: func(salary float64, year int) float64 {
				return math.RoundToEven(salary * float64(year*12) * ClassicRate)
			},
		},
		{
			Name: "Poker",
			method: func(salary float64, year int) float64 {
				sum := 0.0
				for i := 0; i < year*12; i++ {
					rate := PokerMinRate + rand.Float64()*(PokerMaxRate-PokerMinRate)
					sum += salary * rate
				}

				return math.RoundToEven(sum)
			},
		},
		{
			Name: "Progressive",
			method: func(salary float64, year int) float64 {
				rate := ProgressiveRate
				sum := 0.0
				for i := 1; i < year*12; i++ {
					rate += ProgressiveIncrease
					sum += salary * rate
				}

				return math.RoundToEven(sum)
			},
		},
	}

	fmt.Println("Salary: ", Salary)
	fmt.Println("Year: ", year)
	fmt.Println("Result : ")

	for i := range savingMethods {
		fmt.Printf("- %s: %.2f€ \n", savingMethods[i].Name, savingMethods[i].method(Salary, year))
	}

}
