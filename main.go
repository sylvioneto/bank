package main

import "fmt"

type CurrentAccount struct {
	customerName string
	branchNo     string
	accountNo    string
	balance      float32
}

func main() {
	sylvioAccount := CurrentAccount{customerName: "sylvio", branchNo: "0001", accountNo: "12345", balance: 100.01}
	fmt.Println(sylvioAccount)

	netoAccount := CurrentAccount{"neto", "0001", "6789", 200.02}
	fmt.Println(netoAccount)

	var johnAccount *CurrentAccount
	johnAccount = new(CurrentAccount)
	johnAccount.customerName = "john"
	fmt.Println(*johnAccount)

	sylvioAccount2 := CurrentAccount{customerName: "sylvio", branchNo: "0001", accountNo: "12345", balance: 100.01}
	fmt.Println(sylvioAccount2)

	// compare two accounts with same values
	fmt.Println(sylvioAccount == sylvioAccount2)
}
