package model

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// Tenant is the aggregate root object used to provide the abstraction of a tenant.
//
// Fields of this struct should be considered as read-only. They are made public only to allow the persistence layer to
// keep well separated domain and persistence layer.
// Every direct change to Tenant fields can produce unexpected results.
type Tenant struct {
	id          interface{}
	version     int
	tenantID    *TenantID
	name        string
	description string
	active      bool
}

// NewTenant will create a new tenant from provided parameters.
func NewTenant(tenantID *TenantID, name, description string, active bool) (*Tenant, error) {
	t := new(Tenant)
	if err := t.setTenantID(tenantID); err != nil {
		return nil, err
	}
	if err := t.setName(name); err != nil {
		return nil, err
	}
	if err := t.setDescription(description); err != nil {
		return nil, err
	}
	if err := t.setActive(active); err != nil {
		return nil, err
	}
	return t, nil
}

// TenantID will return the tenant id
func (t *Tenant) TenantID() *TenantID {
	return t.tenantID
}

// Name will retrieve the tenant name
func (t *Tenant) Name() string {
	return t.name
}

// Description will retrieve the tenant description
func (t *Tenant) Description() string {
	return t.description
}

// Active is the active status of tenant
func (t *Tenant) Active() bool {
	return t.active
}

// Activate will activate the receiver tenant if it's not already active.
func (t *Tenant) Activate() {
	if !t.Active() {
		t.setActive(true)
		// TODO raise TenantActivated event
	}
}

// Deactivate will deactivate the receiver tenant if it's not already active.
func (t *Tenant) Deactivate() {
	if t.Active() {
		t.setActive(false)
		// TODO raise TenantDeactivated event
	}
}

func (t *Tenant) setTenantID(tenantID *TenantID) error {
	if tenantID == nil {
		return errors.New("tenantID is required")
	}
	t.tenantID = tenantID
	return nil
}

func (t *Tenant) setName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	t.name = name
	return nil
}

func (t *Tenant) setDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("description is required")
	}
	t.description = description
	return nil
}

func (t *Tenant) setActive(active bool) error {
	t.active = active
	return nil
}

// Equals will check if this tenant is equal to provided object.
func (t *Tenant) Equals(other interface{}) bool {
	ot, ok := other.(*Tenant)
	return ok && t.tenantID.Equals(ot.tenantID) && t.name == ot.name
}

func (t *Tenant) String() string {
	return fmt.Sprintf(
		"Tenant [tenantID=%s, name=%s, description=%s, active=%t]",
		t.tenantID,
		t.name,
		t.description,
		t.active,
	)
}