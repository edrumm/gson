package gson

import "fmt"

/*
	JSON object struct
*/
type JSONObject = map[string]interface{}

/*
	Parser object
*/
type Parser struct {
	pos int

	/*
	   TODO
	*/
}

/*
	Create a new parser
*/
func newParser() *Parser {

	/*
		TODO
	*/

	return &Parser{0}
}

func (p *Parser) parseNull(s string) (bool, *parserError) {

	if s == "null" {
		p.pos += 4
		return true, nil
	}

	return false, &parserError{"Bad symbol", s, "null", s}
}

func (p *Parser) parseBool(s string) (bool, *parserError) {

	if s == "true" {
		p.pos += 4
		return true, nil

	} else if s == "false" {
		p.pos += 5
		return false, nil
	}

	return false, &parserError{"Bad symbol", s, "[true|false]", s}
}

func (p *Parser) parseString() {
	/*
		TODO
	*/
}

func (p *Parser) parseNumber() {
	/*
		TODO
	*/
}

func (p *Parser) parseArray() {
	/*
		TODO
	*/
}

func (p *Parser) parseObject() {
	/*
		TODO
	*/
}

func (p *Parser) parseKey(s string) *parserError {

	if JSON_Key.MatchString(s) {
		p.pos += len(s)
		return nil
	}

	return &parserError{"Invalid key", s, "", ""}
}

func Parse(s string) (JSONObject, *parserError) {
	p := newParser()
	p.pos = 0

	/*
	   TODO
	*/

	fmt.Println("Under construction...")

	return nil, nil
}

func ToString(o JSONObject) string {

	/*
		TODO
	*/

	return ""
}
