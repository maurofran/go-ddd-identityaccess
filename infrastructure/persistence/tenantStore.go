package persistence

import (
	"github.com/maurofran/go-ddd-identityaccess/domain/model"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const tenants = "tenants"

// TenantStore is the mongo db related store for tenants.
type TenantStore struct {
	Db *mgo.Database `inject:""`
}

func (ts *TenantStore) tenants() *mgo.Collection {
	return ts.Db.C(tenants)
}

// Insert will insert a new tenant into the store.
func (ts *TenantStore) Insert(tenantID string, name, description string, active bool) (interface{}, error) {
	id := bson.NewObjectId()
	err := ts.tenants().Insert(bson.M{
		"_id":         id,
		"_v":          0,
		"tenantId":    tenantID,
		"name":        name,
		"description": description,
		"active":      active,
	})
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred while adding tenant to store")
	}
	return id, nil
}

// Update will update an existing tenant actually stored into repository, returning an error if the operation fails.
func (ts *TenantStore) Update(id interface{}, version int, name, description string, active bool) (int, error) {
	err := ts.tenants().Update(
		bson.M{
			"_id": id,
			"_v":  version,
		},
		bson.M{
			"$inc": bson.M{
				"_v": 1,
			},
			"$set": bson.M{
				"name":        name,
				"description": description,
				"active":      active,
			},
		},
	)
	if err != nil {
		return version, errors.Wrap(err, "an error occurred while updating tenant")
	}
	return version + 1, nil
}

// Delete will remove an existing tenant from repository, returning an error if the operation fails.
func (ts *TenantStore) Delete(id interface{}, version int) error {
	err := ts.tenants().Remove(bson.M{
		"_id": id,
		"_v":  version,
	})
	if err != nil {
		return errors.Wrap(err, "an error occurred while removing tenant from store")
	}
	return nil
}

func (ts *TenantStore) FindOneByTenantID(tenantID string) (model.TenantStoreItem, error) {
	res := tenantStoreItem{}
	err := ts.tenants().Find(bson.M{"tenantId": tenantID}).One(&res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "an error occurred while retrieving tenant from store")
	}
	return res, nil
}

func (ts *TenantStore) FindOneByName(name string) (model.TenantStoreItem, error) {
	res := tenantStoreItem{}
	err := ts.tenants().Find(bson.M{"name": name}).One(&res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "an error occurred while retrieving tenant from store")
	}
	return res, nil
}

type tenantStoreItem map[string]interface{}

func (tsi tenantStoreItem) ID() interface{} {
	return tsi["_id"]
}

func (tsi tenantStoreItem) Version() int {
	return tsi["_v"].(int)
}

func (tsi tenantStoreItem) TenantID() string {
	return tsi["tenantId"].(string)
}

func (tsi tenantStoreItem) Name() string {
	return tsi["name"].(string)
}

func (tsi tenantStoreItem) Description() string {
	return tsi["description"].(string)
}

func (tsi tenantStoreItem) Active() bool {
	return tsi["active"].(bool)
}
