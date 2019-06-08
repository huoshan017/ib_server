package main

import (
	"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
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
	acc_item := account_mgr.Get(account, true)
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
		account_mgr.Add(account, password)
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
}
