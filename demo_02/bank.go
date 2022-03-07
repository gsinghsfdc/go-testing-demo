package main

import (
	"errors"
	"fmt"
)

type TransactionManger interface {
	TransactionReader
	TransactionWriter
}

type TransactionReader interface {
	ReadBalance(int) (int, error)
}

type TransactionWriter interface {
	UpdateBalance(int, int) error
}

type SqlTransactionManager struct{}

func (stm SqlTransactionManager) ReadBalance(amount int) (int, error) {
	fmt.Print("Represents reading a real db")
	return 10, nil
}

func (stm SqlTransactionManager) UpdateBalance(amount, account int) error {
	fmt.Print("represents updating real db")
	return nil
}

type Account interface {
	Balance(TransactionReader) (int, error)
	Withdraw(int, TransactionManger) error
	Deposit(int, TransactionManger) error
}

type ChequingAccount struct {
	AccountNumber int
}

func (ca ChequingAccount) Balance(tr TransactionReader) (int, error) {
	// just for tests
	if ca.AccountNumber == 22 {
		return 0, fmt.Errorf("account %v is invalid", ca.AccountNumber)
	}
	balance, err := tr.ReadBalance(ca.AccountNumber)
	if err != nil {
		return 0, fmt.Errorf("error occured reading balance, %v", err)
	}
	return balance, nil
}

func (ca ChequingAccount) Deposit(amount int, tm TransactionManger) error {
	balance, err := ca.Balance(tm)
	if err != nil {
		return errors.New("failed to get balance for update")
	}
	err = tm.UpdateBalance(ca.AccountNumber, balance+amount)
	if err != nil {
		return errors.New("error occured updating balance")
	}
	fmt.Printf("Deposited %v", amount)
	return nil
}

func (ca ChequingAccount) Withdraw(amount int, tm TransactionManger) error {
	balance, err := ca.Balance(tm)
	if err != nil {
		return errors.New("failed to get balance for update")
	}
	if amount > balance {
		return errors.New("amount greated than balance")
	}
	tm.UpdateBalance(ca.AccountNumber, balance-amount)
	err = tm.UpdateBalance(ca.AccountNumber, balance+amount)
	if err != nil {
		return errors.New("error occured updating balance")
	}
	fmt.Printf("Withdrawn %v", amount)
	return nil
}
