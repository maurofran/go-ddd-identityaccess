package model

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

var numberPattern = regexp.MustCompile("^(?:\\+\\d{2})\\d{3,}$")

// Telephone is the value object for telephone number.
//
// NOTE: Number is intended to be read-only. The property is exposed only for persistence purpose and it should NEVER
// be modified directly.
type Telephone struct {
	Number string
}

// NewTelephone is the factory function to create new telephone instances.
func NewTelephone(number string) (*Telephone, error) {
	if strings.TrimSpace(number) == "" {
		return nil, errors.New("number is required")
	}
	if !numberPattern.MatchString(number) {
		return nil, errors.New("number format is invalid")
	}
	t := new(Telephone)
	t.Number = number
	return t, nil
}

func (t *Telephone) Equals(other interface{}) bool {
	ot, ok := other.(*Telephone)
	return ok && ot.Number == t.Number
}

func (t *Telephone) String() string {
	return fmt.Sprintf("Telephone [number=%s]", t.Number)
}
