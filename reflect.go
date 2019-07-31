package dbid

import (
	"errors"
	"reflect"
)

var refLocType = reflect.TypeOf((*SchemaLocator)(nil)).Elem()

// extract method extracts schema information from provided interface
func extract(target interface{}) (schema string, single bool, err error) {
	if target == nil {
		err = errors.New("empty receiver")
	}

	refType := reflect.TypeOf(target)
	if refType.Kind() != reflect.Ptr {
		err = errors.New("expected pointer receiver")
		return
	}
	refType = refType.Elem()

	if refKind := refType.Kind(); refKind == reflect.Slice || refKind == reflect.Array {
		refType = refType.Elem()
	} else {
		single = true
	}

	if !refType.Implements(refLocType) {
		err = errors.New("unsupported type")
		return
	}

	schema = (reflect.New(refType)).Interface().(SchemaLocator).Schema()

	return
}
