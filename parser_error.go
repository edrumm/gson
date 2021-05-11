package gson

import (
	"fmt"
	"os"
)

type parserError struct {
	err          string
	invalidToken string
	expect       string
	actual       string
}

func (e parserError) Error() string {
	if e.expect == "" || e.actual == "" {
		return fmt.Sprintf("%s at token %s", e.err, e.invalidToken)
	}

	return fmt.Sprintf("%s at token %s, expected %s but found %s", e.err, e.invalidToken, e.expect, e.actual)
}

/*
	Prints error to stderr
*/
func (e parserError) PrintError() {
	fmt.Fprintf(os.Stderr, "%s\n", e.Error())
}
