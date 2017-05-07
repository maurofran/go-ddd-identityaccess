package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
)

var _ = Describe("NewTenantID", func() {
	It("Should return an error if id is empty", func () {
		_, err := NewTenantID("")
		Expect(err).Should(MatchError("id is required"))
	})
	It("Should return an error if id is blank", func () {
		_, err := NewTenantID("   ")
		Expect(err).Should(MatchError("id is required"))
	})
	It("Should return a tenant id if id is valid", func () {
		tid, err := NewTenantID(uuid.NewRandom().String())
		Expect(tid).ToNot(BeNil())
		Expect(err).ShouldNot(HaveOccurred())
	})
})

var _ = Describe("TenantID", func() {
	var (
		id string
		tenantID *TenantID
	)

	BeforeEach(func () {
		id = uuid.NewRandom().String()
		tenantID, _ = NewTenantID(id)
	})

	Describe("Equals", func () {
		It("Should be equals to itself", func () {
			Expect(tenantID.Equals(tenantID)).To(BeTrue())
		})
		It("Should be equals to another tenantID with same id", func () {
			other, _ := NewTenantID(id)
			Expect(tenantID.Equals(other)).To(BeTrue())
		})
		It("Should not be equal to nil", func () {
			Expect(tenantID.Equals(nil)).To(BeFalse())
		})
		It("Should not be equal to another tenantID with different id", func () {
			other, _ := NewTenantID(uuid.NewRandom().String())
			Expect(tenantID.Equals(other)).To(BeFalse())
		})
		It("Should not be equal to another type", func () {
			Expect(tenantID.Equals("other")).To(BeFalse())
		})
	})

	Describe("String", func () {
		It("Should return the text representation of tenant id", func () {
			Expect(tenantID.String()).ToNot(BeEmpty())
		})
	})
})
