package utils

import (
	"encoding/json"
	"net/http"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Rows  interface{}
	Total interface{}
}

func Rep(w http.ResponseWriter,code int,data interface{}, msg string){
	w.Header().Set("Content-Type","appliaction/json")
	w.WriteHeader(200)
	h := H{
		Code: code,
		Data: data,
		Msg:  msg, 
	}
	ans ,err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	w.Write(ans)
}

func RepList(w http.ResponseWriter,code int,data interface{},total interface{}) {
	w.Header().Set("Content-Type","appliaction/json")
	w.WriteHeader(200)
	h := H{
		Code: code,
		Rows: data,
		Total: total,
	}
	ans ,err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	w.Write(ans)
}

func RepFail(w http.ResponseWriter,msg string) {
	Rep(w ,-1, nil , msg)
}

func RepOK(w http.ResponseWriter,data interface{},msg string) {
	Rep(w , 0, data, msg)
}

func RepOKList(w http.ResponseWriter,data interface{},total interface{}) {
	RepList(w , 0, data, total)
}


