package main

import (
	"github.com/arunraghunath/koanf"
)

func main() {
	k := koanf.New(".")
	fp := koanf.Provider("../mock/mock.json")
	pa := koanf.Parser()
	k.Load(fp, pa)
}
