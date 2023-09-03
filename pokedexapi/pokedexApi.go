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

type AreasApiResponse struct {
  Count int
  Next *string
  Previous *string
  Results []Location
}

type Pokemon struct {
	Name string
	Url string
}

type PokemonEncounter struct {
	Pokemon Pokemon
}

type AreasPokemonApiResponse struct {
  Pokemon_encounters []PokemonEncounter
}

type PokemonApiResponse struct {
  Name string
	Height int
	Weight int
	Base_experience int
}

func FetchAreasForUrl(pokeCache pokecache.PokeCache, url string)(response AreasApiResponse) {
	cacheResponse, ok := pokeCache.Get(url)
	if ok {
		responseStruct := AreasApiResponse{}
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
  responseStruct := AreasApiResponse{}
  err = json.Unmarshal(body, &responseStruct)
  if err != nil {
    log.Fatal(err)
  }
  return responseStruct
}


func FetchPokemonsForUrl(pokeCache pokecache.PokeCache, url string)(response AreasPokemonApiResponse) {
	cacheResponse, ok := pokeCache.Get(url)
	if ok {
		responseStruct := AreasPokemonApiResponse{}
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
  responseStruct := AreasPokemonApiResponse{}
  err = json.Unmarshal(body, &responseStruct)
  if err != nil {
    log.Fatal(err)
  }
  return responseStruct
}

func FetchPokemonsInfoForUrl(pokeCache pokecache.PokeCache, url string)(response PokemonApiResponse) {
	cacheResponse, ok := pokeCache.Get(url)
	if ok {
		responseStruct := PokemonApiResponse{}
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
  responseStruct := PokemonApiResponse{}
  err = json.Unmarshal(body, &responseStruct)
  if err != nil {
    log.Fatal(err)
  }
  return responseStruct
}
