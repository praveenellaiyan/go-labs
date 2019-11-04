# ENCAPSULATION IN GO PROGRAMMING

Encapsulation is nothing but the binding up of fields and methods. In addition to that data hiding is leveraged through encapsulation meaning the fields were accessed only through the exported methods.

## How it can be achieved in Go programming?
1. Define struct with fields not exported [fields starts with lower case]
2. Define exported methods through which the fields were manipulated [methods starts with upper case]

## Programmatic demonstration

```go
type Student struct {
	id   int
	name string
}

func New() *Student {
	return &Student{
		id:   100,
		name: "John",
	}
}

func (s *Student) SetStudentID(id int) {
	s.id = id
}

func (s Student) GetStudentID() int {
	return s.id
}
```