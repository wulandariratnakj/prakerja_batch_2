package main

import "fmt"

type Student struct {
	name  string
	score int
}

const numStudents = 5

func main() {
	fmt.Println("Input : ")

	var students []Student
	var total int

	var minScore = 100
	var maxScore = 0
	var minStudent, maxStudent Student

	for i := 0; i < numStudents; i++ {
		fmt.Printf("Input %d Student's Name: ", i+1)
		var name string
		fmt.Scanln(&name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		var score int
		fmt.Scanln(&score)

		students = append(students, Student{name, score})
		total += score

		if score < minScore {
			minScore = score
			minStudent = students[i]
		}

		if score > maxScore {
			maxScore = score
			maxStudent = students[i]
		}
	}

	fmt.Println("Output:")
	fmt.Printf("Average Score: %.2f\n", float64(total)/float64(numStudents))
	fmt.Printf("Min score of Students: %s (%d)\n", minStudent.name, minScore)
	fmt.Printf("Max score of Students: %s (%d)\n", maxStudent.name, maxScore)
}
