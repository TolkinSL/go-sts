package main

import "fmt"

type CoffeeMachine struct {
	model     string
	status    string
	workHours int
}

func needService(machine *CoffeeMachine) {
	fmt.Printf("Func: Status: %s\n", machine.status)
	machine.status = "Service"
	machine.workHours = 0
	fmt.Printf("Func: Status: %s\n", machine.status)
}

func main() {
	machine1 := CoffeeMachine {
		model: "Tefal",
		status: "Work",
		workHours: 15,
	}

	needService(&machine1)

	fmt.Println()
	fmt.Printf("Main: %#v\n", machine1)
}
