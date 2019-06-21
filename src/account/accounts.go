package main

import (
	"github.com/huoshan017/ib_server/src/account/account_db"
)

var (
	DEFAULT_ACCOUNT_NUM_LOAD = 100000
)

var account_mgr *account_db.T_AccountRecordMgr = account_db.NewT_AccountRecordMgr(DEFAULT_ACCOUNT_NUM_LOAD)

func select_all_accounts() (map[string]*account_db.T_Account, error) {
	return server.db_proxy.GetAccountTable().SelectAllRecordsMap()
}

func select_account(key string) (*account_db.T_Account, error) {
	return server.db_proxy.GetAccountTable().SelectByPrimaryField(key)
}
