package models

type Session struct {
	ID          string `bson:"_id"`
	UserID      string `json:"userID"`
	DeviceToken string `bson:"deviceToken"`
	//Epoch seconds
	EndTime int `bson:"endTime"`
	// AppNames []string `bson:"appNames"`
}

type SessionCreateOpt struct {
	UserID  string `json:"userID"`
	EndTime int    `json:"endTime"`
	// AppNames    []string `json:"appNames"`
	DeviceToken string `json:"deviceToken"`
}
type SessionCreateRes struct {
	Success   bool   `json:"success"`
	SessionID string `json:"sessionID"`
}

type SessionCancelOpt struct {
	UserID    string `json:"userID"`
	SessionID string `json:"sessionID"`
}
