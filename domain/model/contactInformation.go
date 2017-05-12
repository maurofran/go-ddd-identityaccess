package model

import (
	"fmt"
	"github.com/pkg/errors"
)

// ContactInformation is a value object wrapping contact related information.
type ContactInformation struct {
	emailAddress       *EmailAddress
	postalAddress      *PostalAddress
	primaryTelephone   *Telephone
	secondaryTelephone *Telephone
}

// NewContactInformation will create a new contact information, returning an error if some parameter is not valid
func NewContactInformation(
	emailAddress *EmailAddress,
	postalAddress *PostalAddress,
	primaryTelephone,
	secondaryTelephone *Telephone,
) (*ContactInformation, error) {
	if emailAddress == nil {
		return nil, errors.New("emailAddress is required")
	}
	if postalAddress == nil {
		return nil, errors.New("postalAddress is required")
	}
	if primaryTelephone == nil {
		return nil, errors.New("primaryTelephone is required")
	}
	ci := new(ContactInformation)
	ci.emailAddress = emailAddress
	ci.postalAddress = postalAddress
	ci.primaryTelephone = primaryTelephone
	ci.secondaryTelephone = secondaryTelephone
	return ci, nil
}

// EmailAddress of contact information.
func (ci *ContactInformation) EmailAddress() *EmailAddress {
	return ci.emailAddress
}

// PostalAddress of contact information.
func (ci *ContactInformation) PostalAddress() *PostalAddress {
	return ci.postalAddress
}

// PrimaryTelephone of contact information.
func (ci *ContactInformation) PrimaryTelephone() *Telephone {
	return ci.primaryTelephone
}

// SecondaryTelephone of contact information.
func (ci *ContactInformation) SecondaryTelephone() *Telephone {
	return ci.secondaryTelephone
}

// ChangeEmailAddress will create a new contact information with changed e-mail address.
func (ci *ContactInformation) ChangeEmailAddress(emailAddress *EmailAddress) (*ContactInformation, error) {
	if ci.emailAddress.Equals(emailAddress) {
		return ci, nil
	}
	return NewContactInformation(emailAddress, ci.postalAddress, ci.primaryTelephone, ci.secondaryTelephone)
}

// ChangePostalAddress will create a new contact information with changed postal address.
func (ci *ContactInformation) ChangePostalAddress(postalAddress *PostalAddress) (*ContactInformation, error) {
	if ci.postalAddress.Equals(postalAddress) {
		return ci, nil
	}
	return NewContactInformation(ci.emailAddress, postalAddress, ci.primaryTelephone, ci.secondaryTelephone)
}

// ChangePrimaryTelephone will create a new contact information with changed primary telephone number.
func (ci *ContactInformation) ChangePrimaryTelephone(primaryTelephone *Telephone) (*ContactInformation, error) {
	if ci.primaryTelephone.Equals(primaryTelephone) {
		return ci, nil
	}
	return NewContactInformation(ci.emailAddress, ci.postalAddress, primaryTelephone, ci.secondaryTelephone)
}

// ChangeSecondaryTelephone will create a new contact information with changed secondary telephone number.
func (ci *ContactInformation) ChangeSecondaryTelephone(secondaryTelephone *Telephone) (*ContactInformation, error) {
	if (ci.secondaryTelephone == nil && secondaryTelephone == nil) ||
		(ci.secondaryTelephone != nil && ci.secondaryTelephone.Equals(secondaryTelephone)) {
		return ci, nil
	}
	return NewContactInformation(ci.emailAddress, ci.postalAddress, ci.primaryTelephone, secondaryTelephone)
}

// Equals will check if this contact information is equal to provided object.
func (ci *ContactInformation) Equals(other interface{}) bool {
	oci, ok := other.(*ContactInformation)
	return ok && oci != nil && ci.emailAddress.Equals(oci.emailAddress) && ci.postalAddress.Equals(oci.postalAddress) &&
		ci.primaryTelephone.Equals(oci.primaryTelephone) &&
		((ci.secondaryTelephone == nil && oci.secondaryTelephone == nil) ||
			(ci.secondaryTelephone != nil && ci.secondaryTelephone.Equals(oci.secondaryTelephone)))
}

func (ci *ContactInformation) String() string {
	return fmt.Sprintf(
		"ContactInformation [emailAddress=%s, postalAddress=%s, primaryTelephone=%s, secondaryTelephone=%s]",
		ci.emailAddress,
		ci.postalAddress,
		ci.primaryTelephone,
		ci.secondaryTelephone,
	)
}
