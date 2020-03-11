package main

import (
	"fmt"

	"com.sylvioneto.bank/pkg/accounts"
	"com.sylvioneto.bank/pkg/customers"
)

func main() {
	fmt.Println("Accounts created for testing")
	customer := customers.Individual{Name: "sylvio", Id: 11}
	account := accounts.CurrentAccount{Customer: customer, BranchNo: "0001", AccountNo: "12345"}
	account2 := accounts.CurrentAccount{Customer: customer, BranchNo: "0001", AccountNo: "22222"}
	account.Deposit(100)
	account2.Deposit(200)
	fmt.Println(account, account2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 1: transfer with sufficient funds")
	err := account.FundsTransfer(50, &account2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account, account2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 2: transfer insufficient Balance")
	err = account.FundsTransfer(1000, &account2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account, account2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 3: invalid amount")
	err = account.FundsTransfer(-1, &account2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account, account2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 4: Print balance")
	fmt.Println(account.GetBalance(), account2.GetBalance())
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 5: Create a Savings Accounts")
	savings := accounts.SavingsAccount{}
	savings.Deposit(200)
	fmt.Println(savings)
	fmt.Println("-------------------------------------------")

}
