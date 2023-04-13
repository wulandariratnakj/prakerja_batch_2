package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) CalculateRange() float64 {
	return c.FuelIn * 1.5
}

func main() {
	car := Car{
		Type:   "SUV",
		FuelIn: 30.0,
	}

	rangeInKM := car.CalculateRange()
	fmt.Printf("Car type: %s, Fuel In: %.2f L\n", car.Type, car.FuelIn)
	fmt.Printf("Range: %.2f KM\n", rangeInKM)
}
