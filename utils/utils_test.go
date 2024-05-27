package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidEmail(t *testing.T) {
	email := "akinwumikaliyanu@gmail.com"
	b := IsValidEmail(email)
	assert.True(t, b, "Expected emails to be valid")
}

func TestInvalidEmail(t *testing.T) {
	email := "$a$kinwumi$iyanu@"
	b := IsValidEmail(email)
	assert.False(t, b, "Expected emails to be invalid")
}

func TestEmptyField(t *testing.T) {
	field := ""
	isEmpty := IsEmpty(field)
	assert.True(t, isEmpty, "Expected empty return")
}

func TestFieldNotEmpty(t *testing.T) {
	field := "I am not empty"
	isEmpty := IsEmpty(field)
	assert.False(t, isEmpty, "Expected no empty return")
}
