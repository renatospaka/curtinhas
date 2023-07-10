package main

import (
	"fmt"
	"sync"
)

type User struct {
	lock sync.RWMutex
	Name string
}

func doSomething(u *User) {
	u.lock.RLock()
	defer u.lock.RUnlock()

	// do something with `u`
	u.Name = "Name"
	fmt.Printf("Other name: %s\n", u.Name)
}

func main() {
	u := &User{Name: "John"}
	doSomething(u)
	fmt.Printf("Name: %s\n", u.Name)
}
