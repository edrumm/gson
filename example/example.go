package main

import (
	"fmt"
	"github.com/edrumm/gson"
)

func main() {
	// create a new parser
	parser := gson.NewParser()

	// parse from string
	// (struct JSONObject, Error)
	obj, err := parser.Parse(`"language": "Go", "best": true`)

	if err != nil {
		fmt.Println(err.Error())
	}

	/* json object to string
	str := */parser.ToString(obj)
}
