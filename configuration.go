package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/gcfg.v1"
)

var (
	DefaultConfigFile = "config.gcfg"
)

const (
	// The example file kept in version control. We'll copy and load from this
	// by default.
	CONFIG_EXAMPLE = `; gilgamesh chat bot
[discord-bot "gilgamesh"]
email = 
password = 

plugin = clone1018/test
`
)

type DiscordConfig struct {
	Email    string
	Password string
	Plugin   []string
}

type Configuration struct {
	Discord_Bot map[string]*DiscordConfig
}

// Reads the configuration from the config file, copying a config into
// place from the example if one does not yet exist.
func (c *Configuration) load(file string) error {
	err := c.ensureConfigExists(file)
	if err != nil {
		return err
	}

	return gcfg.ReadFileInto(c, file)
}

// Creates the config.gcfg if it does not exist.
func (c *Configuration) ensureConfigExists(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return ioutil.WriteFile(file, []byte(CONFIG_EXAMPLE), 0644)
	} else {
		return nil
	}
}
