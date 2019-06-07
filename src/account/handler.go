package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
)

type VerifyResult struct {
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

	if !account_mgr.Has(account) {

	}

	/*ret, err := w.Write(data)
	if nil != err {
		//_send_error(w, 0, -1)
		log.Printf("verify handler Write err %v, ret %v\n", err.Error(), ret)
		return
	}
	w.WriteHeader(200)*/
}

func _verify(data []byte) (ret_data []byte, err error) {
	return
}

func register_handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Printf("%v\n", reflect.TypeOf(err))
		}
	}()
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("register handler ReadAll err %v\n", err.Error())
		return
	}

	var ret int
	ret, err = w.Write(data)
	if err != nil {
		log.Printf("register handler Write err %v, ret %v\n", err.Error(), ret)
		return
	}
}
