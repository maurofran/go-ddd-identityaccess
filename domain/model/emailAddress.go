package model

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

// EmailAddress is the value object representing an e-mail address.
//
// NOTE: Address is intended to be read-only. The property is exposed only for persistence purpose and it should NEVER
// be modified directly.
type EmailAddress struct {
	Address string
}

// NewEmailAddress is the function used to create new e-mail addresses.
func NewEmailAddress(address string) (*EmailAddress, error) {
	if strings.TrimSpace(address) == "" {
		return nil, errors.New("address is required")
	}
	if !emailPattern.MatchString(address) {
		return nil, errors.New("address is not a valid e-mail address")
	}
	email := new(EmailAddress)
	email.Address = address
	return email, nil
}

// Equals will check if provided object is equal to receiver email address.
func (email *EmailAddress) Equals(other interface{}) bool {
	oemail, ok := other.(*EmailAddress)
	return ok && email.Address == oemail.Address
}

func (email *EmailAddress) String() string {
	return fmt.Sprintf("EmailAddress [address=%s]", email.Address)
}
