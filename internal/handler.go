package flymetarutility

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	icao := query.Get("icao")
	infoType := query.Get("type")

	var response string

	if icao == "" {
		response = "Please provide an ICAO code"
	} else {
		metar := GetMetar(icao)
		rawMetar := metar.RawOb.(string)
		rawTaf := metar.RawTaf.(string)

		switch infoType {
		case "taf":
			response = rawTaf
		case "metar":
			response = rawMetar
		default:
			response = fmt.Sprintf("%s\n%s", rawMetar, rawTaf)
		}
	}

	w.Write([]byte(response))
}
