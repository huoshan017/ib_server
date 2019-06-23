package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	phttp "github.com/huoshan017/ponu/http"
)

type Config struct {
	Id                int32
	Name              string
	ListenAddr        string
	AccountServerAddr string
}

func (this *Config) Init(config_path string) bool {
	data, err := ioutil.ReadFile(config_path)
	if err != nil {
		log.Printf("read config file err: %v\n", config_path, err.Error())
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		log.Printf("json unmarshal err: %v\n", err.Error())
		return false
	}
	return true
}

type Server struct {
	http_service phttp.Service
	config       *Config
}

var server Server

func (this *Server) Init(config *Config) bool {
	this.http_service.HandleFunc("/login", login_handler)
	this.config = config
	return true
}

func (this *Server) Run() {
	this.http_service.GoRun(this.config.ListenAddr)
	for {
		time.Sleep(time.Millisecond * 100)
	}
}
