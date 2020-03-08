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
	sylvioAccount := CurrentAccount{customerName: "sylvio", branchNo: "0001", accountNo: "12345", balance: 100.01}
	fmt.Println(sylvioAccount)

	err := sylvioAccount.Withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)

	err = sylvioAccount.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)

	err = sylvioAccount.Withdraw(-50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)

	err = sylvioAccount.Deposit(-50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)

	err = sylvioAccount.Deposit(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sylvioAccount)
}

// Withdraw validates balance and updates it
func (c *CurrentAccount) Withdraw(amount float64) error {
	var err error
	hasBalance := c.balance >= amount && amount > 0
	if hasBalance {
		c.balance -= amount
	} else {
		err = fmt.Errorf("Insuficient balance: %.2f, or Invalid amount %.2f", c.balance, amount)
	}
	return err
}

// Deposit validates amount and updates the deposit
func (c *CurrentAccount) Deposit(amount float64) error {
	var err error
	if amount > 0 {
		c.balance += amount
		fmt.Println("New balance is", c.balance)
	} else {
		err = fmt.Errorf("Invalid deposit amount: %.2f", amount)
	}
	return err
}
