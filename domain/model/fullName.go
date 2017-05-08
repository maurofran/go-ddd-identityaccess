package model

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

var (
	firstNamePattern = regexp.MustCompile("[A-Z][a-z]*")
	lastNamePattern  = regexp.MustCompile("^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?")
)

// FullName is the value object representing a full person name.
//
// NOTE: FirstName and LastName are intended to be read-only. The property are exposed only for persistence purpose and
// they should NEVER be modified directly.
type FullName struct {
	firstName string
	lastName  string
}

// NewFullName will create a new full name with provided first and last names.
func NewFullName(firstName, lastName string) (*FullName, error) {
	if strings.TrimSpace(firstName) == "" {
		return nil, errors.New("firstName is required")
	}
	if !firstNamePattern.MatchString(firstName) {
		return nil, errors.New("firstName format is invalid")
	}
	if strings.TrimSpace(lastName) == "" {
		return nil, errors.New("lastName is required")
	}
	if !lastNamePattern.MatchString(lastName) {
		return nil, errors.New("lastName format is invalid")
	}
	fn := new(FullName)
	fn.firstName = firstName
	fn.lastName = lastName
	return fn, nil
}

// FirstName is the first name part of full name.
func (fn *FullName) FirstName() string {
	return fn.firstName
}

// LastName is the last name of full name.
func (fn *FullName) LastName() string {
	return fn.lastName
}

// AsFormattedName will return the full name formatted as "firstName lastName".
func (fn *FullName) AsFormattedName() string {
	return fmt.Sprintf("%s %s", fn.firstName, fn.lastName)
}

// WithChangedFirstName will create a new full name with changed first name.
func (fn *FullName) WithChangedFirstName(firstName string) (*FullName, error) {
	return NewFullName(firstName, fn.lastName)
}

// WithChangedLastName will create a new full name with changed last name.
func (fn *FullName) WithChangedLastName(lastName string) (*FullName, error) {
	return NewFullName(fn.firstName, lastName)
}

// Equals will check if provided other object is equal to the receiver one.
func (fn *FullName) Equals(other interface{}) bool {
	ofn, ok := other.(*FullName)
	return ok && fn.firstName == ofn.firstName && fn.lastName == ofn.lastName
}

func (fn *FullName) String() string {
	return fmt.Sprintf("FullName [firstName=%s, lastName=%s]", fn.firstName, fn.lastName)
}
