package tools

import (
	"reflect"
	"strings"
)

func TagNameFormatter(fld reflect.StructField) string {
	var name string

	jsonTag := fld.Tag.Get("json")
	formTag := fld.Tag.Get("form")
	queryTag := fld.Tag.Get("query")
	paramTag := fld.Tag.Get("param")

	if jsonTag != "" {
		name = strings.SplitN(jsonTag, ",", 2)[0]
	} else if formTag != "" {
		name = strings.SplitN(formTag, ",", 2)[0]
	} else if queryTag != "" {
		name = strings.SplitN(queryTag, ",", 2)[0]
	} else if paramTag != "" {
		name = strings.SplitN(paramTag, ",", 2)[0]
	}

	if name == "-" {
		return ""
	}

	return name
}
