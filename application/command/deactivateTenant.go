package command

// DeactivateTenant will deactivate the tenant with provided tenant id.
type DeactivateTenant struct {
	TenantID string `json:"tenantId" validate:"required"`
}