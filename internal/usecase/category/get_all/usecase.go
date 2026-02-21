package get_all

type Usecase struct {
	categoryRepo categoryRepo
}

func New(categoryRepo categoryRepo) *Usecase {
	return &Usecase{categoryRepo: categoryRepo}
}
