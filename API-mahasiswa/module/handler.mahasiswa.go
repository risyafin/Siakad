package module

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetMahasiswas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mahasiswa, err := handler.Usecase.GetMahasiswas()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
