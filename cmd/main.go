package main

import (
	"alcoscanxlsx/app"
	"alcoscanxlsx/checkdbg"
	"alcoscanxlsx/config"
	"alcoscanxlsx/domain/models/application"
	"alcoscanxlsx/gui"
	"alcoscanxlsx/processing"
	"alcoscanxlsx/reductor"
	"alcoscanxlsx/utility"
	"alcoscanxlsx/zaplog"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const modError = "main"

// var version = "0.0.0"
var fileExe string
var dir string

// если local true то папка создается локально
var local = flag.Bool("local", false, "")
var file = flag.String("file", "", "file to parse xlsx")

func init() {
	flag.Parse()
	fileExe = os.Args[0]
	var err error
	dir, err = filepath.Abs(filepath.Dir(fileExe))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %v\n", err)
		os.Exit(1)
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory: %v\n", err)
		os.Exit(1)
	}
}

func errMessageExit(title string, errDescription string) {
	utility.MessageBox(title, errDescription)
	os.Exit(-1)
}

func main() {
	cfg, err := config.New("", !*local)
	if err != nil {
		errMessageExit("ошибка конфигурации", err.Error())
	}

	var logsOutConfig = map[string][]string{
		"logger":   {"stdout", filepath.Join(cfg.LogPath(), config.Name)},
		"echo":     {filepath.Join(cfg.LogPath(), "echo")},
		"reductor": {filepath.Join(cfg.LogPath(), "reductor")},
		"true":     {filepath.Join(cfg.LogPath(), "true")},
	}
	zl, err := zaplog.New(logsOutConfig, true)
	if err != nil {
		errMessageExit("ошибка создания логера", err.Error())
	}
	defer zl.Shutdown()

	lg, err := zl.GetLogger("logger")
	if err != nil {
		errMessageExit("ошибка получения логера", err.Error())
	}
	loger := lg.Sugar()
	loger.Debug("zaplog started")
	loger.Infof("mode = %s", config.Mode)
	if cfg.Warning() != "" {
		loger.Infof("pkg:config warning %s", cfg.Warning())
	}

	errProcessExit := func(title string, errDescription string) {
		loger.Errorf("%s %s", title, errDescription)
		errMessageExit(title, errDescription)
	}
	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, loger, dir)
	// инициализируем пути необходимые приложению
	app.CreatePath()
	// создаем редуктор для хранения моделей приложения
	reductorLogger, err := zl.GetLogger("reductor")
	if err != nil {
		errProcessExit("Ошибка получения логера для редуктора", err.Error())
	}

	if err := reductor.New(reductorLogger.Sugar()); err != nil {
		errProcessExit("Ошибка создания редуктора", err.Error())
	}

	loger.Info("start repo")
	// инициализируем REPO
	// TODO изменить получение путей из конфига
	// dbPath := cfg.DbPath()
	// repoStart := repo.New(app, dbPath)
	// if len(repoStart.Errors()) > 0 {
	// 	fullErr := strings.Join(repoStart.Errors(), "\n")
	// 	errProcessExit("Ошибки запуска репозитория", fullErr)
	// }
	// app.SetRepo(repoStart)

	appModel, err := application.New(app, *file)
	if err != nil {
		errProcessExit("Ошибка получения логера для редуктора", err.Error())
	}
	if err := reductor.Instance().SetModel(appModel, false); err != nil {
		errProcessExit("Ошибка редуктора", err.Error())
	}
	// тесты
	if err := checkdbg.NewChecks(app).Run(); err != nil {
		loger.Errorf("check error %v", err)
		errProcessExit("Check failed", err.Error())
	}

	// GUI
	k, err := processing.New(app)
	if err != nil {
		errProcessExit("запуск обработки с ошибкой ", err.Error())
	}
	guiApp, err := gui.New(k, app)
	if err != nil {
		errProcessExit("создание gui с ошибкой ", err.Error())
	}
	guiApp.Run()
}
