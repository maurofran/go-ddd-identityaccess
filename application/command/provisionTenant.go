package command

// ProvisionTenant is the command sent to provision a new tenant.
type ProvisionTenant struct {
	TenantName        string `json:"tenantName" validate:"notempty"`
	TenantDescription string `json:"tenantDescription" validate:"notempty"`

	AdministratorFirstName string `json:"administratorFirstName" validate:"notempty"`
	AdministratorLastName  string `json:"administratorLastName" validate:"notempty"`
	EmailAddress           string `json:"emailAddress" validate:"notempty,email"`
	PrimaryTelephone       string `json:"primaryTelephone" validate:"notempty"`
	SecondaryTelephone     string `json:"secondaryTelephone"`
	AddressStreetName      string `json:"addressStreetName" validate:"notempty"`
	AddressBuildingNumber  string `json:"addressBuildingNumber"`
	AddressPostalCode      string `json:"addressPostalCode" validate:"notempty"`
	AddressCity            string `json:"addressCity" validate:"notempty"`
	AddressStateProvince   string `json:"addressStateProvince" validate:"notempty"`
	AddressCountryCode     string `json:"addressCountryCode" validate:"notempty,length=2"`
}
