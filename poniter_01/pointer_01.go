package main

import "errors"

type BankAcc struct {
	Balance int
}

func (b *BankAcc) Add(amount int) error {
	b.Balance += amount
	return nil
}

func (b *BankAcc) Decrease(amount int) error {
	if b.Balance < amount {
		return errors.New("insufficient balance")
	}
	b.Balance -= amount
	return nil
}
