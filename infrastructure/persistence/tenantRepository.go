package persistence

import (
	"context"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const tenants = "tenants"

// tenant is the external data type for BSON serialization.
type tenant struct {
	ID          interface{} `bson:"_id,omitempty"`
	Version     int         `bson:"_v"`
	TenantID    string      `bson:"tenantId"`
	Name        string      `bson:"name"`
	Description string      `bson:"description"`
	Active      bool        `bson:"active"`
}

// TenantRepository is the mongo db related repository for tenants.
type TenantRepository struct {
	c *mgo.Collection
}

// NewTenantRepository will create a new TenantRepository instance.
func NewTenantRepository(db *mgo.Database) *TenantRepository {
	tr := new(TenantRepository)
	tr.c = db.C(tenants)
	return tr
}

// Add will add a new tenant to repository, returning an error if the operation fails.
func (tr *TenantRepository) Add(ctx context.Context, tenant *model.Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	id := bson.NewObjectId()
	err := tr.c.Insert(bson.M{
		"_id":         id,
		"_v":          0,
		"tenantId":    tenant.TenantID.ID,
		"name":        tenant.Name,
		"description": tenant.Description,
		"active":      tenant.Active,
	})
	if err != nil {
		return errors.Wrapf(err, "an error occurred while adding tenant %s to repository", tenant)
	}
	tenant.ID = id
	tenant.Version = 0
	return nil
}

// Update will update an existing tenant actually stored into repository, returning an error if the operation fails.
func (tr *TenantRepository) Update(ctx context.Context, tenant *model.Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	err := tr.c.Update(
		bson.M{
			"_id": tenant.ID,
			"_v":  tenant.Version,
		},
		bson.M{
			"$inc": bson.M{
				"_v": 1,
			},
			"$set": bson.M{
				"name":        tenant.Name,
				"description": tenant.Description,
				"active":      tenant.Active,
			},
		},
	)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while updating tenant %s", tenant)
	}
	return nil
}

// Remove will remove an existing tenant from repository, returning an error if the operation fails.
func (tr *TenantRepository) Remove(ctx context.Context, tenant *model.Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	err := tr.c.Remove(bson.M{
		"_id": tenant.ID,
		"_v":  tenant.Version,
	})
	if err != nil {
		return errors.Wrapf(err, "an error occurred while removing tenant %s from repository", tenant)
	}
	return nil
}

// TenantOfId will retrieve a tenant for a given unique tenant id, returning the tenant or an error if the operation
// fails.
func (tr *TenantRepository) TenantOfId(ctx context.Context, id *model.TenantID) (*model.Tenant, error) {
	if id == nil {
		return nil, errors.New("tenantId is required")
	}
	t := new(tenant)
	if err := tr.c.Find(bson.M{"tenantId": id.ID}).One(t); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving tenant by id %s", id)
	}
	return &model.Tenant{
		ID:          t.ID,
		Version:     t.Version,
		TenantID:    &model.TenantID{ID: t.TenantID},
		Name:        t.Name,
		Description: t.Description,
		Active:      t.Active,
	}, nil
}

// TenantNamed will retrieve the tenant for a given name, returning the tenant or an error if the operation fails.
func (tr *TenantRepository) TenantNamed(ctx context.Context, name string) (*model.Tenant, error) {
	t := new(tenant)
	if err := tr.c.Find(bson.M{"name": name}).One(t); err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "An error occurred while retrieving tenant by name %s", name)
	}
	return &model.Tenant{
		ID:          t.ID,
		Version:     t.Version,
		TenantID:    &model.TenantID{ID: t.TenantID},
		Name:        t.Name,
		Description: t.Description,
		Active:      t.Active,
	}, nil
}
