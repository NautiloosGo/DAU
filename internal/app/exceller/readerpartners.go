package exceller

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (c *Chans) ReadAndFormatPartners(fileName string) {
	NextRow := Dau{
		Sourse:      "Partners",
		PartnerName: "",
		Date:        0,
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
		NextRow.Date = day + month*100 + year*10000
		NextRow.Dau, err = strconv.Atoi(record[4])
		if err != nil {
			fmt.Println(err)
		}
		c.IncomingPartners <- NextRow
	}
}
