package pkg

type CardValidationResult struct {
	Valid     bool
	CardType  string
	Carderror error
	CardNo    string
}

func (cvr *CardValidationResult) IsValid() string {
	if cvr.getError() != nil {
		return "invalid"
	}
	if cvr.Valid {
		return "valid"
	}
	return "invalid"
}

func (cvr *CardValidationResult) getCardType() string {
	return cvr.CardType
}

func (cvr *CardValidationResult) getError() error {
	return cvr.Carderror
}

func (cvr *CardValidationResult) getCardNo() string {
	return cvr.CardNo
}
