package pkg

import (
	"errors"
	"regexp"
)

// const Visa string = "^4[0-9]{12}(?:[0-9]{3})?$"
// const Master string = "^5[1-5][0-9]{14}$"
// const AMEX string = "^3[47][0-9]{13}$"
// const Diners string = "^3(?:0[0-5]|[68][0-9])[0-9]{11}$"
// const Discover string = "^6(?:011|5[0-9]{2})[0-9]{12}$"
// const JCB string = "^(?:2131|1800|35\\d{3})\\d{11}$"

var CardCompanies = map[string]string{
	"Visa":     "^4[0-9]{12}(?:[0-9]{3})?$",
	"Master":   "^5[1-5][0-9]{14}$",
	"AMEX":     "^3[47][0-9]{13}$",
	"Diners":   "^3(?:0[0-5]|[68][0-9])[0-9]{11}$",
	"Discover": "^6(?:011|5[0-9]{2})[0-9]{12}$",
	"JCB":      "^(?:2131|1800|35\\d{3})\\d{11}$",
}

type CreditCardCompany struct {
}

func (cc *CreditCardCompany) GleanCompany(cardno string) (string, error) {
	for k, v := range CardCompanies {
		m, _ := regexp.MatchString(v, cardno)
		if m {
			return k, nil
		}
	}

	return "", errors.New("invalid credit card")
}
