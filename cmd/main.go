package main

import (
	"log"

	"github.com/spf13/viper"

	"bot/config"
	"bot/pkg/server"
)

func main(){
	if err:=config.Init(); err!=nil{
		log.Fatalf("%s", err.Error())
	}

	if err:=server.Launch(viper.GetString("httpapi")); err!=nil{
		log.Fatalf("%s", err.Error())
	}
}