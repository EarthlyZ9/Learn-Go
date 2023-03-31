package main

import (
	"fmt"
	"log"

	"github.com/EarthlyZ9/learngo/banking/accounts"
)

func main() {
	account := accounts.NewBankAccount("Jisoo")

	account.Deposit(1000)
	fmt.Println(account.GetBalance())

	// Go does not have 'Exceptions' - should be handled manually
	err := account.Withdraw(80)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.GetBalance())

	account.ChangeOwner("Z9")
	fmt.Println(account.GetOwner())

	fmt.Println(account)

}
