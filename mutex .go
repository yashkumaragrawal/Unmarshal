package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    bal int
)

func init() {
    bal = 1000
}

func deposit(value int, done chan bool) {
    mutex.Lock()
    fmt.Printf("Depositing %d to account with balance: %d\n", value, bal)
    bal += value
    mutex.Unlock()
    done <- true
}

func withdraw(value int, done chan bool) {
    mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, bal)
    bal -= value
    mutex.Unlock()
    done <- true
}

func main() {
    fmt.Printf("In starting balance is %d\n",bal)   
    done := make(chan bool)
    go withdraw(700, done)
    go deposit(500, done)
    <-done
    <-done

    fmt.Printf("New Balance is %d\n", bal)
}
