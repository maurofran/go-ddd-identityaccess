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
	StreetName     string
	BuildingNumber string
	PostalCode     string
	City           string
	StateProvince  string
	CountryCode    string
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
	pa.StreetName = streetName
	pa.BuildingNumber = buildingNumber
	pa.PostalCode = postalCode
	pa.City = city
	pa.StateProvince = stateProvince
	pa.CountryCode = countryCode
	return pa, nil
}

func (pa *PostalAddress) Equals(other interface{}) bool {
	opa, ok := other.(*PostalAddress)
	return ok && pa.StreetName == opa.StreetName && pa.BuildingNumber == opa.BuildingNumber &&
		pa.PostalCode == opa.PostalCode && pa.City == opa.City && pa.StateProvince == opa.StateProvince &&
		pa.CountryCode == opa.CountryCode
}

func (pa *PostalAddress) String() string {
	return fmt.Sprintf(
		"PostalAddress [streetName=%s, buildingNumber=%s, postalCode=%s, city=%s, stateProvince=%s, countryCode=%s]",
		pa.StreetName,
		pa.BuildingNumber,
		pa.PostalCode,
		pa.City,
		pa.StateProvince,
		pa.CountryCode,
	)
}
