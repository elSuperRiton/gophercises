package json

import (
	"encoding/json"
	"fmt"
)

// parseJSON takes in a slice of byte and attempts to parse it
// into a []map[string]string using the standard golang JSON package
func parse(in []byte) ([]map[string]string, error) {
	var u []map[string]string
	err := json.Unmarshal(in, &u)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON : %v", err)
	}

	return u, nil
}
