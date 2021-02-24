package paymentgateway

import "fmt"

type CreditCard struct{}
type CheckingAccount struct{}

func (c *CreditCard) ProcessPayment(float32) bool {
	fmt.Println("process credit card payment")
	return true
}

func (c *CheckingAccount) ProcessPayment(float32) bool {
	fmt.Println("process checking payment")
	return true
}
