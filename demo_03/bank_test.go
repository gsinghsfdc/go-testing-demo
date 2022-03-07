package main

import (
	"errors"
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

func TestBalanceEndpoint(t *testing.T) {

}

func TestBalanceEndpointWithTestServer(t *testing.T) {

}
