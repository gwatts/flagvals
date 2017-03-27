package flagvals

import (
	"fmt"
	"strings"
)

// OneOfString holds a string that must be one of the defined choices.
type OneOfString struct {
	Choices []string // The acceptable choices the user may pass to the flag
	Value   string   // the current value of the flag
}

// Set implements the flag.Value interface.
func (so *OneOfString) Set(v string) error {
	for _, c := range so.Choices {
		if c == v {
			so.Value = v
			return nil
		}
	}
	return fmt.Errorf("invalid choice; must be one of %s", strings.Join(so.Choices, ","))
}

// String returns the current value of the flag.
func (so *OneOfString) String() string {
	return so.Value
}

func NewOneOfString(choices ...string) *OneOfString {
	return &OneOfString{
		Choices: choices,
	}
}
