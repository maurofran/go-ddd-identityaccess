package application

import (
	"context"
	"github.com/maurofran/go-ddd-identityaccess/application/command"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/pkg/errors"
)

// IdentityApplicationService is the application service used to manage identities.
type IdentityApplicationService struct {
	tenantRepository model.TenantRepository
}

// NewIdentityApplicationService will create a new identity application service instance.
func NewIdentityApplicationService(tenantRepository model.TenantRepository) *IdentityApplicationService {
	ias := new(IdentityApplicationService)
	ias.tenantRepository = tenantRepository
	return ias
}

// ProvisionTenant will provision a new tenant.
func (ias *IdentityApplicationService) ProvisionTenant(ctx context.Context, command *command.ProvisionTenant) error {
	return nil
}

// ActivateTenant will activate the tenant with id provided in command.
func (ias *IdentityApplicationService) ActivateTenant(ctx context.Context, command *command.ActivateTenant) error {
	tenant, err := ias.existingTenant(ctx, command.TenantID)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while activating tenant with id %s", command.TenantID)
	}
	tenant.Activate()
	if err = ias.tenantRepository.Update(ctx, tenant); err != nil {
		return errors.Wrapf(err, "an error occurred while activating tenant with id %s", command.TenantID)
	}
	return nil
}

// DeactivateTenant will deactivate the tenant with id provided in command.
func (ias *IdentityApplicationService) DeactivateTenant(ctx context.Context, command *command.DeactivateTenant) error {
	tenant, err := ias.existingTenant(ctx, command.TenantID)
	if err != nil {
		return errors.Wrapf(err, "an error occurred while deactivating tenant with id %s", command.TenantID)
	}
	tenant.Deactivate()
	if err = ias.tenantRepository.Update(ctx, tenant); err != nil {
		return errors.Wrapf(err, "an error occurred while deactivating tenant with id %s", command.TenantID)
	}
	return nil
}

func (ias *IdentityApplicationService) existingTenant(ctx context.Context, tenantID string) (*model.Tenant, error) {
	tid, err := model.NewTenantID(tenantID)
	if err != nil {
		return nil, err
	}
	return ias.tenantRepository.TenantOfId(ctx, tid)
}