package main
import(
  "fmt"
  // "commandHandler"
)

func main() {
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
    command.callback()
    if userInput == "exit" {
      break
    }
  }
}
