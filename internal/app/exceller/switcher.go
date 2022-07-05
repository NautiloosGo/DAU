package exceller

import (
	"fmt"
	"path/filepath"
	"strings"
)

var keyDAU = "DAU"
var keyPartners = "Partners"

type Dau struct {
	Sourse      string
	PartnerName string
	Date        int
	Dau         int
}
type Chans struct {
	IncomingDau      chan Dau
	IncomingPartners chan Dau
}

func Switcher(c *Chans, dir string) {
	m, err := filepath.Glob(dir)
	if err != nil {
		panic(err)
	}
	for _, filename := range m {
		if strings.Contains(filename, keyDAU) {
			fmt.Println("DAU File found. goroutine shoud start. File name: ", filename)
			go c.ReadAndFormatDAU(filename)
		}
		if strings.Contains(filename, keyPartners) {
			fmt.Println("Partner File found. goroutine shoud start File name: ", filename)
			go c.ReadAndFormatPartners(filename)
		}
	}

}
