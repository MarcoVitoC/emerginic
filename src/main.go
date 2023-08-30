package main

import "fmt"

type Patient struct {
	name string
	age int
	symptoms string
	severity string
	code int
	next, prev *Patient // by default, it is already nil
}

var head, tail *Patient


func createPatient(name string, age int, symptoms string, severity string, code int) *Patient {
	return &Patient{
		name: name,
		age: age,
		symptoms: symptoms,
		severity: severity,
		code: code,
	}
}

func createNewPatient(name string, age int, symptoms string, severity string, code int) {
	newPatient := createPatient(name, age, symptoms, severity, code)

	if head == nil {
		head = newPatient
		tail = newPatient
	} else if newPatient.code >= head.code { // push head
		newPatient.next = head
		head.prev = newPatient
		head = newPatient
	} else if newPatient.code <= tail.code { // push tail
		tail.next = newPatient
		newPatient.prev = tail
		tail = newPatient
	} else { // push mid
		curr := head

		for newPatient.code <= head.code {
			curr = curr.next
		}

		newPatient.prev = curr.prev
		newPatient.next = curr
		curr.prev.next = newPatient
		curr.prev = newPatient
	}
}

func insert() {
	var name string
	var age int
	var symptoms string
	var severity string
	var code int

	for {
		fmt.Print("Input patient name[4-25]: ")
		fmt.Scan(&name)

		if (len(name) >= 4 && len(name) <= 25) {
			break
		}
	}

	for {
		fmt.Print("Input patient age[> 0]: ")
		fmt.Scan(&age)

		if (age > 0) {
			break
		}
	}

	for {
		fmt.Print("Input patient symptoms[>= 6]: ")
		fmt.Scan(&symptoms)

		if (len(symptoms) >= 6) {
			break
		}
	}

	for {
		fmt.Print("Input severity[Severe|Moderate|Mild]: ")
		fmt.Scan(&severity)

		if (severity == "Severe" || severity == "Moderate" || severity == "Mild") {
			break
		}
	}

	if severity == "Severe" {
		code = len(severity) - 3
	} else if severity == "Moderate" {
		code = len(severity) - 6
	} else {
		code = len(severity) - 3
	}

	createNewPatient(name, age, symptoms, severity, code)
	menu()
}

func view() {
	if head == nil {
		fmt.Println("There is no patient!")
		fmt.Println("Press enter to continue...")
		fmt.Scanln()
		menu()
	}

	curr := head
	for curr != nil {
		fmt.Println("Name: ", curr.name)
		fmt.Println("Age: ", curr.age)
		fmt.Println("Symptoms: ", curr.symptoms)
		fmt.Println("Severity: ", curr.severity)
		fmt.Println()

		curr = curr.next
	}
}

func menu() {
	var menu int

	for (menu < 1 || menu > 5) {
		fmt.Println("🏥 Emerginic")
		fmt.Println("=============")
		fmt.Println("1. Insert")
		fmt.Println("2. View")
		fmt.Println("3. Update")
		fmt.Println("4. Next Queue")
		fmt.Println("5. Exit")
		fmt.Print(">> ")
		fmt.Scan(&menu)
	}

	switch (menu) {
		case 1:
			insert()
		case 2:
			view()
		case 3:
			// update()
		case 4:
			// nextQueue()
		case 5:
			// exit()
	}
}

func main() {
	menu()
}