package module

import (
	"time"
)

type Dosen struct {
	IdDosen         int       `json:"iddosen"`
	NamaDosen       string    `json:"namadosen"`
	Nidn            string    `json:"nidn"`
	Nip             string    `json:"nip"`
	JenisKelamin    string    `json:"jeniskelamin"`
	IdAgama         int       `json:"idagama"`
	NamaAgama       string    `json:"namaagama"`
	TanggalLahir    time.Time `json:"tanggallahir"`
	IdStatusAktif   bool      `json:"idstatusaktif"`
	NamaStatusAktif string    `json:"namastatusaktif"`
}
