package main

type Dau struct {
	Sourse      string
	PartnerName string
	Date        uint64
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
	Channels.db.FindAndProccessFiles()
}
