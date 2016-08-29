package usdlregex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * Test AL with valid value
 */
func TestValidateDriversLicenseValidAL(t *testing.T) {
    match, err := ValidateDriversLicense("AL", "1234567")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AL with invalid value (8 numbers and max is 7)
 */
func TestValidateDriversLicenseInvalidAL(t *testing.T) {
    match, err := ValidateDriversLicense("AL", "12345678")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value
 */
func TestValidateDriversLicenseValidAZ(t *testing.T) {
    match, err := ValidateDriversLicense("AZ", "A12345678")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with invalid value (One letter and no numbers)
 */
func TestValidateDriversLicenseInvalidAZNoNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("AZ", "A")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with invalid value (One letter and nine numbers)
 */
func TestValidateDriversLicenseInvalidAZNineNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("AZ", "A123456789")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value (Two letters and three numbers)
 */
func TestValidateDriversLicenseValidAZTwoLettersThreeNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("AZ", "AB123")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test AZ with valid value (No letters and nine numbers)
 */
func TestValidateDriversLicenseValidAZNoLettersNineNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("AZ", "123456789")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test CT with valid value (Nine numbers)
 */
func TestValidateDriversLicenseValidCTNineNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("CT", "123456789")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test CT with valid value (Nine numbers)
 */
func TestValidateDriversLicenseInvalidCTOneLetterEightNumbers(t *testing.T) {
    match, err := ValidateDriversLicense("CT", "A12345678")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test KS with valid value (1 Letter + 1 Number + 1 Letter + 1 Number + 1 Letter)
 */
func TestValidateDriversLicenseValidKSToggleLetterNumber(t *testing.T) {
    match, err := ValidateDriversLicense("KS", "A1B1A")

	assert.True(t, match)
	assert.Nil(t, err)
}

/**
 * Test KS with invalid value (1 Letter + 1 Number + 1 Letter)
 */
func TestValidateDriversLicenseInvalidKSToggleLetterNumber(t *testing.T) {
    match, err := ValidateDriversLicense("KS", "A1B")

	assert.False(t, match)
	assert.Nil(t, err)
}

/**
 * Test invalid state
 */
func TestValidateDriversLicenseInvalidState(t *testing.T) {
	var invalidState string = "ZZ"
    match, err := ValidateDriversLicense(invalidState, "1234")

	assert.False(t, match)
	assert.NotNil(t, err)
	assert.Equal(t, "No rules exist for " + invalidState + "!", err.Error())
}