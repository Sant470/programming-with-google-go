package main
import (
  "fmt"
  "os"
  "strings"
  "bufio"
)

type Animal interface {
  Eat()
  Move()
  Speak()
}

// Cow type
type Cow struct {
  food string
  locomotion string
  noise string
}

func (cow Cow) Eat() { fmt.Println(cow.food) }

func (cow Cow) Move() { fmt.Println(cow.locomotion) }

func (cow Cow) Speak() { fmt.Println(cow.noise) }


// Bird type
type Bird struct {
  food string
  locomotion string
  noise string
}

func (bird Bird) Eat() { fmt.Println(bird.food) }

func (bird Bird) Move() { fmt.Println(bird.locomotion) }

func (bird Bird) Speak() { fmt.Println(bird.noise) }


// Snake type
type Snake struct {
  food string
  locomotion string
  noise string
}

func (snake Snake) Eat() { fmt.Println(snake.food) }

func (snake Snake) Move() { fmt.Println(snake.locomotion) }

func (snake Snake) Speak() { fmt.Println(snake.noise) }



func main() {
  animals := make(map[string]Animal)
  for {
    fmt.Print("> ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    inputs := strings.Split(input[0:len(input)-1], " ")
    if inputs[0] == "newanimal" {
      name := inputs[1]
      animalType := inputs[2]
      if animalType == "cow" || animalType == "bird" || animalType == "snake" {
          var animal Animal
          if animalType == "cow" { animal = Cow{"grass", "walk", "moo"} }
          if animalType == "bird" { animal = Bird{"worms", "fly", "peep"} }
          if animalType == "snake" {animal = Snake{"mice", "slither", "hsss"} }
          animals[name] = animal
          fmt.Println("Created it!")
      }
    }
    if inputs[0] == "query" {
      name := inputs[1]
      information := inputs[2]
      if information == "eat" {animals[name].Eat()}
      if information == "move" {animals[name].Move()}
      if information == "speak" {animals[name].Speak()}
    }
  }
}
