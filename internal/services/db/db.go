package db
var KeyPhrases []string

const FileDir="./Original_Data/*.csv"


func  (c *Chans) FindAndProccessFiles(){
	c.Switcher(FileDir)
}

func  (c *Chans) Dau(){
	c.Switcher(FileDir, KeyPhrases)
}