package processing

import (
	"alcoscanxlsx/domain"
	"alcoscanxlsx/domain/models/application"
	"alcoscanxlsx/reductor"
	"fmt"
)

// const startSSCC = "1462709225" // gs1 rus id zapivkom для памяти запивком

type Processing struct {
	domain.Apper
	warnings []string
	errors   []string
}

func New(app domain.Apper) (*Processing, error) {
	return nil, nil
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
