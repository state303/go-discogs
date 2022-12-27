package helper

import (
	"reflect"
	"regexp"
	"strings"
)

var (
	pkRegExp  = regexp.MustCompile(".*primaryKey.*")
	colRegExp = regexp.MustCompile("^column:.*")
)

func ExtractGormPKColumns(i interface{}) []string {
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
				if matched := colRegExp.MatchString(part); !matched {
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
