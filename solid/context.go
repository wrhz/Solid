package solid

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/goccy/go-reflect"
	"github.com/gorilla/mux"
)

type Context struct {
	Writer http.ResponseWriter
	Request *http.Request
}

func parseType (value string, targetType reflect.Kind) (any, error) {
	var err error
	var r any

	switch targetType {
	case reflect.String:
		r = value
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r, err = strconv.Atoi(value)
	case reflect.Float32, reflect.Float64:
		r, err = strconv.ParseFloat(value, 64)
	case reflect.Bool:
		r, err = strconv.ParseBool(value)
	default:
		err = fmt.Errorf("unsupported type: %v", targetType)
	}

	return r, err
}

func (c *Context) BindQuery(s any) error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("BindQuery: expected struct, got %v", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		pathTag := field.Tag.Get("path")

		if pathTag != "" {
			paramValue := c.Request.URL.Query()[pathTag]

			paramType := field.Type.Kind()

			if paramType != reflect.Slice {
				if len(paramValue) == 0 {
					continue
				}
				
				data, err := parseType(paramValue[0], paramType)

				if err != nil {
					return fmt.Errorf("parse field %q as %v: %w", pathTag, paramType, err)
				}

				v.Field(i).Set(reflect.ValueOf(data))
			} else {
				elementType := field.Type.Elem().Kind()

				sliceValue := reflect.MakeSlice(field.Type, 0, len(paramValue))

				for _, valStr := range paramValue {
					elemValue, err := parseType(valStr, elementType)
					if err != nil {
						return fmt.Errorf("parse slice element %q as %v: %w", valStr, elementType, err)
					}
					sliceValue = reflect.Append(sliceValue, reflect.ValueOf(elemValue))
				}

				v.Field(i).Set(sliceValue)
			}
		}
	}
	return nil
}

func (c *Context) BindParams(s any) error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("BindParams: expected struct, got %v", v.Kind())
	}

	params := mux.Vars(c.Request)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		pramTag := field.Tag.Get("param")

		if pramTag != "" {
			data := params[pramTag]

			paramType := field.Type.Kind()

			value, err := parseType(data, paramType)

			if err != nil {
				return fmt.Errorf("parse field %q as %v: %w", pramTag, paramType, err)
			}

			v.Field(i).Set(reflect.ValueOf(value))
		}
	}

	return nil
}