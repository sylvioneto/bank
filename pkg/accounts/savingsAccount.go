package accounts

import (
	"com.sylvioneto.bank/pkg/customers"
	"fmt"
)

type SavingsAccount struct {
	Customer  customers.Individual
	BranchNo  string
	AccountNo string
	balance   float64
}

// Withdraw validates balance and updates it
func (c *SavingsAccount) Withdraw(amount float64) error {
	var err error
	fmt.Printf("Withdraw amount: %.2f \n", amount)
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
func (c *SavingsAccount) Deposit(amount float64) error {
	var err error
	fmt.Printf("Deposit amount: %.2f \n", amount)
	if amount > 0 {
		c.balance += amount
	} else {
		err = fmt.Errorf("Invalid deposit amount: %.2f", amount)
	}
	return err
}

// FundsTransfer transfer funds from source to target account
func (c *SavingsAccount) FundsTransfer(amount float64, targetAcc *SavingsAccount) error {
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

// GetBalance returns the balance
func (c *SavingsAccount) GetBalance() float64{
	return c.balance
}
