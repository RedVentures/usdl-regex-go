package usdl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Test AL with valid value
 */
func TestValidateValidAL(t *testing.T) {
    match, err := Validate("AL", "1234567")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AL with invalid value (8 numbers and max is 7)
 */
func TestValidateInvalidAL(t *testing.T) {
    match, err := Validate("AL", "12345678")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value
 */
func TestValidateValidAZ(t *testing.T) {
    match, err := Validate("AZ", "A12345678")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with invalid value (One letter and no numbers)
 */
func TestValidateInvalidAZNoNumbers(t *testing.T) {
    match, err := Validate("AZ", "A")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with invalid value (One letter and nine numbers)
 */
func TestValidateInvalidAZNineNumbers(t *testing.T) {
    match, err := Validate("AZ", "A123456789")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value (Two letters and three numbers)
 */
func TestValidateValidAZTwoLettersThreeNumbers(t *testing.T) {
    match, err := Validate("AZ", "AB123")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value (No letters and nine numbers)
 */
func TestValidateValidAZNoLettersNineNumbers(t *testing.T) {
    match, err := Validate("AZ", "123456789")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test CT with valid value (Nine numbers)
 */
func TestValidateValidCTNineNumbers(t *testing.T) {
    match, err := Validate("CT", "123456789")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test CT with valid value (Nine numbers)
 */
func TestValidateInvalidCTOneLetterEightNumbers(t *testing.T) {
    match, err := Validate("CT", "A12345678")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test KS with valid value (1 Letter + 1 Number + 1 Letter + 1 Number + 1 Letter)
 */
func TestValidateValidKSToggleLetterNumber(t *testing.T) {
    match, err := Validate("KS", "A1B1A")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test KS with invalid value (1 Letter + 1 Number + 1 Letter)
 */
func TestValidateInvalidKSToggleLetterNumber(t *testing.T) {
    match, err := Validate("KS", "A1B")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test invalid state
 */
func TestValidateInvalidState(t *testing.T) {
	var invalidState string = "ZZ"
    match, err := Validate(invalidState, "1234")

	assert.False(t, match)
	assert.NotNil(t, err)
	assert.Equal(t, "No rules exist for " + invalidState + "!", err.Error())
}