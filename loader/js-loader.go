package loader

import (
	"io/ioutil"
	"path"

	"github.com/robertkrimen/otto"
)

type JsLoader struct {
	Dir         string
	EntrySource []byte
	Vm          *otto.Otto
}

func (l *JsLoader) Load() error {
	_, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	l.EntrySource, err = ioutil.ReadFile(path.Join(l.Dir, "main.js"))
	if err != nil {
		return err
	}

	l.Vm = otto.New()

	return nil
}

func (l *JsLoader) Run() (otto.Value, error) {
	return l.Vm.Run(l.EntrySource)
}
