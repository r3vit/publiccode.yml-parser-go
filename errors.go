package publiccode

import (
	"fmt"
	"strings"
)

type ErrorInvalidKey struct {
	Key string
}

func (e ErrorInvalidKey) Error() string {
	if BaseDir != "" {
		return fmt.Sprintf("invalid remote key: %s", e.Key)
	}
	return fmt.Sprintf("invalid key: %s", e.Key)
}

type ErrorInvalidValue struct {
	Key    string
	Reason string
}

func (e ErrorInvalidValue) Error() string {
	if BaseDir != "" {
		return fmt.Sprintf("wrong value on remote for key %s: %s", e.Key, e.Reason)
	}
	return fmt.Sprintf("wrong value for key %s: %s", e.Key, e.Reason)
}

func newErrorInvalidValue(key string, reason string, args ...interface{}) ErrorInvalidValue {
	return ErrorInvalidValue{Key: key, Reason: fmt.Sprintf(reason, args...)}
}

type ErrorParseMulti []error

func (es ErrorParseMulti) Error() string {
	var ss []string
	for _, e := range es {
		ss = append(ss, e.Error())
	}
	return strings.Join(ss, "\n")
}
