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

type iniConfig struct{}

var gIniConfig iniConfig

func (p *iniConfig) parseFile(name string) (ConfigContainer, error) {
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

	cfg := &iniConfigContainer{data: make(map[string]section)}
	buf := bufio.NewReader(file)
	var sectionName string
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
			sectionName = matchs[1]
			cfg.data[sectionName] = make(section)
			continue
		}

		matchs = sKeyValueRegexp.FindStringSubmatch(line)
		if len(matchs) == 3 {
			if cfg.data[sectionName] != nil {
				cfg.data[sectionName][matchs[1]] = matchs[2]
			}
			continue
		}

		matchs = sKeyRegexp.FindStringSubmatch(line)
		if len(matchs) == 2 {
			cfg.data[sectionName][matchs[1]] = ""
			continue
		}
	}
	return cfg, nil
}

//container
type iniConfigContainer struct {
	data map[string]section
}

func (p *iniConfigContainer) Section(section string) (Sectioner, error) {
	if m, ok := p.data[section]; ok {
		return &m, nil
	}

	return nil, errors.New("no such section")
}
func (p *iniConfigContainer) Int(section, key string) (int64, error) {
	s, err := p.Section(section)
	if err != nil {
		return 0, err
	}

	return s.Int(key)
}

func (p *iniConfigContainer) String(section, key string) (string, error) {
	s, err := p.Section(section)
	if err != nil {
		return "", err
	}

	return s.String(key)
}
func (p *iniConfigContainer) Float64(section, key string) (float64, error) {
	s, err := p.Section(section)
	if err != nil {
		return 0, err
	}

	return s.Float64(key)
}

//section
type section map[string]string

func (p *section) String(key string) (string, error) {
	if v, ok := (*p)[key]; ok {
		return v, nil
	}

	return "", errors.New("no suck key")
}
func (p *section) Float64(key string) (float64, error) {
	if v, ok := (*p)[key]; ok {
		return strconv.ParseFloat(v, 64)

	}

	return 0, errors.New("no suck key")
}
func (p *section) Int(key string) (int64, error) {
	if v, ok := (*p)[key]; ok {
		return string2int(v)
	}

	return 0, errors.New("no suck key")
}

func init() {
	Register("ini", &gIniConfig)
}
