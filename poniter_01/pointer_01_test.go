package main

import (
	"testing"
)

func TestBankAcc(t *testing.T) {
	t.Run("add balance", func(t *testing.T) {
		acc := BankAcc{10}
		err := acc.Add(20)
		chackError(t, err)
		chackBalance(t, acc, BankAcc{30})
	})

	t.Run("Decrease balance", func(t *testing.T) {
		acc := BankAcc{10}
		err := acc.Decrease(2)
		chackError(t, err)
		chackBalance(t, acc, BankAcc{8})
	})
}

func chackError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func chackBalance(t *testing.T, acc BankAcc, want BankAcc) {
	t.Helper()
	got := acc
	if got != want {
		t.Errorf("got %d, want %d", got.Balance, want.Balance)
	}
}
