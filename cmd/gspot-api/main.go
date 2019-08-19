package main

import (
	"github.com/justinas/alice"
	"log"
	"net/http"
)

func main() {
	appCtx := controller.ProvideAppContext()
	dataInitCtx := controller.ProvideDataInitContext(appCtx)
	eventCtx := controller.ProvideEventContext(appCtx)
	commonHandlers := alice.New(controller.RecoverHandler)

	router := controller.NewRouter()

	router.Get("/admin/create_indexes", commonHandlers.ThenFunc(dataInitCtx.CreateIndexesHandler))
	router.Get("/admin/create_trees", commonHandlers.ThenFunc(dataInitCtx.CreateTreesHandler))

	router.Post("/event/weather", commonHandlers.ThenFunc(eventCtx.CreateWeatherHandler))
	router.Post("/event/status", commonHandlers.ThenFunc(eventCtx.CreateStatusHandler))
	router.Post("/event/airquality", commonHandlers.ThenFunc(eventCtx.CreateAirQualityHandler))

	var err = http.ListenAndServe(":"+config.API_SERVER_PORT, router)
	if err != nil {
		log.Fatal(err)
	}
}
