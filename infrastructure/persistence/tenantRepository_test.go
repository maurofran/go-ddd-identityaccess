package persistence_test

import (
	"context"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	. "github.com/maurofran/go-ddd-identityaccess/infrastructure/persistence"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var _ = Describe("TenantRepository", func() {
	var (
		session *mgo.Session
		db      *mgo.Database
		coll    *mgo.Collection
		fixture *TenantRepository
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
		fixture = NewTenantRepository(db)
	})

	AfterEach(func() {
		coll.RemoveAll(bson.M{})
	})

	Describe("Add", func() {
		tid, _ := model.NewTenantID(uuid.NewRandom().String())
		t, _ := model.NewTenant(tid, "test", "Test", true)

		It("Should add a new tenant", func() {
			err := fixture.Add(context.TODO(), t)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return an error if tenant is null", func() {
			err := fixture.Add(context.TODO(), nil)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(0))
		})
		It("Should return an error if tenant with same id is already present", func() {
			fixture.Add(context.TODO(), t)
			other, _ := model.NewTenant(tid, "other", "Other", true)
			err := fixture.Add(context.TODO(), other)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return an error if tenant with same name is already present", func() {
			fixture.Add(context.TODO(), t)
			otherId, _ := model.NewTenantID(uuid.NewRandom().String())
			other, _ := model.NewTenant(otherId, "test", "Other", true)
			err := fixture.Add(context.TODO(), other)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
	})
	Describe("Update", func() {
		tid, _ := model.NewTenantID(uuid.NewRandom().String())
		t, _ := model.NewTenant(tid, "test", "Test", true)

		BeforeEach(func() {
			fixture.Add(context.TODO(), t)
		})

		It("Should update an existing tenant", func() {
			t.Deactivate()
			err := fixture.Update(context.TODO(), t)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
			var el struct {
				Version int `bson:"_v"`
				Active bool `bson:"active"`
			}
			coll.FindId(t.ID).One(&el)
			Expect(el.Version).To(Equal(1))
			Expect(el.Active).To(BeFalse())
		})
		It("Should return an error with missing id", func() {
			coll.RemoveId(t.ID)
			err := fixture.Update(context.TODO(), t)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(0))
		})
		It("Should return an error with nil tenant", func() {
			err := fixture.Update(context.TODO(), nil)
			Expect(err).Should(HaveOccurred())
		})
		It("Should return an error with non persisted tenant", func() {
			other, _ := model.NewTenant(tid, "other", "Other", false)
			err := fixture.Update(context.TODO(), other)
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("Remove", func() {
		tid, _ := model.NewTenantID(uuid.NewRandom().String())
		t, _ := model.NewTenant(tid, "test", "Test", true)

		BeforeEach(func() {
			fixture.Add(context.TODO(), t)
		})

		It("Should remove existing tenant", func() {
			err := fixture.Remove(context.TODO(), t)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(coll.Count()).To(Equal(0))
		})
		It("Should return error if removing a missing tenant", func() {
			coll.RemoveId(t.ID)
			err := fixture.Remove(context.TODO(), t)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(0))
		})
		It("Should return error if removing a nil tenant", func() {
			err := fixture.Remove(context.TODO(), nil)
			Expect(err).Should(HaveOccurred())
			Expect(coll.Count()).To(Equal(1))
		})
		It("Should return error if removing a non persistent tenant", func() {
			other, _ := model.NewTenant(tid, "other", "Other", false)
			err := fixture.Remove(context.TODO(), other)
			Expect(err).Should(HaveOccurred())
		})
	})
	Describe("TenantOfId", func() {
		tid, _ := model.NewTenantID(uuid.NewRandom().String())
		t, _ := model.NewTenant(tid, "test", "Test", true)

		BeforeEach(func() {
			fixture.Add(context.TODO(), t)
		})

		It("Should return the tenant for provided id", func() {
			res, err := fixture.TenantOfId(context.TODO(), tid)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.Equals(t)).To(BeTrue())
		})
		It("Should return nil if no tenant is found for provided id", func() {
			otherId, _ := model.NewTenantID(uuid.NewRandom().String())
			res, err := fixture.TenantOfId(context.TODO(), otherId)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(BeNil())
		})
		It("Should return an error if tenantId is nil", func() {
			res, err := fixture.TenantOfId(context.TODO(), nil)
			Expect(err).Should(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
	Describe("TenantNamed", func() {
		tid, _ := model.NewTenantID(uuid.NewRandom().String())
		t, _ := model.NewTenant(tid, "test", "Test", true)

		BeforeEach(func() {
			fixture.Add(context.TODO(), t)
		})

		It("Should return the tenant for provided name", func() {
			res, err := fixture.TenantNamed(context.TODO(), "test")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.Equals(t)).To(BeTrue())
		})
		It("Should return nil if no tenant is found for provided name", func() {
			res, err := fixture.TenantNamed(context.TODO(), "other")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
})
