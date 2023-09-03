package main

import "fmt"
import pokedexapi "github.com/Srijan1998/pokedexcli/pokedexapi"
import pokecache "github.com/Srijan1998/pokedexcli/pokecache"

type commandConfig struct {
  userInput string
  nextUrl *string
  prevUrl *string
  pokeCache pokecache.PokeCache
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
  response := pokedexapi.FetchForUrl(config.pokeCache, *config.nextUrl)
  config.nextUrl = response.Next
  config.prevUrl = response.Previous
  printLocations(response)
  return config
}

func mapBackCommand(config commandConfig)(commandConfig){
  if config.prevUrl == nil {
    fmt.Println("No previous areas")
  } else {
    response := pokedexapi.FetchForUrl(config.pokeCache, *config.prevUrl)
    config.nextUrl = response.Next
    config.prevUrl = response.Previous
    printLocations(response)
  }
  return config
}

func printLocations(response pokedexapi.ApiResponse) {
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
  }
}
