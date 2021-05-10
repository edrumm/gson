package gson

import (
	"fmt"
	"os"
)

type parserError struct {
	err    string
	expect string
	actual string
}

func (e parserError) Error() string {
	if e.expect == "" || e.actual == "" {
		return fmt.Sprintf("Error while formatting: %s", e.err)
	}

	return fmt.Sprintf("ParserError: %s, expected %s but found %s", e.err, e.expect, e.actual)
}

/*
	Prints error to stderr
*/
func (e parserError) PrintError() {
	fmt.Fprintf(os.Stderr, "%s\n", e.Error())
}
