package main

import (
	"fmt"
	"os"

	"github.com/huoshan017/ib_server/src/account/account_db"
)

var (
	DEFAULT_ACCOUNT_NUM_LOAD = 100000
)

var account_mgr *account_db.T_AccountRecordMgr = account_db.NewT_AccountRecordMgr(DEFAULT_ACCOUNT_NUM_LOAD)

func init_account_records() bool {
	primary_map, err := server.db_proxy.GetAccountTable().SelectAllPrimaryFieldMap()
	if err != nil {
		fmt.Fprintf(os.Stdout, "select all primary field map err: %v\n", err.Error())
		return false
	}

	if primary_map != nil {
		for acc, _ := range primary_map {
			res := server.redis_cluster.HSet("account", acc, 1)
			if res.Err() != nil {
				fmt.Fprintf(os.Stderr, "redis cluster HSET err: %v\n", res.Err().Error())
				return false
			}
		}
		fmt.Fprintf(os.Stdout, "account primary key loaded\n")
	}
	return true
}
