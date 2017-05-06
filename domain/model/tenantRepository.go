package model

import "context"

// TenantRepository is the interface that represents the collection of tenants.
type TenantRepository interface {
	// Add will add a new tenant to repository, returning an error if the operation fails.
	Add(ctx context.Context, tenant *Tenant) error

	// Update will update an existing tenant actually stored into repository, returning an error if the operation fails.
	Update(ctx context.Context, tenant *Tenant) error

	// Remove will remove an existing tenant from repository, returning an error if the operation fails.
	Remove(ctx context.Context, tenant *Tenant) error

	// TenantOfId will retrieve a tenant for a given unique tenant id, returning the tenant or an error if the operation
	// fails.
	TenantOfId(ctx context.Context, id *TenantID) (*Tenant, error)

	// TenantNamed will retrieve the tenant for a given name, returning the tenant or an error if the operation fails.
	TenantNamed(ctx context.Context, name string) (*Tenant, error)
}
