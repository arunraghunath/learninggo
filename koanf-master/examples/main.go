package main

import (
	"github.com/arunraghunath/koanf"
	"github.com/arunraghunath/koanf/providers/file"
	"github.com/arunraghunath/learninggo/koanf/parsers/json"
)

func main() {
	k := koanf.New(".")
	fp := file.Provider("../mock/mock.json")
	pa := json.Parser()
	k.Load(fp, pa)
}
