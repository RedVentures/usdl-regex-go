package usdl

import (
	"regexp"
)

var rules = map[string]*regexp.Regexp{
	// 1-7 Numeric
	"AL": regexp.MustCompile("^[0-9]{1,7}$"),
	// 1-7 Numeric
	"AK": regexp.MustCompile("^[0-9]{1,7}$"),
	// 1 Alpha + 1-8 Numeric
	// 2 Alpha + 2-5 Numeric
	// 9 Numeric
	"AZ": regexp.MustCompile("(^[A-Z]{1}[0-9]{1,8}$)|(^[A-Z]{2}[0-9]{2,5}$)|(^[0-9]{9}$)"),
	// 4-9 Numeric
	"AR": regexp.MustCompile("^[0-9]{4,9}$"),
	// 1 Alpha + 7 Numeric
	"CA": regexp.MustCompile("^[A-Z]{1}[0-9]{7}$"),
	// 9 Numeric
	// 1 Alpha + 3-6 Numeric
	// 2 Alpha + 2-5 Numeric
	"CO": regexp.MustCompile("(^[0-9]{9}$)|(^[A-Z]{1}[0-9]{3,6}$)|(^[A-Z]{2}[0-9]{2,5}$)"),
	// 9 Numeric
	"CT": regexp.MustCompile("^[0-9]{9}$"),
	// 1-7 Numeric
	"DE": regexp.MustCompile("^[0-9]{1,7}$"),
	// 9 Numeric
	"DC": regexp.MustCompile("(^[0-9]{7}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 12 Numeric
	"FL": regexp.MustCompile("^[A-Z]{1}[0-9]{12}$"),
	// 7-9 Numeric
	"GA": regexp.MustCompile("^[0-9]{7,9}$"),
	// 1 Alpha + 14 Numeric
	"GU": regexp.MustCompile("^[A-Z]{1}[0-9]{14}$"),
	// 1 Alpha + 8 Numeric
	// 9 Numeric
	"HI": regexp.MustCompile("(^[A-Z]{1}[0-9]{8}$)|(^[0-9]{9}$)"),
	// 2 Alpha + 6 Numeric + 1 Alpha
	// 9 Numeric
	"ID": regexp.MustCompile("(^[A-Z]{2}[0-9]{6}[A-Z]{1}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 11-12 Numeric
	"IL": regexp.MustCompile("^[A-Z]{1}[0-9]{11,12}$"),
	// 1 Alpha + 9 Numeric
	// 9-10 Numeric
	"IN": regexp.MustCompile("(^[A-Z]{1}[0-9]{9}$)|(^[0-9]{9,10}$)"),
	// 9 Numeric
	// 3 Numeric + 2 Alpha + 4 Numeric
	"IA": regexp.MustCompile("^([0-9]{9}|([0-9]{3}[A-Z]{2}[0-9]{4}))$"),
	// 1 Alpha + 1 Numeric + 1 Alpha + 1 Numeric + 1 Alpha
	// 1 Alpha + 8 Numeric
	// 9 Numeric
	"KS": regexp.MustCompile("(^([A-Z]{1}[0-9]{1}){2}[A-Z]{1}$)|(^[A-Z]{1}[0-9]{8}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 8-9 Numeric
	// 9 Numeric
	"KY": regexp.MustCompile("(^[A_Z]{1}[0-9]{8,9}$)|(^[0-9]{9}$)"),
	// 1-9 Numeric
	"LA": regexp.MustCompile("^[0-9]{1,9}$"),
	// 7-8 Numeric
	// 7 Numeric + 1 Alpha
	"ME": regexp.MustCompile("(^[0-9]{7,8}$)|(^[0-9]{7}[A-Z]{1}$)"),
	// 1Alpha+12Numeric
	"MD": regexp.MustCompile("^[A-Z]{1}[0-9]{12}$"),
	// 1 Alpha + 8 Numeric
	// 9 Numeric
	"MA": regexp.MustCompile("(^[A-Z]{1}[0-9]{8}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 10 Numeric
	// 1 Alpha + 12 Numeric
	"MI": regexp.MustCompile("(^[A-Z]{1}[0-9]{10}$)|(^[A-Z]{1}[0-9]{12}$)"),
	// 1 Alpha + 12 Numeric
	"MN": regexp.MustCompile("^[A-Z]{1}[0-9]{12}$"),
	// 9 Numeric
	"MS": regexp.MustCompile("^[0-9]{9}$"),
	// 1 Alpha + 5-9 Numeric
	// 1 Alpha + 6 Numeric + 'R'
	// 8 Numeric + 2 Alpha
	// 9 Numeric + 1 Alpha
	// 9 Numeric
	"MO": regexp.MustCompile("(^[A-Z]{1}[0-9]{5,9}$)|(^[A-Z]{1}[0-9]{6}[R]{1}$)|(^[0-9]{8}[A-Z]{2}$)|(^[0-9]{9}[A-Z]{1}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 8 Numeric
	// 13 Numeric
	// 9 Numeric
	// 14 Numeric
	"MT": regexp.MustCompile("(^[A-Z]{1}[0-9]{8}$)|(^[0-9]{13}$)|(^[0-9]{9}$)|(^[0-9]{14}$)"),
	// 1-7 Numeric
	"NE": regexp.MustCompile("^[0-9]{1,7}$"),
	// 9 Numeric
	// 10 Numeric
	// 12 Numeric
	// 'X' + 8 Numeric
	"NV": regexp.MustCompile("(^[0-9]{9,10}$)|(^[0-9]{12}$)|(^[X]{1}[0-9]{8}$)"),
	// 2 Numeric + 3 Alpha + 5 Numeric
	"NH": regexp.MustCompile("^[0-9]{2}[A-Z]{3}[0-9]{5}$"),
	// 1 Alpha + 14 Numeric
	"NJ": regexp.MustCompile("^[A-Z]{1}[0-9]{14}$"),
	// 8 Numeric
	// 9 Numeric
	"NM": regexp.MustCompile("^[0-9]{8,9}$"),
	// 1 Alpha + 7 Numeric
	// 1 Alpha + 18 Numeric
	// 8 Numeric
	// 9 Numeric
	// 16 Numeric
	// 8 Alpha
	"NY": regexp.MustCompile("(^[A-Z]{1}[0-9]{7}$)|(^[A-Z]{1}[0-9]{18}$)|(^[0-9]{8}$)|(^[0-9]{9}$)|(^[0-9]{16}$)|(^[A-Z]{8}$)"),
	// 1-12 Numeric
	"NC": regexp.MustCompile("^[0-9]{1,12}$"),
	// 3 Alpha + 6 Numeric
	// 9 Numeric
	"ND": regexp.MustCompile("(^[A-Z]{3}[0-9]{6}$)|(^[0-9]{9}$)"),
	// 1 Alpha + 4-8 Numeric
	// 2 Alpha + 3-7 Numeric
	// 8 Numeric
	"OH": regexp.MustCompile("(^[A-Z]{1}[0-9]{4,8}$)|(^[A-Z]{2}[0-9]{3,7}$)|(^[0-9]{8}$)"),
	// 1 Alpha + 9 Numeric
	// 9 Numeric
	"OK": regexp.MustCompile("(^[A-Z]{1}[0-9]{9}$)|(^[0-9]{9}$)"),
	// 1-9 Numeric
	"OR": regexp.MustCompile("^[0-9]{1,9}$"),
	// 8 Numeric
	"PA": regexp.MustCompile("^[0-9]{8}$"),
	// 5-7 Numeric
	// 9 Numeric
	"PR": regexp.MustCompile("(^[0-9]{9}$)|(^[0-9]{5,7}$)"),
	// 7 Numeric
	// 1 Alpha + 6 Numeric
	"RI": regexp.MustCompile("^([0-9]{7}$)|(^[A-Z]{1}[0-9]{6}$)"),
	// 5-11 Numeric
	"SC": regexp.MustCompile("^[0-9]{5,11}$"),
	// 6-10 Numeric
	// 12 Numeric
	"SD": regexp.MustCompile("(^[0-9]{6,10}$)|(^[0-9]{12}$)"),
	// 7-9 Numeric
	"TN": regexp.MustCompile("^[0-9]{7,9}$"),
	// 7-8 Numeric
	"TX": regexp.MustCompile("^[0-9]{7,8}$"),
	// 4-10 Numeric
	"UT": regexp.MustCompile("^[0-9]{4,10}$"),
	// 8 Numeric
	// 7 Numeric + 'A'
	"VT": regexp.MustCompile("(^[0-9]{8}$)|(^[0-9]{7}[A]$)"),
	// 1 Alpha + 8 Numeric
	// 1 Alpha + 9 Numeric
	// 1 Alpha + 10 Numeric
	// 1 Alpha + 11 Numeric
	// 9 Numeric
	"VA": regexp.MustCompile("(^[A-Z]{1}[0-9]{8,11}$)|(^[0-9]{9}$)"),
	// 1-7 Alpha + any combination of Alpha, Numeric, and * for a total of 12 characters
	// "WA": regexp.MustCompile("^(?=.{12}$)[A-Z]{1,7}[A-Z0-9\\*]{4,11}$"),
	// 7 Numeric
	// 1-2 Alpha + 5-6 Numeric
	"WV": regexp.MustCompile("(^[0-9]{7}$)|(^[A-Z]{1,2}[0-9]{5,6}$)"),
	// 1 Alpha + 13 Numeric
	"WI": regexp.MustCompile("^[A-Z]{1}[0-9]{13}$"),
	// 9-10 Numeric
	"WY": regexp.MustCompile("^[0-9]{9,10}$"),
}
