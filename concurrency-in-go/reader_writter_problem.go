/*
It's an example of what possibly could go wrong if we have more than one thread writing to the shared variable at a time
*/
package main
import (
  "fmt"
  "sync"
)

var i int = 0
var wg sync.WaitGroup


func incr() {
  i = i + 1
  wg.Done()
}

func main() {
  wg.Add(100)
  for j:=0; j< 100; j++ {
    go incr()
  }
  wg.Wait()
  fmt.Println(i) // expected value is 100
}
