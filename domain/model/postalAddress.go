package model

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// PostalAddress is the value object for the postal address.
//
// NOTE: The properties of this struct are intended to be read-only. The property are exposed only for persistence
// purpose and they should NEVER be modified directly.
type PostalAddress struct {
	streetName     string
	buildingNumber string
	postalCode     string
	city           string
	stateProvince  string
	countryCode    string
}

// NewPostalAddress will create a new postal address instance.
func NewPostalAddress(streetName, buildingNumber, postalCode, city, stateProvince, countryCode string) (*PostalAddress, error) {
	if strings.TrimSpace(streetName) == "" {
		return nil, errors.New("streetName is required")
	}
	if strings.TrimSpace(postalCode) == "" {
		return nil, errors.New("postalCode is required")
	}
	if strings.TrimSpace(city) == "" {
		return nil, errors.New("city is required")
	}
	if strings.TrimSpace(stateProvince) == "" {
		return nil, errors.New("stateProvince is required")
	}
	if strings.TrimSpace(countryCode) == "" {
		return nil, errors.New("countryCode is required")
	}
	if len(countryCode) != 2 {
		return nil, errors.New("countryCode must be 2 characters")
	}
	pa := new(PostalAddress)
	pa.streetName = streetName
	pa.buildingNumber = buildingNumber
	pa.postalCode = postalCode
	pa.city = city
	pa.stateProvince = stateProvince
	pa.countryCode = countryCode
	return pa, nil
}

// StreetName is the name of the street.
func (pa *PostalAddress) StreetName() string {
	return pa.streetName
}

// BuildingNumber is the number of building.
func (pa *PostalAddress) BuildingNumber() string {
	return pa.buildingNumber
}

// PostalCode is the postal code.
func (pa *PostalAddress) PostalCode() string {
	return pa.postalCode
}

// City of address.
func (pa *PostalAddress) City() string {
	return pa.city
}

// StateProvince part of address.
func (pa *PostalAddress) StateProvince() string {
	return pa.stateProvince
}

// CountryCode part of address.
func (pa *PostalAddress) CountryCode() string {
	return pa.countryCode
}

func (pa *PostalAddress) Equals(other interface{}) bool {
	opa, ok := other.(*PostalAddress)
	return ok && pa.streetName == opa.streetName && pa.buildingNumber == opa.buildingNumber &&
		pa.postalCode == opa.postalCode && pa.city == opa.city && pa.stateProvince == opa.stateProvince &&
		pa.countryCode == opa.countryCode
}

func (pa *PostalAddress) String() string {
	return fmt.Sprintf(
		"PostalAddress [streetName=%s, buildingNumber=%s, postalCode=%s, city=%s, stateProvince=%s, countryCode=%s]",
		pa.streetName,
		pa.buildingNumber,
		pa.postalCode,
		pa.city,
		pa.stateProvince,
		pa.countryCode,
	)
}
