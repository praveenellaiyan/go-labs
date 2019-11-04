package student

/*
 * Define struct which holds properties of student
 */
type Student struct {
	id   int
	name string
}

/*
 * exported method that initializes student type
 */
func New() *Student {
	return &Student{
		id:   100,
		name: "John",
	}
}

/*
 * update value of student id using exported setter function
 */
func (s *Student) SetStudentID(id int) {
	s.id = id
}

/*
 * update value of student name using exported setter function
 */
func (s *Student) SetStudentName(name string) {
	s.name = name
}

/*
 * retrieve value of student name using exported getter function
 */
func (s Student) GetStudentID() int {
	return s.id
}

/*
 * update value of student name using exported getter function
 */
func (s Student) GetStudentName() string {
	return s.name
}
