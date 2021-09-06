package main
import (
  "fmt"
  "os"
  "strings"
  "bufio"
)

type Animal struct {
  food string
  locomotion string
  noise string
}


func (animal *Animal) Eat() {
  fmt.Println(animal.food)
}

func (animal *Animal) Move() {
  fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
  fmt.Println(animal.noise)
}


func main() {
  for {
    fmt.Print("> ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    inputs := strings.Split(input[0:len(input)-1], " ")
    if inputs[0] == "cow" || inputs[0] == "bird" || inputs[0] == "snake" {
      if inputs[1] == "eat" || inputs[1] == "move" || inputs[1] == "speak" {
        var animal Animal
        if inputs[0] == "cow" { animal = Animal{"grass", "walk", "moo"} }
        if inputs[0] == "bird" { animal = Animal{"worms", "fly", "peep"} }
        if inputs[0] == "snake" {animal = Animal{"mice", "slither", "hsss"} }
        if inputs[1] == "eat" { animal.Eat() }
        if inputs[1] == "move" {animal.Move() }
        if inputs[1] == "speak" {animal.Speak() }
      }
    } else {
      fmt.Println("Wrong input !!!")
    }
  }
}
