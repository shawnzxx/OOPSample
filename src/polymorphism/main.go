package main

import "github.com/shawnzxx/OOPSample/polymorphism/paymentgateway"

func main() {

	//Locate interface at point of consumption layer
	//especially for implement 3rd party library, standard is not publishing library define the interface but rather consuming libraries to define their interface
	//it can easily been removed when you don't need it anymore
	type PaymentOption interface {
		ProcessPayment(float32) bool
	}

	//below PaymentOption can assign different object
	//we can transparent use those type through the interface
	var option PaymentOption

	option = &paymentgateway.CheckingAccount{}

	option.ProcessPayment(10)

	option = &paymentgateway.CreditCard{}

	option.ProcessPayment(20)
}
