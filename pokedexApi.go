package main
import (
	"io"
	"net/http"
  "log"
  "encoding/json"
)

type location struct {
  Name string
  Url string
}

type apiResponse struct {
  Count int
  Next *string
  Previous *string
  Results []location
}

func fetchForUrl(url string)(response apiResponse) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
  responseStruct := apiResponse{}
  err = json.Unmarshal(body, &responseStruct)
  if err != nil {
    log.Fatal(err)
  }
  return responseStruct
}
