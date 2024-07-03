package utils

import (
	"reflect"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func IsEmpty(data interface{}) bool {
	if data == nil {
		return true
	}
	d := reflect.ValueOf(data)
	switch d.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return d.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return d.IsNil()
	}
	return false
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
