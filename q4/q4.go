package q4

import "errors"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if basePrice <= 0 {
		return 0, errors.New("invalid base price")
	}

	var tax, stateTax float64
	switch taxCode {
	case 1:
		tax = .05
	case 2:
		tax = .1
	case 3:
		tax = .15
	default:
		return 0, errors.New("invalid tax code")
	}

	switch state {
	case "SP":
		stateTax = .1
	case "RJ":
		stateTax = .15
	case "MG":
		stateTax = .2
	case "ES":
		stateTax = .25
	default:
		stateTax = .3
	}

	return basePrice + (basePrice * tax) + (basePrice * stateTax), nil
}
