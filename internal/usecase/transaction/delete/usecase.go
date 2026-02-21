package delete

type Usecase struct {
	transactionRepo transactionRepo
}

func New(transactionRepo transactionRepo) *Usecase {
	return &Usecase{transactionRepo: transactionRepo}
}
