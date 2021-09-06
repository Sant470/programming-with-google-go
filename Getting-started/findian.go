package main
import (
  "fmt"
  s "strings"
)


func main() {
  var input string
  fmt.Scan(&input)
  comparableInput := s.ToLower(input)
  if(s.HasPrefix(comparableInput, "i") && s.HasSuffix(comparableInput, "n") && s.Contains(comparableInput, "a")) {
    fmt.Println("Found!")
  } else {
    fmt.Println("Not Found!")
  }
}
