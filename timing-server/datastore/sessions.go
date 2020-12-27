package datastore

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"
	m "github.com/lukejmann/detach2-backend/timing-server/models"

	// m "../models"

	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (s *SessionsStore) CompletedSessions(beforeTime int) (sessions []m.Session, err error) {
	ctx := context.TODO()
	filter := bson.M{}
	cur, err := s.usersC.Find(ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user m.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		for _, session := range user.Sessions {
			if session.EndTime < beforeTime {
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}

func (s *SessionsStore) Cancel(opt m.SessionCancelOpt) (success bool, err error) {
	filter := bson.M{"_id": opt.UserID}
	update := bson.M{"$pull": bson.M{"sessions": bson.M{"_id": opt.SessionID}}}

	res, err := s.usersC.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}

	pp.Println("res:", res)

	success = res.ModifiedCount == 1
	fmt.Printf("Session deleted in Store. success: %v. userID: %v. sessionID %v\n", success, opt.UserID, opt.SessionID)

	return success, nil
}
