package module

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetMahasiswas() ([]Mahasiswa, error) {
	var mahasiswa []Mahasiswa

	result := repo.DB.Find(&mahasiswa)
	return mahasiswa, result.Error

}
