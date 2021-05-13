package gson

import "fmt"

/*
	Type alias for a JSON object
*/
type JSONObject = map[string]interface{}

/*
	parser object
*/
type parser struct {
	pos int

	/*
	   TODO
	*/
}

/*
	Create a new parser
*/
func newParser() *parser {

	/*
		TODO
	*/

	return &parser{0}
}

func (p *parser) parseNull(s string) {
	/*
		TODO
	*/
}

func (p *parser) parseBool(s string) {
	/*
		TODO
	*/
}

func (p *parser) parseString() {
	/*
		TODO
	*/
}

func (p *parser) parseNumber() {
	/*
		TODO
	*/
}

func (p *parser) parseArray() {
	/*
		TODO
	*/
}

func (p *parser) parseObject() {
	/*
		TODO
	*/
}

func (p *parser) parseKey(s string) *parserError {

	/*
		TODO
	*/

	return &parserError{"Invalid key", s, "", ""}
}

func Parse(s string) (*JSONObject, *parserError) {
	p := newParser()

	/*
		Avoid unused error
		TO BE REMOVED
	*/
	p.pos = 0

	/*
	   TODO
	*/

	fmt.Println("gson is currently UNDER CONSTRUCTION...")
	fmt.Println(" * check the repository @ https://github.com/edrumm/gson for updates")
	fmt.Println(" * v0 coming soon")

	return nil, nil
}

func ToString(jo JSONObject) string {

	/*
		TODO
	*/

	return ""
}
