package validators

import (
	"fmt"
	"strings"
)

type ValidationErrors struct {
	Errors map[string]string
}

func (v ValidationErrors) Error() string {
	errMsgs := []string{}
	for field, msg := range v.Errors {
		errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", field, msg))
	}
	return strings.Join(errMsgs, "; ")
}


func (v ValidationErrors) HasErrors() bool {
	return len(v.Errors) > 0
}
