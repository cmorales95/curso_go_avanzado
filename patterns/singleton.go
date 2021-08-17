/*Creational Pattern*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("creating singleton for database")
	time.Sleep(2 * time.Second)
	fmt.Println("creation done")
}

var db *Database
var lock sync.Mutex
var once sync.Once

func GetDatabaseInstance() *Database {
	//lock.Lock()
	//defer lock.Unlock()
	//if db == nil {
	once.Do(func() {
		fmt.Println("Creating DB connection")
		db = &Database{}
		db.CreateSingleConnection()
	})
	//} else {
	//	fmt.Println("DB Already Created")
	//}
	return db
}

func TestInstanceDatabase() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			GetDatabaseInstance()
		}(&wg)
	}
	wg.Wait()
}
