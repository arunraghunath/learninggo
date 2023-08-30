package yaml

import "gopkg.in/yaml.v3"

type YAML struct {
}

func Parser() *YAML {
	return &YAML{}
}

func (pa *YAML) Unmarshal(b []byte) (map[string]interface{}, error) {
	var out map[string]interface{}
	err := yaml.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
