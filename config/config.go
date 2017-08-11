package config

import (
	"fmt"
	"strings"
	"sync"

	hc "github.com/endeveit/go-snippets/cli"
	c "github.com/robfig/config"
)

var (
	conf *c.Config
	once sync.Once
)

const (
	SECTION_DEFAULT = "DEFAULT"
)

func initConfig(paths []string) {
	if len(paths) == 1 {
		once.Do(func() {
			var err error

			conf, err = c.ReadDefault(paths[0])
			hc.CheckError(err)
		})
	}
}

// Reads configuration file from provided path and returns its representation
// You must provide path to the configuration file on first call:
//  package main
//  c := config.Instance("/path/to/file.cfg")
// Then you can call function without arguments:
//  package otherpackage
//  c := config.Instance()
func Instance(paths ...string) *c.Config {
	initConfig(paths)

	return conf
}

// Dumps configuration to string
func Dump(conf *c.Config) string {
	var result []string

	for _, section := range conf.Sections() {
		if section != SECTION_DEFAULT {
			result = append(result, "["+section+"]")
		}

		if options, err := conf.SectionOptions(section); err == nil {
			for _, option := range options {
				if value, err := conf.RawString(section, option); err == nil {
					result = append(result, fmt.Sprintf("%s = %s", option, value))
				}
			}
		}
	}

	return strings.Join(result, "\r\n")
}
