package main
import(
  "fmt"
  pokeCache "github.com/Srijan1998/pokedexcli/pokecache"
  pokedexapi "github.com/Srijan1998/pokedexcli/pokedexapi"
)

func main() {
  baseUrl := "https://pokeapi.co/api/v2/location-area"
  pokeCache := pokeCache.NewCache(10)
  config := commandConfig{"", &baseUrl, nil, pokeCache, nil, make(map[string]pokedexapi.PokemonApiResponse)}
  for {
    var commandStr string
    var secondArgument string
    fmt.Print("pokedex > ")
    fmt.Scanln(&commandStr, &secondArgument)
    fmt.Println("")
    command, ok := GetCliCommandsMap()[commandStr]
    if !ok {
      fmt.Println("Invalid command")
      continue
    }

    config.commandStr = commandStr
    config.secondArgument = &secondArgument
    config = command.callback(config)
    if commandStr == "exit" {
      break
    }
  }
}
