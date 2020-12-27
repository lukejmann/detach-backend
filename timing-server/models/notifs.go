package models

type Notif struct {
	DeviceToken string `bson:"deviceToken"`
	SessionID   string
	// UserID      int
	// SentTime int64 `bson:"sentTime"`

	//optional
	Text string
}
