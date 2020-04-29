package utils

import (
	"reflect"
)

func GetMapIntKeys(m interface{}) []int {

	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		panic("Not a map!")
	}

	valueKeys := v.MapKeys()
	keys := make([]int, len(valueKeys))
	for i := range valueKeys {
		keys[i] = valueKeys[i].Interface().(int)
	}
	return keys
}
