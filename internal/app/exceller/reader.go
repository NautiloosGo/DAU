package exceller

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/NautiloosGo/ga/cmd/ga"
)

func (c *Chans) ReadAndFormatPartners(fileName string) {

	NextRow := ga.Dau{
		Sourse:      "Partners",
		PartnerName: "",
		Date:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Dau:         0,
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'
	reader.Read()
	reader.Read()

	defer file.Close()

	for {

		record, err := reader.Read()

		if err == io.EOF {
			// close(c.IncomingChan)
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		rname := []rune{}
		if record[0] == "/partners/" || record[0] == "/partners" || strings.Count(record[0], "/partners?") != 0 {
			NextRow.PartnerName = "Catalogue"
		} else {
			_, err := fmt.Sscanf(record[0], "/partners/%s", &record[0])
			if err != nil {
				record[0] = "ERROR"
				fmt.Println(err)
			}
			for _, a := range record[0] {
				if a == '/' {
					break
				} else {
					rname = append(rname, a)
				}
			}
			NextRow.PartnerName = string(rname)
		}

		day, err := strconv.Atoi(record[3])
		if err != nil {
			fmt.Println(err)
		}
		month, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println(err)
		}
		year, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println(err)
		}
		NextRow.Date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		NextRow.Dau, err = strconv.Atoi(record[4])
		if err != nil {
			fmt.Println(err)
		}
		c.IncomingChan <- NextRow
	}
}

func (c *Chans) ReadAndFormatDAU(filename string) {
	NextRow := ga.Dau{
		Sourse:      "DAU",
		PartnerName: "",
		Date:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		Dau:         0,
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.Comment = '#'
	reader.Read()
	reader.Read()

	defer file.Close()

	rname := []rune{}
	for {
		record2, err := reader.Read()
		if err == io.EOF {
			//close(c.IncomingChan)
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		rname = []rune{}
		for _, a := range record2[5] {
			if unicode.IsNumber(a) {
				rname = append(rname, a)
			}
		}
		NextRow.Dau, err = strconv.Atoi(string(rname))
		if err != nil {
			fmt.Println(err)
		}

		month, err := strconv.Atoi(string(record2[3]))
		if err != nil {
			fmt.Println(err)
		}
		day, err := strconv.Atoi(string(record2[2]))
		if err != nil {
			fmt.Println(err)
		}
		year, err := strconv.Atoi(string(record2[4]))
		if err != nil {
			fmt.Println(err)
		}
		NextRow.Date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		c.IncomingChan <- NextRow

	}

}
