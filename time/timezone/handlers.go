package timezone

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const timeAPI = "https://timeapi.io/api/TimeZone"

type TimeAPIresponse struct {
	TimeZone         string `json:timeZone xml:timeZone`
	CurrentLocalTime string `json:currentLocalTime xml:timeZone`
}

func getTime(w http.ResponseWriter, r *http.Request) {

	tzs := strings.Split(r.URL.Query().Get("zone"), ",")

	//edge case, no timezone provided => default to UTC
	if len(tzs[0]) == 0 {
		tzs = []string{"UTC"}
	}

	result := map[string]string{}

	for i := 0; i < len(tzs); i++ {

		resp, err := http.Get(timeAPI + "/zone?timeZone=" + tzs[i])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Timezone not found")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprint(w, "unable to parse response")
		}

		var resBody TimeAPIresponse

		if err := json.Unmarshal(body, &resBody); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprint(w, "unable to parse response")
		}

		if len(resBody.CurrentLocalTime) != 0 {
			result[tzs[i]] = resBody.CurrentLocalTime
		}

	}
	//handle bad time zone provided
	if len(result) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Timezone not found")
	} else {

		if r.Header.Get("Content-Type") == "application/xml" {
			xmlOutput := []TimeAPIresponse{}
			for tz, time := range result {
				el := TimeAPIresponse{
					TimeZone:         tz,
					CurrentLocalTime: time,
				}
				xmlOutput = append(xmlOutput, el)
			}

			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(xmlOutput)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		}
	}
}

/*
example usage: http://localhost:8000/api/time?zone=Europe/Amsterdam,America/New_York,Asia/Kolkata
returns json
{
    "America/New_York": "2024-03-29T18:22:24.261251",
    "Asia/Kolkata": "2024-03-30T03:52:24.3845979",
    "Europe/Amsterdam": "2024-03-29T23:22:24.1332437"
}
returns xml
<TimeAPIresponse>
    <TimeZone>Europe/Amsterdam</TimeZone>
    <CurrentLocalTime>2024-03-29T23:33:01.6197779</CurrentLocalTime>
</TimeAPIresponse>
<TimeAPIresponse>
    <TimeZone>America/New_York</TimeZone>
    <CurrentLocalTime>2024-03-29T18:33:01.7892384</CurrentLocalTime>
</TimeAPIresponse>
<TimeAPIresponse>
    <TimeZone>Asia/Kolkata</TimeZone>
    <CurrentLocalTime>2024-03-30T04:03:01.9616736</CurrentLocalTime>
</TimeAPIresponse>

if no timezone is provided, returns UTC
*/
