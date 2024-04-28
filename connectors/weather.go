package connectors

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	config "github.com/joaogabsoaresf/go-weather/config"
)

type Weather struct {
	City        string
	Temperature int64
}

const URL string = "https://api.openweathermap.org/data/2.5/weather?&units=metric"

func GetWethear(city string) (map[string]interface{}, error) {
	u := parseUrl(URL, "q", city)
	return getWethearApiRequest(u)
}

func getWethearApiRequest(url string) (map[string]interface{}, error) {
	secrets := config.GetSecrets()
	u := parseUrl(url, "appid", secrets.OPEN_WEATHER_API)
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decodificando o JSON da resposta
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func parseUrl(u string, param string, value string) string {
	newU, err := url.Parse(u)
	if err != nil {
		return "erro"
	}
	q := newU.Query()
	q.Set(param, value)
	newU.RawQuery = q.Encode()
	return newU.String()
}
