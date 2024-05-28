package module

import "golang.org/x/crypto/bcrypt"

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetAdmins() ([]Admin, error) {
	admins, err := usecase.Repo.GetAdmin()
	return admins, err
}

func (usecase Usecase) GetMahasiswas() ([]Mahasiswa, error) {
	mahasiswas, err := usecase.Repo.GetMahasiswa()
	return mahasiswas, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func (usecase Usecase) Registration(admin *Admin) error {
	hash, _ := HashPassword(admin.Password)
	admin.Password = hash
	err := usecase.Repo.RegistrationAdmin(admin)
	return err
}

func (usecase Usecase) GetDosens() ([]Dosen, error) {
	dosens, err := usecase.Repo.GetDosens()
	return dosens, err
}
