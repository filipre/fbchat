package fbchat

import (
	"errors"
	"net/http"
)

var (
	ErrNoCUser          = errors.New("CUser is not provided")
	ErrUncompleteCookie = errors.New("Cookie is not fully provided")
	ErrNoHTTPClient     = errors.New("HTTP Client is not provided")
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

func (c Cookie) Check() error {
	if c.Datr == "" || c.Xs == "" {
		return ErrUncompleteCookie
	}
	return nil
}

func NewClient(id CUser, j Cookie, c *http.Client) (*Client, error) {
	if id == "" {
		return nil, ErrNoCUser
	}
	if err := j.Check(); err != nil {
		return nil, err
	}
	if c == nil {
		return nil, ErrNoHTTPClient
	}
	return &Client{CUser: id, Cookie: j, Http: c}, nil
}
