package fbchat

import (
	"net/http"
	"testing"
)

func TestNewClientNoCUser(t *testing.T) {
	cookie := Cookie{Datr: "datr", Xs: "xs"}
	client := &http.Client{}

	//No CUser provided
	if _, err := NewClient("", cookie, client); err != ErrNoCUser {
		t.Errorf("Expected Err %s; got Err %s", ErrNoCUser, err)
	}
}

func TestNewClientBadCookie(t *testing.T) {
	cUser := CUser("10001")
	client := &http.Client{}

	//only Datr provided
	if _, err1 := NewClient(cUser, Cookie{Datr: "datr"}, client); err1 != ErrUncompleteCookie {
		t.Errorf("Expected Err %s; got Err %s", ErrUncompleteCookie, err1)
	}

	//only Xs provided
	if _, err2 := NewClient(cUser, Cookie{Xs: "xs"}, client); err2 != ErrUncompleteCookie {
		t.Errorf("Expected Err %s; got Err %s", ErrUncompleteCookie, err2)
	}

	//nothing provided
	if _, err3 := NewClient(cUser, Cookie{}, client); err3 != ErrUncompleteCookie {
		t.Errorf("Expected Err %s; got Err %s", ErrUncompleteCookie, err3)
	}
}

func TestNewClientNoHTTPClient(t *testing.T) {
	cUser := CUser("10001")
	cookie := Cookie{Datr: "datr", Xs: "xs"}

	//No CUser provided
	if _, err := NewClient(cUser, cookie, nil); err != ErrNoHTTPClient {
		t.Errorf("Expected Err %s; got Err %s", ErrNoHTTPClient, err)
	}
}
