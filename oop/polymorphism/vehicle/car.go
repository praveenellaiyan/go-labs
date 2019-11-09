package vehicle

import (
	"fmt"
)

/*
 * type car with their related props
 */
type Car struct {
	model                    string
	numOfGears, currrentGear int
	currrentSpeed            float32
	isStarted                bool
}

/*
 * instantiate car which is ready to be operated
 */
func NewCar(carModel string, totalGears int) *Car {
	return &Car{
		model:         carModel,
		numOfGears:    totalGears,
		currrentGear:  0,
		currrentSpeed: 0,
		isStarted:     false,
	}
}

/*
 * power on the car
 */
func (c *Car) Start() {
	c.isStarted = true
	fmt.Println(c.model, " started!")
}

/*
 * switch gear to one level higher and increase speed by 20kmph
 */
func (c *Car) GearUp() {
	if c.currrentGear < c.numOfGears && c.isStarted {
		c.currrentGear++
		c.currrentSpeed += 20
		fmt.Println("Speed increased by 20kmph and ", c.model, " is moving at ", c.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear up!")
	}
}

/*
 * switch gear to one level lower and decrease speed by 20kmph
 */
func (c *Car) GearDown() {
	if c.currrentGear > 0 {
		c.currrentGear--
		c.currrentSpeed -= 20
		fmt.Println("Speed decreased by 20kmph and ", c.model, " is moving at ", c.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear down!")
	}
}

/*
 * stop the car
 */
func (c *Car) Stop() {
	c.currrentGear = 0
	c.currrentSpeed = 0
	c.isStarted = false
	fmt.Println(c.model, " stopped!")
	fmt.Println("----------------------------")
}
