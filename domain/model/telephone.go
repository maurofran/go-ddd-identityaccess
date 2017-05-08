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
// NOTE: number is intended to be read-only. The property is exposed only for persistence purpose and it should NEVER
// be modified directly.
type Telephone struct {
	number string
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
	t.number = number
	return t, nil
}

// Number is the telephone number.
func (t *Telephone) Number() string {
	return t.number
}

func (t *Telephone) Equals(other interface{}) bool {
	ot, ok := other.(*Telephone)
	return ok && ot.number == t.number
}

func (t *Telephone) String() string {
	return fmt.Sprintf("Telephone [number=%s]", t.number)
}
