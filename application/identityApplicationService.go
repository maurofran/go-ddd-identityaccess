package application

import (
	"context"
	"github.com/maurofran/go-ddd-identityaccess/application/command"
	"github.com/maurofran/go-ddd-identityaccess/application/representation"
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

// IdentityApplicationService is the application service used to manage identities.
type IdentityApplicationService struct {
	validate                  *validator.Validate
	tenantRepository          model.TenantRepository
	tenantProvisioningService model.TenantProvisioningService
}

// NewIdentityApplicationService will create a new identity application service instance.
func NewIdentityApplicationService(
	tenantRepository model.TenantRepository,
	tenantProvisioningService model.TenantProvisioningService,
) *IdentityApplicationService {
	ias := new(IdentityApplicationService)
	ias.validate = validator.New()
	ias.tenantRepository = tenantRepository
	ias.tenantProvisioningService = tenantProvisioningService
	return ias
}

// Tenant will retrieve the representation of tenant with provided id
func (ias *IdentityApplicationService) Tenant(ctx context.Context, tenantId string) (*representation.Tenant, error) {
	tenant, err := ias.existingTenant(ctx, tenantId)
	if err != nil {
		return nil, err
	}
	if tenant == nil {
		return nil, nil
	}
	return &representation.Tenant{
		TenantID:    tenant.TenantID.ID,
		Name:        tenant.Name,
		Description: tenant.Description,
		Active:      tenant.Active,
	}, nil
}

// ProvisionTenant will provision a new tenant.
func (ias *IdentityApplicationService) ProvisionTenant(ctx context.Context, command *command.ProvisionTenant) error {
	if err := ias.validate.Struct(command); err != nil {
		return err
	}
	administratorName, err := model.NewFullName(command.AdministratorFirstName, command.AdministratorLastName)
	if err != nil {
		return errors.Wrap(err, "an error occurred while creating administrator full name")
	}
	emailAddress, err := model.NewEmailAddress(command.EmailAddress)
	if err != nil {
		return errors.Wrap(err, "an error occurred while creating administrator email address")
	}
	postalAddress, err := model.NewPostalAddress(command.AddressStreetName, command.AddressBuildingNumber,
		command.AddressPostalCode, command.AddressCity, command.AddressStateProvince, command.AddressCountryCode)
	if err != nil {
		return errors.Wrap(err, "an error occurred while creating administrator postal address")
	}
	primaryTelephone, err := model.NewTelephone(command.PrimaryTelephone)
	if err != nil {
		return errors.Wrap(err, "an error occurred while creating administrator primary telephone")
	}
	var secondaryTelephone *model.Telephone
	if strings.TrimSpace(command.SecondaryTelephone) != "" {
		secondaryTelephone, err = model.NewTelephone(command.SecondaryTelephone)
		if err != nil {
			return errors.Wrap(err, "an error occurred while creating administrator secondary telephone")
		}
	}
	_, err = ias.tenantProvisioningService.ProvisionTenant(
		ctx,
		command.TenantName,
		command.TenantDescription,
		administratorName,
		emailAddress,
		postalAddress,
		primaryTelephone,
		secondaryTelephone,
	)
	return err
}

// ActivateTenant will activate the tenant with id provided in command.
func (ias *IdentityApplicationService) ActivateTenant(ctx context.Context, command *command.ActivateTenant) error {
	if err := ias.validate.Struct(command); err != nil {
		return err
	}
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
	if err := ias.validate.Struct(command); err != nil {
		return err
	}
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
