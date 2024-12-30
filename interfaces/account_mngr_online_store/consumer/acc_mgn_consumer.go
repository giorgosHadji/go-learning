package main

import (
	"account_manager"
	"fmt"
)

func main() {

	e := account_manager.Employee{account_manager.Account{"kokos", "esi"}, 32}
	e.AddCredits(21)
	fmt.Println(e)
	e.ChangeName("kostakis")
	fmt.Println(e)
	fmt.Println(e.CheckCredits())
}
