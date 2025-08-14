package controller

import (
	"net/http"

	"github.com/ziya2209/goproject/phonedb/internal/controller/handler"
	"github.com/ziya2209/goproject/phonedb/internal/db"
	"github.com/ziya2209/goproject/phonedb/internal/repo"
)

func Start() error {

	db, err := db.DbContion()
	if err != nil {
		return err
	}
	dao := repo.NewDao(db)

	mux := http.NewServeMux()

	mux.Handle("/reset", handler.ResetDatabase(dao))
	mux.HandleFunc("/add", handler.AddPhoneNumber(dao))
	mux.HandleFunc("/get", handler.GetAllPhoneNumbers(dao))
	return http.ListenAndServe(":8080", mux)
}
