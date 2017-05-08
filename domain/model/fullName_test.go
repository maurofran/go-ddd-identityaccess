package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewFullName", func() {
	It("Should return a new full name", func() {
		fn, err := NewFullName("Foo", "Baz")
		Expect(fn).ToNot(BeNil())
		Expect(fn.FirstName()).To(Equal("Foo"))
		Expect(fn.LastName()).To(Equal("Baz"))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should return an error if first name is empty", func() {
		_, err := NewFullName("", "Baz")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if first name format is invalid", func() {
		_, err := NewFullName("780", "Baz")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if last name is empty", func() {
		_, err := NewFullName("Foo", "")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if last name format is invalid", func() {
		_, err := NewFullName("Foo", "780")
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("FullName", func() {
	var fixture *FullName

	BeforeEach(func() {
		fixture, _ = NewFullName("Foo", "Baz")
	})

	Describe("FirstName", func() {
		It("Should return the first name", func() {
			Expect(fixture.FirstName()).To(Equal("Foo"))
		})
	})
	Describe("LastName", func() {
		It("Should return the last name", func() {
			Expect(fixture.LastName()).To(Equal("Baz"))
		})
	})
	Describe("AsFormattedName", func() {
		It("Should return the formatted name", func() {
			Expect(fixture.AsFormattedName()).To(Equal("Foo Baz"))
		})
	})
	Describe("WithChangedFirstName", func() {
		It("Should return a new full name with changed first name", func() {
			nfm, err := fixture.WithChangedFirstName("Bar")
			Expect(nfm).ToNot(BeNil())
			Expect(nfm.FirstName()).To(Equal("Bar"))
			Expect(nfm.LastName()).To(Equal("Baz"))
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should return an error with empty first name", func() {
			nfm, err := fixture.WithChangedFirstName("")
			Expect(nfm).To(BeNil())
			Expect(err).Should(HaveOccurred())
		})
		It("Should return an error with invalid first name", func() {
			nfm, err := fixture.WithChangedFirstName("789")
			Expect(nfm).To(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("WithChangedLastName", func() {
		It("Should return a new full name with changed last name", func() {
			nfm, err := fixture.WithChangedLastName("Bar")
			Expect(nfm).ToNot(BeNil())
			Expect(nfm.FirstName()).To(Equal("Foo"))
			Expect(nfm.LastName()).To(Equal("Bar"))
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Should return an error with empty first name", func() {
			nfm, err := fixture.WithChangedLastName("")
			Expect(nfm).To(BeNil())
			Expect(err).Should(HaveOccurred())
		})
		It("Should return an error with invalid first name", func() {
			nfm, err := fixture.WithChangedLastName("789")
			Expect(nfm).To(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("Equals", func() {
		It("Should be equals to itself", func() {
			Expect(fixture.Equals(fixture)).To(BeTrue())
		})
		It("Should be equals to a full name with same first and last name", func() {
			other, _ := NewFullName("Foo", "Baz")
			Expect(fixture.Equals(other)).To(BeTrue())
		})
		It("Should not be equals to a full name with different first name", func() {
			other, _ := NewFullName("Other", "Baz")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should not be equals to full name with different last name", func() {
			other, _ := NewFullName("Foo", "Other")
			Expect(fixture.Equals(other)).To(BeFalse())
		})
		It("Should return false if equals to nil", func() {
			Expect(fixture.Equals(nil)).To(BeFalse())
		})
		It("Should return false if equals to other type", func() {
			Expect(fixture.Equals("wrong")).To(BeFalse())
		})
	})
})
