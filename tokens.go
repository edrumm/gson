package gson

import "regexp"

const (
	JSON_Quote        = `"`
	JSON_Escape_Quote = "\""
	JSON_Colon        = ":"
	JSON_Brac_L       = "("
	JSON_Brac_R       = ")"
	JSON_Sq_Brac_L    = "["
	JSON_Sq_Brac_R    = "]"
	JSON_Brace_L      = "{"
	JSON_Brace_R      = "}"
	JSON_Comma        = ","
	Whitespace        = " "
)

var (
	JSON_Number_Int   = regexp.MustCompile("[0-9]|[1-9]+")
	JSON_Number_Float = regexp.MustCompile("[-]?([0-9]\\.[0-9]+|[1-9]+\\.[0-9]+)")
	JSON_Key          = regexp.MustCompile("^([a-zA-Z]([a-zA-Z0-9]|[\\-|_])*)+$")
)
