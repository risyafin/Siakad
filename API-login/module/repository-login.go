package module

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) RegistrationAdmin(admin *Admin) error {
	result := repo.DB.Select("Username", "Password").Create(&admin)
	return result.Error
}

func (repo Repository) LoginAdmin(username string, password string) (*Admin, error) {
	var admin Admin
	result := repo.DB.Model(&admin).Where("username =? AND passord =?", username, password).First(&admin)
	return &admin, result.Error
}


func (repo Repository) GetAdmin() ([]Admin, error) {
	var admin []Admin
	result := repo.DB.Find(&admin)
	return admin, result.Error

}

func (repo Repository) GetMahasiswa() ([]Mahasiswa, error) {
	var mahasiswa []Mahasiswa
	result := repo.DB.Find(&mahasiswa)
	return mahasiswa, result.Error
}

func (repo Repository) GetDosens() ([]Dosen, error) {
	var dosen []Dosen
	result := repo.DB.Find(&dosen)
	return dosen, result.Error
}
