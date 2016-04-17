package loader

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

type ModubotLoader struct {
	Dir string
	Vm  *otto.Otto
}

func (l *ModubotLoader) Load() error {
	_, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	return nil
}

func (l *ModubotLoader) Run() error {
	l.Vm = otto.New()

	return nil
}

func (l *ModubotLoader) setup() error {

	return nil
}
