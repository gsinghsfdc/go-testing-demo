package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var transactionReader TransactionReader = SqlTransactionManager{}

func main() {
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/balance", balanceHandler)
	return r
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	account := r.URL.Query().Get("account")
	if account == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	accountNumeric, err := strconv.Atoi(account)
	if err != nil {
		http.Error(w, "unable to parse account value", http.StatusBadRequest)
	}

	acc := ChequingAccount{AccountNumber: accountNumeric}
	balance, err := acc.Balance(transactionReader)
	if err != nil {
		http.Error(w, "an error occured serving your request", http.StatusInternalServerError)
	}

	fmt.Fprintln(w, balance)
}
