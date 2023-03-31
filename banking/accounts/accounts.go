package accounts

import (
	"errors"
	"fmt"
)

// bankAccount struct
type bankAccount struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("NOT ENOUGH BALANCE")

// constructor function for bankAccount
func NewBankAccount(owner string) *bankAccount {
	account := bankAccount{owner: owner, balance: 0}
	return &account
}

// Deposit X amount
func (b *bankAccount) Deposit(amount int) {
	// b is actually a copy of bank account without a asterik
	b.balance += amount
}

// Withdraw X amount
func (b *bankAccount) Withdraw(amount int) error {
	if b.balance < amount {
		return errNoMoney
	}
	b.balance -= amount
	return nil
}

// Print balance
func (b bankAccount) GetBalance() int {
	return b.balance
}

// Change owner of the bank account
func (b *bankAccount) ChangeOwner(newOwner string) {
	b.owner = newOwner
}

// Get bank account owner
func (b bankAccount) GetOwner() string {
	return b.owner
}

// String()
func (b bankAccount) String() string {
	return fmt.Sprint(b.GetOwner(), "'s account: ", b.GetBalance())
}
