package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

//PaymentOption interface used to provide encapsulation by abstract away the service being provided when we invoke method through the interface
type PaymentOption interface {
	ProcessPayment(float32) bool
}

//CreditCard type
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

//CreateCreditCard Constructor of CreditCard
func CreateCreditCard(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
		availableCredit: 5000,
	}
}

func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

//OwnerName return CreditCard's ownerName filed
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

//SetOwnerName set CreditCard's ownerName filed
func (c *CreditCard) SetOwnerName(ownerName string) error {
	if len(ownerName) == 0 {
		return errors.New("Invalid owner name")
	}
	c.ownerName = ownerName
	return nil
}

//CardNumber return CreditCard's cardNumber filed
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

//SetCardNumber set CreditCard's cardNumber filed
func (c *CreditCard) SetCardNumber(cardNumber string) error {
	if !cardNumberPattern.Match([]byte(cardNumber)) {
		return errors.New("Invalid credit card number")
	}
	c.cardNumber = cardNumber
	return nil
}

//ExpirationDay return CreditCard's expirationMonth and expirationYear fileds
func (c CreditCard) ExpirationDay() (int, int) {
	return c.expirationMonth, c.expirationYear
}

//SetExpirationDay set CreditCard's expirationMonth and expirationYear fileds
func (c *CreditCard) SetExpirationDay(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}
	c.expirationMonth = month
	c.expirationYear = year
	return nil
}

//SecurityCode get CreditCard's CVV
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

//SetSecurityCode set CreditCard's CVV
func (c *CreditCard) SetSecurityCode(cvv int) error {
	if cvv < 100 || cvv > 999 {
		return errors.New("cvv must be three digit numbers")
	}
	return nil
}

//AvailableCredit get CreditCard's availableCredit, this can come from a web service, client doesn't know or care
func (c CreditCard) AvailableCredit() float32 {
	return c.availableCredit
}
