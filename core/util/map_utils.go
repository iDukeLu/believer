package util

func GetMapKeys(m map[interface{}]interface{}) []interface{} {
	keys := make([]interface{}, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetMapValues(m map[interface{}]interface{}) []interface{} {
	values := make([]interface{}, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
