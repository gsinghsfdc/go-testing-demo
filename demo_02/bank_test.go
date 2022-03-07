package main

import (
	"errors"
	"fmt"
	"testing"
)

var ErrNotImplemented = errors.New("not implemented")

type TransactionMangerStub struct {
	HandleReadBalance   func(int) (int, error)
	HandleUpdateBalance func(int, int) error
}

func (tms TransactionMangerStub) ReadBalance(account int) (int, error) {
	if tms.HandleReadBalance == nil {
		return 0, ErrNotImplemented
	}

	return tms.HandleReadBalance(account)
}

func (tms TransactionMangerStub) UpdateBalance(amount, account int) error {
	if tms.HandleUpdateBalance == nil {
		return ErrNotImplemented
	}

	return tms.HandleUpdateBalance(amount, account)
}

func TestReadBalanceEdgeCases(t *testing.T) {
	tt := map[string]struct {
		accountNumber             int
		transactionReaderBehavior func(int) (int, error)
	}{
		"AccountNumberInvalid": {
			accountNumber:             22,
			transactionReaderBehavior: successfulRead100},
		"FailedRead": {
			accountNumber:             1,
			transactionReaderBehavior: func(int) (int, error) { return 0, errors.New("failed to read balance") }},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tm := TransactionMangerStub{HandleReadBalance: test.transactionReaderBehavior}
			account := ChequingAccount{AccountNumber: test.accountNumber}

			_, err := account.Balance(tm)

			if err == nil {
				t.Fatal("Failure expected, but got no error")
			}
		})
	}
}

func TestReadBalance(t *testing.T) {
	tm := TransactionMangerStub{HandleReadBalance: successfulRead100}

	account := ChequingAccount{AccountNumber: 1}
	balance, err := account.Balance(tm)
	if err != nil {
		t.Fatal("error not expected when reading balance")
	}

	if balance != 100 {
		t.Fatal(fmt.Sprintf("expected balance 100, got %v", balance))
	}
}

func TestDeposit(t *testing.T) {
	tm := TransactionMangerStub{HandleReadBalance: successfulRead100, HandleUpdateBalance: func(int, int) error { return nil }}

	account := ChequingAccount{AccountNumber: 1}
	err := account.Deposit(100, tm)

	if err != nil {
		t.Fatal(fmt.Sprintf("expected no error on deposit, got %v", err))
	}
}

func successfulRead100(int) (int, error) {
	return 100, nil
}
