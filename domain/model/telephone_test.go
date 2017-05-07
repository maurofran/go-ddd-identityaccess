package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewTelephone", func() {
	It("Should create a new telephone", func() {
		telephone, err := NewTelephone("+39123456")
		Expect(telephone).ShouldNot(BeNil())
		Expect(telephone.Number).To(Equal("+39123456"))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("Should return an error if number is empty", func() {
		_, err := NewTelephone("")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if number is blank", func() {
		_, err := NewTelephone("   ")
		Expect(err).Should(HaveOccurred())
	})
	It("Should return an error if number is not valid", func() {
		_, err := NewTelephone("+391234ABC")
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("Telephone", func() {
	var fixture *Telephone

	BeforeEach(func() {
		fixture, _ = NewTelephone("+3912345678")
	})

	Describe("Equals", func() {
		It("Should be equal to itself", func() {
			Expect(fixture.Equals(fixture)).To(BeTrue())
		})
		It("Should be equal to an telephone with the same address", func() {
			other, _ := NewTelephone("+3912345678")
			Expect(fixture.Equals(other)).To(BeTrue())
		})
		It("Should not be equal to an telephone with different address", func() {
			other, _ := NewTelephone("+3987654321")
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
		It("Should return a textual representation of telephone number", func() {
			Expect(fixture.String()).ToNot(BeEmpty())
		})
	})
})
