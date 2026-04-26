package solid

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/goccy/go-json"
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
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r, err = strconv.ParseUint(value, 10, 64)
	case reflect.TypeOf(time.Time{}).Kind():
		r, err = time.Parse(time.RFC3339, value)
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

func (c *Context) BindForm(s any) error {
	if err := c.Request.ParseForm(); err != nil {
		return fmt.Errorf("parse form: %w", err)
	}

	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("BindParams: expected struct, got %v", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		formTag := field.Tag.Get("form")

		if formTag != "" {
			paramType := field.Type

			if paramType == reflect.TypeOf(multipart.FileHeader{}) {
				files := c.Request.MultipartForm.File[formTag]

				if len(files) == 0 {
					continue
				}

				fh := files[0]
				if paramType == reflect.TypeOf(multipart.FileHeader{}) {
					v.Field(i).Set(reflect.ValueOf(*fh))
				} else if paramType == reflect.TypeOf((*multipart.FileHeader)(nil)) {
					v.Field(i).Set(reflect.ValueOf(fh))
				} else {
					return errors.New("unsupported file field type")
				}
			} else if paramType == reflect.TypeOf([]multipart.FileHeader{}) {
				files := c.Request.MultipartForm.File[formTag]

				sliceValue := reflect.MakeSlice(paramType, 0, len(files))

				for _, fh := range files {
					if paramType.Elem() == reflect.TypeOf(multipart.FileHeader{}) {
						sliceValue = reflect.Append(sliceValue, reflect.ValueOf(*fh))
					} else if paramType.Elem() == reflect.TypeOf((*multipart.FileHeader)(nil)) {
						sliceValue = reflect.Append(sliceValue, reflect.ValueOf(fh))
					} else {
						return errors.New("unsupported file slice field type")
					}
				}

				v.Field(i).Set(sliceValue)
			} else {
				data := c.Request.Form[formTag]
				paramTypeKind := paramType.Kind()

				if paramTypeKind == reflect.Slice {
					elementType := field.Type.Elem().Kind()

					sliceValue := reflect.MakeSlice(field.Type, 0, len(data))

					for _, valStr := range data {
						
						elemValue, err := parseType(valStr, elementType)

						if err != nil {
							return fmt.Errorf("parse slice element %q as %v: %w", valStr, elementType, err)
						}

						sliceValue = reflect.Append(sliceValue, reflect.ValueOf(elemValue))
					}

					v.Field(i).Set(sliceValue)
				} else {
					if len(data) == 0 {
						continue
					}

					value, err := parseType(data[0], paramTypeKind)

					if err != nil {
						return fmt.Errorf("parse field %q as %v: %w", formTag, paramType, err)
					}

					v.Field(i).Set(reflect.ValueOf(value))
				}
			}
		} else {
			return fmt.Errorf("field %q does not have form tag", field.Name)
		}
	}

	return nil
}

func (c *Context) BindJson(s any) error {
	defer c.Request.Body.Close()

    bodyBytes, err := io.ReadAll(c.Request.Body)

    if err != nil {
        return fmt.Errorf("Failed to read: %w", err)
    }

	if err := json.Unmarshal(bodyBytes, &s); err != nil {
		return fmt.Errorf("Failed to unmarshal: %w", err)
	}

	return nil
}

func (c *Context) BindXml(s any) error {
	defer c.Request.Body.Close()

	bodyBytes, err := io.ReadAll(c.Request.Body)

	if err != nil {
		return fmt.Errorf("Failed to read: %w", err)
	}

	if err := xml.Unmarshal(bodyBytes, &s); err != nil {
		return fmt.Errorf("Failed to unmarshal: %w", err)
	}

	return nil
}