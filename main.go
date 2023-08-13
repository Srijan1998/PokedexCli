package main
import(
  "fmt"
)

func main() {
  for {
    var userInput string
    fmt.Print("pokedex > ")
    fmt.Scanln(&userInput)
    fmt.Println("")
    command, ok := getCliCommandsMap(helpCommand, exitCommand)[userInput]
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

func exitCommand() {
  fmt.Println("Exiting")
}

func helpCommand(){
  commands := getCliCommandsMap(helpCommand, exitCommand)
  for _, element := range commands {
    fmt.Printf("%v: %v\n", element.name, element.description)
  }
}

type cliCommand struct {
  name string
  description string
  callback func()
}

func getCliCommandsMap(helpCallback, exitCallback func())(map[string]cliCommand) {
  return map[string]cliCommand {
    "help": {
      name: "help",
      description: "Prints help section",
      callback: helpCallback,
    },
    "exit": {
      name: "exit",
      description: "Exits program",
      callback: exitCallback,
    },
  }
}
