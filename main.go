package main

import (
	"fmt"
)

type CurrentAccount struct {
	customerName string
	branchNo     string
	accountNo    string
	balance      float64
}

func main() {
	fmt.Println("Accounts:")
	sylvioAccount := CurrentAccount{customerName: "sylvio", branchNo: "0001", accountNo: "12345", balance: 100.01}
	fmt.Println(sylvioAccount)

	sylvioAccount2 := CurrentAccount{customerName: "sylvio", branchNo: "0001", accountNo: "22222", balance: 200.01}
	fmt.Println(sylvioAccount2)

	fmt.Println("Test Case 1: transfer with sufficient funds")
	err := sylvioAccount.FundsTransfer(50, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)
	fmt.Println(sylvioAccount2)

	fmt.Println("Test Case 2: transfer insufficient balance")
	err = sylvioAccount.FundsTransfer(1000, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)
	fmt.Println(sylvioAccount2)

	fmt.Println("Test Case 3: invalid amount")
	err = sylvioAccount.FundsTransfer(-1, &sylvioAccount2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)
	fmt.Println(sylvioAccount2)
}

// Withdraw validates balance and updates it
func (c *CurrentAccount) Withdraw(amount float64) error {
	var err error

	if amount <= 0 {
		err = fmt.Errorf("Invalid amount: %.2f", c.balance)
	} else if c.balance <= amount {
		err = fmt.Errorf("Insuficient balance: %.2f", c.balance)
	}

	if err == nil {
		c.balance -= amount
	}

	return err
}

// Deposit validates amount and updates the deposit
func (c *CurrentAccount) Deposit(amount float64) error {
	var err error
	if amount > 0 {
		c.balance += amount
	} else {
		err = fmt.Errorf("Invalid deposit amount: %.2f", amount)
	}
	return err
}

// FundsTransfer transfer funds from source to target account
func (c *CurrentAccount) FundsTransfer(amount float64, targetAcc *CurrentAccount) error {
	var err error
	if amount > 0 {
		err = c.Withdraw(amount)
	} else {
		err = fmt.Errorf("Invalid transfer amount %.2f", amount)
	}

	if err == nil {
		err = targetAcc.Deposit(amount)
	}
	return err
}
