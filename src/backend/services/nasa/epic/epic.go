// https://epic.gsfc.nasa.gov/about/api
package epic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Epic fields ...
// https://api.nasa.gov/EPIC/api/natural?api_key=DEMO_KEY
type EpicItem struct {
	Identifier          string              `json:"identifier"`
	Caption             string              `json:"caption"`
	Image               string              `json:"image"`
	Version             string              `json:"version"`
	CentroidCoordinates CentroidCoordinates `json:"centroid_coordinates"`
	DscovrJ2000Position DscovrJ2000Position `json:"dscovr_j2000_position"`
	LunarJ2000Position  LunarJ2000Position  `json:"lunar_j2000_position"`
	SunJ2000Position    SunJ2000Position    `json:"sun_j2000_position"`
	AttitudeQuaternions AttitudeQuaternions `json:"attitude_quaternions"`
	Date                string              `json:"date"`
	Coords              Coords              `json:"coords"`
	// Optional fields
	ImageUrl string `json:"image_url"`
}

type CentroidCoordinates struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
type DscovrJ2000Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type LunarJ2000Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type SunJ2000Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}
type AttitudeQuaternions struct {
	Q0 float32 `json:"q0"`
	Q1 float32 `json:"q1"`
	Q2 float32 `json:"q2"`
	Q3 float32 `json:"q3"`
}
type Coords struct {
	CentroidCoordinates CentroidCoordinates `json:"centroid_coordinates"`
	DscovrJ2000Position DscovrJ2000Position `json:"dscovr_j2000_position"`
	LunarJ2000Position  LunarJ2000Position  `json:"lunar_j2000_position"`
	SunJ2000Position    SunJ2000Position    `json:"sun_j2000_position"`
	AttitudeQuaternions AttitudeQuaternions `json:"attitude_quaternions"`
}

// EpicResponseItems cache ...
var EpicResponseItems []EpicItem

// Gets the epic data
func GetEpics() []EpicItem {
	items := fetchEpicAPI()
	for i, _ := range items {
		dateParts := strings.Split(items[i].Date, " ")
		datePart := strings.Replace(dateParts[0], "-", "/", -1)
		items[i].ImageUrl = fmt.Sprintf("https://epic.gsfc.nasa.gov/archive/natural/%s/png/%s.png", datePart, items[i].Image)
		fmt.Println(items[i].ImageUrl)
	}
	return items
}

func fetchEpicAPI() []EpicItem {
	//apiKey := ""
	apiEndpointQueryUrl := "https://api.nasa.gov/EPIC/api/natural?api_key=DEMO_KEY"
	EpicResponseItems := getJsonDataFromAPIUrl(apiEndpointQueryUrl)

	return EpicResponseItems
}

// @see: https://blog.alexellis.io/golang-json-api-client/
func getJsonDataFromAPIUrl(url string) []EpicItem {
	httpClient := http.Client{
		Timeout: time.Second * 15,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var responseData []EpicItem
	jsonErr := json.Unmarshal(body, &responseData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return responseData
}
