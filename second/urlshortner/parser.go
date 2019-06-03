package urlshortner

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type urls []map[string]string

// parseYAML takes in a slice of byte and attempts to parse it
// into a []map[string]string using the yaml.v2 package
func parseYAML(in []byte) (urls, error) {
	var u urls
	err := yaml.Unmarshal(in, &u)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML : %v", err)
	}

	return u, nil
}

// parseJSON takes in a slice of byte and attempts to parse it
// into a []map[string]string using the standard golang JSON package
func parseJSON(in []byte) (urls, error) {
	var u urls
	err := json.Unmarshal(in, &u)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON : %v", err)
	}

	return u, nil
}

// buildMap takes in a []map[string]string and maps
// the short path values to their corresponding full path
func buildMap(u urls) map[string]string {
	mappedURLS := make(map[string]string)
	for _, url := range u {
		mappedURLS[url["path"]] = url["url"]
	}

	return mappedURLS
}
