package main

import (
	"fmt"

	"github.com/shawnzxx/OOPSample/composition/account"
)

func main() {
	ca := &account.CreditAccount{}
	fmt.Printf("Amount: %f\n", ca.AvaliableAmount())
	ca.TopupAmount()
	fmt.Printf("Amount: %f\n", ca.AvaliableAmount())
	ca.ProcessPayment(12.2)
	fmt.Printf("Amount: %f\n", ca.AvaliableAmount())

	ch := &account.CheckingAccount{}
	fmt.Printf("Amount: %f\n", ch.AvaliableAmount())
	ch.TopupAmount()
	fmt.Printf("Amount: %f\n", ch.AvaliableAmount())
	ch.ProcessPayment(15.5)
	fmt.Printf("Amount: %f\n", ch.AvaliableAmount())
}
