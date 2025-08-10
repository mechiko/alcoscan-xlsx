package processing

import (
	"alcoscanxlsx/domain"
	"alcoscanxlsx/domain/models/application"
	"alcoscanxlsx/reductor"
	"alcoscanxlsx/utility"
	"fmt"
)

type Processing struct {
	domain.Apper
	warnings []string
	errors   []string

	files     []string
	outDir    string
	Korob     map[string][]*utility.CisInfo
	Palet     map[string][]string
	PaletSort []string
}

func New(app domain.Apper) (*Processing, error) {
	k := &Processing{
		Apper: app,
	}
	k.Reset()
	return k, nil
}

func (k *Processing) AddWarn(warn string) {
	k.warnings = append(k.warnings, warn)
}

func (k *Processing) Warnings() []string {
	return k.warnings
}

func (k *Processing) AddError(err string) {
	k.errors = append(k.errors, err)
}

func (k *Processing) Errors() []string {
	return k.errors
}

func (k *Processing) Reset() {
	k.Korob = make(map[string][]*utility.CisInfo)
	k.Palet = make(map[string][]string)
	k.PaletSort = make([]string, 0)
	k.errors = make([]string, 0)
	k.warnings = make([]string, 0)
}

func (k *Processing) GetModel() (*application.Application, error) {
	modelReductor, err := reductor.Instance().Model(domain.Application)
	if err != nil {
		return nil, fmt.Errorf("processing write %w", err)
	}
	model, ok := modelReductor.(*application.Application)
	if !ok {
		return nil, fmt.Errorf("processing write %w", err)
	}
	return model, nil
}
