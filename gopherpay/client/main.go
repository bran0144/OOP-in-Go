package main

import "../../gopherpay/payment"

func main() {
	var option payment.PaymentOption

	option = &payment.CreditCard{}
	option.ProcessPayment(500)

	option = &payment.CheckingAccount{}

	option.ProcessPayment(500)
}
