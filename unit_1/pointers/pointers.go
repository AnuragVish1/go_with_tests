package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float32

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrorMessage = errors.New("Insufficient Balance")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	}
	return errors.New("Insufficient Balance")
}

func main() {

}
