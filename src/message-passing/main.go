package main

import (
	"fmt"

	"github.com/shawnzxx/OOPSample/messagePassing/payment"
)

func main() {
	chargeCh := make(chan float32)
	//main package don't know what is implementation of CreditCardAccount doing, only through channel to passing data
	//how to process payment let consumer take care

	//messaging passing not like basica encapsulation provide define a method on object and outside invoke that object to manipulate the private fields
	//it's further more abstracting away from the invocation the call from the sender and what service is actually called
	payment.CreateCreditCardAccount(chargeCh)
	chargeCh <- 50.13
	var a string
	fmt.Scanln(&a)
}
