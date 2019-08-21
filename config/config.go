package config

var (
	INDEXES          = [4]string{"greenspot", "status", "weather", "airquality"}
	TREE_NAMES       = [8]string{"Acacia", "Tilia", "Acer", "Abies", "Pinus", "Salix", "Quercus", "Corylus"}
	TREE_IDS         = [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	ELASTIC_URL      = "http://localhost:9200"
	GS_INDEX         = "greenspot"
	WEATHER_INDEX    = "weather"
	STATUS_INDEX     = "status"
	AIRQUALITY_INDEX = "airquality"

	API_SERVER_PORT = "8888"
	API_URL         = "http://localhost:8888"
)
