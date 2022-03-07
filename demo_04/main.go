package main

import (
	"fmt"
)

func main() {
	fmt.Print("hello world!")

	tm := SqlTransactionManager{}

	acc := ChequingAccount{AccountNumber: 1}

	balance, _ := acc.BalanceCAD("", tm)

	fmt.Print(balance)
}
