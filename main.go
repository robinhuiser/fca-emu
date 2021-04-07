/*
 * Trexis Cloud API
 *
 * The public facing API through which connectors are exposed as a single abstract API
 *
 * API version: v1.5
 * Contact: support@trexis.net
 *
 * Author: Robin Huiser robin@technisys.com
 */

package main

import (
	"context"
	"embed"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/robinhuiser/finite-mock-server/generator"
	finite "github.com/robinhuiser/finite-mock-server/server"
	"github.com/robinhuiser/finite-mock-server/util"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed swagger-ui/* api/openapi.yaml
var staticFiles embed.FS

const (
	// Defaults
	DB_VENDOR             = "sqlite3"
	APP_LISTEN_ADDRESS    = "0.0.0.0"
	APP_LISTEN_PORT       = "8080"
	NR_GENERATED_ACCOUNTS = 5
)

func main() {
	// We can ignore if there is no .dotenv (container runtime)
	godotenv.Load()

	// Set application variables
	appListenAddress := util.GetEnvString("MOCK_SERVER_LISTEN_ADDRESS", APP_LISTEN_ADDRESS)
	appListenPort := util.GetEnvString("MOCK_SERVER_LISTEN_PORT", APP_LISTEN_PORT)

	// Connect to database
	client, err := util.GetDatabaseClient(util.GetEnvString("DB_VENDOR", DB_VENDOR))
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	log.Printf("connected to %s database", util.GetEnvString("DB_VENDOR", DB_VENDOR))
	defer client.Close()

	// Run the migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Printf("database migration run successfully")

	// Generate Mock data
	if err := generator.Accounts(util.GetEnvInt("TEST_DATA_ACCOUNTS", NR_GENERATED_ACCOUNTS), client); err != nil {
		log.Fatalf("failed creating testdata: %v", err)
	}

	AccountsApiService := finite.NewAccountsApiService(client)
	AccountsApiController := finite.NewAccountsApiController(AccountsApiService)

	CacheApiService := finite.NewCacheApiService()
	CacheApiController := finite.NewCacheApiController(CacheApiService)

	CardsApiService := finite.NewCardsApiService()
	CardsApiController := finite.NewCardsApiController(CardsApiService)

	EntityApiService := finite.NewEntityApiService()
	EntityApiController := finite.NewEntityApiController(EntityApiService)

	ExchangeApiService := finite.NewExchangeApiService()
	ExchangeApiController := finite.NewExchangeApiController(ExchangeApiService)

	ProductsApiService := finite.NewProductsApiService()
	ProductsApiController := finite.NewProductsApiController(ProductsApiService)

	StatementApiService := finite.NewStatementApiService()
	StatementApiController := finite.NewStatementApiController(StatementApiService)

	StatementsApiService := finite.NewStatementsApiService()
	StatementsApiController := finite.NewStatementsApiController(StatementsApiService)

	TransactionsApiService := finite.NewTransactionsApiService()
	TransactionsApiController := finite.NewTransactionsApiController(TransactionsApiService)

	router := finite.NewRouter(AccountsApiController,
		CacheApiController,
		CardsApiController,
		EntityApiController,
		ExchangeApiController,
		ProductsApiController,
		StatementApiController,
		StatementsApiController,
		TransactionsApiController)

	// Embed the Swagger UI within Go binary
	router.PathPrefix("/").Handler(http.FileServer(http.FS(staticFiles)))

	log.Printf("mock server started on %s:%s", appListenAddress, appListenPort)
	log.Printf("specs available on http://%s:%s/swagger-ui", appListenAddress, appListenPort)
	log.Fatal(http.ListenAndServe(appListenAddress+":"+appListenPort, router))
}
