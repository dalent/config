package config

import "errors"

type Sectioner interface {
	Int(key string) (int, error)
	String(key string) (string, error)
	Float64(key string) (float64, error)
}

type ConfigContainer interface {
	Int(section, key string) (int, error)
	String(section, key string) (string, error)
	Float64(section, key string) (float64, error)
	Section(section string) (Sectioner, error)
}

type Configer interface {
	ParseFile(name string) (ConfigContainer, error)
}

//type  ini
var adapters = make(map[string]Configer)

func Register(name string, config Configer) {
	if config == nil {
		panic("config nil:" + name)
	}

	if _, ok := adapters[name]; ok {
		panic("register twice:" + name)
	}
	adapters[name] = config
}

func NewConfiger(adpaterName string, fileName string) (ConfigContainer, error) {
	adpater, ok := adapters[adpaterName]
	if !ok {
		return nil, errors.New("no such type config ")
	}

	return adpater.ParseFile(fileName)
}
