package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewContactInformation", func() {
	emailAddress, _ := NewEmailAddress("foo.baz@test.com")
	postalAddress, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
	primaryTelephone, _ := NewTelephone("+39123456")
	secondaryTelephone, _ := NewTelephone("+01654321")

	It("Should create new contact information", func() {
		ci, err := NewContactInformation(emailAddress, postalAddress, primaryTelephone, secondaryTelephone)
		Expect(ci).ToNot(BeNil())
		Expect(ci.EmailAddress()).To(Equal(emailAddress))
		Expect(ci.PostalAddress()).To(Equal(postalAddress))
		Expect(ci.PrimaryTelephone()).To(Equal(primaryTelephone))
		Expect(ci.SecondaryTelephone()).To(Equal(secondaryTelephone))
		Expect(err).ToNot(HaveOccurred())
	})

	It("Should return an error if emailAddress is nil", func() {
		_, err := NewContactInformation(nil, postalAddress, primaryTelephone, secondaryTelephone)
		Expect(err).To(HaveOccurred())
	})

	It("Should return an error if postalAddress is nil", func() {
		_, err := NewContactInformation(emailAddress, nil, primaryTelephone, secondaryTelephone)
		Expect(err).To(HaveOccurred())
	})

	It("Should return an error if primaryTelephone is nil", func() {
		_, err := NewContactInformation(emailAddress, postalAddress, nil, secondaryTelephone)
		Expect(err).To(HaveOccurred())
	})

	It("Should create new contact information if secondary telephone is nil", func() {
		ci, err := NewContactInformation(emailAddress, postalAddress, primaryTelephone, nil)
		Expect(ci).ToNot(BeNil())
		Expect(ci.EmailAddress()).To(Equal(emailAddress))
		Expect(ci.PostalAddress()).To(Equal(postalAddress))
		Expect(ci.PrimaryTelephone()).To(Equal(primaryTelephone))
		Expect(ci.SecondaryTelephone()).To(BeNil())
		Expect(err).ToNot(HaveOccurred())
	})
})

var _ = Describe("ContactInformation", func() {
	emailAddress, _ := NewEmailAddress("foo.baz@test.com")
	postalAddress, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
	primaryTelephone, _ := NewTelephone("+39123456")
	secondaryTelephone, _ := NewTelephone("+01654321")
	fixture, _ := NewContactInformation(emailAddress, postalAddress, primaryTelephone, secondaryTelephone)

	Describe("EmailAddress", func() {
		It("Should return the email address", func() {
			Expect(fixture.EmailAddress()).To(Equal(emailAddress))
		})
	})
	Describe("PostalAddress", func() {
		It("Should return the postal address", func() {
			Expect(fixture.PostalAddress()).To(Equal(postalAddress))
		})
	})
	Describe("PrimaryTelephone", func() {
		It("Should return the primary telephone", func() {
			Expect(fixture.PrimaryTelephone()).To(Equal(primaryTelephone))
		})
	})
	Describe("SecondaryTelephone", func() {
		It("Should return the secondary telephone", func() {
			Expect(fixture.SecondaryTelephone()).To(Equal(secondaryTelephone))
		})
	})

	Describe("ChangeEmailAddress", func() {
		It("Should return a new contact information with changed email address", func() {
			otherEmail, _ := NewEmailAddress("foo.baz+1@test.com")
			ci, err := fixture.ChangeEmailAddress(otherEmail)
			Expect(ci).ToNot(BeNil())
			Expect(ci).ToNot(Equal(fixture))
			Expect(ci.EmailAddress()).To(Equal(otherEmail))
			Expect(ci.PostalAddress()).To(Equal(fixture.PostalAddress()))
			Expect(ci.PrimaryTelephone()).To(Equal(fixture.PrimaryTelephone()))
			Expect(ci.SecondaryTelephone()).To(Equal(fixture.SecondaryTelephone()))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return the fixture if not changing email address", func() {
			otherEmail, _ := NewEmailAddress("foo.baz@test.com")
			ci, err := fixture.ChangeEmailAddress(otherEmail)
			Expect(ci).ToNot(BeNil())
			Expect(ci).To(Equal(fixture))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return an error if email address is nil", func() {
			_, err := fixture.ChangeEmailAddress(nil)
			Expect(err).To(HaveOccurred())
		})
	})
	Describe("ChangePostalAddress", func() {
		It("Should return a new contact information with changed postal address", func() {
			otherAddress, _ := NewPostalAddress("Goldfield Rd.", "4", "96815", "Honolulu", "HI", "US")
			ci, err := fixture.ChangePostalAddress(otherAddress)
			Expect(ci).ToNot(BeNil())
			Expect(ci).ToNot(Equal(fixture))
			Expect(ci.EmailAddress()).To(Equal(fixture.EmailAddress()))
			Expect(ci.PostalAddress()).To(Equal(otherAddress))
			Expect(ci.PrimaryTelephone()).To(Equal(fixture.PrimaryTelephone()))
			Expect(ci.SecondaryTelephone()).To(Equal(fixture.SecondaryTelephone()))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return the fixture if not changing postal address", func() {
			otherAddress, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
			ci, err := fixture.ChangePostalAddress(otherAddress)
			Expect(ci).ToNot(BeNil())
			Expect(ci).To(Equal(fixture))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return an error if postal address is nil", func() {
			_, err := fixture.ChangePostalAddress(nil)
			Expect(err).To(HaveOccurred())
		})
	})
	Describe("ChangePrimaryTelephone", func() {
		It("Should return a new contact information with changed primary telephone", func() {
			otherPhone, _ := NewTelephone("+560987654")
			ci, err := fixture.ChangePrimaryTelephone(otherPhone)
			Expect(ci).ToNot(BeNil())
			Expect(ci).ToNot(Equal(fixture))
			Expect(ci.EmailAddress()).To(Equal(fixture.EmailAddress()))
			Expect(ci.PostalAddress()).To(Equal(fixture.PostalAddress()))
			Expect(ci.PrimaryTelephone()).To(Equal(otherPhone))
			Expect(ci.SecondaryTelephone()).To(Equal(fixture.SecondaryTelephone()))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return the fixture if not changing primary telephone", func() {
			otherPhone, _ := NewTelephone("+39123456")
			ci, err := fixture.ChangePrimaryTelephone(otherPhone)
			Expect(ci).ToNot(BeNil())
			Expect(ci).To(Equal(fixture))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return an error if primary telephone is nil", func() {
			_, err := fixture.ChangePrimaryTelephone(nil)
			Expect(err).To(HaveOccurred())
		})
	})
	Describe("ChangeSecondaryTelephone", func() {
		It("Should return a new contact information with changed secondary telephone", func() {
			otherPhone, _ := NewTelephone("+560987654")
			ci, err := fixture.ChangeSecondaryTelephone(otherPhone)
			Expect(ci).ToNot(BeNil())
			Expect(ci).ToNot(Equal(fixture))
			Expect(ci.EmailAddress()).To(Equal(fixture.EmailAddress()))
			Expect(ci.PostalAddress()).To(Equal(fixture.PostalAddress()))
			Expect(ci.PrimaryTelephone()).To(Equal(fixture.PrimaryTelephone()))
			Expect(ci.SecondaryTelephone()).To(Equal(otherPhone))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return the fixture if not changing secondary telephone", func() {
			otherPhone, _ := NewTelephone("+01654321")
			ci, err := fixture.ChangeSecondaryTelephone(otherPhone)
			Expect(ci).ToNot(BeNil())
			Expect(ci).To(Equal(fixture))
			Expect(err).ToNot(HaveOccurred())
		})
		It("Should return a new contact information if secondary telephone is nil", func() {
			ci, err := fixture.ChangeSecondaryTelephone(nil)
			Expect(ci).ToNot(BeNil())
			Expect(ci).ToNot(Equal(fixture))
			Expect(ci.EmailAddress()).To(Equal(fixture.EmailAddress()))
			Expect(ci.PostalAddress()).To(Equal(fixture.PostalAddress()))
			Expect(ci.PrimaryTelephone()).To(Equal(fixture.PrimaryTelephone()))
			Expect(ci.SecondaryTelephone()).To(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
