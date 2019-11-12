# COMPOSITION IN GO PROGRAMMING
Go is not a purely object oriented programming.
</br>Still in place of inheritance composition can implemented which depicts <b>has-a</b> relationship.

## How it can be achieved in Go programming?
In Go programming, Composition is acheieved through embedded struct.
<br>[i.e] Nesting one struct into another one.

Real world example:
1. car has engine, wheels, seats, wiper etc
2. mobile has sensors, display etc

## Programmatic demonstration

```go
// address of the employee
type Address struct {
	city string
	zip  int
}

type Employee struct {
	id      int
	name    string
	dept    string
	address Address
}
```