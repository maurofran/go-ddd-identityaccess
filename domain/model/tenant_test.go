package model_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/domain/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
)

var _ = Describe("NewTenant", func() {
	var (
		id       string
		tenantId *TenantID
	)

	BeforeEach(func() {
		id = uuid.NewRandom().String()
		tenantId, _ = NewTenantID(id)
	})

	It("Should return an error if tenant id is null", func() {
		_, err := NewTenant(nil, "streamtune", "Streamtune", true)
		Expect(err).Should(MatchError("tenantID is required"))
	})
	It("Should return an error if name is empty", func() {
		_, err := NewTenant(tenantId, "", "Streamtune", true)
		Expect(err).Should(MatchError("name is required"))
	})
	It("Should return an error if name is blank", func() {
		_, err := NewTenant(tenantId, "   ", "Streamtune", true)
		Expect(err).Should(MatchError("name is required"))
	})
	It("Should return an error if description is empty", func() {
		_, err := NewTenant(tenantId, "streamtune", "", true)
		Expect(err).Should(MatchError("description is required"))
	})
	It("Should return an error if description is blank", func() {
		_, err := NewTenant(tenantId, "streamtune", "   ", true)
		Expect(err).Should(MatchError("description is required"))
	})
	It("Should return a tenant if parameters are valid", func() {
		t, err := NewTenant(tenantId, "streamtune", "Streamtune", true)
		Expect(t).ToNot(BeNil())
		Expect(err).To(BeNil())
	})
})

var _ = Describe("Tenant", func() {

})
