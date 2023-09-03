package main

import "fmt"
import pokedexapi "github.com/Srijan1998/pokedexcli/pokedexapi"
import pokecache "github.com/Srijan1998/pokedexcli/pokecache"
import "math/rand"
import "math"

type commandConfig struct {
  commandStr string
  nextUrl *string
  prevUrl *string
  pokeCache pokecache.PokeCache
  secondArgument *string
  caughtPokemons map[string]pokedexapi.PokemonApiResponse
}

func exitCommand(config commandConfig)(commandConfig) {
  fmt.Println("Exiting")
  return config
}

func helpCommand(config commandConfig)(commandConfig){
  commands := GetCliCommandsMap()
  for _, element := range commands {
    fmt.Printf("%v: %v\n", element.name, element.description)
  }
  return config
}

func mapCommand(config commandConfig)(commandConfig){
  response := pokedexapi.FetchAreasForUrl(config.pokeCache, *config.nextUrl)
  config.nextUrl = response.Next
  config.prevUrl = response.Previous
  printLocations(response)
  return config
}

func mapBackCommand(config commandConfig)(commandConfig){
  if config.prevUrl == nil {
    fmt.Println("No previous areas")
  } else {
    response := pokedexapi.FetchAreasForUrl(config.pokeCache, *config.prevUrl)
    config.nextUrl = response.Next
    config.prevUrl = response.Previous
    printLocations(response)
  }
  return config
}

func exploreCommand(config commandConfig)(commandConfig) {
  response := pokedexapi.FetchPokemonsForUrl(config.pokeCache, "https://pokeapi.co/api/v2/location-area/" + *config.secondArgument)
  printPokemons(response)
  return config
}

func catchCommand(config commandConfig)(commandConfig) {
  response := pokedexapi.FetchPokemonsInfoForUrl(config.pokeCache, "https://pokeapi.co/api/v2/pokemon/" + *config.secondArgument)
  probability := math.Max(1/float64((response.Base_experience - 100)), 0.2)
  randNum := rand.Float64()
   if probability > randNum {
     fmt.Println("Caught")
     config.caughtPokemons[response.Name] = response
   } else {
     fmt.Println("Escaped")
   }
  return config
}

func inspectCommand(config commandConfig)(commandConfig) {
  val, ok := config.caughtPokemons[*config.secondArgument]
  if ok {
    fmt.Println(val)
  } else {
    fmt.Printf("Haven't caught %s\n", *config.secondArgument)
  }
  return config
}

func pokedexCommand(config commandConfig)(commandConfig) {
  for name, _ := range config.caughtPokemons {
    fmt.Println(name)
  }
  return config
}

func printPokemons(response pokedexapi.AreasPokemonApiResponse) {
  for _, encounter := range response.Pokemon_encounters {
    fmt.Println(encounter.Pokemon.Name)
  }
}

func printLocations(response pokedexapi.AreasApiResponse) {
  for _, location := range response.Results {
    fmt.Println(location.Name)
  }
}

type cliCommand struct {
  name string
  description string
  callback func(config commandConfig)(commandConfig)
}

func GetCliCommandsMap()(map[string]cliCommand) {
  return map[string]cliCommand {
    "help": {
      name: "help",
      description: "Prints help section",
      callback: helpCommand,
    },
    "exit": {
      name: "exit",
      description: "Exits program",
      callback: exitCommand,
    },
    "map": {
      name: "map",
      description: "Displays next location areas",
      callback: mapCommand,
    },
    "mapback": {
      name: "Map Back",
      description: "Displays previous location areas",
      callback: mapBackCommand,
    },
    "explore": {
      name: "Explore",
      description: "Explore area for pokemons",
      callback: exploreCommand,
    },
    "catch": {
      name: "Catch",
      description: "Catch a pokemon",
      callback: catchCommand,
    },
    "inspect": {
      name: "Inspect",
      description: "Inspect a caught pokemon",
      callback: inspectCommand,
    },
    "pokedex": {
      name: "Pokedex",
      description: "Lists all caught pokemon",
      callback: pokedexCommand,
    },
  }
}
