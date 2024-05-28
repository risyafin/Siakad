package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("mahasiswa", forwardRequest("http://localhost:8081/mahasiswas")).Methods("GET")
	r.HandleFunc("/dosen", forwardRequest("http://localhost:8080/dosens")).Methods("GET")
	r.HandleFunc("/dosen/{nidn}",forwardRequest("http://localhost/dosen/{nidn}")).Methods("GET")

	http.ListenAndServe(":8888", r)
}

func forwardRequest(targetUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(targetUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)

		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
