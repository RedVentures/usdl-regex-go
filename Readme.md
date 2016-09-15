# USDL-REGEX-GO

USDL-REGEX-GO is a driver's license regex validation library.  It is a golang port from the USDLRegex [github.com/adambullmer/USDLRegex](https://github.com/adambullmer/USDLRegex) PHP project.  Like USDLRegex, it is sourced from https://ntsi.com/drivers-license-format/

## Use cases

Anytime you need to validate a driver's license!

## Usage

The state_rules.json included in the library contains all of the regex rules for each state.  All you need to do is import this library and provide a state number along with a drivers license number

Example:

Import:

```
import (
	"github.com/redventures/usdl-regex-go"
)
```

In Use:

```
match, err = usdl.Validate(stateCode, licenseNumber)
```
