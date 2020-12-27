package actions

import (
	"github.com/lukejmann/detach2-backend/timing-server/datastore"
	n "github.com/lukejmann/detach2-backend/timing-server/notifications"
	"github.com/segmentio/ksuid"
)

var store *datastore.Datastore

var notifManager *n.NotificationsManager

// var timingManager *timing.TimingManager

func GetDatastore() (*datastore.Datastore, error) {
	if store != nil {
		return store, nil
	}
	var err error
	store, err = datastore.NewDatastore()
	if err != nil {
		return nil, err
	}
	return store, nil
}

func GetNotificationManager() (*n.NotificationsManager, error) {
	if notifManager != nil {
		return notifManager, nil
	}
	var err error
	notifManager, err = n.NewNotificationsManager()
	if err != nil {
		return nil, err
	}
	return notifManager, nil
}

func genNotifID() string {
	id := ksuid.New()
	return id.String()
}
