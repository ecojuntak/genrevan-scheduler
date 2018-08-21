package model

import (
	"errors"
	"regexp"
)

func ValidateLXCName(s string) error {
	if len(s) < 1 || len(s) > 63 {
		return errors.New("LXC name isn't a valid hostname")
	}

	r, _ := regexp.Compile("^[a-zA-Z][a-zA-Z0-9-]*[a-zA-Z0-9]$")

	if !r.MatchString(s) {
		return errors.New("LXC name isn't a valid hostname")
	}

	return nil
}

func ParseImageName(s string) string {
	r := regexp.MustCompile("[0-9]+\\.[0-9]+")
	return r.FindString(s)
}
