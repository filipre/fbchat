package timelogger

import (
	"time"

	"github.com/filipre/fbchat"
)

type TimeLogger struct {
	Done Done
	Jobs Jobs
}

type Done map[fbchat.CUser]chan bool
type Jobs map[fbchat.CUser]bool

func New() *TimeLogger {
	return &TimeLogger{
		Done: make(Done),
		Jobs: make(Jobs),
	}
}

func (t *TimeLogger) Start(id fbchat.CUser, cookie fbchat.Cookie, interval time.Duration, saveCh chan<- *Onlinetime, errorCh chan<- error) {

	if t.Jobs[id] == true {
		t.Done[id] <- true
	}

	t.Jobs[id] = true
	t.Done[id] = make(chan bool)

	metaErrCh := make(chan error)
	go observe(id, cookie, interval, saveCh, t.Done[id], metaErrCh)

	select {
	case err := <-metaErrCh:
		t.Jobs[id] = false
		errorCh <- err
		return
	}
}

func (t *TimeLogger) Stop(id fbchat.CUser) {

	if t.Jobs[id] == false {
		return
	}

	t.Done[id] <- true
	t.Jobs[id] = false
}

type Onlinetime struct {
	CUser     fbchat.CUser
	Client    string
	Status    string
	LoggedIn  Date
	LoggedOut Date
	LoggedBy  fbchat.CUser
}

type Date string

const DateLayout = "2006-01-02 15:04:05" //here is some weird bug. see http://play.golang....tood..

func Now() Date {
	return Date(time.Now().Format(DateLayout))
}
