package db

// "github.com/NautiloosGo/ga/internal/app/exceller"
const FileDir = "./originaldata/*.csv"

type Dau struct {
	Sourse      string
	PartnerName string
	Date        uint64
	Dau         uint64
}
type Chans struct {
	IncomingChan chan Dau
}

func (c *Chans) FindAndProccessFiles() {
	exceller.c.Switcher(FileDir)
	// exceller.c.Switcher(FileDir)
}
