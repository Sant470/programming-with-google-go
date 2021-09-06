package main
import (
  "fmt"
  "strconv"
  "sort"
  "os"
  "bufio"
  "strings"
  "sync"
)

func sortSlice(sli []int, wg *sync.WaitGroup) {
  fmt.Println(sli)
  sort.Ints(sli)
  wg.Done()
}

func mergeSortedArray(sli1 []int, sli2 []int) []int {
  sortedSlice := make([]int, 0, len(sli1)+len(sli2))
  i, j:= 0, 0
  for i < len(sli1) && j < len(sli2) {
    if(sli1[i] < sli2[j]) {
      sortedSlice = append(sortedSlice, sli1[i])
      i++
    } else {
      sortedSlice = append(sortedSlice, sli2[j])
      j++
    }
  }
  for i < len(sli1) {
    sortedSlice = append(sortedSlice, sli1[i])
    i++
  }
  for j < len(sli2) {
    sortedSlice = append(sortedSlice, sli2[j])
    j++
  }
  return sortedSlice
}

func main() {
  var wg sync.WaitGroup
  fmt.Println("Please enter integers to sort: ")
  reader := bufio.NewReader(os.Stdin)
  input, _ := reader.ReadString('\n')
  inputs := strings.Split(input[0:len(input)-1], " ")
  inputSlice := make([]int, 0, 20)
  for _, input := range inputs {
    curr, err := strconv.Atoi(input)
    if err != nil { continue }
    inputSlice = append(inputSlice, curr)
  }
  subSliceSize := len(inputSlice)/4
  for i :=0 ; i < len(inputSlice); {
    max := i + subSliceSize
    if(max + subSliceSize > len(inputSlice)) { max = len(inputSlice) }
    wg.Add(1)
    go sortSlice(inputSlice[i: max], &wg)
    i = max
  }
  wg.Wait()
  sortedSlice1 := mergeSortedArray(inputSlice[0: subSliceSize], inputSlice[subSliceSize: 2*subSliceSize])
  sortedSlice2 := mergeSortedArray(inputSlice[2*subSliceSize: 3*subSliceSize], inputSlice[3*subSliceSize: len(inputSlice)])
  sortedArray := mergeSortedArray(sortedSlice1, sortedSlice2)
  fmt.Println(sortedArray)
}
