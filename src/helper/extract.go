package helper

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var pkRegExp = regexp.MustCompile(".*primaryKey.*")

func ExtractGormPKColumns(i interface{}) []string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("failed during... %+v\n", reflect.TypeOf(i))
		}
	}()

	val := reflect.TypeOf(i)

	columns := make([]string, 0)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	fields := reflect.VisibleFields(val)
	for _, field := range fields {
		if tag, ok := field.Tag.Lookup("gorm"); !ok || !pkRegExp.MatchString(tag) {
			continue
		} else {
			for _, part := range strings.Split(tag, ";") {
				if matched, _ := regexp.MatchString("^column:.*", part); !matched {
					continue
				}
				columnName := strings.ReplaceAll(part, "column:", "")
				if columnName != "updated_at" {
					columns = append(columns, columnName)
				}
			}
		}
	}
	return columns
}