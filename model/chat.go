package model

import "time"

type Chat struct {
	SendId   int
	RecvId   int
	SendDate time.Time
	Msg      string
}
