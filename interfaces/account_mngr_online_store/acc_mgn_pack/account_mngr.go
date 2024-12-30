package account_manager

import "fmt"

type Account struct {
	FirstName string
	Lastname  string
}

func (acc Account) String() string {
	return fmt.Sprintf("%v %v", acc.FirstName, acc.Lastname)
}

type Employee struct {
	Account
	NumCredits float64
}

func (e *Employee) AddCredits(credits float64) {
	e.NumCredits += credits
}

func (e *Employee) RemoveCredits(credits float64) {
	e.NumCredits -= credits
}
func (e *Employee) CheckCredits() float64 {
	return e.NumCredits
}
func (acc *Account) ChangeName(s string) {
	acc.FirstName = s
}
