package main

import "fmt"
import "./min"
import "./max"

func main(){
  fmt.Println("1,2,3 Min is ", min.Min(1,2,3))
  fmt.Println("1,2,3 Max is ", max.Max(1,2,3))
}
