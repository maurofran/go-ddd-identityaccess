package command

// ProvisionTenant is the command sent to provision a new tenant.
type ProvisionTenant struct {
	TenantName        string `json:"tenantName" validate:"required"`
	TenantDescription string `json:"tenantDescription" validate:"required"`

	AdministratorFirstName string `json:"administratorFirstName" validate:"required"`
	AdministratorLastName  string `json:"administratorLastName" validate:"required"`
	EmailAddress           string `json:"emailAddress" validate:"required,email"`
	PrimaryTelephone       string `json:"primaryTelephone" validate:"required"`
	SecondaryTelephone     string `json:"secondaryTelephone"`
	AddressStreetName      string `json:"addressStreetName" validate:"required"`
	AddressBuildingNumber  string `json:"addressBuildingNumber"`
	AddressPostalCode      string `json:"addressPostalCode" validate:"required"`
	AddressCity            string `json:"addressCity" validate:"required"`
	AddressStateProvince   string `json:"addressStateProvince" validate:"required"`
	AddressCountryCode     string `json:"addressCountryCode" validate:"required,alpha,length=2"`
}
