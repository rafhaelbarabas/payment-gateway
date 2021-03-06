package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_ValidateAmountValue(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "no limit for this transaction", err.Error())

	transaction.Amount = 0
	err = transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "amount must be greater then one", err.Error())
}
