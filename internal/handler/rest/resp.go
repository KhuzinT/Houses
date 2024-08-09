package rest

import (
	"clevergo.tech/jsend"
	"log"
	"net/http"
)

const (
	SuccessStatus = "success"
	FailStatus    = "fail"
	ErrorStatus   = "error"
)

func SendResp(w http.ResponseWriter, status string, data interface{}, code int, msg string) {
	if msg != "" {
		log.Println("Send Response: ", msg)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var err error
	switch status {
	case SuccessStatus:
		err = jsend.Success(w, data, code)
	case FailStatus:
		err = jsend.Fail(w, data, code)
	case ErrorStatus:
		err = jsend.Error(w, msg, code)
	}

	if err != nil {
		panic(err)
	}
}
