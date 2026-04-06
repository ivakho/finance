package get_income

type Usecase struct {
	transactionRepo transactionRepo
}

func New(transactionRepo transactionRepo) *Usecase {
	return &Usecase{transactionRepo: transactionRepo}
}
