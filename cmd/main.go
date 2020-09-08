package main

import (
	"github.com/mitchellh/mapstructure"
	"github.com/yukiz97/cls-general-services/apiservices"
	"github.com/yukiz97/cls-general-services/lcservices"
	"github.com/yukiz97/utils/config"
	"log"
)

type configuration struct {
	MYSQLHost     string
	MYSQLUsername string
	MYSQLPassword string
	MYSQLDB       string
	ListenPort    int
}

func main() {
	configuration := configuration{}
	mapConfig := config.ParseJSONConfigToMap("D:\\DevApps\\_Workspace\\Golang\\.mydata\\cls-services\\config-general.json")
	err := mapstructure.Decode(mapConfig, &configuration)

	if err != nil {
		log.Fatal(err)
	}

	lcservices.InitLocalServices(configuration.MYSQLHost, configuration.MYSQLUsername, configuration.MYSQLPassword, configuration.MYSQLDB)
	apiservices.InitRestfulAPIServices(configuration.ListenPort)
}
