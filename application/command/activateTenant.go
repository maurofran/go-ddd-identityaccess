package command

// ActivateTenant is the command issued to activate the tenant with provided id.
type ActivateTenant struct {
	TenantID string `json:"tenantId" validate:"notempty"`
}