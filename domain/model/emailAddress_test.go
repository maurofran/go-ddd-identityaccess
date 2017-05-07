package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewEmailAddress", func() {
	It("Should create a new email address", func() {
		email, err := NewEmailAddress("foo.baz@test.com")
		Expect(email).ShouldNot(BeNil())
		Expect(email.Address).To(Equal("foo.baz@test.com"))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should return an error if email is null", func() {
		_, err := NewEmailAddress("")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if email is blank", func() {
		_, err := NewEmailAddress("   ")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if email is not valid", func() {
		_, err := NewEmailAddress("foo@baz.")
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("EmailAddress", func() {
	var fixture *EmailAddress

	BeforeEach(func() {
		fixture, _ = NewEmailAddress("foo.baz@test.com")
	})

	Describe("Equals", func() {
		It("Should be equal to itself", func() {
			Expect(fixture.Equals(fixture)).To(BeTrue())
		})
		It("Should be equal to an email with the same address", func() {
			other, _ := NewEmailAddress("foo.baz@test.com")
			Expect(fixture.Equals(other)).To(BeTrue())
		})
		It("Should not be equal to an email with different address", func() {
			other, _ := NewEmailAddress("oter@test.com")
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
		It("Should return a textual representation of email address", func() {
			Expect(fixture.String()).ToNot(BeEmpty())
		})
	})
})
