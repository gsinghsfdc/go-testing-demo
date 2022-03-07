package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ErrNotImplemented = errors.New("not implemented")

type TransactionReaderStub struct {
	HandleReadBalance func(int) (int, error)
}

func (tms TransactionReaderStub) ReadBalance(account int) (int, error) {
	if tms.HandleReadBalance == nil {
		return 0, ErrNotImplemented
	}

	return tms.HandleReadBalance(account)
}

func TestReadBalance(t *testing.T) {
	tm := TransactionReaderStub{HandleReadBalance: successfulRead100}

	account := ChequingAccount{AccountNumber: 1}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"rates":{"CAD":2.0}}`))
	}))
	defer server.Close()

	balance, err := account.BalanceCAD(server.URL, tm)
	if err != nil {
		t.Fatal("error not expected when reading balance")
	}

	if balance < 100 {
		t.Fatal(fmt.Sprintf("expected balance 100, got %v", balance))
	}
}

func successfulRead100(int) (int, error) {
	return 100, nil
}

func Test01(t *testing.T) {
	fmt.Print("hello world!")

	tm := SqlTransactionManager{}

	acc := ChequingAccount{AccountNumber: 1}

	balance, _ := acc.BalanceCAD("", tm)

	fmt.Print(balance)
}
