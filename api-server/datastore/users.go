package datastore

import (
	"context"
	"fmt"

	m "github.com/lukejmann/detach2-backend/api-server/models"

	// m "../models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersStore struct {
	usersC *mongo.Collection
}

func (s *UsersStore) Create(userID string, email string) (success bool, err error) {
	user := m.User{
		UserID: userID,
		Email:  email,
		// AppleReciept: "",
		Sessions: []m.Session{},
	}

	res, err := s.usersC.InsertOne(context.TODO(), user)
	if err != nil {
		return false, err
	}

	success = res.InsertedID != nil
	fmt.Println("User created in Store. userID: ", res.InsertedID)

	return success, nil
}

func (s *UsersStore) Exists(userID string) (bool, error) {
	filter := bson.M{"_id": userID}
	var user m.User
	err := s.usersC.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *UsersStore) UpdateEmail(userID string, email string) (bool, error) {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"email": email}}
	res, err := s.usersC.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	fmt.Println("User email updated in store. userID: %v email: %v. updated: %v\n", userID, email, res.MatchedCount == 1)
	return res.MatchedCount == 1, nil
}

// func (s *UsersStore) GetReceipt(userID string) (string, error) {
// 	filter := bson.M{"_id": userID}
// 	var user m.User
// 	err := s.usersC.FindOne(context.TODO(), filter).Decode(&user)
// 	if err != nil {
// 		return "", err
// 	}
// 	return user.AppleReciept, nil
// }

// func (s *UsersStore) UpdateSubStatus(userID string, subStatus m.SubStatus) (bool, error) {
// 	filter := bson.M{"_id": userID}
// 	update := bson.M{"$set": bson.M{"subStatus": subStatus}}
// 	res, err := s.usersC.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		return false, err
// 	}
// 	fmt.Printf("User sub status updated in store. userID: %v email: %v. subStatus: %v\n", userID, subStatus, res.MatchedCount == 1)
// 	return res.MatchedCount == 1, nil
// }
