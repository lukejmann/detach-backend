package api

import (
	"github.com/lukejmann/detach2-backend/api-server/datastore"
)

var store *datastore.Datastore

func GetDatastore() (*datastore.Datastore, error) {
	if store != nil {
		return store, nil
	}
	var err error
	store, err = datastore.NewDatastore()
	if err != nil {
		return nil, err
	}
	// fmt.Printf("")
	return store, nil
}
