package main

import (
	"Preeti_Challange/creditcardchecker/pkg"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type CardValidator struct{}

func (cv *CardValidator) Validate(cardno string) pkg.CardValidationResult {
	re := regexp.MustCompile("[^0-9]+")
	card := re.ReplaceAllString(cardno, "")
	if card == "" || len(card) < 13 || len(card) > 19 {
		return pkg.CardValidationResult{
			CardNo:    card,
			Carderror: errors.New("card length check failed"),
		}
	}
	// fmt.Println("CardNo after cleaning: " + card)

	if !cv.luhnCheck(card) {
		return pkg.CardValidationResult{
			CardNo:    card,
			Carderror: errors.New("failed luhn check"),
		}
	}

	cardCompany := pkg.CreditCardCompany{}
	cc, err := cardCompany.GleanCompany(card)
	if err != nil {
		return pkg.CardValidationResult{
			CardNo:    card,
			Carderror: err,
		}
	}

	return pkg.CardValidationResult{
		CardNo:    card,
		CardType:  cc,
		Valid:     true,
		Carderror: nil,
	}

}

func (cv *CardValidator) luhnCheck(cardno string) bool {
	digits := len(cardno)
	oddOrEven := digits & 1
	sum := 0

	for count := 0; count < digits; count++ {
		t := string(cardno[count])
		digit, err := strconv.Atoi(t)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return false
		}

		if (count&1)^oddOrEven == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	if sum == 0 {
		return false
	}

	return sum%10 == 0
}

func printTest(cardno string) {
	result := CardValidator{}
	cvr := result.Validate(cardno)
	if !cvr.Valid {
		fmt.Println("invalid")
		return
	}
	fmt.Printf(cvr.IsValid() + " : " + cvr.CardType + "\n")
}

func main() {
	visa := "4444444444444448"
	master := "5500005555555559"
	amex := "371449635398431"
	diners := "36438936438936"
	discover := "6011016011016011"
	jcb := "3566003566003566"
	luhnFail := "1111111111111111"
	invalid := "abcdabcdabcdabcd"

	printTest(visa)
	printTest(master)
	printTest(amex)
	printTest(diners)
	printTest(discover)
	printTest(jcb)
	printTest(invalid)
	printTest(luhnFail)
}
