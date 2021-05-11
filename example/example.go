package main

import "github.com/edrumm/gson"

func main() {
	// parse from string
	// (map[string]interface{}, *parserError)
	obj, err := gson.Parse(`"language": "Go", "best": true`)

	if err != nil {
		// error message
		// (string)
		err.Error()

		// print error message to os.Stderr
		err.PrintError()
	}

	/* json object to string
	   (string)
	str := */gson.ToString(obj)
}
