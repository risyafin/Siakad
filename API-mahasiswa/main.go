package main

import (
	"net/http"
	"siakad-dosen/API-mahasiswa/module"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/mahasiswa?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Failed Connect To Databases")
	}
	MahasiswaRepo := module.Repository{DB: db}
	MahasiswaUsecase := module.Usecase{Repo: MahasiswaRepo}
	MahasiswaHandler := module.Handler{Usecase: MahasiswaUsecase}

	r := mux.NewRouter()

	r.HandleFunc("/mahasiswas", MahasiswaHandler.GetMahasiswas).Methods("GET")
	http.ListenAndServe(":9090", r)
}
