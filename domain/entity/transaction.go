package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	err := t.validateAmountGreaterThanOneThousand()
	if err != nil {
		return err
	}
	err = t.validateAmountGreaterThanOne()
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) validateAmountGreaterThanOneThousand() error {
	if t.Amount > 1000 {
		return errors.New("no limit for this transaction")
	}
	return nil
}

func (t *Transaction) validateAmountGreaterThanOne() error {
	if t.Amount < 1 {
		return errors.New("amount must be greater then one")
	}
	return nil
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}
