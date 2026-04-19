package get_expense_total

type Usecase struct {
	categoryRepo categoryRepo
}

func New(categoryRepo categoryRepo) *Usecase {
	return &Usecase{categoryRepo: categoryRepo}
}
