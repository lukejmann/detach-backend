package timing

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lukejmann/detach2-backend/timing-server/actions"
	"github.com/lukejmann/detach2-backend/timing-server/datastore"
	n "github.com/lukejmann/detach2-backend/timing-server/notifications"

	"github.com/k0kubun/pp"
)

const REFRESH_INTERVAL = 500 * time.Millisecond
const SUB_REFRESH_INTERVAL = 15

type TimingManager struct {
	store        *datastore.Datastore
	notifManager *n.NotificationsManager
	stopChan     chan bool
}

func NewTimingManager() (*TimingManager, error) {
	store, err := actions.GetDatastore()
	if err != nil {
		return nil, err
	}

	nM, err := actions.GetNotificationManager()
	if err != nil {
		return nil, err
	}

	tM := &TimingManager{store: store, notifManager: nM}

	if err != nil {
		return nil, err
	}
	time.LoadLocation("UTC")

	return tM, nil
}

func (tM *TimingManager) Start() error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var i = 0

	go func() {
		ticker := time.NewTicker(REFRESH_INTERVAL)

		for {
			select {
			case <-ticker.C:
				startTime := time.Now()
				//every 10 sec
				if i%20 == 0 {
					pp.Printf("Tick at %v\n", startTime)
				}

				//every 1 sec
				if i%2 == 0 {
					endTime := startTime.Add(REFRESH_INTERVAL * 2)
					tM.updateDueSessions(startTime, endTime)
				}

				//every 2 sec
				// if i%4 == 0 {
				// 	tM.updateDueNotifs(startTime)
				// }

				i += 1
			case <-sigChan:
				tM.stopChan <- true
			case <-tM.stopChan:
				close(tM.stopChan)
			}
		}
	}()

	return nil
}
