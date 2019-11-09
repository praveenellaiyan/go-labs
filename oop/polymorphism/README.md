# POLYMORPHISM IN GO PROGRAMMING

Many Forms.
Ability for different objects to respond differently to the same message is known as polymorphism.

## How it can be achieved in Go programming?
In Go programming, polymorphism is acheieved through interface.
<br>Unlike java here we don't need to explicitly tell the type is implementing an interface.
<br>If a type define implementation for all methods from an interface then it implicitly tells a type is implements respective interface.

1. Define an interface.
2. Define a type with implementation for all methods from an interface that you wanna implement.

## Programmatic demonstration

```go
type Vehicle interface {
	Start()
}

type Car struct {
	model	string
}

func (c *Car) Start() {
	fmt.Println(c.model, " started!")
}

type Bike struct {
	model	string
}

func (b *Bike) Start() {
	fmt.Println(b.model, " started!")
}
```