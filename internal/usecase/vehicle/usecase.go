package vehicle

type Usecase struct {
	repository Repository
}

func NewUsecase(r Repository) Usecase {
	return Usecase{
		repository: r,
	}
}
