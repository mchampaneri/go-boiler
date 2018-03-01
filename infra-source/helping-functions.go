package main

import (
	"strings"
)

func Slugy(inputs []string) string{
	return strings.TrimSpace(strings.ToLower(strings.Join(inputs,"-")))
}

