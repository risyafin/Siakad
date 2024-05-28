package module

import "time"

type Mahasiswa struct {
	Nama_Mahasiswa          int       `json:"nama_mahasiswa"`
	Jenis_Kelamin           string    `json:"jenis_kelamin"`
	Tanggal_Lahir           time.Time `json:"Tanggal_Lahir"`
	Id_Perguruan_Tinggi     int       `json:"id_perguruan_tinggi"`
	Id_Mahasiswa            int       `json:"id_mahasiswa"`
	Id_Agama                int       `json:"id_agama"`
	Nama_Agama              string    `json:"nama_agama"`
	Id_Prodi                string    `json:"id_prodi"`
	Nama_Program_Studi      string    `json:"nama_program_studi"`
	Nama_Status_Mahasiswa   string    `json:"nama_status_mahasiswa"`
	Nim                     string    `json:"nim"`
	Id_Periode              string    `json:"id_periode"`
	Nama_Periode_Masuk      string    `json:"nama_periode_masuk"`
	Id_Registrasi_Mahasiswa int       `json:"id_registrasi_mahasiswa"`
}
