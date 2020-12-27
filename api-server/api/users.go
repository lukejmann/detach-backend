package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/lukejmann/detach2-backend/api-server/models"
)

func serveLogin(w http.ResponseWriter, r *http.Request) error {
	userID := mux.Vars(r)["userID"]
	email := mux.Vars(r)["email"]

	store, err := GetDatastore()
	if err != nil {
		return err
	}

	userExists, err := store.Users.Exists(userID)
	if err != nil {
		return err
	}

	if !userExists {
		created, err := store.Users.Create(userID, email)
		if err != nil {
			return err
		}
		fmt.Printf("Created user %v with email %v\n", userID, email)
		// subStatus := m.SubStatus{
		// 	Status:  "invalid",
		// 	SubDate: 0,
		// }
		res := m.LoginResult{
			Success: created,
			// SubStatus: subStatus,
		}
		return writeJSON(w, res)

	} else {
		updated, err := store.Users.UpdateEmail(userID, email)
		if err != nil {
			return err
		}
		fmt.Printf("Logged in user %v with email %v\n", userID, email)
		// appleReceipt, err := store.Users.GetReceipt(userID)
		// if err != nil {
		// 	return err
		// }
		// subStatus, err := receiptToStatus(appleReceipt)
		// if err != nil {
		// 	return err
		// }
		// if subStatus == nil {
		// 	*subStatus = m.SubStatus{
		// 		Status:  "invalid",
		// 		SubDate: 0,
		// 	}
		// }
		res := m.LoginResult{
			Success: updated,
			// SubStatus: *subStatus,
		}
		return writeJSON(w, res)
	}
}

//sessionID is "NA" if invalid
// func serveCheckReceipt(w http.ResponseWriter, r *http.Request) error {
// 	var opt m.CheckReceiptOpt
// 	err := json.NewDecoder(r.Body).Decode(&opt)
// 	if err != nil {
// 		return err
// 	}

// 	// pp.Println("serveCheckReceipt opt:", opt)

// 	subStatus, err := receiptToStatus(opt.AppleReciept)
// 	if err != nil {
// 		return err
// 	}

// 	// pp.Println("serveCheckReceipt subStatus:", subStatus)

// 	store, err := GetDatastore()
// 	if err != nil {
// 		return err
// 	}

// 	success, err := store.Users.UpdateSubStatus(opt.UserID, *subStatus)
// 	if success {
// 		fmt.Printf("API updated user subscription. userID: %v\n", opt.UserID)
// 	} else {
// 		fmt.Printf("API failed to update user subscription. userID: %v. err: %v\n", opt.UserID, err.Error())
// 	}
// 	res := m.CheckReceiptRes{Success: success, SubStatus: *subStatus}
// 	return writeJSON(w, res)
// }

// func receiptToStatus(receipt string) (*m.SubStatus, error) {
// 	if receipt == "" {
// 		status := m.SubStatus{
// 			Status:  "inactive",
// 			SubDate: 0,
// 		}
// 		return &status, nil
// 	}
// 	client := appstore.New()
// 	password := "f3a457be4499482ba0dfa50d83649c18"
// 	req := appstore.IAPRequest{
// 		ReceiptData: receipt,
// 		Password:    password,
// 	}
// 	resp := &appstore.IAPResponse{}
// 	ctx := context.Background()
// 	err := client.Verify(ctx, req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	greatestExpirMS := 0
// 	var currentReceipt appstore.InApp
// 	r := false
// 	for _, receipt := range resp.LatestReceiptInfo {
// 		expirMS, err := strconv.Atoi(receipt.ExpiresDate.ExpiresDateMS)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if expirMS > greatestExpirMS {
// 			currentReceipt = receipt
// 			greatestExpirMS = expirMS
// 			r = true
// 		}
// 	}
// 	if r {
// 		cr := currentReceipt
// 		// isTrial := cr.IsTrialPeriod == "true"
// 		now := time.Now()
// 		nowNano := now.UnixNano()
// 		nowMS := nowNano / 1000000
// 		expirMS, err := strconv.Atoi(cr.ExpiresDate.ExpiresDateMS)
// 		if err != nil {
// 			return nil, err
// 		}
// 		var isExpired bool
// 		// if !isTrial {
// 		// pp.Println("nowMS:", int(nowMS))
// 		isExpired = expirMS < int(nowMS)
// 		// } else {
// 		// 	//add 30 seconds
// 		// 	isExpired = expirMS < (int(nowMS) + 1000*60*30)
// 		// }

// 		var status = "inactive"
// 		if !isExpired {
// 			status = "active"
// 		}
// 		ms, err := strconv.ParseInt(cr.PurchaseDate.PurchaseDateMS, 10, 64)
// 		if err != nil {
// 			return nil, err
// 		}
// 		subDate := int(ms / 1000)

// 		subStatus := m.SubStatus{
// 			Status:  status,
// 			SubDate: subDate,
// 		}
// 		return &subStatus, nil
// 	}
// 	return nil, nil
// }
