package module

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetMahasiswas()([]Mahasiswa, error) {
	result, err := usecase.Repo.GetMahasiswas()
	return result, err
}
