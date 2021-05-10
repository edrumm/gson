package main

import (
	"github.com/edrumm/gson"
)

func main() {
	// create a new parser
	// (* Parser)
	parser := gson.NewParser()

	// parse from string
	// (map[string]interface{}, *parserError)
	obj, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		// error message
		// (string)
		err.Error()

		// print error message to os.Stderr
		err.PrintError()
	}

	/* json object to string
	str := */parser.ToString(obj)
}
