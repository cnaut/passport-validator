package main

import (
	"testing"
)

func TestIsPassportValid(t *testing.T) {
	tables := []struct {
		passport      Passport
		departureDate Date
		isValid       bool
	}{
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 1, Day: 1}}, Date{Year: 2020, Month: 1, Day: 1}, true},
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 1, Day: 1}}, Date{Year: 2022, Month: 1, Day: 1}, false},
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 1, Day: 1}}, Date{Year: 2021, Month: 2, Day: 1}, false},
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 12, Day: 1}}, Date{Year: 2021, Month: 1, Day: 1}, true},
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 7, Day: 2}}, Date{Year: 2021, Month: 1, Day: 1}, true},
		{Passport{Country: "US", ExpirationDate: Date{Year: 2021, Month: 6, Day: 2}}, Date{Year: 2021, Month: 1, Day: 1}, false},
		{Passport{Country: "UK", ExpirationDate: Date{Year: 2021, Month: 1, Day: 1}}, Date{Year: 2020, Month: 1, Day: 1}, false},
	}

	for _, table := range tables {
		isValid := IsPassportValid(table.passport, table.departureDate)
		if isValid != table.isValid {
			t.Errorf("Validity of %s passport expiring on %d/%d/%d was not correct. Expected %t", table.passport.Country, table.passport.ExpirationDate.Month, table.passport.ExpirationDate.Day, table.passport.ExpirationDate.Year, table.isValid)
		}
	}
}
