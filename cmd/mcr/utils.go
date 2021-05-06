package main

import (
	"encoding/json"
	"os"
)

func printJSON(object interface{}) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "   ")
	return encoder.Encode(object)
}
