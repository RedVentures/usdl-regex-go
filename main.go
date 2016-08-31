package usdl

import (
	"errors"
	"regexp"
)

var ErrorNoRules = errors.New("no rules found")

/**
 * Validate a drivers license
 *
 * @param stateCode string The state code for the drivers license
 * @param licenseNumber string The drivers license number
 *
 * @return A boolean indicating if the drivers license is valid or any potential error found
 *
 */
func Validate(stateCode string, licenseNumber string) (match bool, err error) {
	var sr []*regexp.Regexp
	var ok bool

	// Check if any rules found
	if sr, ok = rules[stateCode]; !ok {
		// No rules found, send error indicating this
		return false, ErrorNoRules
	}

	// Iterate over rules to see if the license number
	// does not match the regex
	for _, r := range sr {
		if !r.MatchString(licenseNumber) {
			return false, nil
		}
	}

	// license number matched the regex for the state
	return true, nil
}
