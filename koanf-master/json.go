package koanf

import "encoding/json"

type JSON struct {
}

func Parser() *JSON {
	return &JSON{}
}

func (j *JSON) Unmarshal(b []byte) (map[string]interface{}, error) {
	var out map[string]interface{}

	err := json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
