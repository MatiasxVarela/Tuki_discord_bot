package services

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRandomCat() (io.ReadCloser, error) {
	CAT_API := os.Getenv("CAT_API")

	res, err := http.Get("https://api.thecatapi.com/v1/images/search?api_key=" + CAT_API)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	url := data[0]["url"].(string)
	res, err = http.Get(url)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
