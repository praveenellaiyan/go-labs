package main

import (
	"fmt"

	"github.com/praveenellaiyan/go-labs/oop/encapsulation/student"
)

func main() {

	studentRef := student.New()

	/*
	 * id is not an exported prop of student
	 * so it is hidden outside the package and restrict modification
	 * this depicts encapsulation concept which is of data hiding
	 * umcommentting the below statement throws compilation error
	 */
	//studentRef.Id = 500

	fmt.Println("================")
	fmt.Println("Before Updation")
	fmt.Println("================")

	displayStudentInfo(studentRef)

	fmt.Println("================")
	fmt.Println("After Updation")
	fmt.Println("================")

	updateStudentInfo(studentRef, "Tony", 200)

	displayStudentInfo(studentRef)
	fmt.Println("================")
}

func displayStudentInfo(stud *student.Student) {
	fmt.Println("Student Id >> ", stud.GetStudentID())
	fmt.Println("Student Name >> ", stud.GetStudentName())
}

func updateStudentInfo(stud *student.Student, name string, id int) {
	stud.SetStudentID(id)
	stud.SetStudentName(name)
}
