/*
 * Trading API
 *
 * This API allow to interact with the trading system.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	hktrading_server "github.com/hungknow/hktrading_server/go"
	"github.com/phuslu/log"
	"hungknow.com/blockchain/config"
	"hungknow.com/blockchain/datafeed"
	"hungknow.com/blockchain/db/dbstoresql"
	"hungknow.com/blockchain/logutils"
	"hungknow.com/blockchain/symbols"
)

func NewRouter(routers ...hktrading_server.Router) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			router.Method(route.Method, route.Pattern, handler)
		}
	}

	return router
}

func main() {
	logutils.SetupPhusluLog()

	symbolManager := symbols.NewSymbolManager()

	appConfig, appErr := config.GetConfig("../../config")
	if appErr != nil {
		log.Panic().Msgf("Error getting config: %+v", appErr)
	}
	dbStore, appErr := dbstoresql.NewDBSQLStore(appConfig.DB)
	if appErr != nil {
		log.Panic().Msgf("Error creating db store: %+v", appErr)
	}
	forexDataFeed := datafeed.NewForexDataFeed(dbStore, symbolManager)

	DefaultAPIService := hktrading_server.NewAPIService(symbolManager, forexDataFeed)
	DefaultAPIController := hktrading_server.NewDefaultAPIController(DefaultAPIService)

	router := NewRouter(DefaultAPIController)

	log.Info().Msgf("Server started on port 9001")
	err := http.ListenAndServe(":9001", router)
	if err != nil {
		log.Fatal().Msgf("Error starting server: %+v", err)
	}
}
