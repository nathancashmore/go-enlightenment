/*
A simple wallet in which you can Deposit, Withdraw or obtain a Balance of the
BitCoin held within it.

Provides examples of the following:
 - Pointers (*) otherwise pass by value
 - nil
 - errors
 - github.com/kisielk/errcheck

Remember

& <- makes a pointer from a variable.

* <-  "fetches" the value stored where a pointer points to.

*/
package pointers

import (
	"errors"
	"fmt"
)

const ErrNoFunds = "unable to withdraw more than your current balance"

type Stringer interface {
	String() string
}

type BitCoin int

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.balance {
		return errors.New(ErrNoFunds)
	}
	w.balance -= amount
	return nil
}
