package application

import (
	"alcoscanxlsx/config"
	"alcoscanxlsx/domain"
	"fmt"
)

type Application struct {
	model   domain.Model
	Title   string
	Output  string
	Debug   bool
	License string
	//
	InDir  string
	OutDir string
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper, inDir string) (*Application, error) {
	model := &Application{
		model: domain.Application,
		Title: "Application Title",
		InDir: inDir,
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (a *Application) SyncToStore(app domain.Apper) (err error) {
	// ...
	return err
}

// читаем состояние приложения
func (a *Application) ReadState(app domain.Apper) (err error) {
	a.Output = app.Options().Output
	a.Debug = config.Mode == "development"
	a.License = app.Options().Application.License
	return nil
}

func (a *Application) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *a
	return &dst, nil
}

func (a *Application) Model() domain.Model {
	return a.model
}

func (a *Application) Reset() {
}
