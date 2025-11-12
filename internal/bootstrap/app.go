package bootstrap

import (
	"net/http"

	"github.com/kevalsabhani/splitify/internal/config"
)

type app struct {
	server *server
}

func NewApp() (*app, error) {
	conf, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := initDb(conf.Db.Dsn)
	if err != nil {
		return nil, err
	}

	// TODO: Remove below line once db is used
	_ = db

	// TODO: Change router once you have routes implemented
	server := &server{
		router: http.DefaultServeMux,
		port:   conf.App.Env,
	}

	return &app{
		server: server,
	}, nil
}

func (app *app) Run() error {
	return app.server.Run()
}
