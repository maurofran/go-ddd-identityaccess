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
// NOTE: address is intended to be read-only. The property is exposed only for persistence purpose and it should NEVER
// be modified directly.
type EmailAddress struct {
	address string
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
	email.address = address
	return email, nil
}

// Address will retrieve the address for this email.
func (email *EmailAddress) Address() string {
	return email.address
}

// Equals will check if provided object is equal to receiver email address.
func (email *EmailAddress) Equals(other interface{}) bool {
	oemail, ok := other.(*EmailAddress)
	return ok && email.address == oemail.address
}

func (email *EmailAddress) String() string {
	return fmt.Sprintf("EmailAddress [address=%s]", email.address)
}
