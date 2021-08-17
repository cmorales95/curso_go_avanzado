package main

import "fmt"

func main() {
	fmt.Println("---Factory---")
	laptop, _ := ComputerFactory("Laptop")
	desktop, _ := ComputerFactory("Desktop")

	fmt.Println(laptop)
	fmt.Println(desktop)

	TestInstanceDatabase()
}
