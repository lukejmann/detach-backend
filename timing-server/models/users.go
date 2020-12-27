package models

type User struct {
	UserID string `bson:"_id"`
	Email  string `bson:"email"`
	// AppleReciept string    `bson:"appleReciept"`
	Sessions []Session `bson:"sessions"`
}

type LoginResult struct {
	Success bool `json:"success"`
	// SubStatus SubStatus `json:"subStatus"`
}

// type SubStatus struct {
// 	//"active" or "inactive"
// 	Status  string `json:"status"`
// 	SubDate int    `json:"subDate"`
// }

// type CheckReceiptOpt struct {
// 	UserID       string `json:"userID"`
// 	AppleReciept string `bson:"appleReciept"`
// }

// type CheckReceiptRes struct {
// 	Success   bool      `json:"success"`
// 	SubStatus SubStatus `json:"subStatus"`
// }
