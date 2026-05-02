package solid

import (
	"fmt"

	"github.com/goccy/go-reflect"
)

func (c *Context) SaveCookie(s any, option *CookieOption) error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("BindCookie: expected struct, got %v", v.Kind())
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		cookieTag := field.Tag.Get("cookie")

		if cookieTag != "" {
			c.SetCookie(&Cookie{
				Name: cookieTag,
				Value: fmt.Sprintf("%v", v.Field(i).Interface()),
			}, option)
		}
	}

	return nil
}