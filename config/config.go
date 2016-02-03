package config

import (
	"sync"

	hc "github.com/endeveit/go-snippets/cli"
	c "github.com/robfig/config"
)

var (
	conf *c.Config
	once sync.Once
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
