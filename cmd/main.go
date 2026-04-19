package main

import (
	"context"
	addcategory "finance/internal/api/handler/category/add"
	deletecategory "finance/internal/api/handler/category/delete"
	getallcategory "finance/internal/api/handler/category/get_all"
	getcategorybyid "finance/internal/api/handler/category/get_by_id"
	getexpensetotal "finance/internal/api/handler/category/get_expense_total"
	getincometotal "finance/internal/api/handler/category/get_income_total"
	updatecategory "finance/internal/api/handler/category/update"
	addtransaction "finance/internal/api/handler/transaction/add"
	deletetransaction "finance/internal/api/handler/transaction/delete"
	gettransaction "finance/internal/api/handler/transaction/get"
	gettransactionbyid "finance/internal/api/handler/transaction/get_by_id"
	getexpensetransaction "finance/internal/api/handler/transaction/get_expense"
	getincometransaction "finance/internal/api/handler/transaction/get_income"
	"os"
	"strings"

	updatetransaction "finance/internal/api/handler/transaction/update"
	categoryrepo "finance/internal/repository/category"
	transactionrepo "finance/internal/repository/transaction"
	"finance/internal/storage"
	categorystorage "finance/internal/storage/category"
	transactionstorage "finance/internal/storage/transaction"
	usecasecategoryadd "finance/internal/usecase/category/add"
	usecasecategorydelete "finance/internal/usecase/category/delete"
	usecasecategorygetall "finance/internal/usecase/category/get_all"
	usecasecategorybyid "finance/internal/usecase/category/get_by_id"
	usecasecategoryexpensetotal "finance/internal/usecase/category/get_expense_total"
	usecasecategoryincometotal "finance/internal/usecase/category/get_income_total"
	usecasecategoryupdate "finance/internal/usecase/category/update"
	usecasetransactionadd "finance/internal/usecase/transaction/add"
	usecasetransactiondelete "finance/internal/usecase/transaction/delete"
	usecasetransactionget "finance/internal/usecase/transaction/get"
	usecasetransactiongetbyid "finance/internal/usecase/transaction/get_by_id"
	usecasetransactiongetexpense "finance/internal/usecase/transaction/get_expense"
	usecasetransactiongetincome "finance/internal/usecase/transaction/get_income"
	usecasetransactionupdate "finance/internal/usecase/transaction/update"

	botservice "finance/internal/service/tg_bot"

	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	postgresDB, err := storage.NewPostgresDB(storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		log.Fatal(err)
	}

	botService := botservice.New()
	go func() {
		botService.Handle()
	}()

	categoryStorage := categorystorage.New(postgresDB)
	transactionStorage := transactionstorage.New(postgresDB)

	categoryRepo := categoryrepo.New(categoryStorage)
	transactionRepo := transactionrepo.New(transactionStorage)

	usecaseCategoryAdd := usecasecategoryadd.New(categoryRepo)
	usecaseCategoryGetByID := usecasecategorybyid.New(categoryRepo)
	usecaseCategoryGetAll := usecasecategorygetall.New(categoryRepo)
	usecaseCategoryGetIncomeTotal := usecasecategoryincometotal.New(categoryRepo)
	usecaseCategoryGetExpenseTotal := usecasecategoryexpensetotal.New(categoryRepo)
	usecaseCategoryUpdate := usecasecategoryupdate.New(categoryRepo)
	usecaseCategoryDelete := usecasecategorydelete.New(categoryRepo)

	handlerCategoryAdd := addcategory.New(ctx, usecaseCategoryAdd)
	handlerCategoryGetByID := getcategorybyid.New(ctx, usecaseCategoryGetByID)
	handlerCategoryGetAll := getallcategory.New(ctx, usecaseCategoryGetAll)
	handlerCategoryGetIncomeTotal := getincometotal.New(ctx, usecaseCategoryGetIncomeTotal)
	handlerCategoryGetExpenseTotal := getexpensetotal.New(ctx, usecaseCategoryGetExpenseTotal)
	handlerCategoryUpdate := updatecategory.New(ctx, usecaseCategoryUpdate)
	handlerCategoryDelete := deletecategory.New(ctx, usecaseCategoryDelete)

	usecaseTransactionAdd := usecasetransactionadd.New(transactionRepo)
	usecaseTransactionGetIncome := usecasetransactiongetincome.New(transactionRepo)
	usecaseTransactionGetExpense := usecasetransactiongetexpense.New(transactionRepo)
	usecaseTransactionGet := usecasetransactionget.New(transactionRepo)
	usecaseTransactionGetByID := usecasetransactiongetbyid.New(transactionRepo)
	usecaseTransactionUpdate := usecasetransactionupdate.New(transactionRepo)
	usecaseTransactionDelete := usecasetransactiondelete.New(transactionRepo)

	handlerTransactionAdd := addtransaction.New(ctx, usecaseTransactionAdd)
	handlerTransactionGetIncome := getincometransaction.New(ctx, usecaseTransactionGetIncome)
	handlerTransactionGetExpense := getexpensetransaction.New(ctx, usecaseTransactionGetExpense)
	handlerTransactionGet := gettransaction.New(ctx, usecaseTransactionGet)
	handlerTransactionGetByID := gettransactionbyid.New(ctx, usecaseTransactionGetByID)
	handlerTransactionUpdate := updatetransaction.New(ctx, usecaseTransactionUpdate)
	handlerTransactionDelete := deletetransaction.New(ctx, usecaseTransactionDelete)

	router := gin.New()

	origins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")

	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	category := router.Group("/category")
	{
		category.POST("/", handlerCategoryAdd.AddCategory)
		category.GET("/", handlerCategoryGetAll.GetAllCategory)
		category.GET("/:id", handlerCategoryGetByID.GetCategoryByID)
		category.GET("/getIncome", handlerCategoryGetIncomeTotal.GetCategoryIncomeTotal)
		category.GET("/getExpense", handlerCategoryGetExpenseTotal.GetCategoryExpenseTotal)
		category.PUT("/", handlerCategoryUpdate.UpdateCategory)
		category.DELETE("/:id", handlerCategoryDelete.DeleteCategory)
	}

	transaction := router.Group("/transactions")
	{
		transaction.POST("/", handlerTransactionAdd.AddTransaction)
		transaction.GET("/:id", handlerTransactionGetByID.GetTransactionByID)
		transaction.GET("/getIncome", handlerTransactionGetIncome.GetIncome)
		transaction.GET("/getExpense", handlerTransactionGetExpense.GetExpense)
		transaction.GET("/getAll", handlerTransactionGet.GetTransaction)
		transaction.PUT("/", handlerTransactionUpdate.UpdateTransaction)
		transaction.DELETE("/:id", handlerTransactionDelete.DeleteTransaction)
	}

	router.Run(":" + os.Getenv("APP_PORT"))
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = os.Getenv("APP_PORT")
	// }

	// router.Run("0.0.0.0:" + port)
}
