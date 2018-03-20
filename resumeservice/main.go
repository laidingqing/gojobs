package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/laidingqing/gojobs/common/config"
	"github.com/laidingqing/gojobs/common/tracing"
	"github.com/laidingqing/gojobs/resumeservice/service"
	"github.com/laidingqing/gojobs/resumeservice/storage/mongo"
	"github.com/spf13/viper"
)

var appName = "resumeservice"
var (
	profile         string
	configServerURL string
	storage         string
	port            string
	zipkin          string
)

func init() {
	flag.StringVar(&profile, "profile", "dev", "Environment profile, something similar to spring profiles")
	flag.StringVar(&configServerURL, "configServerUrl", "http://localhost:8888", "Address to config server")

	flag.StringVar(&storage, "storage", "mongo", "Environment storage")
	flag.StringVar(&port, "port", "6768", "service port")
	flag.StringVar(&zipkin, "zipkin", "http://192.168.99.100:9411", "service port")

	flag.Usage = usage
	flag.Parse()

	viper.Set("profile", profile)
	viper.Set("configServerUrl", configServerURL)
	viper.Set("storage", storage)
	viper.Set("port", port)
	viper.Set("zipkin_server_url", zipkin)
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("Starting %v\n", appName)

	config.LoadConfiguration(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"))

	initializeTracing()
	err := initializeStorage()
	if err != nil {
		panic("No 'db' set in configuration, cannot start")
	}

	service.StartWebServer(viper.GetString("port"))
}

func initializeStorage() error {
	switch viper.GetString("storage") {
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
	tracing.InitTracing(viper.GetString("zipkin_server_url"), appName)
}

var usageStr = `
Usage: resumeservice-xx-xx [options]
Server Options:
    --configServerUrl <configServerUrl>   
    --port <port>     
`

func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
