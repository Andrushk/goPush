package entity

import "time"

type Device struct {
	Id         ID
	Token      string
	Registered time.Time
}