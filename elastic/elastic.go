package elastic

import (
	"context"
	"github.com/greenspots/api/config"
	"github.com/greenspots/model"
	"github.com/olivere/elastic/v7"
	"log"
)

const WEATHER_UPDATE_SCRIPT_SOURCE = "\"source\": \"ctx._source.weather = params.weather\","
const STATUS_UPDATE_SCRIPT_SOURCE = "\"source\": \"ctx._source.status = params.status\","
const AIRQUALITY_UPDATE_SCRIPT_SOURCE = "\"source\": \"ctx._source.status = params.airquality\","

type ESApi struct {
	*elastic.Client
}

func NewESApi() *ESApi {
	client, err := elastic.NewSimpleClient(elastic.SetURL(config.ELASTIC_URL))
	if err != nil {
		log.Fatal(err)
	}
	return &ESApi{client}
}

func (api *ESApi) CreateGspotIndexes(index string) {
	exists, _ := api.IndexExists(index).Do(context.Background())
	if !exists {
		_, err := api.CreateIndex(index).Do(context.Background())
		if err != nil {
			log.Println("Cannot create index " + index)
		}
	} else {
		log.Println("Index already exists: " + index)
	}
}

func (api *ESApi) IndexNewGspot(treeName string, treeId string) {
	exists, _ := api.Exists().
		Index(config.GS_INDEX).
		Id(treeId).
		Do(context.Background())
	if !exists {
		gspot := model.Greenspot{Name: treeName}
		_, err := api.
			Index().
			Index(config.GS_INDEX).
			Type("_doc").
			Id(treeId).
			BodyJson(gspot).
			Do(context.Background())
		if err != nil {
			log.Println("Could not index tree in greenspot index. Name: " + treeName)
			log.Println(err)
		}
	} else {
		log.Println("Tree exists in  ID: " + treeId)
	}
}

func (api *ESApi) UpdateWeather(id string, weather model.Weather) {
	updateScript := elastic.NewScript(WEATHER_UPDATE_SCRIPT_SOURCE).Param(config.WEATHER_INDEX, weather)
	ctx := context.Background()
	_, err := api.Update().Index(config.GS_INDEX).Id(id).Type("_doc").Script(updateScript).Do(ctx)
	if err != nil {
		log.Println("Could update weather for tree in greenspot index. Id: " + id)
		log.Fatal(err)
	}
}

func (api *ESApi) UpdateStatus(id string, status model.Status) {
	updateScript := elastic.NewScript(STATUS_UPDATE_SCRIPT_SOURCE).Param(config.STATUS_INDEX, status)
	ctx := context.Background()
	_, err := api.Update().Index(config.GS_INDEX).Id(id).Type("_doc").Script(updateScript).Do(ctx)
	if err != nil {
		log.Println("Could update status for tree in greenspot index. Id: " + id)
		log.Fatal(err)
	}
}

func (api *ESApi) UpdateAirquality(id string, status model.AirQuality) {
	updateScript := elastic.NewScript(AIRQUALITY_UPDATE_SCRIPT_SOURCE).Param(config.AIRQUALITY_INDEX, status)
	ctx := context.Background()
	_, err := api.Update().Index(config.GS_INDEX).Id(id).Type("_doc").Script(updateScript).Do(ctx)
	if err != nil {
		log.Println("Could update airquality for tree in greenspot index. Id: " + id)
		log.Fatal(err)
	}
}

func (api *ESApi) IndexWeather(weather model.Weather) string {
	return api.indexNewDocument(config.WEATHER_INDEX, weather)
}

func (api *ESApi) IndexStatus(status model.Status) string {
	return api.indexNewDocument(config.STATUS_INDEX, status)
}

func (api *ESApi) IndexAirQuality(status model.AirQuality) string {
	return api.indexNewDocument(config.AIRQUALITY_INDEX, status)
}

func (api *ESApi) indexNewDocument(index string, body interface{}) string {
	indexed, err := api.
		Index().
		Index(index).
		Type("_doc").
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		log.Println("Could not index event")
		log.Println(err)
	}
	return indexed.Id
}
