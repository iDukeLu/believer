package util

func Ternary(expression bool, result1 interface{}, result2 interface{}) interface{} {
	if expression {
		return result1
	} else {
		return result2
	}
}

func GetIntDefault(value1 int, value2 int) int {
	return Ternary(value1 == 0, value1, value2).(int)
}

func GetFloat64Default(value1 float64, value2 float64) float64 {
	return Ternary(value1 == 0, value1, value2).(float64)
}

func GetStringDefault(value1 string, value2 string) string {
	return Ternary(value1 == "", value1, value2).(string)
}
