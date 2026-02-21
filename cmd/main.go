package main

import (
	"context"
	addcategory "finance/internal/api/handler/category/add"
	deletecategory "finance/internal/api/handler/category/delete"
	getallcategory "finance/internal/api/handler/category/get_all"
	updatecategory "finance/internal/api/handler/category/update"
	addtransaction "finance/internal/api/handler/transaction/add"
	deletetransaction "finance/internal/api/handler/transaction/delete"
	gettransaction "finance/internal/api/handler/transaction/get"
	getalltransaction "finance/internal/api/handler/transaction/get_all"
	updatetransaction "finance/internal/api/handler/transaction/update"
	categoryrepo "finance/internal/repository/category"
	transactionrepo "finance/internal/repository/transaction"
	"finance/internal/storage"
	categorystorage "finance/internal/storage/category"
	transactionstorage "finance/internal/storage/transaction"
	usecasecategoryadd "finance/internal/usecase/category/add"
	usecasecategorydelete "finance/internal/usecase/category/delete"
	usecasecategorygetall "finance/internal/usecase/category/get_all"
	usecasecategoryupdate "finance/internal/usecase/category/update"
	usecasetransactionadd "finance/internal/usecase/transaction/add"
	usecasetransactiondelete "finance/internal/usecase/transaction/delete"
	usecasetransactionget "finance/internal/usecase/transaction/get"
	usecasetransactiongetall "finance/internal/usecase/transaction/get_all"
	usecasetransactionupdate "finance/internal/usecase/transaction/update"

	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	postgresDB, err := storage.NewPostgresDB(storage.Config{
		Host:     "localhost",
		Port:     "5442",
		Username: "postgres",
		DBName:   "postgres",
		Password: "qwerty",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal(err)
	}

	categoryStorage := categorystorage.New(postgresDB)
	transactionStorage := transactionstorage.New(postgresDB)

	categoryRepo := categoryrepo.New(categoryStorage)
	transactionRepo := transactionrepo.New(transactionStorage)

	usecaseCategoryAdd := usecasecategoryadd.New(categoryRepo)
	usecaseCategoryGetAll := usecasecategorygetall.New(categoryRepo)
	usecaseCategoryUpdate := usecasecategoryupdate.New(categoryRepo)
	usecaseCategoryDelete := usecasecategorydelete.New(categoryRepo)

	handlerCategoryAdd := addcategory.New(ctx, usecaseCategoryAdd)
	handlerCategoryGetAll := getallcategory.New(ctx, usecaseCategoryGetAll)
	handlerCategoryUpdate := updatecategory.New(ctx, usecaseCategoryUpdate)
	handlerCategoryDelete := deletecategory.New(ctx, usecaseCategoryDelete)

	usecaseTransactionAdd := usecasetransactionadd.New(transactionRepo)
	usecaseTransactionGetAll := usecasetransactiongetall.New(transactionRepo)
	usecaseTransactionGet := usecasetransactionget.New(transactionRepo)
	usecaseTransactionUpdate := usecasetransactionupdate.New(transactionRepo)
	usecaseTransactionDelete := usecasetransactiondelete.New(transactionRepo)

	handlerTransactionAdd := addtransaction.New(ctx, usecaseTransactionAdd)
	handlerTransactionGetAll := getalltransaction.New(ctx, usecaseTransactionGetAll)
	handlerTransactionGet := gettransaction.New(ctx, usecaseTransactionGet)
	handlerTransactionUpdate := updatetransaction.New(ctx, usecaseTransactionUpdate)
	handlerTransactionDelete := deletetransaction.New(ctx, usecaseTransactionDelete)

	router := gin.New()

	category := router.Group("/category")
	{
		category.POST("/", handlerCategoryAdd.AddCategory)
		category.GET("/", handlerCategoryGetAll.GetAllCategory)
		category.PUT("/", handlerCategoryUpdate.UpdateCategory)
		category.DELETE("/:id", handlerCategoryDelete.DeleteCategory)
	}

	transaction := router.Group("/transaction")
	{
		transaction.POST("/", handlerTransactionAdd.AddTransaction)
		transaction.GET("/getAll/:id", handlerTransactionGetAll.GetAllTransaction)
		transaction.GET("/:id", handlerTransactionGet.GetTransaction)
		transaction.PUT("/", handlerTransactionUpdate.UpdateTransaction)
		transaction.DELETE("/:id", handlerTransactionDelete.DeleteTransaction)
	}

	router.Run(":8080")
}
