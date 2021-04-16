package gson

import "fmt"

type parserError struct {
	err string
	loc int
}

func (e *parserError) Error() string {
	return fmt.Sprintf("ParserError: %s at %d", e.err, e.loc)
}
