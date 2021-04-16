package gson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type JSONObject = map[string]interface{}

type Parser struct {
	/*
	   TODO
	*/
}

func flatten(str string) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.Compact(buf, []byte(str)); err != nil {
		return nil, &parserError{err.Error(), 0}
	}

	return buf.Bytes(), nil
}

func Parse(str string) (JSONObject, error) {
	obj := make(JSONObject)

	// flatten(str)

	/*
	   TODO
	*/

	fmt.Println("Under Construction...")

	return obj, nil
}
