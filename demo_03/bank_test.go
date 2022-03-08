package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func TestBalanceEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080/balance?account=12", nil)
	rec := httptest.NewRecorder()

	balanceHandler(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %v", result.StatusCode)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("Error reading body")
	}

	// silly comparison, ideally the body will be parsed, this is just for demonstration purposes
	if strBody := string(body); strBody != "10\n" {
		t.Fatalf("expected 10, got %s", strBody)
	}
}

func TestBalanceEndpointWithTestServer(t *testing.T) {
	server := httptest.NewServer(handler())
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/balance?account=12", server.URL))
	if err != nil {
		t.Fatalf("an error occured making request to balance endpoint, err: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading body")
	}

	// silly comparison, ideally the body will be parsed, this is just for demonstration purposes
	if strBody := string(body); strBody != "10\n" {
		t.Fatalf("expected 10, got %s", strBody)
	}
}
