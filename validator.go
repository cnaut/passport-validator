package main

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

type Passport struct {
	Country        string
	ExpirationDate Date
}

func isPassportValid(passport Passport, returnDate Date) bool {
	if passport.Country != "US" {
		return false
	}

	if passport.ExpirationDate.Year < returnDate.Year {
		return false
	}

	// Passport expires in the same year as the return date
	if passport.ExpirationDate.Year == returnDate.Year {
		if passport.ExpirationDate.Month < returnDate.Month {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Welcome to your passport validator")
	passport := Passport{Country: "US", ExpirationDate: Date{Year: 2020}}
	date := Date{Year: 2021}
	fmt.Println(isPassportValid(passport, date))
}
