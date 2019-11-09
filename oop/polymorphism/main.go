package main

import (
	"github.com/praveenellaiyan/go-labs/oop/polymorphism/vehicle"
)

func main() {

	alto := vehicle.NewCar(6)
	operate(alto)

	duke := vehicle.NewBike(4)
	operate(duke)
}

func operate(vehicle vehicle.Vehicle) {
	vehicle.Start()
	vehicle.GearUp()
	vehicle.GearDown()
	vehicle.Stop()
}
