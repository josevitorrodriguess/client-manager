package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	EmailRegex     = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}as~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	CPFRegex       = regexp.MustCompile(`^([0-9]{3}\.?[0-9]{3}\.?[0-9]{3}-?[0-9]{2})$`)
	PhoneRegex     = regexp.MustCompile(`^(0?[0-9]{2})?\s*?([0-9])\s*?([0-9]{4})\s*[-]?\s*([0-9]{4})$`)
	BirthDateRegex = regexp.MustCompile(`^(19|20)\d\d[- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])$|^(0[1-9]|[12][0-9]|3[01])[- /.](0[1-9]|1[012])[- /.](19|20)\d\d$`)
	CEPRegex       = regexp.MustCompile(`^([0-9]{5})-?([0-9]{3})$`)
	CNPJRegex      = regexp.MustCompile(`^([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})$`)
)

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

