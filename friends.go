package fbchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FriendsJSON struct {
	Ar      int `json:"__ar"`
	Payload struct {
		Profiles Friends `json:"profiles"`
	} `json:"payload"`
	Error        int    `json:"error"`
	ErrorSummary string `json:"errorSummary"`
}

type Friends map[CUser]Profile

type Profile struct {
	Id     CUser  `json:"id"`
	Name   string `json:"name"`
	Vanity string `json:"vanity"`
	Gender Gender `json:"gender"`
}

type Gender int

func (g Gender) String() string {
	switch g {
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "other"
	}
}

func (c Client) ReqFriendsJSON(ids ...CUser) (*FriendsJSON, error) {
	//make request url and cookie
	idsStr := ""
	for i, id := range ids {
		idsStr = idsStr + fmt.Sprintf("&ids[%d]=%s", i, id)
	}
	url := fmt.Sprintf("https://www.facebook.com/chat/user_info/?__user=%s&__a=1%s", c.CUser, idsStr)
	cookie := fmt.Sprintf("c_user=%s;datr=%s;xs=%s;", c.CUser, c.Cookie.Datr, c.Cookie.Xs)

	//do the request with
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", cookie)
	res, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	//recieve the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(body[len("for (;;);"):]))
	var f FriendsJSON
	if err = decoder.Decode(&f); err != nil {
		return nil, err
	}
	if f.Error != 0 {
		return nil, fmt.Errorf("Login into FB failed. Maybe wrong cookies? Facebook Error (%d): %s", f.Error, f.ErrorSummary)
	}
	return &f, nil
}

func (c Client) ReqFriends(ids ...CUser) (*Friends, error) {
	f, err := c.ReqFriendsJSON(ids...)
	if err != nil {
		return nil, err
	}
	return &f.Payload.Profiles, nil
}
