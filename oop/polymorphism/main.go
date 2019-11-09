package main

import (
	"github.com/praveenellaiyan/go-labs/oop/polymorphism/vehicle"
)

func main() {

	// construct new car Alto K10 with max of 6 gears
	alto := vehicle.NewCar("Alto K10", 6)
	operate(alto)

	// construct new bike Duke V2 with max of 4 gears
	duke := vehicle.NewBike("Duke V2", 4)
	operate(duke)
}

/* function to operate vehicle that is received */
func operate(vehicle vehicle.Vehicle) {
	vehicle.Start()
	vehicle.GearUp()
	vehicle.GearDown()
	vehicle.Stop()
}
