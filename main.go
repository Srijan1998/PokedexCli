package main
import(
  "fmt"
  // "commandHandler"
)

func main() {
  baseUrl := "https://pokeapi.co/api/v2/location-area"
  config := commandConfig{"", &baseUrl, nil}
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
