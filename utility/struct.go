package utility

import (
	"fmt"
	"reflect"
	"strings"
)

// массив строк имен структуры в input
// exclude строка со списком через запятую имен полей исключаемых
// prefix строка префикс для генерируемых имен например "rule."
// для sqlboiler исключать надо те поля которые в связанных таблицах могут быть NULL но по
// типу они обязательны и когда BIND происходит вылетает ошибка
func StructFieldNames(input interface{}, exclude string, prefix string) (out []string) {
	rValue := reflect.ValueOf(input)
	rType := rValue.Type()
	if rType.Kind() == reflect.Struct {
		out = make([]string, 0, rType.NumField())
		for i := 0; i < rType.NumField(); i++ {
			// fld := rType.Field(i)
			val, ok := rValue.Field(i).Interface().(string)
			if ok {
				if exclude != "" && strings.Contains(exclude, val) {
					// пропускаем если есть
					continue
				}
				s := fmt.Sprintf("%s%v", prefix, val)
				out = append(out, s)
			}
		}
	}
	return out
}
