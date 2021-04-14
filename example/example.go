package main

import (
	"fmt"
	"github.com/edrumm/gson"
)

func main() {
	_, err := gson.Parse("Hello")

	if err != nil {
		fmt.Println(err.Error())
	}
}
