package timelogger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/filipre/fbchat"
)

type onlineKey struct {
	CUser  fbchat.CUser
	Client string
	Status string
}

type onlineMap map[onlineKey]*Onlinetime

func (oMap onlineMap) diff(other onlineMap) onlineMap {
	res := make(onlineMap)
	for key, val := range other {
		if _, ok := oMap[key]; !ok {
			res[key] = val
		}
	}
	return res
}

func (t *TimeLogger) observe(cUser fbchat.CUser, cookie fbchat.Cookie, interval time.Duration, saveCh chan<- *Onlinetime, errorCh chan<- error) {

	c, err := fbchat.NewClient(cUser, cookie, &http.Client{})
	if err != nil {
		errorCh <- err
		t.Jobs[cUser] = false
		return
	}

	onlineBefore := make(onlineMap)
	onlineNow := make(onlineMap)
	fmt.Printf("[%s]: Logging Job started\n", cUser)

	tick := time.NewTicker(interval * time.Second)

	for {
		select {
		case <-tick.C:
			onlineNow = make(onlineMap)

			online, err := c.ReqOnline()
			if err != nil {
				errorCh <- err
				t.Jobs[cUser] = false
				return
			}

			for id, data := range *online {
				for client, status := range data {
					if client == "status" || status == "offline" {
						continue
					}
					key := onlineKey{CUser: id, Client: client, Status: status}
					onlineNow[key] = nil
				}
			}

			loggedInMap := onlineBefore.diff(onlineNow)
			//fmt.Printf("LOGGED IN (%d): -> %+v\n", len(loggedInMap), loggedInMap)

			loggedOutMap := onlineNow.diff(onlineBefore)
			//fmt.Printf("LOGGED OUT (%d): -> %+v\n", len(loggedOutMap), loggedOutMap)

			//New Client found
			for key := range loggedInMap {
				onlineBefore[key] = &Onlinetime{CUser: key.CUser, Client: key.Client, Status: key.Status, LoggedIn: Now(), LoggedBy: cUser}
			}

			//Client changed his status
			for key, inactive := range loggedOutMap {
				inactive.LoggedOut = Now()

				saveCh <- inactive
				delete(onlineBefore, key)
			}

			fmt.Printf("[%s]: New Client (%d), Saved to DB (%d), Observing (%d)\n", cUser, len(loggedInMap), len(loggedOutMap), len(onlineNow))

		case <-t.Done[cUser]:
			fmt.Printf("[%s]: stopped\n", cUser)
			//t.Jobs[cUser] = false TODO...
			return
		}
	}
}
