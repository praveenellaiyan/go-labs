package vehicle

import (
	"fmt"
)

/*
 *
 */
type Bike struct {
	numOfGears, currrentGear int
	currrentSpeed            float32
	isStarted                bool
}

/*
 *
 */
func NewBike(totalGears int) *Bike {
	return &Bike{
		numOfGears:    totalGears,
		currrentGear:  0,
		currrentSpeed: 0,
		isStarted:     false,
	}
}

/*
 *
 */
func (b *Bike) Start() {
	b.isStarted = true
	fmt.Println("Bike started!")
}

/*
 *
 */
func (b *Bike) GearUp() {
	if b.currrentGear < b.numOfGears && b.isStarted {
		b.currrentGear++
		b.currrentSpeed += 10
		fmt.Println("Speed increased by 20kmph and Bike is moving at ", b.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear up!")
	}
}

/*
 *
 */
func (b *Bike) GearDown() {
	if b.currrentGear > 0 {
		b.currrentGear--
		b.currrentSpeed -= 10
		fmt.Println("Speed decreased by 20kmph and Bike is moving at ", b.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear down!")
	}
}

func (b *Bike) Stop() {
	b.currrentGear = 0
	b.currrentSpeed = 0
	b.isStarted = false
	fmt.Println("Bike stopped!")
	fmt.Println("----------------------------")
}
