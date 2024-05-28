package module

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetDosens() ([]Dosen, error) {
	var dosens []Dosen
	err := repo.DB.Find(&dosens).Error
	return dosens, err
}
func (repo Repository) GetDosenByNIDN(nidn string) (*Dosen, error) {
	var dosen *Dosen
	err := repo.DB.Where("NIDN =?", nidn).First(&dosen).Error
	return dosen, err
}
