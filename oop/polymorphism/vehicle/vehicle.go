package vehicle

/*
 * Vehicle interface that holds common behavior in real time
 */
type Vehicle interface {
	Start()
	GearUp()
	GearDown()
	Stop()
}
