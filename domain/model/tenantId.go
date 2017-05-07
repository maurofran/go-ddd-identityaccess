package model

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// TenantID is a value object representing the unique identifier of a tenant.
type TenantID struct {
	ID string
}

// NewTenantID is used to create a new TenantID instance.
func NewTenantID(id string) (*TenantID, error) {
	tid := new(TenantID)
	if err := tid.setID(id); err != nil {
		return nil, err
	}
	return tid, nil
}

func (tid *TenantID) setID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id is required")
	}
	tid.ID = id
	return nil
}

// Equals check if provided object is equal to this tenant identifier.
func (tid *TenantID) Equals(other interface{}) bool {
	otid, ok := other.(*TenantID)
	return ok && tid.ID == otid.ID
}

func (tid *TenantID) String() string {
	return fmt.Sprintf("TenantID [id=%s]", tid.ID)
}
