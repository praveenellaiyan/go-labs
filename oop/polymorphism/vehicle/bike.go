package vehicle

import (
	"fmt"
)

/*
 * type bike with their related props
 */
type Bike struct {
	model                    string
	numOfGears, currrentGear int
	currrentSpeed            float32
	isStarted                bool
}

/*
 * instantiate bike which is ready to be operated
 */
func NewBike(bikeModel string, totalGears int) *Bike {
	return &Bike{
		model:         bikeModel,
		numOfGears:    totalGears,
		currrentGear:  0,
		currrentSpeed: 0,
		isStarted:     false,
	}
}

/*
 * power on the bike
 */
func (b *Bike) Start() {
	b.isStarted = true
	fmt.Println(b.model, " started!")
}

/*
 * switch gear to one level higher and increase speed by 20kmph
 */
func (b *Bike) GearUp() {
	if b.currrentGear < b.numOfGears && b.isStarted {
		b.currrentGear++
		b.currrentSpeed += 10
		fmt.Println("Speed increased by 20kmph and ", b.model, " is moving at ", b.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear up!")
	}
}

/*
 * switch gear to one level lower and decrease speed by 20kmph
 */
func (b *Bike) GearDown() {
	if b.currrentGear > 0 {
		b.currrentGear--
		b.currrentSpeed -= 10
		fmt.Println("Speed decreased by 20kmph and ", b.model, " is moving at ", b.currrentSpeed, "kmph!")
	} else {
		fmt.Println("You can't gear down!")
	}
}

/*
 * stop the bike
 */
func (b *Bike) Stop() {
	b.currrentGear = 0
	b.currrentSpeed = 0
	b.isStarted = false
	fmt.Println(b.model, " stopped!")
	fmt.Println("----------------------------")
}
