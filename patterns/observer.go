/*behavior pattern*/
package main

import "fmt"

type Topic interface {
	Register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateVale(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{name: name}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateVale(i.name)
	}
}

func (i *Item) Register(observer Observer) {
	i.observers = append(i.observers, observer)
}

type Email struct {
	id string
}

func (e *Email) getId() string {
	return e.id
}

func (e *Email) updateVale(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n\n", value, e.id)
}
