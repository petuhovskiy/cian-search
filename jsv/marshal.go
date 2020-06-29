package jsv

import (
	"encoding/json"
	"fmt"
)

func Marshal(obj interface{}) (map[string]string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	err = traverse("", data, result)
	return result, err
}

func traverse(prefix string, data map[string]interface{}, result map[string]string) error {
	for k, v := range data {
		if nextData, ok := v.(map[string]interface{}); ok {
			err := traverse(prefix+k+".", nextData, result)
			if err != nil {
				return err
			}
			continue
		}

		realKey := prefix + k
		realValue := fmt.Sprintf("%v", v)

		if _, ok := result[realKey]; ok {
			return fmt.Errorf("duplicate key found: %v", realKey)
		}

		result[realKey] = realValue
	}

	return nil
}
