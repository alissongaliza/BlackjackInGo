package utils

import (
	"fmt"
	"reflect"
)

func GetMapIntKeys(m interface{}) ([]int, error) {

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return nil, fmt.Errorf("Not a map!")
	}

	valueKeys := v.MapKeys()
	keys := make([]int, len(valueKeys))
	for i := range valueKeys {
		keys[i] = valueKeys[i].Interface().(int)
	}
	return keys, nil
}
