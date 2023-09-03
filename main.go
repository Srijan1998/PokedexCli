package main
import(
  "fmt"
  pokeCache "github.com/Srijan1998/pokedexcli/pokecache"
)

func main() {
  baseUrl := "https://pokeapi.co/api/v2/location-area"
  pokeCache := pokeCache.NewCache(10)
  config := commandConfig{"", &baseUrl, nil, pokeCache, nil}
  for {
    var commandStr string
    var areaStr string
    fmt.Print("pokedex > ")
    fmt.Scanln(&commandStr, &areaStr)
    fmt.Println("")
    command, ok := GetCliCommandsMap()[commandStr]
    if !ok {
      fmt.Println("Invalid command")
      continue
    }

    config.commandStr = commandStr
    config.areaStr = &areaStr
    config = command.callback(config)
    if commandStr == "exit" {
      break
    }
  }
}
