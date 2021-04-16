package gson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
	JSON object struct
	represents a key : value pair
*/
type JSONObject = map[string]interface{}

/*
	Parser struct object
*/
type Parser struct {
	/*
	   TODO
	*/
}

/*
	Create a new parser
*/
func NewParser() *Parser {

	/*
		TODO
	*/

	return &Parser{}
}

/*
	Removes whitespace and newlines
	Converts to []byte for easier parsing
*/
func flatten(str string) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.Compact(buf, []byte(str)); err != nil {
		return nil, &parserError{err.Error(), -1}
	}

	return buf.Bytes(), nil
}

func (p *Parser) Parse(str string) (JSONObject, error) {
	obj := make(JSONObject)

	// tokens, err := flatten(str)

	/*
	   TODO
	*/

	fmt.Println("Under Construction...")

	return obj, nil
}

func (p *Parser) ToString(obj JSONObject) string {

	/*
		TODO
	*/

	return ""
}
