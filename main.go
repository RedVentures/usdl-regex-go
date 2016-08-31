package usdl

import (
	"errors"
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
	// Assume no rule will be found
	var isRuleFound = false

	// Grab the array of rules based on the state code
	// and iterate over them
	for _, rule := range rules[stateCode] {
		isRuleFound = true

		// See if a match exists for this rule
		match, err = rule.MatchString(licenseNumber), nil

		// Return if an error or match is found
		// Otherwise try the next rule
		if err != nil || match {
			return
		}
	}

	// Check to see if no rule was found
	if !isRuleFound {
		// If no match is found for any rules return
		// an error indicating this
		return isRuleFound, ErrorNoRules
	}

	// No match was found
	return false, nil
}
