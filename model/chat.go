package model

import "time"

type Chat struct {
	Cid      int
	SendId   int
	RecvId   int
	SendDate time.Time
	Msg      string
	UnRead   bool
}
