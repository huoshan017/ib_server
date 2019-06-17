package main

import (
	"log"

	"github.com/huoshan017/ib_server/src/account/account_db"
	"github.com/huoshan017/mysql-go/proxy/client"
)

type DBProxy struct {
	db           mysql_proxy.DB
	table_proxys account_db.TablesProxyManager
}

func (this *DBProxy) Connect(proxy_addr string, db_host_id int32, db_host_alias, db_name string) bool {
	err := this.db.Connect(proxy_addr, db_host_id, db_host_alias, db_name)
	if err != nil {
		log.Printf("db proxy connect err: %v\n", err.Error())
		return false
	}
	this.table_proxys.Init(&this.db)
	return true
}

func (this *DBProxy) GoRun() {
	this.db.GoRun()
}

func (this *DBProxy) Save() {
	this.db.Save()
}

func (this *DBProxy) End() {
	this.db.Close()
}

func (this *DBProxy) GetTableManager() *account_db.TablesProxyManager {
	return &this.table_proxys
}
