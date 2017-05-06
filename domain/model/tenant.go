package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// Tenant is the aggregate root object used to provide the abstraction of a tenant.
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
	t.id = nil
	t.version = 0
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
	if !t.active {
		t.active = true
		// TODO raise TenantActivated event
	}
}

// Deactivate will deactivate the receiver tenant if it's not already active.
func (t *Tenant) Deactivate() {
	if t.active {
		t.active = false
		// TODO raise TenantDeactivated event
	}
}

// TenantID is the unique id of this tenant.
func (t *Tenant) TenantID() *TenantID {
	return t.tenantID
}

// Name is the unique name of this tenant.
func (t *Tenant) Name() string {
	return t.name
}

// Description is the description of this tenant.
func (t *Tenant) Description() string {
	return t.description
}

// Active is the value of active flag.
func (t *Tenant) Active() bool {
	return t.active
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

// GetBSON is the implementation of bson.Getter interface.
func (t *Tenant) GetBSON() (interface{}, error) {
	return &tenant{
		ID:          t.id,
		Version:     t.version,
		TenantID:    t.tenantID.id,
		Name:        t.name,
		Description: t.description,
		Active:      t.active,
	}, nil
}

// SetBSON is the implementation of bson.Setter interface.
func (t *Tenant) SetBSON(raw bson.Raw) error {
	decoded := new(tenant)
	if err := raw.Unmarshal(decoded); err != nil {
		return err
	}
	tenantID, _ := NewTenantID(decoded.TenantID)
	t.id = decoded.ID
	t.version = decoded.Version
	t.tenantID = tenantID
	t.name = decoded.Name
	t.description = decoded.Description
	t.active = decoded.Active
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

// tenant is the external data type for BSON serialization.
type tenant struct {
	ID          interface{} `bson:"_id,omitempty"`
	Version     int         `bson:"_v"`
	TenantID    string      `bson:"tenantId"`
	Name        string      `bson:"name"`
	Description string      `bson:"description"`
	Active      bool        `bson:"active"`
}
