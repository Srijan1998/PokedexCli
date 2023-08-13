package main

import "fmt"

func exitCommand() {
  fmt.Println("Exiting")
}

func helpCommand(){
  commands := GetCliCommandsMap()
  for _, element := range commands {
    fmt.Printf("%v: %v\n", element.name, element.description)
  }
}

type cliCommand struct {
  name string
  description string
  callback func()
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
  }
}
