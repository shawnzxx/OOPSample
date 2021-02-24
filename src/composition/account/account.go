package account

import "fmt"

type Account struct {
	amount float32
}

func (a *Account) AvaliableAmount() float32 {
	fmt.Println("Retrieve avaliable amount from db")
	return a.amount
}

func (a *Account) TopupAmount() {
	fmt.Println("topup 100 dollar")
	a.amount = 100.00
}

func (a *Account) ProcessPayment(amount float32) {
	fmt.Printf("Process payment amount %f", amount)
	a.amount -= amount
}

//2nd strategy if you don't want expose base type outside you need to provide the filed name
//and write function explicityly so that outside can call your childfunction and compiler will calling parent function
type CreditAccount struct {
	account Account
}

func (c *CreditAccount) AvaliableAmount() float32 {
	return c.account.AvaliableAmount()
}
func (c *CreditAccount) TopupAmount() {
	c.account.TopupAmount()
}
func (c *CreditAccount) ProcessPayment(amount float32) {
	c.account.ProcessPayment(amount)
}

//1st strategy: list the type not giving field name
//compiler will auto delegate the calls to the embedded methods when parent type does not have method
type CheckingAccount struct {
	Account
}
