package usdl

import (
	"testing"
)

/**
 * Test validation logic with valid
 * states and drivers license numbers
 */
func TestValidateValidValues(t *testing.T) {
	dls := []struct {
		state string
		dl    string
	}{
		{"AL", "1234567"},
		{"AZ", "A12345678"},
		// Two letters and three numbers
		{"AZ", "AB123"},
		// No letters and nine numbers
		{"AZ", "AB123"},
		// Nine numbers
		{"CT", "123456789"},
		// 1 Letter + 1 Number + 1 Letter + 1 Number + 1 Letter
		{"KS", "A1B1A"},
		// add more states here
	}

	for _, d := range dls {
		match, err := Validate(d.state, d.dl)
		if !match {
			t.Error("%s did not validate", d.state)
		}

		if err != nil {
			t.Error("%s had an error: %s", d.state, err)
		}
	}
}

/**
 * Test validation logic with invalid
 * states and drivers license numbers
 */
func TestValidateInvalidValues(t *testing.T) {
	dls := []struct {
		state string
		dl    string
	}{
		{"AL", "12345678"},
		// One letter and no numbers
		{"AZ", "A"},
		// One letter and nine numbers
		{"AZ", "A123456789"},
		// One letter eight numbers
		{"CT", "A12345678"},
		// 1 Letter + 1 Number + 1 Letter
		{"KS", "A1B"},
	}

	for _, d := range dls {
		match, err := Validate(d.state, d.dl)
		if match {
			t.Error("%s should not have found a match", d.state)
		}

		if err != nil {
			t.Error("%s had an error: %s", d.state, err)
		}
	}
}

/**
 * Test invalid state
 */
func TestValidateInvalidState(t *testing.T) {
	var invalidState string = "ZZ"
	match, err := Validate(invalidState, "1234")

	if match {
		t.Error("Should be false")
	}

	if err == nil {
		t.Error("Expected value not to be nil")
	}

	if ErrorNoRules != err {
		t.Error("Expected no rules to be found")
	}
}
