package gson

import "fmt"

type parserError struct {
	err string
	loc int
}

func (e *parserError) Error() string {
	if e.loc == -1 {
		return fmt.Sprintf("Error while formatting: %s", e.err)
	}

	return fmt.Sprintf("ParserError: %s at %d", e.err, e.loc)
}
