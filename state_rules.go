package usdlregex

type StateRules struct {
	Rules			Rules		`json:"rules"`
}

type Rules []struct {
	State 			string 		`json:"state"`
    Regex 			string 		`json:"regex"`	
}
