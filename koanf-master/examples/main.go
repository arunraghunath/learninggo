package main

import (
	"github.com/arunraghunath/koanf"
	"github.com/arunraghunath/koanf/parsers/json"
	"github.com/arunraghunath/koanf/providers/file"
)

func main() {
	k := koanf.New(".")
	fp := file.Provider("../mock/mock.json")
	pa := json.Parser()
	k.Load(fp, pa)
}
