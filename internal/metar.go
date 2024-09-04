package flymetarutility

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type METAR struct {
	MetarID     interface{} `json:"metar_id"`
	IcaoID      interface{} `json:"icaoId"`
	ReceiptTime interface{} `json:"receiptTime"`
	ObsTime     interface{} `json:"obsTime"`
	ReportTime  interface{} `json:"reportTime"`
	Temp        interface{} `json:"temp"`
	Dewp        interface{} `json:"dewp"`
	Wdir        interface{} `json:"wdir"`
	Wspd        interface{} `json:"wspd"`
	Wgst        interface{} `json:"wgst"`
	Visib       interface{} `json:"visib"`
	Altim       interface{} `json:"altim"`
	Slp         interface{} `json:"slp"`
	QcField     interface{} `json:"qcField"`
	WxString    interface{} `json:"wxString"`
	PresTend    interface{} `json:"presTend"`
	MaxT        interface{} `json:"maxT"`
	MinT        interface{} `json:"minT"`
	MaxT24      interface{} `json:"maxT24"`
	MinT24      interface{} `json:"minT24"`
	Precip      interface{} `json:"precip"`
	Pcp3Hr      interface{} `json:"pcp3hr"`
	Pcp6Hr      interface{} `json:"pcp6hr"`
	Pcp24Hr     interface{} `json:"pcp24hr"`
	Snow        interface{} `json:"snow"`
	VertVis     interface{} `json:"vertVis"`
	MetarType   interface{} `json:"metarType"`
	RawOb       interface{} `json:"rawOb"`
	MostRecent  interface{} `json:"mostRecent"`
	Lat         interface{} `json:"lat"`
	Lon         interface{} `json:"lon"`
	Elev        interface{} `json:"elev"`
	Prior       interface{} `json:"prior"`
	Name        interface{} `json:"name"`
	Clouds      []struct {
		Cover interface{} `json:"cover"`
		Base  interface{} `json:"base"`
	} `json:"clouds"`
	RawTaf interface{} `json:"rawTaf"`
}

func GetMetar(icao string) METAR {
	url := fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%s&format=json&taf=true", icao)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var metar []METAR
	if err := json.Unmarshal(body, &metar); err != nil {
		log.Fatal(err)
	}

	return metar[0]
}
