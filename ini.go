package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
)

var (
	gCommentRegular  = `^[ \t]*[;#]`
	gSectionRegular  = `^[ \t]*\[(.*)\][ \t]*`
	gKeyValueRegular = `^[ \t]*([^ =\t]+)[ \t=]+([^ =\t\n]+)`
	gKeyRegular      = `^[ \t]*([^ =\t\n]+)`
	gLineBreak       = byte('\n')
)

type IniConfig struct{}

var gIniConfig IniConfig

func (p *IniConfig) ParseFile(name string) (ConfigContainer, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, errors.New("config file not exist:")
	}

	defer file.Close()

	//panic if error
	//we can add it ot init?
	sCommentRegexp := regexp.MustCompile(gCommentRegular)
	sSectionRegexp := regexp.MustCompile(gSectionRegular)
	sKeyValueRegexp := regexp.MustCompile(gKeyValueRegular)
	sKeyRegexp := regexp.MustCompile(gKeyRegular)

	cfg := &IniConfigContainer{data: make(map[string]Section)}
	buf := bufio.NewReader(file)
	var section string
	for {
		line, err := buf.ReadString(gLineBreak)
		if err == io.EOF {
			break
		}

		if sCommentRegexp.FindString(line) != "" {
			//we don't want to keep it
			continue
		}

		matchs := sSectionRegexp.FindStringSubmatch(line)
		if len(matchs) == 2 {
			section = matchs[1]
			cfg.data[section] = make(Section)
			continue
		}

		matchs = sKeyValueRegexp.FindStringSubmatch(line)
		if len(matchs) == 3 {
			if cfg.data[section] != nil {
				cfg.data[section][matchs[1]] = matchs[2]
			}
			continue
		}

		matchs = sKeyRegexp.FindStringSubmatch(line)
		if len(matchs) == 2 {
			cfg.data[section][matchs[1]] = ""
			continue
		}
	}
	return cfg, nil
}

//container
type IniConfigContainer struct {
	data map[string]Section
}

func (p *IniConfigContainer) Section(section string) (Sectioner, error) {
	if m, ok := p.data[section]; ok {
		return &m, nil
	}

	return nil, errors.New("no such section")
}
func (p *IniConfigContainer) Int(section, key string) (int, error) {
	s, err := p.Section(section)
	if err != nil {
		return 0, err
	}

	return s.Int(key)
}

func (p *IniConfigContainer) String(section, key string) (string, error) {
	s, err := p.Section(section)
	if err != nil {
		return "", err
	}

	return s.String(key)
}
func (p *IniConfigContainer) Float64(section, key string) (float64, error) {
	s, err := p.Section(section)
	if err != nil {
		return 0, err
	}

	return s.Float64(key)
}

//section
type Section map[string]string

func (p *Section) String(key string) (string, error) {
	if v, ok := (*p)[key]; ok {
		return v, nil
	}

	return "", errors.New("no suck key")
}
func (p *Section) Float64(key string) (float64, error) {
	if v, ok := (*p)[key]; ok {
		return strconv.ParseFloat(v, 64)

	}

	return 0, errors.New("no suck key")
}
func (p *Section) Int(key string) (int, error) {
	if v, ok := (*p)[key]; ok {
		return strconv.Atoi(v)
	}

	return 0, errors.New("no suck key")
}

func init() {
	Register("ini", &gIniConfig)
}
