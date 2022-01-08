package main

import (
	repoHttp "LineManWongNaiInternShip/miniproject/Repository"
	"LineManWongNaiInternShip/miniproject/app"
	"LineManWongNaiInternShip/miniproject/config"
	userService "LineManWongNaiInternShip/miniproject/service/user/implement"
	"github.com/sirupsen/logrus"
	"log"
)

func newApp(appConfig *config.Config) *app.App {
	userRepo, _ := repoHttp.New(appConfig.PublicIP)
	user := userService.New(userRepo)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

const (
	NETWORK = "tcp"
)

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}
