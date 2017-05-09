package persistence_test

import (
	. "github.com/maurofran/go-ddd-identityaccess/infrastructure/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("TenantStore", func() {
	var (
		session *mgo.Session
		db      *mgo.Database
		coll    *mgo.Collection
		fixture *TenantStore
	)

	BeforeSuite(func() {
		session, _ = mgo.Dial("mongodb://localhost:27017")
		db = session.DB("identityaccess")
		coll = db.C("tenants")
		coll.Create(&mgo.CollectionInfo{})
		coll.EnsureIndex(mgo.Index{
			Name:   "ixu_tenantId",
			Key:    []string{"tenantId"},
			Unique: true,
		})
		coll.EnsureIndex(mgo.Index{
			Name:   "ixu_name",
			Key:    []string{"name"},
			Unique: true,
		})
	})

	AfterSuite(func() {
		coll.DropCollection()
		session.Close()
	})

	BeforeEach(func() {
		fixture = &TenantStore{Db: db}
	})

	AfterEach(func() {
		coll.RemoveAll(bson.M{})
	})

	Describe("Insert", func() {
		tenantId := uuid.NewRandom().String()

		It("Should add a new tenant", func() {
			id, err := fixture.Insert(tenantId, "test", "Test", true)
			Expect(id).ToNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return an error if tenant with same id is already present", func() {
			fixture.Insert(tenantId, "test", "Test", true)
			_, err := fixture.Insert(tenantId, "other", "Other", false)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return an error if tenant with same name is already present", func() {
			fixture.Insert(tenantId, "test", "Test", true)
			_, err := fixture.Insert(uuid.NewRandom().String(), "test", "Test", true)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
	})
	Describe("Update", func() {
		var (
			tenantId string
			id       interface{}
		)

		BeforeEach(func() {
			tenantId = uuid.NewRandom().String()
			id, _ = fixture.Insert(tenantId, "test", "Test", true)
		})

		It("Should update an existing tenant", func() {
			v, err := fixture.Update(id, 0, "test", "Other", false)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
			el := bson.M{}
			coll.FindId(id).One(&el)
			Expect(v).To(Equal(1))
			Expect(el["name"]).To(Equal("test"))
			Expect(el["description"]).To(Equal("Other"))
			Expect(el["active"]).To(Equal(false))
		})
		It("Should return an error with missing id", func() {
			newId := bson.NewObjectId()
			_, err := fixture.Update(newId, 0, "test", "Other", false)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return an error with wrong version", func() {
			_, err := fixture.Update(id, 1, "test", "Other", false)
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("Delete", func() {
		var (
			tenantId string
			id       interface{}
		)

		BeforeEach(func() {
			tenantId = uuid.NewRandom().String()
			id, _ = fixture.Insert(tenantId, "test", "Test", true)
		})

		It("Should remove existing tenant", func() {
			err := fixture.Delete(id, 0)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(0))
		})
		It("Should return error if removing a missing tenant", func() {
			err := fixture.Delete(bson.NewObjectId(), 0)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return error if removing a tenant with wrong version", func() {
			err := fixture.Delete(id, 1)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
	})
	Describe("FindOneByTenantID", func() {
		var (
			tenantId string
			id       interface{}
		)

		BeforeEach(func() {
			tenantId = uuid.NewRandom().String()
			id, _ = fixture.Insert(tenantId, "test", "Test", true)
		})

		It("Should return the tenant for provided id", func() {
			res, err := fixture.FindOneByTenantID(tenantId)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.ID()).To(Equal(id))
			Expect(res.Version()).To(Equal(0))
			Expect(res.TenantID()).To(Equal(tenantId))
			Expect(res.Name()).To(Equal("test"))
			Expect(res.Description()).To(Equal("Test"))
			Expect(res.Active()).To(Equal(true))
		})
		It("Should return nil if no tenant is found for provided id", func() {
			res, err := fixture.FindOneByTenantID(uuid.NewRandom().String())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
	Describe("FindOneByName", func() {
		var (
			tenantId string
			id       interface{}
		)

		BeforeEach(func() {
			tenantId = uuid.NewRandom().String()
			id, _ = fixture.Insert(tenantId, "test", "Test", true)
		})

		It("Should return the tenant for provided name", func() {
			res, err := fixture.FindOneByName("test")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.ID()).To(Equal(id))
			Expect(res.Version()).To(Equal(0))
			Expect(res.TenantID()).To(Equal(tenantId))
			Expect(res.Name()).To(Equal("test"))
			Expect(res.Description()).To(Equal("Test"))
			Expect(res.Active()).To(Equal(true))
		})
		It("Should return nil if no tenant is found for provided name", func() {
			res, err := fixture.FindOneByName("other")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
})
