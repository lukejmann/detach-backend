package timing

import (
	"log"
	"time"

	"github.com/k0kubun/pp"
	m "github.com/lukejmann/detach2-backend/timing-server/models"
)

//TODO: remove fatal
func (tM *TimingManager) updateDueSessions(startTime time.Time, endTime time.Time) error {

	// startTimeUnix := int(startTime.Unix())
	endTimeUnix := int(endTime.Unix())
	// fmt.Printf("In update due sessions. starttime: %v. endtime: %v\n", startTimeUnix, endTimeUnix)

	completedSessions, err := tM.store.Sessions.CompletedSessions(endTimeUnix)
	if err != nil {
		log.Fatal(err)
	}

	// pp.Println("completedSessions: ", completedSessions)
	for _, session := range completedSessions {
		//send notif (session)
		notif := m.Notif{
			DeviceToken: session.DeviceToken,
			SessionID:   session.ID,
			// SentTime:
			Text: "Blocking Session Completed.",
		}
		tM.notifManager.SendRefreshForegroundNotification(notif)

		opt := m.SessionCancelOpt{
			UserID:    session.UserID,
			SessionID: session.ID,
		}
		success, err := tM.store.Sessions.Cancel(opt)
		if err != nil {
			pp.Printf("failed to cancel session! err: %v. session: %v\n", err.Error(), session)

		}
		if !success {
			pp.Printf("failed to cancel session! session: %v\n", session)
		}

	}

	return nil
}
