package config

import "errors"

type sectioner interface {
	Int(key string) (int64, error)
	String(key string) (string, error)
	Float64(key string) (float64, error)
}

type configContainer interface {
	Int(section, key string) (int64, error)
	String(section, key string) (string, error)
	Float64(section, key string) (float64, error)
	Section(section string) (sectioner, error)
}

type configer interface {
	parseFile(name string) (configContainer, error)
}

//type  ini
var adapters = make(map[string]configer)

func Register(name string, config configer) {
	if config == nil {
		panic("config nil:" + name)
	}

	if _, ok := adapters[name]; ok {
		panic("register twice:" + name)
	}
	adapters[name] = config
}

func NewConfiger(adpaterName string, fileName string) (configContainer, error) {
	adpater, ok := adapters[adpaterName]
	if !ok {
		return nil, errors.New("no such type config ")
	}

	return adpater.parseFile(fileName)
}
