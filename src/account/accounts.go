package main

import (
	"log"
	"sync"
	"time"

	"github.com/huoshan017/ib_server/src/account/account_db"

	"github.com/hashicorp/golang-lru/simplelru"
)

var (
	DEFAULT_ACCOUNT_NUM_LOAD = 100000
)

type AccountMgr struct {
	account_table *account_db.T_Account_Table_Proxy
	accounts_have map[string]bool
	accounts_load *simplelru.LRU
	locker        sync.RWMutex
}

var account_mgr AccountMgr

func (this *AccountMgr) Init(account_table *account_db.T_Account_Table_Proxy) error {
	accounts, err := simplelru.NewLRU(DEFAULT_ACCOUNT_NUM_LOAD, nil)
	if err != nil {
		return err
	}
	this.accounts_load = accounts
	this.account_table = account_table
	log.Printf("Loading accounts from db ...\n")
	accounts_have := this.account_table.SelectAllPrimaryFieldMap()
	log.Printf("Loaded accounts: %v\n", accounts_have)

	if accounts_have == nil {
		accounts_have = make(map[string]bool)
	}
	this.accounts_have = accounts_have
	return nil
}

func (this *AccountMgr) Add(acc, pwd string) bool {
	this.locker.Lock()
	defer this.locker.Unlock()

	if _, o := this.accounts_have[acc]; o {
		return false
	}

	if !this.accounts_load.Contains(acc) {
		account := account_db.Create_T_Account()
		account.Lock()
		account.Set_account(acc)
		account.Set_password(pwd)
		account.Set_register_time(int32(time.Now().Unix()))
		account.Unlock()
		this.account_table.Insert(account)
		this.accounts_load.Add(acc, account)
		this.accounts_have[acc] = true
	}

	return true
}

func (this *AccountMgr) Has(acc string) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()
	_, o := this.accounts_have[acc]
	return o
}

func (this *AccountMgr) IsLoad(acc string) bool {
	this.locker.RLock()
	defer this.locker.RUnlock()
	return this.accounts_load.Contains(acc)
}

func (this *AccountMgr) Get(acc string, is_load bool) *account_db.T_Account {
	var account *account_db.T_Account

	this.locker.RLock()
	if this.accounts_load.Contains(acc) {
		acc_inter, o := this.accounts_load.Get(acc)
		if o && acc_inter != nil {
			account = acc_inter.(*account_db.T_Account)
		}
	}
	this.locker.RUnlock()

	if account != nil {
		return account
	}

	if is_load {
		account = this.account_table.SelectByPrimaryField(acc)
		if account == nil {
			return nil
		}
		this.locker.Lock()
		this.accounts_load.Add(acc, account)
		this.locker.Unlock()
	}

	return account
}
