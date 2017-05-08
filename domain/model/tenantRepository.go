package model

import (
	"context"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// TenantRepository is the interface that represents the collection of tenants.
type TenantRepository struct {
	store TenantStore
}

// NewTenantRepository will create a new tenant repository instance with provided store as backend.
func NewTenantRepository(store TenantStore) *TenantRepository {
	tr := new(TenantRepository)
	tr.store = store
	return tr
}

// NextIdentity will generate a new tenant identifier.
func (tr *TenantRepository) NextIdentity() (*TenantID, error) {
	return NewTenantID(uuid.NewRandom().String())
}

// Add will add a new tenant to repository, returning an error if the operation fails.
func (tr *TenantRepository) Add(ctx context.Context, tenant *Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	id, err := tr.store.Insert(tenant.tenantID.id, tenant.name, tenant.description, tenant.active)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while adding tenant %s to repository", tenant)
	}
	tenant.id = id
	tenant.version = 0
	return nil
}

// Update will update an existing tenant actually stored into repository, returning an error if the operation fails.
func (tr *TenantRepository) Update(ctx context.Context, tenant *Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	version, err := tr.store.Update(tenant.id, tenant.version, tenant.name, tenant.description, tenant.active)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while updating tenant %s", tenant)
	}
	tenant.version = version
	return nil
}

// Remove will remove an existing tenant from repository, returning an error if the operation fails.
func (tr *TenantRepository) Remove(ctx context.Context, tenant *Tenant) error {
	if tenant == nil {
		return errors.New("tenant is required")
	}
	err := tr.store.Delete(tenant.id, tenant.version)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while removing tenant %s from repository", tenant)
	}
	return nil
}

// TenantOfId will retrieve a tenant for a given unique tenant id, returning the tenant or an error if the operation
// fails.
func (tr *TenantRepository) TenantOfId(ctx context.Context, id *TenantID) (*Tenant, error) {
	if id == nil {
		return nil, errors.New("tenantId is required")
	}
	t, err := tr.store.FindOneByTenantID(id.id)
	if err != nil {
		return nil, errors.Wrapf(err, "An error occurred while retrieving tenant by id %s", id)
	}
	if t == nil {
		return nil, nil
	}
	return &Tenant{
		id:          t.ID(),
		version:     t.Version(),
		tenantID:    &TenantID{id: t.TenantID()},
		name:        t.Name(),
		description: t.Description(),
		active:      t.Active(),
	}, nil
}

// TenantNamed will retrieve the tenant for a given name, returning the tenant or an error if the operation fails.
func (tr *TenantRepository) TenantNamed(ctx context.Context, name string) (*Tenant, error) {
	t, err := tr.store.FindOneByName(name)
	if err != nil {
		return nil, errors.Wrapf(err, "An error occurred while retrieving tenant by name %s", name)
	}
	if t == nil {
		return nil, nil
	}
	return &Tenant{
		id:          t.ID(),
		version:     t.Version(),
		tenantID:    &TenantID{id: t.TenantID()},
		name:        t.Name(),
		description: t.Description(),
		active:      t.Active(),
	}, nil
}

// TenantStore is the interface that must be exposed by those object who provide tenant persistence.
type TenantStore interface {
	Insert(tenantID string, name, description string, active bool) (id interface{}, err error)
	Update(id interface{}, version int, name, description string, active bool) (newVersion int, err error)
	Delete(id interface{}, version int) error
	FindOneByTenantID(tenantID string) (TenantStoreItem, error)
	FindOneByName(name string) (TenantStoreItem, error)
}

// TenantStoreItem is the representation of a tenant stored into TenantStore.
type TenantStoreItem interface {
	ID() interface{}
	Version() int
	TenantID() string
	Name() string
	Description() string
	Active() bool
}
