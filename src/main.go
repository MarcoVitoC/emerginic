package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
)

type Patient struct {
	name string
	age int
	symptoms string
	severity string
	code int
	next, prev *Patient // by default, it is already nil
}

var front, back *Patient

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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

	if front == nil {
		front = newPatient
		back = newPatient
	} else if newPatient.code >= front.code { // push head
		newPatient.next = front
		front.prev = newPatient
		front = newPatient
	} else if newPatient.code <= back.code { // push tail
		back.next = newPatient
		newPatient.prev = back
		back = newPatient
	} else { // push mid
		curr := front

		for newPatient.code <= front.code {
			curr = curr.next
		}

		newPatient.prev = curr.prev
		newPatient.next = curr
		curr.prev.next = newPatient
		curr.prev = newPatient
	}
}

func insertPatient() {
	reader := bufio.NewReader(os.Stdin)
	
	var name, symptoms, severity string
	var age, code int

	for {
		fmt.Print("Input patient name[4-25]: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if (len(name) >= 4 && len(name) <= 25) {
			break
		} else {
			fmt.Println("Name must be greater than 4 characters and less than 25 characters")
		}
	}

	for {
		fmt.Print("Input patient age[> 0]: ")
		_, err := fmt.Scan(&age)
		fmt.Scanln()

		if (err != nil) {
			fmt.Println("Invalid input. Please enter a number.")

			var invalidInput string
			fmt.Scanln(&invalidInput)
		} else if (age <= 0) {
			fmt.Println("Age must be greater than 0")
		} else {
			break
		}
	}

	for {
		fmt.Print("Input patient symptoms[>= 6]: ")
		symptoms, _ = reader.ReadString('\n')
		symptoms = strings.TrimSpace(symptoms)

		if (len(symptoms) >= 6) {
			break
		} else {
			fmt.Println("Symptoms description must be greater than 6 characters")
		}
	}

	for {
		fmt.Print("Input severity[Severe|Moderate|Mild]: ")
		fmt.Scan(&severity)

		if (severity == "Severe" || severity == "Moderate" || severity == "Mild") {
			break
		}
	}

	if severity == "Severe" || severity == "Mild" {
		code = len(severity) - 3
	} else {
		code = len(severity) - 6
	}

	createNewPatient(name, age, symptoms, severity, code)

	fmt.Println()
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
	menu()
}

func viewPatients() {
	if front == nil {
		fmt.Println("There is no patient!")
	} else {
		fmt.Println("Patient List:")
		fmt.Println("------------------------------------------------------------------------------------------------------")
		fmt.Printf("|%-3s| %-20s| %-4s| %-55s| %-10s|\n", "No", "Name", "Age", "Symptoms", "Severity")
		fmt.Println("------------------------------------------------------------------------------------------------------")
		
		var iteration int = 1
		curr := front
		for curr != nil {
			fmt.Printf("|%-3d| %-20s| %-4d| %-55s| %-10s|\n", iteration, curr.name, curr.age, curr.symptoms, curr.severity)
			iteration++
			curr = curr.next
		}
		fmt.Println("------------------------------------------------------------------------------------------------------")
	}

	fmt.Println()
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
	menu()
}

func nextPatient() {
	if front == nil {
		fmt.Println("There is no patient!")
	} else {
		fmt.Println("Name: ", front.name)
		fmt.Println("Age: ", front.age)
		fmt.Println("Symptoms: ", front.symptoms)
		fmt.Println("Severity: ", front.severity)

		if front == back {
			front = nil
			back = nil
		} else {
			curr := front

			front = curr.next
			curr.next = nil
			front.prev = nil
		}
	}

	fmt.Println()
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
	menu()
}

func menu() {
	var menu int

	for (menu < 1 || menu > 5) {
		cls()
		fmt.Println("🏥 Emerginic")
		fmt.Println("=============")
		fmt.Println("1. Insert Patient")
		fmt.Println("2. View Patients")
		fmt.Println("3. Next Patient")
		fmt.Println("4. Exit")
		fmt.Print(">> ")
		fmt.Scan(&menu)
		fmt.Scanln()
	}

	switch (menu) {
		case 1:
			cls()
			insertPatient()
		case 2:
			cls()
			viewPatients()
		case 3:
			cls()
			nextPatient()
		case 4:
			cls()
			fmt.Println("Thank you for using this app!")
			os.Exit(0)
	}
}

func main() {
	menu()
}