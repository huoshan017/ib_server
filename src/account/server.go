package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	phttp "github.com/huoshan017/ponu/http"
)

type Config struct {
	Id                int32
	Name              string
	ListenAddr        string
	DBProxyServerAddr string
	DBHostId          int32
	DBHostAlias       string
	DBName            string
	RedisClusterAddrs []string
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
	db_proxy      DBProxy
	http_service  phttp.Service
	config        *Config
	redis_cluster *redis.ClusterClient
}

var server Server

func (this *Server) Init(config *Config) bool {
	if !this.db_proxy.Connect(config.DBProxyServerAddr, config.DBHostId, config.DBHostAlias, config.DBName) {
		return false
	}
	this.http_service.HandleFunc("/account_verify", verify_handler)
	this.http_service.HandleFunc("/account_register", register_handler)
	this.config = config

	init_account_records()

	this.redis_cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: config.RedisClusterAddrs,
	})

	if this.redis_cluster == nil {
		return false
	}

	fmt.Fprintf(os.Stdout, "redis cluster created\n")

	return true
}

func (this *Server) Run() {
	this.db_proxy.GoRun()
	this.http_service.GoRun(this.config.ListenAddr)
	for {
		time.Sleep(time.Millisecond * 100)
	}
}
