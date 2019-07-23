package main

import (
	"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"

	"github.com/huoshan017/ib_server/src/account/account_db"
)

const (
	ERR_NONE               = iota
	ERR_INTERNAL           = 1
	ERR_ACCOUNT_NOT_FOUND  = 2
	ERR_PASSWORD_INCORRECT = 3
	ERR_ALREADY_REGISTERED = 4
)

type VerifyResult struct {
	ErrCode int32
}

type RegisterResult struct {
	ErrCode int32
}

func verify_handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Printf("%v\n", reflect.TypeOf(err))
		}
	}()
	defer r.Body.Close()

	querys := r.URL.Query()
	account := querys.Get("account")
	password := querys.Get("password")

	var result VerifyResult
	acc_item := account_mgr.Get(account)
	if acc_item != nil {
		if acc_item.Get_password() != password {
			result.ErrCode = ERR_PASSWORD_INCORRECT
		}
	} else {
		result.ErrCode = ERR_ACCOUNT_NOT_FOUND
	}

	data, err := json.Marshal(&result)
	if err != nil {
		log.Printf("verify handler json marshal err %v\n", err.Error())
		return
	}

	var ret int
	ret, err = w.Write(data)
	if err != nil {
		log.Printf("verify handler Write err %v, ret %v\n", err.Error(), ret)
		return
	}

	log.Printf("verified account %v\n", account)
}

func register_handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Printf("%v\n", reflect.TypeOf(err))
		}
	}()
	defer r.Body.Close()

	query := r.URL.Query()
	account := query.Get("account")
	password := query.Get("password")

	var result RegisterResult
	if account_mgr.Has(account) {
		result.ErrCode = ERR_ALREADY_REGISTERED
	} else {
		acc := &account_db.T_Account{}
		acc.Set_account(account)
		acc.Set_password(password)
		account_mgr.Add(acc)
		server.db_proxy.GetTableManager().GetT_AccountTableProxy().Insert(acc)
	}

	data, err := json.Marshal(&result)
	if err != nil {
		log.Printf("register handler json marshal err %v\n", err.Error())
		return
	}

	var ret int
	ret, err = w.Write(data)
	if err != nil {
		log.Printf("register handler Write err %v, ret %v\n", err.Error(), ret)
		return
	}

	log.Printf("registered account %v\n", account)
}
