package integertoroman

func intToRoman(num int) string {
	roman := ""

	for num > 0 {
		if num >= 1000 {
			roman += "M"
			num -= 1000
		}
		if num >= 900 && num < 1000 {
			roman += "CM"
			num -= 900
		}
		if num >= 500 && num < 900 {
			roman += "D"
			num -= 500
		}
		if num >= 400 && num < 500 {
			roman += "CD"
			num -= 400
		}
		if num >= 100 && num < 400 {
			roman += "C"
			num -= 100
		}
		if num >= 90 && num < 100 {
			roman += "XC"
			num -= 90
		}
		if num >= 50 && num < 90 {
			roman += "L"
			num -= 50
		}
		if num >= 40 && num < 50 {
			roman += "XL"
			num -= 40
		}
		if num >= 10 && num < 40 {
			roman += "X"
			num -= 10
		}
		if num == 9 {
			roman += "IX"
			num -= 9
		}
		if num >= 5 && num < 9 {
			roman += "V"
			num -= 5
		}
		if num == 4 {
			roman += "IV"
			num -= 4
		}
		if num >= 1 && num < 4 {
			roman += "I"
			num -= 1
		}
	}

	return roman
}
