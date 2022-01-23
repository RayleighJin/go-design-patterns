package facade

import (
	"fmt"
)

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Printf("wallet balance added with %d successfully\n", amount)
	return
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance not sufficient")
	}
	w.balance -= amount
	return nil
}
