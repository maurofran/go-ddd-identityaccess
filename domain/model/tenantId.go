package model

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// TenantID is a value object representing the unique identifier of a tenant.
type TenantID struct {
	id string
}

// NewTenantID is used to create a new TenantID instance.
func NewTenantID(id string) (*TenantID, error) {
	tid := new(TenantID)
	if err := tid.setID(id); err != nil {
		return nil, err
	}
	return tid, nil
}

// ID is the unique id of tenantID instance.
func (tid *TenantID) ID() string {
	return tid.id
}

func (tid *TenantID) setID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id is required")
	}
	tid.id = id
	return nil
}

// Equals check if provided object is equal to this tenant identifier.
func (tid *TenantID) Equals(other interface{}) bool {
	otid, ok := other.(*TenantID)
	return ok && tid.id == otid.id
}

func (tid *TenantID) String() string {
	return fmt.Sprintf("TenantID [id=%s]", tid.id)
}
