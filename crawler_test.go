package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestObjectToString(t *testing.T) {
	object := generateObject()
	s := Object(object).String()
	assert.Equal(t, "1;2", s)
}

func generateObject() map[string]interface{} {
	m := make(map[string]interface{})
	m["a"] = 1
	m["b"] = 2
	return m
}
