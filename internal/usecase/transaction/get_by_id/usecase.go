package get_by_id

type Usecase struct {
	transactionRepo transactionRepo
}

func New(transactionRepo transactionRepo) *Usecase {
	return &Usecase{transactionRepo: transactionRepo}
}
