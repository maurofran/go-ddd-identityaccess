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
	ID          interface{}
	Version     int
	TenantID    *TenantID
	Name        string
	Description string
	Active      bool
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

// Activate will activate the receiver tenant if it's not already active.
func (t *Tenant) Activate() {
	if !t.Active {
		t.Active = true
		// TODO raise TenantActivated event
	}
}

// Deactivate will deactivate the receiver tenant if it's not already active.
func (t *Tenant) Deactivate() {
	if t.Active {
		t.Active = false
		// TODO raise TenantDeactivated event
	}
}

func (t *Tenant) setTenantID(tenantID *TenantID) error {
	if tenantID == nil {
		return errors.New("tenantID is required")
	}
	t.TenantID = tenantID
	return nil
}

func (t *Tenant) setName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	t.Name = name
	return nil
}

func (t *Tenant) setDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return errors.New("description is required")
	}
	t.Description = description
	return nil
}

func (t *Tenant) setActive(active bool) error {
	t.Active = active
	return nil
}

// Equals will check if this tenant is equal to provided object.
func (t *Tenant) Equals(other interface{}) bool {
	ot, ok := other.(*Tenant)
	return ok && t.TenantID.Equals(ot.TenantID) && t.Name == ot.Name
}

func (t *Tenant) String() string {
	return fmt.Sprintf(
		"Tenant [tenantID=%s, name=%s, description=%s, active=%t]",
		t.TenantID,
		t.Name,
		t.Description,
		t.Active,
	)
}