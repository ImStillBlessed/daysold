package main

import (
	"flag"
	"fmt"
	"time"
)

type Birthday struct {
	Day, Month, Year int
	Date             time.Time
}
type ageBreakdown struct {
	Years, ExtraMonths, TotalMonths, Weeks, Days int
}

func calculateAge(dob Birthday, age *ageBreakdown) {
	current := time.Now()

	years := current.Year() - dob.Date.Year()
	months := int(current.Month()) - int(dob.Date.Month())
	days := current.Day() - dob.Date.Day()

	if days < 0 {
		prevMonth := current.AddDate(0, -1, 0)
		days += daysInMonth(prevMonth.Year(), prevMonth.Month())
		months--
	}

	if months < 0 {
		months += 12
		years--
	}

	duration := current.Sub(dob.Date)
	totalDays := int(duration.Hours() / 24)
	totalWeeks := totalDays / 7

	age.Years = years
	age.ExtraMonths = months
	age.TotalMonths = years*12 + months
	age.Days = totalDays
	age.Weeks = totalWeeks
}

func daysInMonth(year int, month time.Month) int {
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

func validateDate(dob Birthday) (time.Time, error) {
	date := time.Date(dob.Year, time.Month(dob.Month), dob.Day, 0, 0, 0, 0, time.UTC)
	if date.After(time.Now()) {
		return time.Time{}, fmt.Errorf("date cannot be in the future")

	}
	if date.Day() != dob.Day || int(date.Month()) != dob.Month || date.Year() != dob.Year {
		return time.Time{}, fmt.Errorf("invalid Date")
	}

	return date, nil
}

func askInput(dob *Birthday) {
	var err error
	for {
		fmt.Printf("Enter your date of birth (e.g. 25 12 1990): ")
		fmt.Scan(&dob.Day, &dob.Month, &dob.Year)

		if dob.Date, err = validateDate(*dob); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
}

func main() {
	dobFlag := flag.String("dob", "", "Your date of birth in DD-MM-YYYY format")
	flag.Parse()

	var dob Birthday
	var age ageBreakdown
	fallback := false

	if *dobFlag != "" {
		_, err := fmt.Sscanf(*dobFlag, "%d-%d-%d", &dob.Day, &dob.Month, &dob.Year)
		if err != nil {
			fmt.Println("Invalid date from flag:", err)
			fallback = true
		} else {
			dob.Date, err = validateDate(dob)
			if err != nil {
				fmt.Println("Invalid date from flag:", err)
				fallback = true
			}
		}
	} else {
		fallback = true
	}

	if fallback {
		askInput(&dob)
	}

	fmt.Printf("Born on the %d of %s, %d\n", dob.Day, dob.Date.Month(), dob.Year)
	calculateAge(dob, &age)
	fmt.Printf("\nYou have lived for:\n")
	fmt.Printf("🟢 %d years and %d months\n", age.Years, age.ExtraMonths)
	fmt.Printf("📆 %d total months\n", age.TotalMonths)
	fmt.Printf("📅 %d weeks\n", age.Weeks)
	fmt.Printf("🕒 %d days\n", age.Days)
	fmt.Println()
}
