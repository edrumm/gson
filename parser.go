package gson

import (
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

func parseSequence(input, expect string) (bool, *parserError) {
	ex := []rune(expect)

	if len(input) != len(expect) {
		return false, &parserError{"Bad symbol", expect, input}
	}

	for i, c := range input {
		if c != ex[i] {
			return false, &parserError{"Bad symbol", expect, input}
		}
	}

	return true, nil
}

func parseNull() {
	/*
		TODO
	*/
}

func parseBool() {
	/*
		TODO
	*/
}

func parseString() {
	/*
		TODO
	*/
}

func parseNumber() {
	/*
		TODO
	*/
}

func parseArray() {
	/*
		TODO
	*/
}

func parseObject() {
	/*
		TODO
	*/
}

func parseKey(key string) bool {
	/*
		TODO
	*/

	return JSON_Key.MatchString(key)
}

func (p *Parser) Parse(str string) (JSONObject, *parserError) {
	obj := make(JSONObject)

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
