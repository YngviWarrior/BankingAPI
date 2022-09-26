package utils

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func QueryFormatPagination(values []any) (pagination string, err error) {
	pagination = "LIMIT "

	for _, val := range values {
		t := reflect.TypeOf(val).Kind()
		if t != reflect.Int {
			err = errors.New("pagination value is not an integer")
			return
		} else {
			pagination += fmt.Sprintf(`%v,`, val)
		}
	}

	pagination = pagination[:len(pagination)-1]

	return pagination, err
}

func QueryFormatWheres(column []string, values []any) string {
	var wheres string = "WHERE "
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

	return wheres
}

func QueryFormatUpdates(field []string, fieldValue []any) string {
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

	return updates
}

func QueryFormatter(field []string, fieldValue []any, columns []string, values []any, pageValues []any, order string) (paginations string, wheres string, updates string) {
	var errList []error
	var err error
	if len(pageValues) > 0 {
		paginations, err = QueryFormatPagination(pageValues)

		if err != nil {
			errList = append(errList, err)
		}
	}

	if len(columns) > 0 && len(values) > 0 && len(columns) == len(values) {
		wheres = QueryFormatWheres(columns, values)
	} else if len(columns) != len(values) {
		errList = append(errList, errors.New("wrong parameters at query wheres formatter"))
	}

	if len(field) > 0 && len(fieldValue) > 0 && len(field) == len(fieldValue) {
		updates = QueryFormatWheres(field, fieldValue)
	} else if len(field) != len(fieldValue) {
		errList = append(errList, errors.New("wrong parameters at query updates formatter"))
	}

	if len(errList) > 0 {
		log.Printf("CRLBC 01: %v", errList)
	}

	return
}
