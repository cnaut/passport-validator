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

func IsPassportValid(passport Passport, returnDate Date) bool {
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

		// Passport expires within 6 months of departure date
		if passport.ExpirationDate.Month < (returnDate.Month + 6) {
			return false
		}

		if passport.ExpirationDate.Month == (returnDate.Month+6) && passport.ExpirationDate.Day < returnDate.Day {
			return false
		}
	} else if passport.ExpirationDate.Year == returnDate.Year+1 {
		if (passport.ExpirationDate.Month + 12) < (returnDate.Month + 6) {
			return false
		}

		if (passport.ExpirationDate.Month+12) == (returnDate.Month+6) && passport.ExpirationDate.Day < returnDate.Day {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Welcome to your passport validator")
	passport := Passport{Country: "US", ExpirationDate: Date{Year: 2020, Month: 4, Day: 14}}
	date := Date{Year: 2019, Month: 10, Day: 13}
	fmt.Println(IsPassportValid(passport, date))
}
