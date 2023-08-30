package koanf

import (
	"fmt"
	"log"
	"os"
)

type Koanf struct {
	conf Conf
}

type Conf struct {
	Delim string
}

func New(delim string) *Koanf {
	return NewWithConf(Conf{
		Delim: delim,
	})
}

func NewWithConf(conf Conf) *Koanf {
	return &Koanf{
		conf: conf,
	}
}

func (k *Koanf) Load(fp *File, pa *JSON) {
	b, err := fp.ReadBytes()
	if err != nil {
		log.Fatal("Error reading file", err)
		os.Exit(1)
	}
	conf, err := pa.Unmarshal(b)
	fmt.Println(conf)

}
