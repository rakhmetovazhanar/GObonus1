package main

import (
	"fmt"
)

type Exam struct {
	subject      string
	teacher_name string
	class        int
	student_num  int
}

type Students struct {
	name   string
	grades int
}

func (s *Exam) setStudent_num(student_num int) {
	s.student_num = student_num
}

func (g Students) result() {
	if g.grades > 5 {
		fmt.Println(g.name + ", you pass!")
		fmt.Printf("Your grade is %d\n", g.grades)
	} else {
		fmt.Println(g.name + ", you have to pass exam again!")
		fmt.Printf("Your grade is %d\n", g.grades)
	}
}

func main() {

	stud1 := Students{
		name:   "Zhanar",
		grades: 10,
	}

	stud2 := Students{
		name:   "Amina",
		grades: 4,
	}

	stud3 := Students{
		name:   "Aya",
		grades: 6,
	}

	final := Exam{
		subject:      "Math",
		teacher_name: "Akerke",
		class:        11,
		student_num:  4,
	}

	final.setStudent_num(3)

	fmt.Printf("Subject name: %s\n", final.subject)
	fmt.Printf("Teacher name: %s\n", final.teacher_name)
	fmt.Printf("Class: %d\n", final.class)
	fmt.Printf("Student number: %d\n", final.student_num)
	stud1.result()
	stud2.result()
	stud3.result()
}
