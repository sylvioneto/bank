package accounts

import (
	"fmt"
)

type CurrentAccount struct {
	CustomerName string
	BranchNo     string
	AccountNo    string
	Balance      float64
}

// Withdraw validates Balance and updates it
func (c *CurrentAccount) Withdraw(amount float64) error {
	var err error
	fmt.Printf("Withdraw amount: %.2f \n", amount)
	if amount <= 0 {
		err = fmt.Errorf("Invalid amount: %.2f", c.Balance)
	} else if c.Balance <= amount {
		err = fmt.Errorf("Insuficient Balance: %.2f", c.Balance)
	}

	if err == nil {
		c.Balance -= amount
	}

	return err
}

// Deposit validates amount and updates the deposit
func (c *CurrentAccount) Deposit(amount float64) error {
	var err error
	fmt.Printf("Deposit amount: %.2f \n", amount)
	if amount > 0 {
		c.Balance += amount
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
