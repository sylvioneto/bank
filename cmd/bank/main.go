package main

import (
	"fmt"

	"com.sylvioneto.bank/pkg/accounts"
	"com.sylvioneto.bank/pkg/customers"
)

func main() {
	fmt.Println("Accounts created for testing")
	sylvioCustomer := customers.Individual{Name: "sylvio", Id: 11}
	sylvioAccount := accounts.CurrentAccount{Customer: sylvioCustomer, BranchNo: "0001", AccountNo: "12345"}
	sylvioAccount2 := accounts.CurrentAccount{Customer: sylvioCustomer, BranchNo: "0001", AccountNo: "22222"}
	sylvioAccount.Deposit(100)
	sylvioAccount2.Deposit(200)
	fmt.Println(sylvioAccount, sylvioAccount2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 1: transfer with sufficient funds")
	err := sylvioAccount.FundsTransfer(50, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount, sylvioAccount2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 2: transfer insufficient Balance")
	err = sylvioAccount.FundsTransfer(1000, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount, sylvioAccount2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 3: invalid amount")
	err = sylvioAccount.FundsTransfer(-1, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount, sylvioAccount2)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 4: Print balance")
	fmt.Println(sylvioAccount.GetBalance(), sylvioAccount2.GetBalance())
	fmt.Println("-------------------------------------------")
}
