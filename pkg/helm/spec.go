package helm

import (
	"gopkg.in/yaml.v3"
)

// GetValuesMap
//
// Return the mapped out values of a chart
func (spec *ChartSpec) GetValuesMap() (map[string]interface{}, error) {
	var values map[string]interface{}

	err := yaml.Unmarshal([]byte(spec.ValuesYaml), &values)
	if err != nil {
		return nil, err
	}

	return values, nil
}
