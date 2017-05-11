package resource

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maurofran/go-ddd-identityaccess/application"
	"net/http"
)

// TenantResource is the struct handling the methods for Tenants.
type TenantResource struct {
	IdentityService *application.IdentityService `inject:""`
}

func (tr *TenantResource) init(router *mux.Router) {
	router.Handle("/tenants", Handler(tr.GetTenant)).Methods("GET")
}

// GetTenant will send the tenant with provided tenant id.
func (tr *TenantResource) GetTenant(w http.ResponseWriter, r *http.Request) error {
	tenantId := mux.Vars(r)["tenantId"]
	t, err := tr.IdentityService.Tenant(r.Context(), tenantId)
	if err != nil {
		return err
	}
	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", iasMediaType)
	json.NewEncoder(w).Encode(t)
	return nil
}
