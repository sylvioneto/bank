package main

import (
	"fmt"

	"com.sylvioneto.bank/accounts"
)

func main() {
	fmt.Println("Accounts created for testing")
	sylvioAccount := accounts.CurrentAccount{CustomerName: "sylvio", BranchNo: "0001", AccountNo: "12345", Balance: 100.01}
	sylvioAccount2 := accounts.CurrentAccount{CustomerName: "sylvio", BranchNo: "0001", AccountNo: "22222", Balance: 200.01}
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
}
