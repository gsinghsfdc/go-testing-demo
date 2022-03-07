package main

import (
	"fmt"
)

func main() {
	fmt.Print("hello world!")

	tm := SqlTransactionManager{}

	acc := ChequingAccount{AccountNumber: 1}
	acc.Deposit(100, tm)
}
