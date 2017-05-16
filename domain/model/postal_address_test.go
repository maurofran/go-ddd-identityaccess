package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewPostalAddress", func() {
	It("Should create a new postal address", func() {
		postalAddress, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
		Expect(postalAddress).ShouldNot(BeNil())
		Expect(postalAddress.StreetName()).To(Equal("6th St."))
		Expect(postalAddress.BuildingNumber()).To(Equal("123"))
		Expect(postalAddress.PostalCode()).To(Equal("32904"))
		Expect(postalAddress.City()).To(Equal("Melbourne"))
		Expect(postalAddress.StateProvince()).To(Equal("FL"))
		Expect(postalAddress.CountryCode()).To(Equal("US"))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should return an error if streetName is empty", func() {
		_, err := NewPostalAddress("", "123", "32904", "Melbourne", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if streetName is blank", func() {
		_, err := NewPostalAddress("   ", "123", "32904", "Melbourne", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if postalCode is empty", func() {
		_, err := NewPostalAddress("6th St.", "123", "", "Melbourne", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if postalCode is blank", func() {
		_, err := NewPostalAddress("6th St.", "123", "   ", "Melbourne", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if city is empty", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if city is blank", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "   ", "FL", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if stateProvince is empty", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if stateProvince is blank", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "   ", "US")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if countryCode is empty", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if countryCode is blank", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "  ")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if countryCode length is different from 2", func() {
		_, err := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "USA")
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("PostalAddress", func() {
	var fixture *PostalAddress

	BeforeEach(func() {
		fixture, _ = NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
	})

	Describe("StreetName", func() {
		It("Should return the street name", func() {
			Expect(fixture.StreetName()).To(Equal("6th St."))
		})
	})
	Describe("BuildingNumber", func() {
		It("Should return the building number", func() {
			Expect(fixture.BuildingNumber()).To(Equal("123"))
		})
	})
	Describe("PostalCode", func() {
		It("Should return the postal code", func() {
			Expect(fixture.PostalCode()).To(Equal("32904"))
		})
	})
	Describe("City", func() {
		It("Should return the city", func() {
			Expect(fixture.City()).To(Equal("Melbourne"))
		})
	})
	Describe("StateProvince", func() {
		It("Should return the state/province", func() {
			Expect(fixture.StateProvince()).To(Equal("FL"))
		})
	})
	Describe("CountryCode", func() {
		It("Should return the country code", func() {
			Expect(fixture.CountryCode()).To(Equal("US"))
		})
	})
	Describe("Equals", func() {
		It("Should be equal to itself", func() {
			Expect(fixture.Equals(fixture)).To(BeTrue())
		})
		It("Should be equal to an postal address with the same address data", func() {
			other, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "US")
			Expect(fixture.Equals(other)).To(BeTrue())
		})
		It("Should not be equal to a postal address with different streetName", func() {
			other, _ := NewPostalAddress("7th St.", "123", "32904", "Melbourne", "FL", "US")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to a postal address with different buildingNumber", func() {
			other, _ := NewPostalAddress("6th St.", "124", "32904", "Melbourne", "FL", "US")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to a postal address with different postalCode", func() {
			other, _ := NewPostalAddress("6th St.", "123", "32905", "Melbourne", "FL", "US")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to a postal address with different city", func() {
			other, _ := NewPostalAddress("6th St.", "123", "32904", "New York", "FL", "US")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to a postal address with different stateProvince", func() {
			other, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "NY", "US")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to a postal address with different countryCode", func() {
			other, _ := NewPostalAddress("6th St.", "123", "32904", "Melbourne", "FL", "IT")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to nil", func() {
			Expect(fixture.Equals(nil)).To(BeFalse())
		})
		It("Should not be equal to different type", func() {
			Expect(fixture.Equals("wrong")).To(BeFalse())
		})
	})
	Describe("String", func() {
		It("Should return a textual representation of postal adddress", func() {
			Expect(fixture.String()).ToNot(BeEmpty())
		})
	})
})
