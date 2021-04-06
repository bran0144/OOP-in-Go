package main

import (
	"fmt"
)

type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment")
}
func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)
	return creditAccount
}
func main() {

	chargeCh := make(chan float32)
	CreateCreditAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)
	// var option payment.PaymentOption

	// option = payment.CreateCreditAccount(
	// 	"Debra Ingram",
	// 	"1111-2222-3333-4444",
	// 	5,
	// 	2021,
	// 	123)

	// option.ProcessPayment(500)

	// option = payment.CreateCashAccount()
	// option.ProcessPayment(500)

	// fmt.Printf("Owner name: %v\n", credit.OwnerName())
	// fmt.Printf("Card Number: %v\n", credit.CardNumber())
	// fmt.Println("Trying to change card number")
	// err := credit.SetCardNumber("invalid")
	// if err != nil {
	// 	fmt.Printf("That didn't work: %v\n", err)
	// }
	// fmt.Printf("Available credit: %v\n", credit.AvailableCredit())
}
