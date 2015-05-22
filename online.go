package fbchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OnlineJSON struct {
	Ar      int `json:"__ar"`
	Payload struct {
		Time      int `json:"time"`
		BuddyList struct {
			NowAvailableList map[CUser]BuddyInformation `json:"nowAvailableList"`
		} `json:"buddy_list"`
	} `json:"payload"`
	ErrorCode    int    `json:"error"`
	ErrorSummary string `json:"errorSummary"`
}

type BuddyInformation struct {
	I bool         `json:"i"`
	A int          `json:"a"`
	C int          `json:"c"`
	P OnlineStatus `json:"p"`
}

type OnlineStatus map[string]string

type Online map[CUser]map[string]string

type OnlineJSONError OnlineJSON

func (j OnlineJSONError) Error() string {
	return fmt.Sprintf("login into FB failed, maybe wrong cookies? Facebook error (%d): %s", j.ErrorCode, j.ErrorSummary)
}

func (c Client) ReqOnlineJSON() (*OnlineJSON, error) {
	url := fmt.Sprintf("https://www.facebook.com/ajax/chat/buddy_list.php?user=%s&__a=1", c.CUser)
	cookie := fmt.Sprintf("datr=%s;xs=%s;c_user=%s;", c.Cookie.Datr, c.Cookie.Xs, c.CUser)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", cookie)
	res, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(body[len("for (;;);"):]))
	var j OnlineJSON
	err = decoder.Decode(&j)
	if err != nil {
		return nil, err
	}
	if j.ErrorCode != 0 {
		return nil, OnlineJSONError(j)
	}
	return &j, nil
}

func (c Client) ReqOnline() (*Online, error) {
	j, err := c.ReqOnlineJSON()
	if err != nil {
		return nil, err
	}
	o := make(Online)
	for id, status := range j.Payload.BuddyList.NowAvailableList {
		o[id] = status.P
	}
	return &o, nil
}
