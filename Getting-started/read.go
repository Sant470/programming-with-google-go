package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

type name struct {
  fname string
  lname string
}

func main() {
  var names = make([] name, 0, 10)
  var filename string
  fmt.Println("Please enter the filename to read!")
  _, err:= fmt.Scan(&filename)
  if(err != nil) {
    fmt.Println("Incorrect input file!!", err)
  }
  file, _ := os.Open(filename)
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    nameArr:= strings.Split(text, " ")
    currentName := name {nameArr[0], nameArr[1]}
    names = append(names, currentName)
  }
  for _, name := range names {
    fmt.Println("fname: ", name.fname, "lname: ", name.lname)
  }
}
