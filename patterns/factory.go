package main

import "errors"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop ",
			stock: 25,
		}}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{Computer: Computer{
		name:  "Desktop Computer",
		stock: 35,
	}}
}

func ComputerFactory(computer string) (IProduct, error) {
	switch computer {
	case "Laptop":
		return NewLaptop(), nil
	case "Desktop":
		return NewDesktop(), nil
	default:
		return nil, errors.New("invalid computer type")
	}
}