package notifications

import (
	"github.com/lukejmann/detach2-backend/timing-server/datastore"
	m "github.com/lukejmann/detach2-backend/timing-server/models"

	"github.com/k0kubun/pp"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"github.com/sideshow/apns2/payload"
)

type NotificationsManager struct {
	client *apns2.Client
}

var store *datastore.Datastore

const topic = "com.detachapp.ios1"

func NewNotificationsManager() (*NotificationsManager, error) {
	cert, err := certificate.FromP12File("notifications/certs/cert5.p12", "")
	if err != nil {
		pp.Printf("Cert err: %v\n", err)
		return nil, err
	}

	client := apns2.NewClient(cert).Development()

	nM := &NotificationsManager{client: client}

	return nM, nil
}

func (nM *NotificationsManager) SendRefreshForegroundNotification(notif m.Notif) {
	pp.Println("deviceToken: ", notif.DeviceToken)
	notification := &apns2.Notification{}
	notification.DeviceToken = notif.DeviceToken
	notification.Topic = topic
	notification.Priority = 10
	notification.PushType = apns2.PushTypeAlert
	// if notifData.Text != nil {
	payload := payload.NewPayload().Alert(notif.Text).ContentAvailable()
	// }

	notification.Payload = payload

	res, err := nM.client.Push(notification)
	if err != nil {
		pp.Println("error in sending Foreground notif!: %v", err)
	}
	if res.StatusCode == 200 {
		pp.Println("successfully sent notif %v\n", notif.Text)
	} else {
		pp.Println("failt to send notif %v. \nres: %v\n", notification, res)
	}
}

func getDatastore() (*datastore.Datastore, error) {
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
