package main

import (
	"net/http"
	"siakad-dosen/API-dosen/module"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/dosen?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to databases")
	}

	dosensRepo := module.Repository{DB: db}
	dosensUsecase := module.Usecase{Repo: dosensRepo}
	dosenHandler := module.Handler{Usecase: dosensUsecase}

	r := mux.NewRouter()

	r.HandleFunc("/dosens", dosenHandler.GetDosens).Methods("GET")
	r.HandleFunc("/dosen/{nidn}", dosenHandler.GetDosenByNIDN).Methods("GET")
	http.ListenAndServe(":8080", r)

}
