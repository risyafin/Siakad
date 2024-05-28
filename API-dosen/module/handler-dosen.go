package module

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetDosens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dosens, err := handler.Usecase.GetDosens()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(dosens)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func (handler Handler) GetDosenByNIDN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	nidn := vars["nidn"]

	dosen, err := handler.Usecase.GetDosenByNIDN(nidn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(dosen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
