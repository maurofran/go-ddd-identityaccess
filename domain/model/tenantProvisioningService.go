package model

import (
	"context"
	"github.com/pkg/errors"
)

// TenantProvisioningService is the domain service for provisioning tenants.
type TenantProvisioningService struct {
	TenantRepository *TenantRepository `inject:""`
}

// ProvisionTenant will provision a new tenant.
func (tps *TenantProvisioningService) ProvisionTenant(
	ctx context.Context,
	tenantName,
	tenantDescription string,
	administratorName *FullName,
	emailAddress *EmailAddress,
	postalAddress *PostalAddress,
	primaryTelephone *Telephone,
	secondaryTelephone *Telephone,
) (DomainEvents, error) {
	events := noEvents()
	tenantID, err := tps.TenantRepository.NextIdentity()
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while generating new tenant ID")
	}
	tenant, err := NewTenant(tenantID, tenantName, tenantDescription, true)
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while creating new tenant")
	}
	if err = tps.TenantRepository.Add(ctx, tenant); err != nil {
		return nil, errors.Wrapf(err, "an error occurred while adding tenant %s to repository", tenant)
	}
	events = events.and(tenantProvisioned(tenantID))

	ev, err := tps.registerAdministratorFor(tenant, administratorName, emailAddress, postalAddress, primaryTelephone,
		secondaryTelephone)
	if err != nil {
		return nil, errors.Wrapf(err, "an error occurred while registering administrator for %s", tenant)
	}
	events = events.and(ev...)
	return events, nil
}

func (tps *TenantProvisioningService) registerAdministratorFor(
	tenant *Tenant,
	administratorName *FullName,
	emailAddress *EmailAddress,
	postalAddress *PostalAddress,
	primaryTelephone *Telephone,
	secondaryTelephone *Telephone,
) (DomainEvents, error) {
	return nil, nil
}
