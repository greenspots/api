package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type (
	EventContext struct {
		AppContext
	}
)

func (context *EventContext) OnStatusEvent(event *model.StatusEvent) string {
	var wg sync.WaitGroup
	wg.Add(2)
	statusChan := make(chan string)
	go func(status model.Status) {
		defer wg.Done()
		statusEventId := context.EsApi.IndexStatus(status)
		statusChan <- statusEventId
	}(event.Status)
	go func(event *model.StatusEvent) {
		defer wg.Done()
		context.EsApi.UpdateStatus(event.Gspotid, event.Status)
	}(event)
	wg.Wait()
	return <-statusChan
}

func (context *EventContext) OnWeatherEvent(event *model.WeatherEvent) string {
	weatherId := context.EsApi.IndexWeather(event.Weather)
	context.EsApi.UpdateWeather(event.Gspotid, event.Weather)
	return weatherId
}

func (context *EventContext) OnAirQualityEvent(event *model.AirQualityEvent) string {
	weatherId := context.EsApi.IndexAirQuality(event.AirQuality)
	context.EsApi.UpdateAirquality(event.Gspotid, event.AirQuality)
	return weatherId

}

func (context *EventContext) CreateStatusHandler(w http.ResponseWriter, r *http.Request) {
	event := &model.StatusEvent{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal(reqBody, event)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	eventId := context.OnStatusEvent(event)
	_, _ = fmt.Fprintf(w, "Status Event Processed. Id: %s", eventId)
}

func (context *EventContext) CreateWeatherHandler(w http.ResponseWriter, r *http.Request) {
	event := &model.WeatherEvent{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal(reqBody, event)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	eventId := context.OnWeatherEvent(event)
	_, _ = fmt.Fprintf(w, "Weather Event Processed. Id: %s", eventId)
}

func (context *EventContext) CreateAirQualityHandler(w http.ResponseWriter, r *http.Request) {
	event := &model.AirQualityEvent{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.Unmarshal(reqBody, event)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	eventId := context.OnAirQualityEvent(event)
	_, _ = fmt.Fprintf(w, "Status Event Processed. Id: %s", eventId)
}
