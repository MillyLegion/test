package main
import "fmt"
func Max(a ... int ) int {
   fmt.Printf("Arrayvalues: %v, %v, %v\n", a, b, c)
   if a > b && a > c {
     return a
   } else if b>a && b > c {
     return b
   }
   return c
}

func main() {
   println(Max(10,20,30))
}
