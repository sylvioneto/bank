package main

import "fmt"

type CurrentAccount struct {
	customerName string
	branchNo     string
	accountNo    string
	balance      float32
}

func main() {
	fmt.Println(CurrentAccount{})
}
