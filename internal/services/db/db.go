package db

import (
	ex "github.com/NautiloosGo/ga/internal/app/exceller"
)

const FileDir = "./originaldata/*.csv"

func FindAndProccessFiles(c *ex.Chans) {
	ex.Switcher(c, FileDir)
}
