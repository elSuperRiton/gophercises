package yaml

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// parseYAML takes in a slice of byte and attempts to parse it
// into a []map[string]string using the yaml.v2 package
func parse(in []byte) ([]map[string]string, error) {
	var u []map[string]string
	err := yaml.Unmarshal(in, &u)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML : %v", err)
	}

	return u, nil
}
