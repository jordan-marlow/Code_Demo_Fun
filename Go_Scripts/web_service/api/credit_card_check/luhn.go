package credit_card_check

import (
	"unicode"
)

func isValidCreditCard(card_number string) bool {
	/*
		This function uses the Luhn algorithm to
		determine if a given credit card number
		is valid or not.
		Reference:  https://en.wikipedia.org/wiki/Luhn_algorithm
	*/
	sum := 0
	x := len(card_number)
	parity := x % 2
	for i := 0; i < x; i++ {
		r := rune(card_number[i])
		if !unicode.IsDigit(r) {
			return false
		}
		value := int(r - '0')
		if i%2 != parity {
			sum += value
		} else if value > 4 {
			sum += 2*value - 9
		} else {
			sum += 2 * value
		}

	}
	return sum%10 == 0
}

func getCreditCardMaker(card_number string) string {
	/*
		Returns the given maker of the credit card
		based off the values of the credit card
		number.
		The logic thusfar only works off the first
		number of the card and its self explanatory
		in the code below.
	*/
	r := rune(card_number[0])
	if !unicode.IsDigit(r) {
		return "Unknown"
	}
	value := int(r - '0')
	switch value {
	case 3:
		return "American Express"
	case 4:
		return "Visa"
	case 5:
		return "MasterCard"
	case 6:
		return "Discover"
	default:
		return "Unknown"
	}
}
