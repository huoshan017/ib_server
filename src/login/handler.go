package main

import (
	//"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"runtime/debug"
)

func login_handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Printf("%v\n", reflect.TypeOf(err))
		}
	}()
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if nil != err {
		_send_error(w, 0, -1)
		log.Printf("login handler ReadAll err[%s]\n", err.Error())
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
