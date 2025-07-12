package main

import "testing"

// ---- getOrdinalDay tests ----
func TestGetOrdinalDay(t *testing.T) {
	cases := map[int]string{
		1: "1st", 2: "2nd", 3: "3rd", 4: "4th",
		11: "11th", 12: "12th", 13: "13th", 14: "14th",
		21: "21st", 22: "22nd", 23: "23rd", 24: "24th",
		31: "31st", 32: "32nd", 33: "33rd", 34: "34th",
	}
	for input, expected := range cases {
		got := getOrdinalDay(input)
		if got != expected {
			t.Errorf("getOrdinalDay(%d) = %s; want %s", input, got, expected)
		}
	}
}

// ---- validateDate tests ----
func TestValidateDate(t *testing.T) {
	cases := []struct {
		dob      Birthday
		expected string
	}{
		{Birthday{Day: 29, Month: 2, Year: 2020}, ""},
		{Birthday{Day: 31, Month: 4, Year: 2021}, "invalid Date"},
		{Birthday{Day: 31, Month: 12, Year: 2021}, ""},
		{Birthday{Day: 30, Month: 2, Year: 2021}, "invalid Date"},
	}

	for _, c := range cases {
		_, err := validateDate(c.dob)
		if err != nil && err.Error() != c.expected {
			t.Errorf("validateDate(%v) = %v; want %s", c.dob, err, c.expected)
		} else if err == nil && c.expected != "" {
			t.Errorf("validateDate(%v) = nil; want %s", c.dob, c.expected)
		}
	}
}
