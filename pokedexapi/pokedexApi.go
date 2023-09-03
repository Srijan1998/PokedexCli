package pokedexapi
import (
	"io"
	"net/http"
  "log"
  "encoding/json"
	pokecache "github.com/Srijan1998/pokedexcli/pokecache"
)

type Location struct {
  Name string
  Url string
}

type ApiResponse struct {
  Count int
  Next *string
  Previous *string
  Results []Location
}

func FetchForUrl(pokeCache pokecache.PokeCache, url string)(response ApiResponse) {
	cacheResponse, ok := pokeCache.Get(url)
	if ok {
		responseStruct := ApiResponse{}
		json.Unmarshal(cacheResponse, &responseStruct)
	  return responseStruct
	}
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
	pokeCache.Add(url, body)
  responseStruct := ApiResponse{}
  err = json.Unmarshal(body, &responseStruct)
  if err != nil {
    log.Fatal(err)
  }
  return responseStruct
}
