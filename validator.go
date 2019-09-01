package main

import (
	"fmt"
	"time"
)

type Passport struct {
	Country         string
	ExpirationYear  int
	ExpirationMonth time.Month
}

func main() {
	fmt.Println("Welcome to your passport validator")
}
