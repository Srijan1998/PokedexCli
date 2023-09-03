package main
import(
  "fmt"
  pokeCache "github.com/Srijan1998/pokedexcli/pokecache"
)

func main() {
  baseUrl := "https://pokeapi.co/api/v2/location-area"
  pokeCache := pokeCache.NewCache(10)
  config := commandConfig{"", &baseUrl, nil, pokeCache}
  for {
    var userInput string
    fmt.Print("pokedex > ")
    fmt.Scanln(&userInput)
    fmt.Println("")
    command, ok := GetCliCommandsMap()[userInput]
    if !ok {
      fmt.Println("Invalid command")
      continue
    }

    config.userInput = userInput
    config = command.callback(config)
    if userInput == "exit" {
      break
    }
  }
}
