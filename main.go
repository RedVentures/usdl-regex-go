package usdl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"regexp"
)

var stateRules StateRules
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
	stateRulesFile, err := ioutil.ReadFile("stateRegex.json")

	if err != nil {
		return
	}

	err = json.Unmarshal(stateRulesFile, &stateRules)

	if err != nil {
		return
	}

	regex, err := getRegex(stateCode)

	if err != nil {
		return
	}

	match, err = regexp.MatchString(regex, licenseNumber)

	return
}

/**
 * Helper method used to get the regex based on state code
 *
 * @param stateCode string The state code for the drivers license
 *
 * @return A regex string for a state or any potential error found
 *
 */
func getRegex(stateCode string) (string, error) {
	var err error

	for _, rule := range stateRules.Rules {
		if stateCode == rule.State {
			return rule.Regex, err
		}
	}

	// No rules found
	return "", ErrorNoRules
}
