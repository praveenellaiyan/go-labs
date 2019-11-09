package vehicle

import (
	"fmt"
)

/*
 *
 */
type Car struct {
	numOfGears, currrentGear int
	currrentSpeed            float32
	isStarted                bool
}

/*
 *
 */
func NewCar(totalGears int) *Car {
	return &Car{
		numOfGears:    totalGears,
		currrentGear:  0,
		currrentSpeed: 0,
		isStarted:     false,
	}
}

/*
 *
 */
func (c *Car) Start() {
	c.isStarted = true
	fmt.Println("Car started!")
}

/*
 *
 */
func (c *Car) GearUp() {
	if c.currrentGear < c.numOfGears && c.isStarted {
		c.currrentGear++
		c.currrentSpeed += 20
		fmt.Println("Speed increased by 20kmph and Car is moving at ", c.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear up!")
	}
}

/*
 *
 */
func (c *Car) GearDown() {
	if c.currrentGear > 0 {
		c.currrentGear--
		c.currrentSpeed -= 20
		fmt.Println("Speed decreased by 20kmph and Car is moving at ", c.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear down!")
	}
}

func (c *Car) Stop() {
	c.currrentGear = 0
	c.currrentSpeed = 0
	c.isStarted = false
	fmt.Println("Car stopped!")
	fmt.Println("----------------------------")
}
