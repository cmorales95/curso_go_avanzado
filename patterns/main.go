package main

import "fmt"

func main() {
	laptop, _ := ComputerFactory("Laptop")
	desktop, _ := ComputerFactory("Desktop")

	fmt.Println(laptop)
	fmt.Println(desktop)

	TestInstanceDatabase()

	cash := &CashPayment{}
	ProcessPayment(cash)

	//bank := &BankPayment{}
	//ProcessPayment(bank)

	bpa := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		BankAccount: 5,
	}
	ProcessPayment(bpa)

	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &Email{id: "12ab"}
	secondObserver := &Email{id: "34dc"}

	nvidiaItem.Register(firstObserver)
	nvidiaItem.Register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
