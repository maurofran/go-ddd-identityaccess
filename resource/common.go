package resource

import "github.com/gorilla/mux"

const iasMediaType = "application/vnd.maurofran.ias+json"

type ApiV1Resources struct {
	Router  *mux.Router     `inject:""`
	Tenants *TenantResource `inject:""`
}

func (api *ApiV1Resources) Init() {
	r := api.Router.PathPrefix("/api/v1").Subrouter()
	api.Tenants.init(r)
}
