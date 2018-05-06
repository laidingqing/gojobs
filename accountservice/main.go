package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/laidingqing/gojobs/accountservice/conf"
	"github.com/laidingqing/gojobs/accountservice/service"
	"github.com/laidingqing/gojobs/accountservice/storage/mongo"
	"github.com/laidingqing/gojobs/common/config"
	"github.com/laidingqing/gojobs/common/tracing"
)

var appName = "accountservice"

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.Parse()
	fmt.Println("config:", configFile)
	var err error
	conf.AccountConf, err = config.LoadConf(configFile)
	if err != nil {
		log.Printf("Load yaml config file error: '%v'", err)
		return
	}
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Starting %v\n", appName)

	initializeTracing()
	err := initializeStorage()
	if err != nil {
		panic("No 'db' set in configuration, cannot start")
	}

	service.StartWebServer(conf.AccountConf.Core.Port)
}

func initializeStorage() error {
	switch conf.AccountConf.Core.Storage {
	case "mongo":
		service.DbStorage = mongo.New()
	default:
		logrus.Errorf("storage error: can't find storage driver")
		return errors.New("can't find storage driver")
	}
	service.DbStorage.OpenSession()
	return nil
}

func initializeTracing() {
	tracing.InitTracing(conf.AccountConf.Core.ZipkinURI, appName)
}
