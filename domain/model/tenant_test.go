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
	var (
		id       string
		tenantId *TenantID
		tenant   *Tenant
	)

	BeforeEach(func() {
		id = uuid.NewRandom().String()
		tenantId, _ = NewTenantID(id)
		tenant, _ = NewTenant(tenantId, "streamtune", "Streamtune", true)
	})

	Describe("TenantID", func() {
		It("Should return the unique tenant id", func() {
			Expect(tenant.TenantID()).To(BeEquivalentTo(tenantId))
		})
	})
	Describe("Name", func() {
		It("Should return the unique tenant name", func() {
			Expect(tenant.Name()).To(Equal("streamtune"))
		})
	})
	Describe("Description", func() {
		It("Should return the tenant description", func() {
			Expect(tenant.Description()).To(Equal("Streamtune"))
		})
	})
	Describe("Active", func() {
		It("Should return the activation status", func() {
			Expect(tenant.Active()).To(BeTrue())
		})
	})
	Describe("Activate", func() {
		It("Should do nothing if tenant is active", func() {
			tenant.Activate()
			Expect(tenant.Active()).To(BeTrue())
		})
		It("Should activate the tenant if it's not active", func() {
			tenant.Deactivate()
			tenant.Activate()
			Expect(tenant.Active()).To(BeTrue())
		})
	})
	Describe("Deactivate", func() {
		It("Should deactivate the tenant if it's active", func() {
			tenant.Deactivate()
			Expect(tenant.Active()).To(BeFalse())
		})
		It("Should do nothing if tenant it's deactivated", func() {
			tenant.Deactivate()
			tenant.Deactivate()
			Expect(tenant.Active()).To(BeFalse())
		})
	})
	Describe("Equals", func() {
		It("Should return true if equals to itself", func() {
			Expect(tenant.Equals(tenant)).To(BeTrue())
		})
		It("Should return true if equals to tenant with same tenant id and name", func() {
			other, _ := NewTenant(tenantId, "streamtune", "Other description", false)
			Expect(tenant.Equals(other)).To(BeTrue())
		})
		It("Should return false if other tenant has different tenant id", func() {
			otherId, _ := NewTenantID(uuid.NewRandom().String())
			other, _ := NewTenant(otherId, "streamtune", "Other description", false)
			Expect(tenant.Equals(other)).To(BeFalse())
		})
		It("Should return false if other tenant has different name", func() {
			other, _ := NewTenant(tenantId, "other", "Other description", false)
			Expect(tenant.Equals(other)).To(BeFalse())
		})
		It("Should return false if other tenant is different type", func() {
			Expect(tenant.Equals("different")).To(BeFalse())
		})
		It("Should return false if other tenant is nil", func() {
			Expect(tenant.Equals(nil)).To(BeFalse())
		})
	})
	Describe("String", func() {
		It("Should return the textual description of tenant", func() {
			Expect(tenant.String()).ToNot(BeEmpty())
		})
	})
})
