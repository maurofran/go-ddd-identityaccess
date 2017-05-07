package model

import (
	"context"
	"github.com/pkg/errors"
)

// TenantProvisioningService is the domain service for provisioning tenants.
type TenantProvisioningService struct {
	tenantRepository TenantRepository
}

// NewTenantProvisioningService will create a new tenant provisioning service instance.
func NewTenantProvisioningService(tenantRepository TenantRepository) *TenantProvisioningService {
	tps := new(TenantProvisioningService)
	tps.tenantRepository = tenantRepository
	return tps
}

// ProvisionTenant will provision a new tenant.
func (tps *TenantProvisioningService) ProvisionTenant(
	ctx context.Context,
	tenantName,
	tenantDescription string,
	administratorName *FullName,
	emailAddress *EmailAddress,
	postalAddress interface{},
	primaryTelephone *Telephone,
	secondaryTelephone *Telephone,
) (*Tenant, error) {
	tenantID, err := tps.tenantRepository.NextIdentity()
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while generating new tenant ID")
	}
	tenant, err := NewTenant(tenantID, tenantName, tenantDescription, true)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while creating new tenant")
	}
	if err = tps.tenantRepository.Add(ctx, tenant); err != nil {
		return nil, errors.Wrapf(err, "an error occurred while adding tenant %s to repository", tenant)
	}
	if err = tps.registerAdministratorFor(tenant, administratorName, emailAddress, postalAddress, primaryTelephone, secondaryTelephone); err != nil {
		return nil, errors.Wrapf(err, "an error occurred while registering administrator for %s", tenant)
	}
	// TODO raise TenantProvisioned event
	return tenant, nil
}

func (tps *TenantProvisioningService) registerAdministratorFor(
	tenant *Tenant,
	administratorName *FullName,
	emailAddress *EmailAddress,
	postalAddress interface{},
	primaryTelephone *Telephone,
	secondaryTelephone *Telephone,
) error {
	return nil
}
