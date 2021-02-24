package main

import (
	"fmt"

	"github.com/shawnzxx/OOPSample/encapsulation/payment"
)

func main() {
	creditCard := payment.CreateCreditCard("Zhang Xiaoxiao", "4444-3333-2222-1111", 10, 2021, 789)
	fmt.Printf("Owner's name: %s\n", creditCard.OwnerName())
	fmt.Printf("Owner's card number: %s\n", creditCard.CardNumber())
	fmt.Println("Owner trying to change card number")
	err := creditCard.SetCardNumber("Invalid Card Number")
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
	}
	fmt.Printf("Owner's credit card balance: %v\n", creditCard.AvailableCredit())

	//Since CreditCard object implemented the interface so can assign it to the interface object
	var option payment.PaymentOption
	option = creditCard
	option.ProcessPayment(500)

	option = payment.CreateCashAccount()
	option.ProcessPayment(500)
}
