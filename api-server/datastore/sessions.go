package datastore

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	m "github.com/lukejmann/detach2-backend/api-server/models"

	// m "../models"

	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionsStore struct {
	usersC *mongo.Collection
}

func genUUID() (string, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

func (s *SessionsStore) Create(opt m.SessionCreateOpt) (res *m.SessionCreateRes, err error) {
	id, err := genUUID()
	if err != nil {
		return nil, err
	}
	session := m.Session{
		UserID:      opt.UserID,
		ID:          id,
		DeviceToken: opt.DeviceToken,
		EndTime:     opt.EndTime,
		// AppNames:    opt.AppNames,
	}

	filter := bson.M{"_id": opt.UserID}
	update := bson.M{"$push": bson.M{"sessions": session}}

	mRes, err := s.usersC.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	pp.Println("mRes:", mRes)

	success := mRes.ModifiedCount == 1
	fmt.Printf("Session created in Store. userID: %v. sessionID %v\n", opt.UserID, id)
	res = &m.SessionCreateRes{
		Success:   success,
		SessionID: id,
	}
	return res, nil
}

func (s *SessionsStore) Cancel(opt m.SessionCancelOpt) (success bool, err error) {
	filter := bson.M{"_id": opt.UserID}
	update := bson.M{"$pull": bson.M{"sessions": bson.M{"_id": opt.SessionID}}}

	res, err := s.usersC.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}

	success = res.ModifiedCount == 1
	fmt.Printf("Session deleted in Store. userID: %v. sessionID %v\n", opt.UserID, opt.SessionID)

	return success, nil
}
