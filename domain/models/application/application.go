package application

import (
	"alcoscanxlsx/config"
	"alcoscanxlsx/domain"
	"fmt"
	"time"
)

type Zakaz struct {
	Magazin  string
	Quantity string
	Gtin     string
	ID       string
	CodeAP   string // ищем в А3 в справочнике для контроля
}

type Application struct {
	model        domain.Model
	Title        string
	Export       string
	Browser      string
	BrowserList  []string
	Output       string
	Debug        bool
	Host         string
	Port         string
	DbLiteDesc   string
	DbConfigDesc string
	DbZnakDesc   string
	DbA3Desc     string
	License      string
	FsrarID      string
	startTime    time.Time
	endTime      time.Time
	period       string
	// специфично для белзаказ
	File              string
	Reestr            map[string][]*Zakaz
	Magazins          []string
	Magazin           string
	SerialNumberType  string
	CisType           string
	ContactPerson     string
	ReleaseMethodType string
	CreateMethodType  string
	PaymentType       string
	TemplateId        string
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper, file string) (*Application, error) {
	model := &Application{
		model:    domain.Application,
		Title:    "Application Title",
		File:     file,
		Reestr:   make(map[string][]*Zakaz),
		Magazins: make([]string, 0),
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (a *Application) SyncToStore(app domain.Apper) (err error) {
	app.SaveOptions("export", a.Export)
	// ...
	return err
}

// читаем состояние приложения
func (a *Application) ReadState(app domain.Apper) (err error) {
	a.Export = app.Options().Export
	a.Browser = app.Options().Browser
	a.Output = app.Options().Output
	a.Host = app.Options().Hostname
	a.Port = app.Options().HostPort
	a.Debug = config.Mode == "development"
	// if repo.Dbs().A3().Exists {
	// 	a.DbA3Desc = fmt.Sprintf("[%s] %s", repo.Dbs().A3().Driver, repo.Dbs().A3().File)
	// }
	// if repo.Dbs().Self().Exists {
	// 	fname := repo.Dbs().Self().File
	// 	if filepath.IsAbs(fname) {
	// 		fname = filepath.Base(fname)
	// 	}
	// 	a.DbLiteDesc = fmt.Sprintf("[%s] %s", repo.Dbs().Self().Driver, fname)
	// }
	// if repo.Dbs().ConfigInfo().Exists {
	// 	a.DbConfigDesc = fmt.Sprintf("[%s] %s", repo.Dbs().ConfigInfo().Driver, repo.Dbs().ConfigInfo().File)
	// }
	// if repo.Dbs().Znak().Exists {
	// 	a.DbZnakDesc = fmt.Sprintf("[%s] %s", repo.Dbs().Znak().Driver, repo.Dbs().Znak().File)
	// }
	a.License = app.Options().Application.License
	a.FsrarID = app.Options().Application.Fsrarid
	a.InitDateMn()
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
	a.Magazins = make([]string, 0)
	// m.Reestr = map[string][]*Zakaz{}
	for key := range a.Reestr {
		delete(a.Reestr, key)
	}
	a.Magazin = ""
}
