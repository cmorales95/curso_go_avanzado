/*Structural Pattern*/
package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (c *CashPayment) Pay() {
	fmt.Println("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (b BankPayment) Pay(bankAccount int) {
	fmt.Println("Pay using Bankaccount %d", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	BankAccount int
}

func (b *BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.BankAccount)
}
