package controller

import (
	"net/http"

	//"github.com/ziya2209/goproject/phonedb/internal/controller/handler"
	"github.com/ziya2209/goproject/phonedb/internal/db"
)

func Start() error {
	// initialize server
	dao, err := db.DbContion()
	if err != nil {
		return err
    }
	_=dao

	// initialize server
	mux := http.NewServeMux()

	// add middlewares
	// 1. logger middleware
	// add routes
	//1. ResetDatabase handler
	//mux.Handle("/reset",handler.ResetDatabase(dao))
	// 2. AddPhoneNumber handler
	//mux.HandleFunc("/add", handler.AddPhoneNumber(dao))
	// 3. GetAllPhoneNumbers handler
	//mux.HandleFunc("/get",handler.GetAllPhoneNumbers(dao))

	// start server
	return http.ListenAndServe(":8080",mux)
}
