package main

import (
	"bank"
	"fmt"
	"log"
	"net/http"
	"strconv"
)
// as this is outside of a function we need to use var and cannot use <:=> to declare the variable
var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}
	accounts[1003] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Georgios",
			Address: "Sunny Mexico",
			Phone:   "(213) 555 44444",
		},
		Number: 1003,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}
func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				fmt.Fprintf(w, account.Statement())
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	numberqsTo := req.URL.Query().Get("numberTo")
	numberqsFrom := req.URL.Query().Get("numberFrom")
	amountqs := req.URL.Query().Get("amount")

	if numberqsTo == "" {
		fmt.Fprintf(w, "Account number <to> is missing!")
		return
	}

	if numberqsFrom == "" {
		fmt.Fprintf(w, "Account number <from> is missing!")
		return
	}
	if numberTo, err := strconv.ParseFloat(numberqsTo, 64); err != nil {
		fmt.Fprintf(w, "Invalid account <to>  number!")
		return
	} else {
		if numberFrom, err := strconv.ParseFloat(numberqsFrom, 64); err != nil {
			fmt.Fprintf(w, "Invalid account <from>  number!")
			return
		} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
			fmt.Fprintf(w, "Invalid amount number!")
		} else {
			accountTo, okTo := accounts[numberTo]
			accountFrom, okFrom := accounts[numberFrom]
			if !okTo && !okFrom {
				fmt.Fprintf(w, "Correct to and from accounts can't be found!")
			} else {
				err := accountFrom.Transfer(amount, accountTo)
				if err != nil {
					fmt.Fprintf(w, "%v", err)
				} else {
					fmt.Fprintf(w, accountFrom.Statement())
				}
			}
		}
	}
}
