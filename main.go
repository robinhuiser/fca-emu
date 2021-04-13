/*
 * Finite Cloud API Emulator
 *
 * Author: Robin Huiser robin@technisys.com
 */

package main

import (
	"bytes"
	"context"
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/dimiro1/banner"
	"github.com/joho/godotenv"
	"github.com/robinhuiser/fca-emu/generator"
	finite "github.com/robinhuiser/fca-emu/server"
	"github.com/robinhuiser/fca-emu/util"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed swagger-ui/* api/openapi.yaml
var staticFiles embed.FS

//go:embed banner.txt
var appBanner string

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

const (
	// Defaults
	DB_VENDOR             = "sqlite3"
	APP_LISTEN_ADDRESS    = "0.0.0.0"
	APP_LISTEN_PORT       = "8080"
	NR_GENERATED_ENTITIES = 20
	NR_BRANCHES_PER_BANK  = 50
)

func main() {
	// We can ignore if there is no .dotenv (container runtime)
	godotenv.Load()

	// Display the banner
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(appBanner))

	log.Printf("version %s, commit %s, built at %s by %s", version, commit, date, builtBy)

	// Set application variables
	appListenAddress := util.GetEnvString("EMULATOR_LISTEN_ADDRESS", APP_LISTEN_ADDRESS)
	appListenPort := util.GetEnvString("EMULATOR_LISTEN_PORT", APP_LISTEN_PORT)

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
	if err := generator.Generate(
		util.GetEnvInt("NUMBER_OF_ENTITIES", NR_GENERATED_ENTITIES),
		util.GetEnvInt("BRANCHES_PER_BANK", NR_BRANCHES_PER_BANK),
		client); err != nil {
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

	// Embed the Swagger UI within the Go binary
	router.PathPrefix("/").Handler(http.FileServer(http.FS(staticFiles)))
	log.Printf("specs available on http://%s:%s/swagger-ui", appListenAddress, appListenPort)
	log.Printf("mock server listening on %s:%s", appListenAddress, appListenPort)

	log.Fatal(http.ListenAndServe(appListenAddress+":"+appListenPort, router))
}
