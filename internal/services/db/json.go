package db

import (
	"encoding/json"
	"fmt"
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
	"io/ioutil"
	"log"
	"os"
)

//const catalogsdirectory = "data/"
const catalogsdirectory = ""

type Catlist struct {
	DAU      []ex.Dau
	Partners []ex.Dau
}

func GetCatalogs(c *Catlist) {
	name := "dau"

	catalogFilename := fmt.Sprintf("%s%s.json", catalogsdirectory, name)
	f, err := os.OpenFile(catalogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	rawDataIn, err := ioutil.ReadFile(catalogFilename)
	if err != nil {
		log.Println("Cannot load catalog:", err)
	}

	err = json.Unmarshal(rawDataIn, &c.DAU)
	if err != nil {
		log.Println("Invalid catalogs format:", err)
	}

	name = "partners"

	catalogFilename = fmt.Sprintf("%s%s.json", catalogsdirectory, name)
	f, err = os.OpenFile(catalogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	rawDataIn, err = ioutil.ReadFile(catalogFilename)
	if err != nil {
		log.Println("Cannot load catalog:", err)
	}

	err = json.Unmarshal(rawDataIn, &c.Partners)
	if err != nil {
		log.Println("Invalid catalogs format:", err)
	}
}

func RewriteStorage(catalog *Catlist) error {
	//заливаем catalog обратно в json
	rawDataOut, err := json.MarshalIndent(&catalog.DAU, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	catalogFilename := fmt.Sprintf("%s%s.json", catalogsdirectory, "dau")
	err = ioutil.WriteFile(catalogFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated catalog file:", err)
	}

	rawDataOut, err = json.MarshalIndent(&catalog.Partners, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	catalogFilename = fmt.Sprintf("%s%s.json", catalogsdirectory, "partners")
	err = ioutil.WriteFile(catalogFilename, rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write updated catalog file:", err)
	}

	return nil
}
