package main
import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
)

func Swap(sli []int, index int) {
  temp := sli[index]
  sli[index] = sli[index+1]
  sli[index+1] = temp
}

func BubbleSort(sli []int) {
  for i:=0 ; i < len(sli); i++ {
    for j:=0; j< len(sli) - 1 - i; j++ {
      if sli[j] > sli[j+1] {
        Swap(sli, j)
      }
    }
  }
}

func main() {
  index :=0
  inputSlice := make([]int, 0, 10)
  fmt.Println("Enter upto 10 integers to sort")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  inputs := strings.Split(input[0:len(input)-1], " ")
  for _, input := range inputs {
    if index >= 10 {
      break;
    }
    input, err := strconv.Atoi(input)
    if err != nil {
      continue
    }
    inputSlice = append(inputSlice, input)
  }
  BubbleSort(inputSlice)
  fmt.Println(inputSlice)
}
