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

var head, tail *Patient

func cls() {
	cmd := exec.Command("clear")
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
	reader := bufio.NewReader(os.Stdin)
	
	var name, symptoms, severity string
	var age, code int

	for {
		fmt.Print("Input patient name[4-25]: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if (len(name) >= 4 && len(name) <= 25) {
			break
		}
	}

	for {
		fmt.Print("Input patient age[> 0]: ")
		fmt.Scan(&age)
		fmt.Scanln()

		if (age > 0) {
			break
		}
	}

	for {
		fmt.Print("Input patient symptoms[>= 6]: ")
		symptoms, _ = reader.ReadString('\n')
		symptoms = strings.TrimSpace(symptoms)

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

	if severity == "Severe" || severity == "Mild" {
		code = len(severity) - 3
	} else {
		code = len(severity) - 6
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

	fmt.Println("Patient List:")
	fmt.Println("------------------------------------------------------------------------------------------------------")
	fmt.Printf("|%-3s| %-20s| %-4s| %-55s| %-10s|\n", "No", "Name", "Age", "Description", "Code")
	fmt.Println("------------------------------------------------------------------------------------------------------")
	
	var iteration int
	curr := head
	for curr != nil {
		fmt.Printf("|%-3d| %-20s| %-4d| %-55s| %-10s|\n", iteration, curr.name, curr.age, curr.symptoms, curr.severity)
		iteration++
		curr = curr.next
	}

	fmt.Println("------------------------------------------------------------------------------------------------------")
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
		fmt.Println("1. Insert")
		fmt.Println("2. View")
		fmt.Println("3. Update")
		fmt.Println("4. Next Queue")
		fmt.Println("5. Exit")
		fmt.Print(">> ")
		fmt.Scan(&menu)
		fmt.Scanln()
	}

	switch (menu) {
		case 1:
			cls()
			insert()
		case 2:
			cls()
			view()
		case 3:
			// update()
		case 4:
			// nextQueue()
		case 5:
			os.Exit(0)
	}
}

func main() {
	menu()
}