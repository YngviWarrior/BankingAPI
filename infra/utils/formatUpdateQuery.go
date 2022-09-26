package utils

import (
	"fmt"
	"reflect"
)

func FormatUpdateQuery(field []string, fieldValue []any, column []string, values []any) (string, string) {
	var updates string
	for i, v := range field {
		for j, val := range fieldValue {
			if i == j {
				t := reflect.TypeOf(val).Kind()
				if t == reflect.String {
					updates += fmt.Sprintf(`%v = "%v",`, v, val)
				} else {
					updates += fmt.Sprintf(`%v = %v,`, v, val)
				}
			}
		}
	}

	updates = updates[:len(updates)-1]

	var wheres string
	for i, v := range column {
		for j, val := range values {
			if i == j {
				t := reflect.TypeOf(val).Kind()
				if t == reflect.String {
					wheres += fmt.Sprintf(`%v = "%v",`, v, val)
				} else {
					wheres += fmt.Sprintf(`%v = %v,`, v, val)
				}
			}
		}
	}

	wheres = wheres[:len(wheres)-1]

	return updates, wheres
}
