package gson

import "regexp"

const (
	json_quote        = `"`
	json_escape_quote = "\""
	json_colon        = ":"
	json_brac_l       = "("
	json_brac_r       = ")"
	json_sq_brac_l    = "["
	json_sq_brac_r    = "]"
	json_brace_l      = "{"
	json_brace_r      = "}"
	json_comma        = ","
	whitespace        = " "
)

var (
	json_number_int   = regexp.MustCompile("[-]?([0-9]|[1-9]+)")
	json_number_float = regexp.MustCompile("[-]?([0-9]\\.[0-9]+|[1-9]+\\.[0-9]+)")
	json_key          = regexp.MustCompile("^([a-zA-Z]([a-zA-Z0-9]|[\\-|_])*)+$")
)
