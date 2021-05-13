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
	jsonNumberInt   = regexp.MustCompile("[-]?([0-9]|[1-9]+)")
	jsonNumberFloat = regexp.MustCompile("[-]?([0-9]\\.[0-9]+|[1-9]+\\.[0-9]+)")
	jsonKey         = regexp.MustCompile("^([a-zA-Z]([a-zA-Z0-9]|[\\-|_])*)+$")
	typeMapping     = map[string]string{
		"number":  "float64",
		"string":  "string",
		"boolean": "bool",
		"null":    "nil",
		"array":   "[]interface{}",
		"object":  "map[string]interface{}",
	}
)
