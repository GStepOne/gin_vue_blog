package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exists := getObj.Elem().FieldByName(e.Field()); exists {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
