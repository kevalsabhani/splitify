package bootstrap

import (
	"fmt"

	"github.com/kevalsabhani/splitify/internal/config"
	"github.com/kevalsabhani/splitify/internal/http"
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

	server := &server{
		router: http.NewRouter(),
		port:   fmt.Sprintf(":%s", conf.App.Port),
	}

	return &app{
		server: server,
	}, nil
}

func (app *app) Run() error {
	return app.server.Run()
}
