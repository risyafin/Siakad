package module

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetDosens() ([]Dosen, error) {
	dosens, err := usecase.Repo.GetDosens()
	return dosens, err
}

func (usecase Usecase) GetDosenByNIDN(nidn string) (*Dosen, error) {
	var dosen *Dosen
	dosen, err := usecase.Repo.GetDosenByNIDN(nidn)
	return dosen, err
}
