package fbchat

import (
	"errors"
	"net/http"
)

type Client struct {
	CUser  CUser
	Cookie Cookie
	Http   *http.Client
}

type CUser string

type Cookie struct {
	Datr string `json:"datr"`
	Xs   string `json:"xs"`
}

func NewClient(id CUser, j Cookie, c *http.Client) (*Client, error) {
	if id == "" {
		return nil, errors.New("CUser is not provided")
	}
	if j.Datr == "" || j.Xs == "" {
		return nil, errors.New("Cookie is not fully provided")
	}
	if c == nil {
		return nil, errors.New("HTTP Client is not provided")
	}
	return &Client{CUser: id, Cookie: j, Http: c}, nil
}
