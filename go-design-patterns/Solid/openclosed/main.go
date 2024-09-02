package main

import "fmt"

// Example for Open Closed Principle from SOLID

type PaymentMethod interface {
	Pay()
}

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard", cc.amount)
}

func main() {
	p := Payment{}
	cc := CreditCard{12.22}

	p.Process(cc)

}