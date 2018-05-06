package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/laidingqing/gojobs/common/config"
	"github.com/laidingqing/gojobs/common/tracing"
	"github.com/laidingqing/gojobs/resumeservice/conf"
	"github.com/laidingqing/gojobs/resumeservice/service"
	"github.com/laidingqing/gojobs/resumeservice/storage/mongo"
)

var appName = "resumeservice"

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.Parse()
	fmt.Println("config:", configFile)
	var err error
	conf.ResumeConf, err = config.LoadConf(configFile)
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

	service.StartWebServer(conf.ResumeConf.Core.Port)
}

func initializeStorage() error {
	switch conf.ResumeConf.Core.Storage {
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
	tracing.InitTracing(conf.ResumeConf.Core.ZipkinURI, appName)
}
