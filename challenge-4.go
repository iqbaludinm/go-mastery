package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	name       string
	address    string
	occupation string
	reason     string
}

type Classroom struct {
	students []Biodata
}

type getterBiodata interface {
	getName() string
	getAdress() string
	getOccupation() string
	getReason() string
}

func (b Biodata) getName() string {
	return b.name
}

func (b Biodata) getAdress() string {
	return b.address
}

func (b Biodata) getOccupation() string {
	return b.occupation
}

func (b Biodata) getReason() string {
	return b.reason
}

func display(index int, listStudents Classroom) {
	if index >= 1 && index <= len(listStudents.students) {
		student := listStudents.students[index-1]
		fmt.Printf("\nName\t\t: %s\nAddress\t\t: %s\nOccupation\t: %s\nReason\t\t: %s\n", student.getName(), student.getAdress(), student.getOccupation(), student.getReason())
	} else {
		fmt.Println("Data not found")
		os.Exit(2)
	}
}

func main() {
	studentsData := []Biodata{
		{ 
			name: "Iqbal", 
			address: "Malang", 
			occupation: "Software Engineer", 
			reason: "Interested in learning Go",
		},
		{ 
			name: "Rizqy", 
			address: "Surabaya", 
			occupation: "DB Engineer", 
			reason: "Wants to improve data analysis skills with Go",
		},
		{ 
			name: "Tatang", 
			address: "Singapore", 
			occupation: "Web Developer", 
			reason: "Excited about Go's web development capabilities",
		},
	}

	arg := os.Args[1]
	i, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(2)
	}

	classroom := Classroom{students: studentsData}
	display(i, classroom)
	
}