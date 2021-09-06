package main
import (
  "fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
  return func(t float64) float64 {
    finalDisp := s0 + v0* t + (a/2)*t*t
    return finalDisp
  }
}


func main() {
  var initDisp, vel, acl, time float64
  fmt.Println("Enter acceleration: ")
  _, err1 := fmt.Scanf("%f", &acl)
  if err1 != nil {
    fmt.Println("Incorrect input for acceleration", err1)
    return
  }
  fmt.Println("Enter Velocity: ")
  _, err2 := fmt.Scanf("%f", &vel)
  if err2 != nil {
    fmt.Println("Incorrect input for Velocity", err2)
    return
  }
  fmt.Println("Enter initial displacement: ")
  _, err3 := fmt.Scanf("%f", &initDisp)
  if err3 != nil {
    fmt.Println("Incorrect input for initial displacement", err3)
    return
  }
  fn := GenDisplaceFn(acl, vel, initDisp);
  fmt.Println("Please Enter time")
  _, err := fmt.Scanf("%f", &time)
  if err != nil {
    fmt.Println("Incorrect input for time", err)
    return
  }
  fmt.Println(fn(time));
}
