package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {

	_, err := NewCreditCard("5000000000000000", "Jose da Silva", 12, 2022, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5458628241493448", "Jose da Silva", 12, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("5458628241493448", "Jose da Silva", 13, 2022, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5458628241493448", "Jose da Silva", 0, 2022, 123)
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5458628241493448", "Jose da Silva", 1, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("5458628241493448", "Jose da Silva", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5458628241493448", "Jose da Silva", 1, time.Now().Year(), 123)
	assert.Nil(t, err)
}
