package roman_to_integer

func RomanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return convert(s)
	}

	if i := convert(s[:2]); i != 0 {
		return i + RomanToInt(s[2:])
	}

	return convert(s[:1]) + RomanToInt(s[1:])
}

func convert(s string) int {
	if s == "IV" {
		return 4
	}

	if s == "IX" {
		return 9
	}

	if s == "XL" {
		return 40
	}

	if s == "XC" {
		return 90
	}

	if s == "CD" {
		return 400
	}

	if s == "CM" {
		return 900
	}

	if s == "M" {
		return 1000
	}

	if s == "D" {
		return 500
	}

	if s == "C" {
		return 100
	}

	if s == "L" {
		return 50
	}

	if s == "X" {
		return 10
	}

	if s == "V" {
		return 5
	}

	if s == "I" {
		return 1
	}

	return 0
}
