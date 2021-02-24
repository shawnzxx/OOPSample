package payment

import "fmt"

type CreditCardAccount struct{}

//processPayment encapsulated only payment internal function can call, outside don't know this function
func (*CreditCardAccount) processPayment(amount float32) {
	fmt.Printf("Process payment with amount %f\n", amount)
}

//CreateCreditCardAccount use channel to further isolated main and payment package
func CreateCreditCardAccount(chargeCh chan float32) *CreditCardAccount {
	creditCardAccount := &CreditCardAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditCardAccount.processPayment(amount)
		}
	}(chargeCh)

	return creditCardAccount
}
