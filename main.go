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
	var r *regexp.Regexp
	var ok bool

	if r, ok = rules[stateCode]; !ok {
		return false, ErrorNoRules
	}

	return r.MatchString(licenseNumber), nil
}
