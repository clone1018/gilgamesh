package loader

import (
	"io/ioutil"
)

type PluginLoader struct {
	Dir string
}

func (l *PluginLoader) Load() error {
	_, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	return nil
}
