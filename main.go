package main

import (
	"fmt"

	"github.com/sylvioneto/bank/pkg/accounts"
	"github.com/sylvioneto/bank/pkg/customers"
)

// accountInterface is a interface to pay bills with Current and Savings accounts
type accountInterface interface {
	Withdraw(amount float64) error
}

func PayBill(account accountInterface, amount float64) error {
	return account.Withdraw(amount)
}

func main() {
	fmt.Println("Accounts created for testing")
	customer := customers.Individual{Name: "Neto", Id: 11, Email: "neto@x.com", HomeAddress: "Toronto"}
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
	savings := accounts.SavingsAccount{Customer: customer, BranchNo: "0001", AccountNo: "33333"}
	savings.Deposit(200)
	fmt.Println(savings)
	fmt.Println("-------------------------------------------")

	fmt.Println("Test Case 6: Pay bills using interfaces")
	fmt.Println(account)
	err = PayBill(&account, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	fmt.Println(savings)
	err = PayBill(&savings, 15)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(savings)
	fmt.Println("-------------------------------------------")
}
