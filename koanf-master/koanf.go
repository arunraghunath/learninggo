package koanf

import (
	"fmt"
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

func (k *Koanf) Load(p Provider, pa Parser) error {
	if p == nil {
		fmt.Errorf("Load received nil provider")
	}
	b, err := p.ReadBytes()
	if err != nil {
		return err
	}
	conf, err := pa.Unmarshal(b)
	if err != nil {
		return err
	}

	fmt.Println(conf)
	return nil
}
