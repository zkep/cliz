package validator

const (
	defaultRequiredMsg = "field is required"
	defaultRangeMsg    = "must be between %v and %v"
	defaultLenMsg      = "must be exactly %v"
	defaultPatternMsg  = "must match pattern '%v'"
	defaultInMsg       = "must be one of %v"
	defaultEqMsg       = "must equal '%v'"
	defaultGtMsg       = "must be greater than %v"
	defaultLtMsg       = "must be less than %v"
	defaultContainsMsg = "must contain '%v'"
	defaultEmailMsg    = "must be a valid email address"
	defaultURLMsg      = "must be a valid URL"
	defaultAlphaMsg    = "must contain only alphabetic characters"
	defaultAlphanumMsg = "must contain only alphanumeric characters"
)

const (
	formatItems   = "%.0f items"
	formatInteger = "%.0f"
	formatFloat   = "%.2f"
	formatBoolean = "%t"
)
