package exceller

import (
	"fmt"
	"path/filepath"
	"strings"
)

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
