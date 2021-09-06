package main
import (
  "fmt"
  s "strconv"
  sort "sort"
)

func main() {
  inputSlice := make([]int, 0, 3)
  for {
    var input string
    fmt.Println("Please enter an integer to add in slice")
    fmt.Scan(&input)
    if(input == "X" || input == "x") {
      break
    }
    num, err := s.Atoi(input)
    if (err != nil) {
      fmt.Println("entered number is", num)
      continue
    }
    inputSlice = append(inputSlice, num)
    sort.Ints(inputSlice)
    fmt.Println("input slice: ", inputSlice)
  }
}
