package datastore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//A Datastore accesses the datastore (in MongoDB).
type Datastore struct {
	Users    UsersStore
	Sessions SessionsStore

	usersC *mongo.Collection

	dbc *mongo.Client
}

// NewDatastore creates a new client for accessing the datastore. If dbh is nil, it uses the global DB handle.
func NewDatastore() (*Datastore, error) {
	err, dbc := connect()
	if err != nil {
		return nil, err
	}
	usersC := dbc.Database("detachOne").Collection("users")

	d := &Datastore{dbc: dbc, usersC: usersC}
	d.Users = UsersStore{usersC: usersC}
	d.Sessions = SessionsStore{usersC: usersC}

	return d, nil
}

func connect() (error, *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return err, nil
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err, nil
	}

	return nil, client
}
