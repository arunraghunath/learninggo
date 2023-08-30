package main

import (
	"github.com/arunraghunath/koanf"
	"github.com/arunraghunath/koanf/parsers/json"
	"github.com/arunraghunath/koanf/parsers/yaml"
	"github.com/arunraghunath/koanf/providers/file"
)

func loadjson() {
	k := koanf.New(".")
	fp := file.Provider("../mock/mock.json")
	pa := json.Parser()
	k.Load(fp, pa)
}

func loadyaml() {
	k := koanf.New(".")
	fp := file.Provider("../mock/mock.yml")
	pa := yaml.Parser()
	k.Load(fp, pa)
}

func main() {
	//loadjson()
	loadyaml()

}
