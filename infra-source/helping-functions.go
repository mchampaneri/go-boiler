package main

import (
	"strings"
)

// Slugy takes arry of string as input and makes
// slug for it by replaceing any blank space to dash
func Slugy(inputs []string) string{
	return strings.TrimSpace(strings.ToLower(strings.Join(inputs,"-")))
}

