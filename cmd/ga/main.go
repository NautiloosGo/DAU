package main

import (
	"time"

	db "github.com/NautiloosGo/ga/internal/services/db"
)

type Dau struct {
	Sourse      string
	PartnerName string
	Date        time.Date
	Dau         uint64
}
type Chans struct {
	IncomingChan chan Dau
}

var inbox chan Dau

var Channels = Chans{
	IncomingChan: make(chan Dau),
}

func main() {
	db.Channels.FindAndProccessFiles()
}
