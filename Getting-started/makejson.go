package main
import (
  "fmt"
  "encoding/json"
)

func main() {
  person := map[string]string {"name": "default","address": "default"}
  var name, address string
  fmt.Println("Please enter your name and address by space separated")
  fmt.Scan(&name)
  fmt.Scan(&address)
  person["name"] = name
  person["address"] = address
  byteArr, _:= json.Marshal(person)
  fmt.Println(string(byteArr))
}
